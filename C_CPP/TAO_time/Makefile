
IDL= tao_idl
IDLFLAGS= -hc "C.hh" -hs "S.hh" -hT "S_T.hh" -cs "C.cc" -ss "S.cc" -sT "S_T.cc"

CC= g++
CCFLAGS=-I$(ACE_ROOT) -I$(TAO_ROOT)

LDFLAGS=-L$(ACE_ROOT)/ace -lACE -lTAO -lTAO_PortableServer

## ------------------------------------------------------------

all : client servant

## -----

timeC.hh : time.idl
	$(IDL) $(IDLFLAGS) time.idl

timeC.o : timeC.cc timeC.hh
	$(CC) $(CCFLAGS) -c timeC.cc

timeS.o : timeS.cc timeS.hh timeC.hh
	$(CC) $(CCFLAGS) -c timeS.cc

## -----

client.o : client.cc timeC.hh
	$(CC) $(CCFLAGS) -c client.cc

client : client.o timeC.o
	$(CC) -o client client.o timeC.o $(LDFLAGS)

## -----

servant.o : servant.cc timeC.hh
	$(CC) $(CCFLAGS) -c servant.cc

servant : servant.o timeC.o timeS.o
	$(CC) -o servant servant.o timeS.o timeC.o $(LDFLAGS)
