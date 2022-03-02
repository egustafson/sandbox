package main

// This is the Main Application in the example.  This package will
// directly manipulate the zerolog package.  This example has the
// following two additional packages:
//
// * clientlib - an example client lib package that logs
// * log - the example and prototype client lib logging package
//

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/egustafson/sandbox/go/zerolog-client/clientlib"
)

func main() {

	clientlib.ClientFuncDebug("log-block-1")
	clientlib.ClientFuncInfo("log-block-1")

	clientlib.ClientUpLevel()

	clientlib.ClientFuncDebug("log-block-2")
	clientlib.ClientFuncInfo("log-block-2")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	fmt.Println("--> set global level to: info")

	clientlib.ClientFuncDebug("log-block-3")
	clientlib.ClientFuncInfo("log-block-3")

	clientlib.ClientFuncDebug("log-block-4")
	clientlib.ClientFuncInfo("log-block-4")

	fmt.Println("done.")
}
