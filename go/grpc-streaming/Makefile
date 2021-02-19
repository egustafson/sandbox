# build Golang gRPC example

PROTOS = pb/demo.proto
PBFILES = ${PROTOS:.proto=.pb.go}

PROTOC = protoc
PROTOC_FLAGS = --go_out=plugins=grpc:. --go_opt=paths=source_relative $*.proto
#
# protoc flags place the generated .pb.go in the same directory as the .proto


.PHONEY: build
build: ${PBFILES}
	go build -o client client.go
	go build -o server server.go

# (re)compile all .proto -> .pg.go
.PHONEY: pbs
pbs: ${PBFILES}


# compile gRPC ProtoBuf files (.pb.go) from .proto's
%.pb.go: %.proto
	${PROTOC} ${PROTOC_FLAGS} $*.proto


.PHONEY: clean
clean:
	go clean ./...
	-rm -f client
	-rm -f server

.PHONEY: realclean
realclean: clean
	@echo "--> cleaning checked in, generated code <--"
	-rm -f demo.pb.go
