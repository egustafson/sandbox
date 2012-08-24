#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>

#define LOG_PATH "./log-socket"
/* #define LOG_PATH "/dev/log-socket-egg"  */

int main( int argc, char** argv ) {

    int                 fd;
    struct sockaddr_un  sa;
    socklen_t           salen;
    char                log_path[255];

    printf("tsock - starting\n");

    if ( argc > 1 ) {
        strcpy( log_path, argv[1] );
    } else {
        strcpy( log_path, LOG_PATH );
    }

    printf("log_path = \"%s\"\n", log_path);

    unlink( log_path ); /* we don't care if this fails */

    fd = socket( PF_UNIX, SOCK_DGRAM, 0 );
    if ( fd < 0 ) {
        fprintf(stderr, "socket() call failed - exiting.\n");
        exit(1);
    }

    memset( &sa, 0, sizeof(sa) );
    sa.sun_family = PF_UNIX;
    strcpy( sa.sun_path, log_path );
    salen = sizeof( sa.sun_family )  + strlen( sa.sun_path );

    if ( bind( fd, (struct sockaddr *)&sa, salen ) < 0 ) {
        fprintf(stderr, "bind() call failed - exiting.\n");
        exit(1);
    }

    if ( chmod( log_path, 0666 ) < 0 ) {
        fprintf(stderr, "chmod() call failed - exiting.\n");
        exit(1);
    }

    printf("tsock - start-up successful, ...\n");
    sleep(10);
    unlink( log_path ); /* clean-up after ourselves */
    printf("exiting.\n");
    return 0;
}
