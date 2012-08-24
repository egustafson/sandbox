/* client.c - client for gizmod (germit)
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


#define KERMIT_PATH "/usr/local/bin/kermit"

/* Allocate a pseudo-terminal pair -              */
/*  Function populates pty_master and pty_slave   */
/*  with the path of the master and slave devices */
/*  and returns the FD of the open master, or -1  */ 
/*  on failure. */
int alloc_pty_pair(char* pty_master, char* pty_slave);

/* Allocate a socket and connect it - */
/*  This function creates the socket  */
/*  connection between this client    */
/*  and the server running on the     */
/*  remote host.  Returns the socket  */
/*  FD or -1 on error.                */
int connect_sock(const char* serv_ip);

void run_kermit(char* slave_tty); /* exec the kermit program */

void relay(int master_fd, int sock_fd); /* relay data from master to socket */

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int main (int argc, char* argv[]) {

    int   pty_fd;               /* master FD */
    int   sock_fd;              /* remote socket */
    int   child_pid;            /* the background worker, the child */
    char  pty_master[25];       /* master's name (path) */
    char  pty_slave[25];        /* slave's name */

    sock_fd = connect_sock(argv[1]);                  /* function failure will exit process */
    pty_fd  = alloc_pty_pair(pty_master, pty_slave);  /* function failure will exit process */

    if ( (child_pid=fork()) < -1 ) {
        fprintf(stderr, "ERROR: failed to fork background worker.\n");
        exit(1);
    } else if ( 0 == child_pid ) { /* child here */
        relay(pty_fd, sock_fd);
    } else {                    /* parent here */
        run_kermit(pty_slave);
    }

}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int alloc_pty_pair(char* pty_master, char* pty_slave) {

    int master_fd;
    char *ptr1, *ptr2;

    strcpy(pty_master, "/dev/ptyXY");
    for ( ptr1 = "pq"; *ptr1 != 0; ptr1++ ) {
        pty_master[8] = *ptr1;
        for ( ptr2 = "0123456789abcdef"; *ptr2 != 0; ptr2++ ) {
            pty_master[9] = *ptr2;

            /* try to open the master */
            if ( (master_fd = open(pty_master, O_RDWR)) >= 0 ) {
                /* found one !! */
                strcpy(pty_slave, pty_master);
                pty_slave[5] = 't';
                return master_fd;
            }
        }
    }
    
    /* didn't find one */
    fprintf(stderr, "ERROR: couldn't open pseudo-terminal.\n");
    exit(1);
    return -1;
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

int connect_sock(const char* serv_ip) {
   
/*     #define REMOTE_SERVER "127.0.0.1" */
    #define REMOTE_SERVER "10.3.4.18"
    #define REMOTE_PORT   2222
                                /* daytime -> 13, echo -> 7 */
    int            sock_fd;
    int            proto;

    struct protoent* my_proto;
    struct sockaddr_in   remote_addr;

    my_proto = getprotobyname("tcp");
    if ( (sock_fd=socket(PF_INET, SOCK_STREAM, my_proto->p_proto )) < 0 ) {
        fprintf(stderr, "ERROR: couldn't allocate a socket()\n");
        exit(1);
    }

    bzero(&remote_addr, sizeof(remote_addr));
    remote_addr.sin_family = AF_INET;
    remote_addr.sin_port   = htons(REMOTE_PORT); /* TCP echo server */
    inet_pton(AF_INET, serv_ip, &remote_addr.sin_addr);

    if ( connect(sock_fd, (struct sockaddr*) &remote_addr, sizeof(remote_addr)) < 0 ) {
        fprintf(stderr, "ERROR: connect() with remote host failed.\n");
        exit(1);
    }
    
    return sock_fd;
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void run_kermit(char* slave_tty) {

    /* put myself in my own process group & make myself the forground process */
/*     setpgid(getpid(), getpid()); */
/*     tcsetpgrp(0, getpid()); */

    execl(KERMIT_PATH, "kermit", "-l", slave_tty, (char*)0);
    fprintf(stderr, "ERROR: exec failed, errno: %d\n", errno);
    exit(1);
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

void relay(int master_fd, int sock_fd) {

    fd_set rset, allset;
    int    maxfdp1;
    int    shutdown;

    char   buf[1024];
    int    rdsize;

    /* we're effectively a daemon process, shutdown our stdin and out */
    close(0);
    close(1);
    close(2);

    shutdown = 0;
    maxfdp1  = ((sock_fd>master_fd?sock_fd:master_fd)+1);
    FD_ZERO(&allset);

    FD_SET(master_fd, &allset);
    FD_SET(sock_fd,   &allset);

    while ( !shutdown ) {
        rset = allset;
        select(maxfdp1, &rset, NULL, NULL, NULL);
        
        if ( FD_ISSET(master_fd, &rset) ) {
            rdsize = read(master_fd, buf, sizeof(buf));
            if ( rdsize > 0 ) {
                /* should really recover from a partial write */
                write(sock_fd, buf, rdsize);
            } else {
                shutdown = 1;
            }
        }

        if ( FD_ISSET(sock_fd, &rset) ) {
            rdsize = read(sock_fd, buf, sizeof(buf));
            if ( rdsize > 0 ) {
                /* should really recover from a partial write */
                write(master_fd, buf, rdsize);
            } else {
                shutdown = 1;
            }
        }
    }

    close(master_fd);
    close(sock_fd);

    exit(0);
}
