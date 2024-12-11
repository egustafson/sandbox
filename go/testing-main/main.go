package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("Starting tests")

	m := testing.MainStart(mockTestDeps{},
		[]testing.InternalTest{{"Demo-Testcase", RunDemoTestSet1}},
		nil, nil, nil,
	)

	result := m.Run()
	fmt.Println(result)

	fmt.Println("done.")
}
