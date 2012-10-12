#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>

#define SVC_PORT   19283
#define MCAST_ADDR "224.0.2.1"
/* #define MCAST_ADDR "10.3.4.1" */

/*
#define SVC_PORT   25826
#define MCAST_ADDR "239.192.74.66"
*/

#define MAXDATA    2048

#define SA struct sockaddr



int main(int argc, char* argv[]) {

    char* send_addr = MCAST_ADDR;
    int send_port   = SVC_PORT;
    int sockfd;
    struct sockaddr_in sendaddr;
    char senddata[MAXDATA];
    
    bzero(&sendaddr, sizeof(sendaddr));
    sendaddr.sin_family = AF_INET;
    sendaddr.sin_port = htons(send_port);
    inet_pton(AF_INET, send_addr, &sendaddr.sin_addr);

    if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
        perror("socket failed to allocate");

    strcpy(senddata, "A short message.");

    sendto(sockfd, senddata, strlen(senddata), 0, (SA*)&sendaddr, sizeof(sendaddr));
}

