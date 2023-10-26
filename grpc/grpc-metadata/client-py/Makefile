# Makefile for Python client - gRPC-TLS-Python-Golang example

PROTODIR = ../proto
PROTOS = ${PROTODIR}/service.proto
PBDIR = .
PBFILES = ${PBDIR}/service_pb2.py

PROTOC = python -m grpc_tools.protoc
PROTOC_FLAGS = -I${PROTODIR} --python_out=. --grpc_python_out=.


.PHONY: build
build: ${PBFILES}

${PBFILES}: ${PROTOS}
	${PROTOC} ${PROTOC_FLAGS} $?

.PHONY: clean
clean:
	@-rm -rf `find . -name __pycache__`

.PHONY: realclean
realclean: clean
	-rm -f *_pb2*.py

.PHONY: grpc-help
grpc-help:
	python -m grpc_tools.protoc -h
