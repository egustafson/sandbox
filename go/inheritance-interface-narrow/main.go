package main

import (
	"fmt"
)

type Animal interface {
	Say() interface{}
}

// --
type CatNoise struct {
	Noise string
}

func (cn CatNoise) String() string { return cn.Noise }

type Cat interface {
	Animal
	CatSay() *CatNoise
}

type cat struct{}

func (c *cat) CatSay() *CatNoise { return &CatNoise{Noise: "meow"} }
func (c *cat) Say() interface{}  { return c.CatSay() }

// --

func obscureType(a Animal) Animal { return a }

func main() {

	a := obscureType(&cat{})
	fmt.Println(a.Say())
}
