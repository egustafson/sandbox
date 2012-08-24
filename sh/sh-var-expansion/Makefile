ALL = 1 2 3 4 5

all:
	export ALL="${ALL}" && \
	./make-test.sh

alt:
	for i in ${ALL}; do \
		echo file$$i-* ;\
	done