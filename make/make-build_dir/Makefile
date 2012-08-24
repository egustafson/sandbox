SHELL=/bin/sh

include .kernel_dir

all: 
	@echo "Using '`cat .kernel_dir`' to build the kernel in."
	@(BLD_DIR=`cat .kernel_dir`; cd /tmp; echo compile $$BLD_DIR)

.kernel_dir:
	@verify_build_dir.sh ${KERNEL_DIR}

bail: foo
	@if [ ! -d bail ]; then \
		mkdir bail; \
		echo "Made bail dir"; \
	fi
	@echo "Ran bail target"

foo:
	touch foo

clean:
	-cd ${KERNEL_DIR}; make clean
	@echo "clean completed successfully."

realclean: distclean
distclean: clean
	-rm -r ${KERNEL_DIR}
	-rm -f .kernel_dir
	@echo "distclean completed successfully."

