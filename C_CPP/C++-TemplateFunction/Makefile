#

CCC = g++

HDRS = FunctionRegistry.hh TemplateFunction.hh
SRCS = FunctionRegistry.cc demoTemplFunc.cc

OBJS = $(SRCS:%.cc=%.o)

all : demoTemplFunc.o FunctionRegistry.o
	$(CCC) -g ${OBJS} -o demoTemplFunc

.cc.o : ${SRCS} ${HDRS}
	$(CCC) -g -c $<