all : udpstorm

udpstorm : main.o
	gcc -o udpstorm main.o

main.o : main.c
	gcc -c -Wall -pedantic main.c

clean :
	rm -rf main.o udpstorm
