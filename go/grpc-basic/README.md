Simple gRPC Example
===================

This example has two goals:
1. To show a _simplistic_ gRPC client and server in Go.
2. To show a possible directory layout and build (Makefile) strategy.

Note that meeting only goal 1 could simplify the example to three
files in a single directory -- this would not build with `go build` as
there would be two executables in the directory.  Rather this example
attempts to also show a starting point for code layout suitable for a
small project.  _Versioned APIs are not tackled here._

Prerequisites
-------------

1. `protoc` installed and in your $PATH.  The protocol buffer compiler.
   - https://developers.google.com/protocol-buffers/docs/downloads

2. `protoc-gen-go` installed and in your $PATH.  The plugin for Golang.
   - https://developers.google.com/protocol-buffers/docs/reference/go-generated

Usage
-----

### Build
```
> make build
(cd cmd/server; go build)
(cd cmd/client; go build)
ln -sf cmd/server/server .
ln -sf cmd/client/client .
>
```

### Running
In two separate windows, _client_ and _server_.

<ins>Server:</ins> (press ctl-C to exit)
```
> ./server
Go gRPC Tutorial
2021/01/20 15:50:51 Receive message body from client: Hello From Client.
```

<ins>Client:</ins>
```
> ./client
2021/01/20 15:50:51 response:  Hello from the Server
>
```


Citations & Reference
---------------------

* Example inspiration:
  https://tutorialedge.net/golang/go-grpc-beginners-tutorial/

* Versioned APIs (in Go):
  https://medium.com/swlh/building-apis-with-grpc-and-go-9a6d369d7ce

### gRPC and `protoc`

* `protoc` compiler flags:
  - (bottom of): https://developers.google.com/protocol-buffers/docs/proto3#generating
  - https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers

* gRPC `protoc`:
  https://grpc.io/docs/languages/go/basics/#generating-client-and-server-code

* Stack Overflow (_Correct format of protoc go\_package_):
  https://stackoverflow.com/questions/61666805/correct-format-of-protoc-go-package



<!--  LocalWords:  Makefile
 -->
