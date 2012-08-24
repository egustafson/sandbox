/* client.c             */

/*
 * Connects to the local host at port 1234
 */

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include "commands.h"

/*
 * Strings we send to the server.
 */

extern int errno;

int main() {

  char                hostname[64];
  int                 s;
  struct hostent*     hp;
  struct sockaddr_in  sin;
  char                buff[1000];
  message_s           msg;

  gethostname(hostname, sizeof(hostname));

  if ( (hp = gethostbyname(hostname)) == NULL ) {
    fprintf(stderr, "%s: unknown host.\n", hostname);
    exit(1);
  }
  
  if ( (s = socket(AF_INET, SOCK_STREAM, 0)) < 0 ) {
    perror("client:  socket");
    exit(1);
  }
  
  sin.sin_family = AF_INET;
  sin.sin_port = htons(1234);
  bcopy(hp->h_addr, &sin.sin_addr, hp->h_length);

  if ( connect(s, (struct sockaddr*)&sin, sizeof(sin)) < 0 ) {
    perror("client:  connect");
    exit(1);
  }

  msg.msgType = ECHO;
  strcpy( msg.buff, "The message to echo." );
  send( s, &msg, sizeof(msg), 0 );
  
  msg.msgType = CLOSE_SOCK;
  send( s, &msg, sizeof(msg), 0 );

  shutdown(s, 1);
  while ( recv(s, buff, sizeof(buff),0) > 0 ) {
      /* do nothing */
  }
  shutdown(s, 2);
  close(s);
  exit(0);
}


