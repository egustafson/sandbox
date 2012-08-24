#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>

#define SERV_PORT 11223
#define NUM_MSGS 4

#define SA struct sockaddr

void send_udp(int sockfd, const SA *pservaddr, socklen_t servlen, int numPkts);

int main(int argc, char* argv[]) 
{
/*     char targetaddr[] = "10.10.6.200"; */ 
    char* targetaddr; 
    int tgtPort = SERV_PORT;
    int pktCnt = NUM_MSGS;
    int sockfd;
    struct sockaddr_in servaddr;


    if ( argc < 2 ) {
        printf("usage: udbstorm <IP addr> [num pkts] [port]\n");
        exit(1);
    }

    targetaddr = argv[1];

    if ( argc > 2 ) {
        pktCnt = atoi(argv[2]);
    }

    if ( argc > 3 ) {
        tgtPort = atoi(argv[3]);
    }

    printf("Targeting %s:%d with %d packets.\n", targetaddr, tgtPort, pktCnt);

    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    servaddr.sin_port   = htons(tgtPort);
    inet_pton(AF_INET, targetaddr, &servaddr.sin_addr);

    sockfd = socket(AF_INET, SOCK_DGRAM, 0);

    send_udp(sockfd, (SA *)&servaddr, sizeof(servaddr), pktCnt);

    return 0;
}

/* ====================================================================== */

#define MAXDATA  4096

void send_udp(int sockfd, const SA *pservaddr, socklen_t servlen, int numPkts)
{
    int ii;
    char  senddata[MAXDATA];
    const int set_on = 1;           /* The 'on' position */

    if ( 0 > setsockopt(sockfd, SOL_SOCKET, SO_BROADCAST, &set_on, sizeof(set_on)) ) {
        perror("setsockopt(): ");
        exit(1);
    }
    strcpy(senddata, "A simple message to send in UDP.");

    for ( ii = 0; ii < numPkts; ++ii ) {
        
        sendto(sockfd, senddata, strlen(senddata), 0, pservaddr, servlen);
    }
}
