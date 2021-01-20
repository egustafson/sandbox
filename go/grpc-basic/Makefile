# build Golang gRPC example

PROTOS = api/chat.proto
PBFILES = ${PROTOS:.proto=.pb.go}

PROTOC = protoc
PROTOC_FLAGS = --go_out=plugins=grpc:. --go_opt=paths=source_relative $*.proto
#
# protoc flags place the generated .pb.go in the same directory as the .proto


.PHONEY: build
build: ${PBFILES}
	(cd cmd/server; go build)
	(cd cmd/client; go build)
	ln -sf cmd/server/server .
	ln -sf cmd/client/client .

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
	-rm -f api/chat.pb.go
