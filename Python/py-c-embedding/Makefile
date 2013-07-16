
CFLAGS = $(shell python-config --cflags)
#CFLAGS = -I/usr/include/python2.7 -Wall

LDFLAGS = $(shell python-config --ldflags)
#LDFLAGS = -lpython2.7

main: main.o
	gcc -o main main.o ${LDFLAGS}

main.o:	main.c
	gcc -c main.c ${CFLAGS}

clean:
	rm -rf *.o main
