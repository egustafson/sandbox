
CFLAGS = -O2

all : main

## --------------------------------------------------

main : main.cc Timer.o
	g++ $(CFLAGS) -o main main.cc Timer.o

Timer.o : Timer.cc Timer.hh
	g++ $(CFLAGS) -c Timer.cc

clean:
	rm *.o main
