package main

import (
	"fmt"
	"os"

	"github.com/aclindsa/ofxgo"
)

// https://github.com/aclindsa/ofxgo

func rdOFXGO() {

	f, err := os.Open(ofxFilename)
	if err != nil {
		panic("unable to open OFX file")
	}
	defer f.Close()

	resp, err := ofxgo.ParseResponse(f)
	if err != nil {
		fmt.Printf("can't parse response: %v\n", err)
		return
	}

	_ = resp
}
