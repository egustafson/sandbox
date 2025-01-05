package main

import (
	"bufio"
	"fmt"
	"os"

	ofx "github.com/rockstardevs/goofx"
)

// const ofxFilename = "samples/qfx/ally-alldates.qfx"
// const ofxFilename = "samples/qfx/wells-checking-20230801-20241231.qfx"
// const ofxFilename = "samples/qfx/citi-mc-20240101-20241231.qfx"
const ofxFilename = "samples/qfx/wealthfront-savings-20221101-20250105.qfx"

func rdOFX() {

	f, err := os.Open(ofxFilename)
	if err != nil {
		panic("unable to open OFX file")
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	document, err := ofx.NewDocumentFromXML(rd, ofx.NewCleaner())
	if err != nil {
		panic("unable to parse OFX file")
	}
	fmt.Printf("%#v", document)
}
