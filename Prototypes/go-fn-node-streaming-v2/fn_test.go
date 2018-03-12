package fnnode

import (
	"fmt"
	"log"
	_ "strings"
	_ "testing"
)

type testMsg struct {
	body string
}

func (tm testMsg) Body() interface{} {
	return tm.body
}

func ExampleUnityFn() {
	s := MakeNilSpec()
	unity, err := MakeUnityFn(s)
	if err != nil {
		log.Fatal("MakeUnityFn() returned an error")
	}
	msg := testMsg{body: "test-body"}
	resp := unity(msg)
	fmt.Println(resp.Body().(string))
	// Output:
	// test-body
}
