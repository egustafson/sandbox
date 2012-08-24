/* gizmod.c - a gizmo tunnel server
 *
 */

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>
#include <syslog.h>
#include <termios.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/uio.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define MSG_GIZMO_BUSY "\n** The gizmo server for this host is in use. **\n\n"
#define MSG_GIZMO_OPEN_ERR "\n** Error: gizmo server can not open gizmo tty **\n\n"
#define MSG_GIZMO_FINISHED "\nSession complete.\n"

#define SA struct sockaddr

void report(char* logmsg, int errnum);
void fatal(char* logmsg, int errnum);
void exit_handler(void);
void daemon_init(void);

/* Report back on the connecting socket any errors or        */
/* abnormal conditions about the server (port busy, ...)     */
void gizmo_error(int sockfd, const char* message);

/* Allocate and setup a server socket, ready for accept call */
int listen_sock(void);

/* Open, and configure the gizmo's console tty.              */
/*  return the open fd, or -1 on error, log the error.       */
int open_gizmo(const char* gizmo_tty);

/* Relay data between the gizmo and the socket in full       */
/* duplex.  When the socket closes return.                   */
void serve_gizmo(int gizmofd, int sockfd);

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int main( int argc, char* argv[] ) {

    int servfd;                 /* listening socket fd */
    int connfd;                 /* connected socket fd */
    int gizmofd;                /* gizmo tty fd */

    struct sockaddr_in clientaddr;
    int caddrlen;

    char buf[] = "Message to gizmo.\r\n";
    int  readsz;

    daemon_init();

    openlog("gizmod", LOG_PID, LOG_DAEMON);
    atexit(&exit_handler);
    report("daemon started.", 0);

    servfd = listen_sock();

    report("waiting for a connection.", 0);
    for ( ; ; ) {               /* forever */
        connfd = accept(servfd, (SA*)&clientaddr, &caddrlen);
        report("connection.", 0);
        if( (gizmofd = open_gizmo(argv[1]))<0 ) {
            gizmo_error(connfd, MSG_GIZMO_OPEN_ERR);
            close(connfd);
            break;
        }

        serve_gizmo(gizmofd, connfd); /* Relay - full duplex loop */

        gizmo_error(connfd, MSG_GIZMO_FINISHED);
        close(gizmofd);
        close(connfd);
        report("connection closed.", 0);
    }

    close(servfd);
    report("exiting.", 0);
    exit(0);
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void daemon_init(void) {

    if ( fork() > 0 ) {
        /* parent - say good night gracie */
        exit(0);
    }

    setsid();                   /* become session leader */
    chdir("/");                 /* set CWD to root */
    umask(0);
    /* close stdin, out, err, we don't use them */
    close(0); 
    close(1); 
    close(2);
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void serve_gizmo(int gizmofd, int sockfd) {
    
    fd_set rset, allset;
    int    maxfdp1;
    int    shutdown;
    
    char   buf[128];
    int    rdsize;

    shutdown = 0;
    maxfdp1  = ((sockfd>gizmofd?sockfd:gizmofd)+1);
    FD_ZERO(&allset);

    FD_SET(gizmofd, &allset);
    FD_SET(sockfd,  &allset);

    while ( !shutdown ) {
        rset = allset;
        select(maxfdp1, &rset, NULL, NULL, NULL);

        if ( FD_ISSET(gizmofd, &rset) ) {
            rdsize = read(gizmofd, buf, sizeof(buf));
            if ( rdsize > 0 ) {
                /* should really recover from a partial write */
                write(sockfd, buf, rdsize);
            } else {
                /* this should never happen, but ... */
                shutdown = 1;
            }
        }
        
        if ( FD_ISSET(sockfd, &rset) ) {
            rdsize = read(sockfd, buf, sizeof(buf));
            if ( rdsize > 0 ) {
                /* should really recover from a partial write */
                write(gizmofd, buf, rdsize);
            } else {
                shutdown = 1;
            }
        }
    }
    
    /* caller will take care of FD close()'s */
    return;
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int open_gizmo(const char* gizmo_tty) {
    
    #define GIZMO_TTY "/dev/tty00"

    int gizmofd;                /* gizmo tty fd */
    struct termios term;

    if ( (gizmofd=open(gizmo_tty, (O_RDWR | O_NONBLOCK), 0))<0 ) {
        report("open(GIZMO_TTY): %s", errno);
        close(gizmofd);
        return -1;
    }
        
    if ( tcgetattr(gizmofd, &term) < 0 ) {
        report("tcgetattr(GIZMO_TTY): %s", errno);
        close(gizmofd);
        return -1;
    }
        
        
    /* set the correct terminal characteristics */
    cfsetispeed( &term, B9600 );
    cfsetospeed( &term, B9600 );
        
    term.c_lflag &= ~(ECHO | ICANON | IEXTEN | ISIG);
    term.c_iflag &= ~(BRKINT | ICRNL | INPCK | IXON );
    term.c_oflag &= ~(OPOST);
    term.c_cflag &= ~(CSIZE | PARENB);
    term.c_cflag |=  (CS8 | CLOCAL);
    term.c_cc[VMIN]  = 1;       /* 1 character at a time */
    term.c_cc[VTIME] = 0;       /* no timeout */
        
    if ( tcsetattr(gizmofd, TCSANOW, &term) < 0 ) {
        report("tcsetattr(GIZMO_TTY): %s", errno);
        close(gizmofd);
        return -1;
    }
        
    if ( tcflush(gizmofd, TCIOFLUSH) < 0 ) {
        report("tcflush(GIZMO_TTY): %s", errno);
        close(gizmofd);
        return -1;
    }

    if ( tcsendbreak(gizmofd, 0) < 0 ) {
        report("tcsendbreak(GIZMO_TTY): %s", errno);
        close(gizmofd);
        return -1;
    }

    return gizmofd;
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int listen_sock(void) {

    #define LISTEN_PORT 2222
    #define LISTEN_QUE  2

    int sock_fd;
    int proto;

    struct protoent* my_proto;
    struct sockaddr_in   servaddr;

    my_proto = getprotobyname("tcp");
    if ( (sock_fd=socket(PF_INET, SOCK_STREAM, my_proto->p_proto )) < 0 ) {
        fatal("socket(): %s", errno);
    }

    bzero(&servaddr, sizeof(servaddr));
    servaddr.sin_family      = AF_INET;
    servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
    servaddr.sin_port        = htons(LISTEN_PORT); 

    if ( bind(sock_fd, (SA*)&servaddr, sizeof(servaddr)) < 0 ) {
        fatal("bind(): %s", errno);
    }

    if ( listen(sock_fd, LISTEN_QUE) < 0 ) { 
        fatal("listen(): %s", errno);
    }
    
    return sock_fd;
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void gizmo_error(int sockfd, const char* message) {

    write(sockfd, message, strlen(message));
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void report(char* logmsg, int errnum) {

    char buf[127];
    int  debug;

    if ( errnum > 0 ) {
        sprintf(buf, logmsg, strerror(errnum));
    } else {
        strcpy(buf, logmsg);
    }
    
    debug = 0;
    if ( debug ) {
        fputs(buf, stderr);
        fputc('\n', stderr);
    } else {
        syslog(LOG_WARNING, buf);
    }
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void fatal(char* logmsg, int errnum) {

    int debug;

    debug = 1;
    if ( debug ) {
        fputs("FATAL - ", stderr);
    }
    report(logmsg, errnum);
    exit(1);
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void exit_handler(void) {

    report("daemon exiting.", 0);
    closelog();
}
