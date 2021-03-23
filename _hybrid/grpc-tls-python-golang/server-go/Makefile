# Makefile for Golang Service - gRPC-TLS-Python-Golang example

PROTODIR = ../proto
PROTOS = ${PROTODIR}/service.proto
PBDIR = pb
PBFILES = ${PBDIR}/service.pb.go

PROTOC = protoc
PROTOC_FLAGS = -I${PROTODIR} --go_out=plugins=grpc:${PBDIR} --go_opt=paths=source_relative

.PHONEY: build
build: server

server: ${PBFILES} server.go
	go build -o server

.PHONEY: pbs
pbs: ${PBFILES}

${PBFILES}: ${PROTOS}
	${PROTOC} ${PROTOC_FLAGS} $<

.PHONEY: clean
clean:
	go clean ./...

.PHONEY: realclean
realclean: clean
	-rm -f pb/*.pb.go
