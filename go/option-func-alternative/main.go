package main

// This example is meant to demonstrate the Golang Option Pattern where the
// options do not form a closure on the object under construction, which is the
// more common example.  This example uses `optionFlag`, a bool as the holder of
// the options state.  A more complex example could use a struct, possibly
// `optionStruct`, if multiple options were needed.
//
// This example came about because of a need to control different behavior
// INSIDE a constructor and not leave traces of that behavior in the returned
// object.  Specifically:  the constructor would exit with error in the normal
// flow, but when WithContinueOnError() was passed as an option, the constructor
// would continue constructing even if an error was determined to exist during
// construction.  The specific error was _possibly_ recoverable after the
// constructed object was returned.

import "fmt"

type Object interface {
	Value() string
}

type object struct {
	value string
}

type ObjOption func(*bool)

func NewObject(options ...ObjOption) Object {
	var optionFlag bool = false // <-- option state INSIDE the constructor, could be a struct
	for _, o := range options {
		o(&optionFlag) // <-- set options based on an object that is NOT the constructed object.
	}
	obj := &object{
		value: "option-flag FALSE",
	}
	if optionFlag { // <-- do something different in construction based on option
		obj.value = "option-flag was set"
	}
	return obj
}

func WithOptionEnabled() ObjOption {
	return func(f *bool) {
		*f = true
	}
}

func (o *object) Value() string {
	return o.value
}

func main() {
	o := NewObject()
	fmt.Printf("1: obj.Value() = %s\n", o.Value())
	o = NewObject(WithOptionEnabled())
	fmt.Printf("2: obj.Value() = %s\n", o.Value())
}
