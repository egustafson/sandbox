# Makefile for Python client - gRPC-TLS-Python-Golang example

PROTODIR = ../proto
PROTOS = ${PROTODIR}/service.proto
PBDIR = .
PBFILES = ${PBDIR}/service_pb2.py

PROTOC = python -m grpc_tools.protoc
PROTOC_FLAGS = -I${PROTODIR} --python_out=. --grpc_python_out=.


.PHONEY: build
build: ${PBFILES}

${PBFILES}: ${PROTOS}
	${PROTOC} ${PROTOC_FLAGS} $?

.PHONEY: clean
clean:
	@-rm -rf `find . -name __pycache__`

.PHONEY: realclean
realclean: clean
	-rm -f *_pb2*.py

.PHONEY: grpc-help
grpc-help:
	python -m grpc_tools.protoc -h
