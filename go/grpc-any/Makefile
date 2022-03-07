# build Golang gRPC example

PROTODIR = proto
PROTOS = ${PROTODIR}/demo.proto
PBDIR = pb
PBFILES = ${PBDIR}/demo.pb.go

PROTOC = protoc
PROTOC_FLAGS = -I${PROTODIR} --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
#
# protoc flags place the generated .pb.go in the same directory as the .proto

.PHONEY: build
build: grpc-any

grpc-any: ${PBFILES} main.go
	go build .

pb:
	mkdir pb

# (re)compile all .proto -> .pg.go
.PHONEY: pbs
pbs: ${PBFILES} pb

${PBFILES}: ${PROTOS} pb
	${PROTOC} ${PROTOC_FLAGS} $<

.PHONEY: clean
clean:
	go clean ./...
	-rm -f grpc-any

.PHONEY: realclean
realclean: clean
	@echo "--> cleaning checked in, generated code <--"
	-rm -f pb
