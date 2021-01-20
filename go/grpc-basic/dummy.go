package main_dummy

// MainDummy package is here to ensure that module dependancies are properly
// calculated in the go.mod.  Use `go mod tidy` to refresh.  (Inspiration from
// etcd source tree)
import (
	_ "github.com/egustafson/sandbox/go/grpc-basic/cmd/client"
	_ "github.com/egustafson/sandbox/go/grpc-basic/cmd/server"
)
