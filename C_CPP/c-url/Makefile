all: main

main: main.o parseurl.o
	gcc -g -o main main.o parseurl.o

.c.o :
	gcc -g -c $<

