package main

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

type Root struct {
	RootID string
	Child1 *ChildA
	Child2 *ChildB
	Value  string
}

type ChildA struct {
	Avalue       string
	Children     []*ChildB
	AnotherValue string
}

type ChildB struct {
	StrValue string
	IntValue int
}

func Walk(root interface{}) {
	wl := make([]reflect.Value, 0, 10)
	wl = append(wl, reflect.ValueOf(root).Elem())
	for len(wl) > 0 {
		fmt.Println("--")
		item := wl[0]
		if item.Kind() == reflect.Ptr {
			fmt.Println("dereferencing ptr")
			item = item.Elem()
		}
		fmt.Printf("Type: %s\n", item.Type())
		if item.Kind() == reflect.Struct {
			for ii := 0; ii < item.NumField(); ii++ {
				name := item.Type().Field(ii).Name
				fmt.Printf("adding: %s\n", name)
				val := item.Field(ii)
				wl = append(wl, val)
			}
		}
		if item.Kind() == reflect.Slice {
			fmt.Println("slice found")
			for ii := 0; ii < item.Len(); ii++ {
				fmt.Printf("adding element: %d\n", ii)
				val := item.Index(ii)
				wl = append(wl, val)
			}
		}
		if item.Kind() == reflect.String {
			fmt.Println("string found")
			if item.CanSet() != true {
				fmt.Println("*** unsetable")
			} else {
				item.SetString("new-value")
			}
		}
		wl = wl[1:]
		fmt.Printf("len(wl) = %d\n", len(wl))
	}
}

func Print(root interface{}) {
	d, err := yaml.Marshal(root)
	if err != nil {
		panic("yaml marshal error")
	}
	fmt.Printf("---\n%s...\n", string(d))
}

func main() {
	// placeholder
	fmt.Println("done.")
}
