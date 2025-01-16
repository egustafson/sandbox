package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

const filename = "input.yaml"

var envvars = map[string]string{
	"TRANSLATE": "translated-value",
	"USERNAME":  "bogus-user",
	"PASSWORD":  "bogus-password",
}

func main() {

	// set envvars
	for k, v := range envvars {
		os.Setenv(k, v)
	}

	// load from YAML file
	bin, err := os.ReadFile("input.yaml")
	if err != nil {
		panic(err)
	}
	ymlObj := make(map[string]any)
	if err = yaml.Unmarshal(bin, ymlObj); err != nil {
		panic(err)
	}

	//
	// Perform Decoration -- the demo.
	//
	if err = Decorate(EnvMapFunc, &ymlObj); err != nil {
		panic(err)
	}

	// print decorated object as YAML
	bout, err := yaml.Marshal(ymlObj)
	if err != nil {
		panic(err)
	}
	fmt.Println("---")
	fmt.Print(string(bout))
	fmt.Println("...")
}

// MapFunc defines a function that maps an input string to an output string.
type MapFunc func(in string) string

// EnvMapFunc returns the value of $ENVIRONMENT_VARIABLE if present and the `in`
// string starts with '$', else it returns `in`
func EnvMapFunc(in string) string {
	if strings.HasPrefix(in, "$") {
		if replacement, ok := os.LookupEnv(in[1:]); ok {
			return replacement
		}
	}
	return in
}

// Decorate will traverse `root` and for each leaf element that is a string, it
// will attempt to modify that string element in situ, using `MapFunc`.
func Decorate(m MapFunc, root any) error {

	const initListSize = 30
	wl := make([]reflect.Value, 0, initListSize) // the "walk list" (stack recursion)
	// wl = append(wl, reflect.ValueOf(root).Elem())
	wl = append(wl, deref(reflect.ValueOf(root)))
	for len(wl) > 0 {
		item := deref(wl[0]) // dereference thru interface{} and ptrs
		wl = wl[1:]          // pop the head

		switch item.Kind() {
		case reflect.Struct:
			for ii := 0; ii < item.NumField(); ii++ { // traverse struct fields
				wl = append(wl, item.Field(ii)) // stack recursion
			}
		case reflect.Map:
			iter := item.MapRange()
			for iter.Next() {
				val := deref(iter.Value()) // dereference thru interface{} and ptrs
				if val.Kind() == reflect.String {
					newStr := m(val.String())
					item.SetMapIndex(iter.Key(), reflect.ValueOf(newStr))
				} else {
					wl = append(wl, iter.Value())
				}
			}
		case reflect.Slice:
			for ii := 0; ii < item.Len(); ii++ { // traverse slice elements
				wl = append(wl, item.Index(ii)) // stack recursion
			}
		case reflect.String:
			if !item.CanSet() {
				// item is likely private, consider this an error
				return errors.New("unable to set (unexported) string field")
			}
			item.SetString(m(item.String())) // decorate with MapFunc
		}
	}
	return nil
}

// deref dereferences through `interface{}` and pointers to return the value
// behind both.  It is recursive
func deref(item reflect.Value) reflect.Value {
	for { // recursion
		switch item.Kind() {
		case reflect.Ptr:
			item = item.Elem()
		case reflect.Interface:
			item = reflect.ValueOf(item.Interface())
		default:
			return item // <-- exit
		}
	}
}
