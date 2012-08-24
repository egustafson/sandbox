###################################
# HP-UX Config
# 
# CCC = aCC
# CCFLAGS = -DACE_LACKS_PRAGMA_ONCE -I$(ACE_ROOT) -I$(TAO_ROOT)
# LDFLAGS = -L/view/ericg-bidir/vobs/tao/pre_1_1a/lib-pa2.0n -lpthread -lACE -lTAO

###################################
# NetBSD Config
#
LD = g++
CPP = g++
CFLAGS = -w -I$(ACE_ROOT) -I$(TAO_ROOT)
LDFLAGS = -L$(ACE_ROOT)/ace -lACE -lTAO


OBJS = timeS.o timeC.o myclient.o myserver.o

all : myserver myclient

myserver : $(OBJS)
	$(LD) -o myserver $(LDFLAGS) timeS.o timeC.o myserver.o

myclient : $(OBJS)
	$(LD) -o myclient $(LDFLAGS) timeC.o myclient.o
