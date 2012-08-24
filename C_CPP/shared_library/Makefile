all: main.shared main.noshare

main.o: main.cc Registry.hh
	g++ -c main.cc

Registry.o: Registry.cc Registry.hh
	g++ -c Registry.cc

B.o: B.cc Registry.hh
	g++ -c B.cc 

libB.so.0.0: B.o 
	-gcc -shared -o libB.so.0.0 B.o

main.shared: main.o Registry.o libB.so.0.0
	g++ -o main.shared main.o Registry.o -L. -lB

main.noshare: main.o Registry.o
	g++ -o main.noshare main.o Registry.o 

clean:
	rm -f *.o lib* main.noshare main.shared
