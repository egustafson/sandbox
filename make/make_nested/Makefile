CFLAGS = -c

# could also be
# CFLAGS ?= -c

DEBUG_CFLAGS = -g -D_DEBUG

# Conditional Assignment (comment out first line)
FOO += -x
FOO ?= -z

all:  build-all

.PHONY: debug
debug:
	make "CFLAGS=${CFLAGS} ${DEBUG_CFLAGS}"

build-all:
	@echo "CFLAGS = ${CFLAGS}"
	@echo "FOO = ${FOO}"
