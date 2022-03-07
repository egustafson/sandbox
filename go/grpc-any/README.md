grpc-any Example
----------------

Use the gRPC (google.protobuf.Any) type to pass an arbitrary type
across a pseudo call boundary.

Note:  there are older, historical ways.  This package attempts to
implement the current best practice, as of 2/2022.

* https://pkg.go.dev/google.golang.org/protobuf@v1.27.1/types/known/anypb

Some learning(s) along the way:

## Install BOTH `protoc` and the associated includes

The protoc compiler (https://github.com/protocolbuffers/protobuf)
comes with two things:

1. `protoc` the IDL compiler
2. Include protobuf files (.proto's)

If you are only compiling IDL you generated then only `protoc` is
required.  IF you wish to add core protobuf types, like `Any`, then
you need the includes as well.

Download the `protobuff-all-x.y.z.tar.gz` file, it includes both.

Docs:  https://grpc.io/docs/protoc-installation/

## Golang - the golang plugin

Each language requires a language specific plugin to generate protobuf
stubs for that language.

* https://developers.google.com/protocol-buffers/docs/reference/go-generated

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

The above will install `protoc-gen-go`, the golang plugin, into your
`$GOBIN` directory.

## Golang libraries

Additional golang libraries are needed.  The
`google.golang.org/protobuf` packages are newer, and preferred over
(the deprecated) `github.com/golang/protobuf` packages.

* https://pkg.go.dev/google.golang.org/protobuf@v1.27.1/types/known/anypb

The `anypb` package contains documentation describing how to marshal
and unmarshal between compiled (concrete) protobuf types and the `any`
type.  (as does this example)
