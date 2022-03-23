package main

import (
	"errors"
	"fmt"
)

type ErrType1 interface {
	error
	GetValue() int
}

type ErrType2 interface {
	error
	GetString() string
}

type ErrType3 interface {
	ErrType2
	GetValue() int
}

type errBase struct {
	msg string
}

type errT1 struct {
	errBase
	val int
}

type errT2 struct {
	errBase
	val string
}

type errT3 struct {
	errT2
	intVal int
}

func (e errBase) Error() string {
	return e.msg
}

func (e errT1) GetValue() int {
	return e.val
}

func (e errT2) GetString() string {
	return e.val
}

func (e errT3) GetValue() int {
	return e.intVal
}

func newE1() error {
	return &errT1{
		errBase: errBase{msg: "I am err type-1"},
		val:     1,
	}
}

func newE2() error {
	return &errT2{
		errBase: errBase{msg: "I am err type-2"},
		val:     "string-val",
	}
}

func newE3() error {
	return &errT3{
		errT2: errT2{
			errBase: errBase{msg: "I am err type-3"},
			val:     "string-val-e3",
		},
		intVal: 3,
	}
}

// --  main()  -------------------------

func main() {

	e1 := newE1()
	e1err := ErrType1(nil)
	if errors.As(e1, &e1err) {
		fmt.Println("e1 is an ErrType1")
	}

	e2 := newE2()
	e2err := ErrType2(nil)
	if errors.As(e2, &e2err) {
		fmt.Println("e2 is an ErrType2")
	}

	e3 := newE3()
	e3err := ErrType3(nil)
	if errors.As(e3, &e3err) {
		fmt.Println("e3 is an ErrType3")
	}
	if errors.As(e3, &e2err) {
		fmt.Println("e3 is an ErrType2 as well")
		fmt.Printf("e3.val: %v\n", e2err.GetString())
	}

}
