#include <stdio.h>
#include <strings.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/uio.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define SA struct sockaddr

int main ( int argc, char **argv ) {

    int    sockfd, n;
    char   recvline[2048 + 1];
    struct sockaddr_in servaddr;

    if ( argc != 2 ) {
        fprintf( stderr, "usage: a.out <IPaddress>\n");
        exit(1);
    }

    if ( (sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0 ) {
        fprintf( stderr, "socket error\n");
        exit(1);
    }

    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(13); /* daytime server */
    if ( inet_pton(AF_INET, argv[1], &servaddr.sin_addr) <= 0 ) {
        fprintf( stderr, "inet_pton error for %s\n", argv[1]);
        exit(1);
    }

    if ( connect(sockfd, (SA *) &servaddr, sizeof(servaddr)) < 0 ) {
        fprintf( stderr, "connect error\n" );
        exit(1);
    }

    while ( ( n=read(sockfd, recvline, 2048) ) > 0) {
        recvline[n] = 0;
        if (fputs(recvline, stdout) == EOF) {
            fprintf(stderr, "fputs error\n");
            exit(1);
        }
    }
    if ( n < 0 ) {
        fprintf(stderr, "read error\n");
        exit(1);
    }
    exit(0);
}
