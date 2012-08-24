/* daytime-udp.c */

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>

#define BUFSZ    256
#define SERVICE  "daytime"

main(int argc, char** argv) {

    int  s, n, len;
    char buf[BUFSZ];
    struct hostent *hp;
    struct servent *sp;
    struct sockaddr_in sin;

    if ((s = socket(AF_INET, SOCK_DGRAM, 0)) < 0) {
        perror("socket");
        exit(1);
    }

    if ((sp = getservbyname(SERVICE, "udp")) == NULL) {
        fprintf(stderr, "%s/udp: unknown service.\n", SERVICE);
        exit(1);
    }

    while (--argc) {
        
        if ((hp = gethostbyname(*++argv)) == NULL) {
            fprintf(stderr, "%s: host unknown.\n", *argv);
            continue;
        }
        
        sin.sin_family = AF_INET;
        sin.sin_port = sp->s_port;
        bcopy(hp->h_addr, &sin.sin_addr, hp->h_length);

        printf("%s: ", *argv);
        fflush(stdout);

        if (sendto(s, buf, BUFSZ, 0, (struct sockaddr*)&sin, sizeof(sin)) < 0 ) {
            perror("sendto");
            continue;
        }

        len = sizeof(sin);
        n = recvfrom(s, buf, sizeof(buf), 0, (struct sockaddr*)&sin, &len);

        if (n < 0) {
            perror("recvfrom");
            continue;
        }

        buf[n] = '\0';
        printf("%s\n", buf);

    }
    close(s);
    exit(0);
}
