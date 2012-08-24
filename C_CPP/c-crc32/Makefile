all : main

main: main.o crc32.o
	$(CC) $(LDFLAGS) main.o crc32.o -o main

.c.o:
	$(CC) $(CFLAGS) -c $<

clean:
	rm -rf *.o main