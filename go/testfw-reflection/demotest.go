package main

import (
	"fmt"

	"github.com/egustafson/go/testfw-reflection/demofwk"
)

type DemoTestSuite struct{}

func init() {
	demofwk.Register(new(DemoTestSuite))
}

func (dts *DemoTestSuite) TestFuncOne() {
	fmt.Println("TestFuncOne")
}
