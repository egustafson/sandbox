all : gclient gizmod 

gclient : gclient.c
	gcc -o gclient gclient.c

gizmod : gizmod.c
	gcc -o gizmod gizmod.c

hname : hname.c
	gcc -o hname hname.c

clean:
	-rm -f *.o gizmod gclient
