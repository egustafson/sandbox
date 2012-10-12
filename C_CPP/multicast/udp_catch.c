#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>
#include <net/if.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>

#define SVC_PORT   19283
#define MCAST_ADDR "224.0.2.1"

/*
#define SVC_PORT   25826
#define MCAST_ADDR "239.192.74.66"
*/

#define MAXDATA    2048

#define SA struct sockaddr



int main(int argc, char* argv[]) {

    char* listen_addr = MCAST_ADDR;
    int listen_port   = SVC_PORT;
    int sockfd;
    struct sockaddr_in recvaddr;
    struct sockaddr_in srcaddr;
    socklen_t srcaddrlen = sizeof(srcaddr);
    char recvdata[MAXDATA];
    
    struct ip_mreq mreq;

    bzero(&recvaddr, sizeof(recvaddr));
    recvaddr.sin_family = AF_INET;
    /* recvaddr.sin_addr.s_addr = htonl(INADDR_ANY); */
    recvaddr.sin_port = htons(listen_port);
    inet_pton(AF_INET, listen_addr, &recvaddr.sin_addr);

    if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
        perror("socket failed to allocate");

    inet_pton(AF_INET, listen_addr, &mreq.imr_multiaddr.s_addr);
    mreq.imr_interface.s_addr = htonl(INADDR_ANY);
    setsockopt(sockfd, IPPROTO_IP, IP_ADD_MEMBERSHIP, &mreq, sizeof(mreq));

    if ( bind(sockfd, (SA*)&recvaddr, sizeof(recvaddr)) < 0 )
        perror("bind failed");

    recvfrom(sockfd, recvdata, MAXDATA, 0, (SA*)&srcaddr, &srcaddrlen);
}

