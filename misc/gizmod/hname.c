/* hname.c - hostname resolution
 *
 */

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <netinet/in.h>
#include <arpa/inet.h>

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int main (int argc, char* argv[] ) {

    char   *ptr, **pptr;
    char   str[100];
    struct hostent *hptr;

    while (--argc > 0) {
        ptr = *++argv;
        if ( (hptr = gethostbyname(ptr)) == NULL ) {
            fprintf(stderr, "gethost error for: %s\n", ptr);
            continue;
        }
        printf("official hostname: %s\n", hptr->h_name);

        for ( pptr = hptr->h_aliases; *pptr != NULL; pptr++ ) {
            printf("\talias: %s\n", *pptr);
        }
        
        switch (hptr->h_addrtype) {
        case AF_INET:
            pptr = hptr->h_addr_list;
            for( ; *pptr != NULL; pptr++ ) {
                printf("\taddress: %s\n",
                       inet_ntop(hptr->h_addrtype, *pptr, str, sizeof(str)));
            }
            break;
        default:
            printf("\taddress: (UNKNOWN)\n");
            break;
        }
    }
    return 0;
}
