/* inet-server.c */

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <unistd.h>

#define NSTRS       3

char *strs[NSTRS] = {
    "This is the first string from the server.\n",
    "This is the second string from the server.\n",
    "This is the third string from the server.\n"
};

extern int errno;

main() {
    char c;
    FILE *fp;
    int  fromlen;
    char hostname[64];
    struct hostent *hp;
    int  i, s, ns;
    struct sockaddr_in sin, fsin;

    gethostname(hostname, sizeof(hostname));

    if ((hp = gethostbyname(hostname)) == NULL) {
        fprintf(stderr, "%s: host unknown.\n", hostname);
        exit(1);
    }

    if ((s = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("server: socket");
        exit(1);
    }

    sin.sin_family = AF_INET;
    sin.sin_port = htons(1234);
    bcopy(hp->h_addr, &sin.sin_addr, hp->h_length);

    if (bind(s, (struct sockaddr*)&sin, sizeof(sin)) < 0) {
        perror("server: bind");
        exit(1);
    }

    if (listen(s, 5) < 0) {
        perror("server: listen");
        exit(1);
    }

    if ((ns = accept(s, (struct sockaddr*)&fsin, &fromlen)) < 0 ) {
        perror("server: accept");
        exit(1);
    }

    fp = fdopen(ns, "r");
    
    for (i=0; i < NSTRS; i++) {
        send(ns, strs[i], strlen(strs[i]), 0);
    }

    sleep(20);

    for (i=0; i < NSTRS; i++) {
        while ((c=fgetc(fp)) != EOF) {
            putchar(c);

            if (c == '\n') {
                break;
            }
        }
    }

    shutdown(s, 2);

    exit(0);
}
