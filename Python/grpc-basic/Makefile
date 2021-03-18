# build Python gRPC example

PROTOS = pb/demo/demo.proto
PBFILES = ${PROTOS:.proto=_pb2.py}



PROTOC = python -m grpc_tools.protoc
PROTOC_FLAGS = -Ipb --python_out=. --grpc_python_out=.



.PHONEY: build
build: ${PBFILES}

.PHONEY: pbs
pbs: ${PBFILES}


# compile gRPC ProtoBuf files (.proto) -> (_pb2.py)   (_pb2_grpc.py is a side effect)
%_pb2.py: %.proto
	${PROTOC} ${PROTOC_FLAGS} $?


pb/demo/demo.proto: pb/demo.proto
	mkdir -p pb/demo
	(cd pb/demo; ln -s ../demo.proto .)

.PHONEY: clean
clean:
	-rm -rf `find . -name __pycache__`

.PHONEY: realclean
realclean: clean
	@echo "--> cleaning checked in, generated code <--"
	-rm -f demo/*_pb2*.py
	-rm -rf pb/demo

.PHONEY: grpc-help
grpc-help:
	python -m grpc_tools.protoc -h
