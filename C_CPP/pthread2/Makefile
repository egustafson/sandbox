all : main demo

main : main.c
	gcc -g -I/usr/pkg/include -o main main.c -L/usr/pkg/lib -lpthread

demo : demo.cc
	g++ -g -I/usr/pkg/include -o demo demo.cc -L/usr/pkg/lib -lpthread
