/* server.c          */

/*
 * Connects to port 1234 on the local host.
 */

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include "commands.h"

extern int errno;
int process( int socketID, int sequence_number );

/* ********** main() ********** */

int main() {

  int                 fromlen;
  char                hostname[64];
  struct hostent*     hp;
  int                 cmd, s, ns, seq_num;
  struct sockaddr_in  sin, fsin;
  char                buff[1000];

  gethostname(hostname, sizeof(hostname));
  
  if ( (hp = gethostbyname(hostname)) == NULL ) {
    fprintf(stderr, "%s: host unknown.\n.", hostname);
    exit(1);
  }

  if ( (s = socket(AF_INET, SOCK_STREAM, 0)) < 0 ) {
    perror("server: socket");
    exit(1);
  }

  sin.sin_family = AF_INET;
  sin.sin_port   = htons(1234);
  bcopy(hp->h_addr, &sin.sin_addr, hp->h_length);

  if ( bind(s, (struct sockaddr*)&sin, sizeof(sin)) < 0 ) {
    perror("server: bind");
    exit(1);
  }

  if ( listen(s, 5) < 0 ) {
    perror("server: listen");
    exit(1);
  }

  printf("\nServer waiting...\n");
  seq_num = 1;
  cmd = 0;
  while ( cmd != SHUTDOWN ) {

    if ( (ns = accept(s, (struct sockaddr*)&fsin, &fromlen)) < 0 ) {
      perror("server: accept");
      exit(1);
    }
    
    cmd = 0;
    while ( cmd != CLOSE_SOCK ) {
        cmd = process( ns, seq_num++ );
        if ( SHUTDOWN == cmd ) {
            break;
        }
    }

    shutdown(ns, 1);
    while ( recv(ns, buff, sizeof(buff), 0) > 0 ) {
        /* do nothing */
    }
    shutdown(ns, 2);
    close(ns);
  }

  close(s);
  exit(0);
}

/* ************************************************** */

int process( int socketID, int sequence_number ) {

    message_s msg;
    int       recv_size;

    recv_size = recv( socketID, &msg, sizeof(msg), 0 );

    if ( recv_size < 0 ) {
        printf("Error, receive failed, shutting down.\n");
        return SHUTDOWN;
    } else {
        printf("<%3d>[%4d bytes] : ", sequence_number, recv_size);
    }

    switch ( msg.msgType ) {
    case PING_REQUEST:
        printf("PING_REQUEST\n");
        msg.msgType = PING_RESPONSE;
        send( socketID, &msg, sizeof(msg) , 0 );
        break;
    case ECHO:
        printf("ECHO  \"%s\"\n", msg.buff);
        break;
    case CLOSE_SOCK:
        printf("CLOSE_SOCK\n");
        break;
    case SHUTDOWN:
        printf("SHUTDOWN\n");
        break;
    default:
        printf("UNKNOWN COMMAND - ( %d )\n", msg.msgType);
        break;
    }
    return msg.msgType;
}
