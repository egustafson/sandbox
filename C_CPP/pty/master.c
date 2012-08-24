#include <fcntl.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <termios.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

int connect_socket();           /* return socket fd */
int open_pty();

int main() {

    int   pty_fd;
    int   sock_fd;
    char  buf[256];
    int   rd_size;

    sock_fd = connect_socket();
    pty_fd  = open_pty();

/*     strcpy(buf, "xyzzy\n\0"); */
/*     write(pty_fd, buf, 7); */
    while ( (rd_size=read(sock_fd, buf, sizeof(buf))) > 0 ) {
/*         fprintf(stderr, "Read %d bytes.\n", rd_size); */
        write(pty_fd, buf, rd_size);
/*         write(2, buf, rd_size); */
    }
    tcdrain(pty_fd);
    close(sock_fd);
    close(pty_fd);
    return 0;
}

int connect_socket() {
    int   sock_fd;
    int   proto;
    int   result;

    struct protoent* my_protoent;
    struct sockaddr_in servaddr;

    my_protoent = getprotobyname("tcp");

    sock_fd = socket(PF_INET, SOCK_STREAM, my_protoent->p_proto);
    if ( sock_fd < 0 ) {
        fprintf(stderr, "failed to alloc socket()\n");
        exit(1);
    }

    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family = AF_INET;
    /* servaddr.sin_port   = htons(7); /* echo server */
    servaddr.sin_port   = htons(13); /* daytime server */
    inet_pton(AF_INET, "127.0.0.1", &servaddr.sin_addr);

    result = connect(sock_fd, 
                     (struct sockaddr*) &servaddr, 
                     sizeof(servaddr) );
    if ( result < 0 ) {
        fprintf(stderr, "failed to connect()\n");
        exit(1);
    }

    return sock_fd;
}

int open_pty() {

    int pty_fd;

    pty_fd = open("/dev/ptyqf", O_RDWR, 0);

    if ( pty_fd < 0 ) {
        fprintf(stderr, "failed to open pty.\n");
        exit(1);
    }
    
    return pty_fd;
}
