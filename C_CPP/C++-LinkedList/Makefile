all: bench

bench: bench.o intlist.o
	gcc -o bench bench.o intlist.o

bench.o: bench.cpp intlist.h
	gcc -c bench.cpp

intlist.o: intlist.cpp intlist.h
	gcc -c intlist.cpp
