# Makefile for Golang Service - gRPC-error-handling example

PROTODIR = ../proto
PROTOS = ${PROTODIR}/service.proto
PBDIR = pb
PBFILES = ${PBDIR}/service.pb.go

#PROTOCGENGO_OLDER = TRUE  ## uncomment for older versions of protoc-gen-go
PROTOC = protoc
ifdef PROTOCGENGO_OLDER
	PROTOC_FLAGS = -I${PROTODIR} --go_out=plugins=grpc:${PBDIR} --go_opt=paths=source_relative
else
	PROTOC_FLAGS = -I${PROTODIR} --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
endif

.PHONEY: build
build: server

server: ${PBFILES} server.go
	go build -o server

pb:
	mkdir pb

.PHONEY: pbs
pbs: ${PBFILES} pb

${PBFILES}: ${PROTOS} pb
	${PROTOC} ${PROTOC_FLAGS} $<

.PHONEY: clean
clean:
	go clean ./...

.PHONEY: realclean
realclean: clean
	-rm -f pb/*.pb.go
