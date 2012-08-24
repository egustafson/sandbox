HDRS = commands.h
EXECS = server echoer pinger stopit spam

all:  $(EXECS)

% : %.c $(HDRS)
	$(CC) -Wall -pedantic -o $* $<

# client: client.c
# 	gcc -o client client.c
# 
# server: server.c
# 	gcc -o server server.c

