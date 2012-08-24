CC        = gcc
CFLAGS    = -pedantic -g
CXXFLAGS  = $(CFLAGS)
LD        = $(CC)
LDFLAGS   = 

MODULES   = ClassDef

HEADERS   = $(MODULES:=.h)
OBJS     = $(MODULES:=.o) main.o

main : $(OBJS) 
	$(LD) -o $@ $(OBJS) $(LDFLAGS)

%.o : %.cpp $(HEADERS)
	$(CC) -c $(CXXFLAGS) -o $@ $<

