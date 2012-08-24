all: domParser domBuilder domWriter

CCC = g++
CCFLAGS = -I/usr/pkg/include -L/usr/pkg/lib -lxerces-c

domParser: domParser.cc
	${CCC} ${CCFLAGS} -o domParser domParser.cc

domBuilder: domBuilder.cc
	${CCC} ${CCFLAGS} -o domBuilder domBuilder.cc

domWriter: domWriter.cc
	${CCC} ${CCFLAGS} -o domWriter domWriter.cc

clean:
	rm -f domParser domBuilder domWriter main *.o
