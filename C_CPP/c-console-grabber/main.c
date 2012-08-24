#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>

#define CON_DEV "/dev/tty02"
#define LOG_DEV "/tmp/log_tty02.log"

char buf[2048];

int main() {

    int con_fd;
    int log_fd;
    int ii;

    con_fd = open(CON_DEV, O_RDONLY, 0);
    if ( con_fd < 0 ) {
        fprintf(stderr, "Failed to open %s, exiting.\n", CON_DEV);
        exit(1);
    }

    log_fd = open(LOG_DEV, (O_WRONLY|O_CREAT|O_TRUNC), 0644);
    if ( log_fd < 0 ) {
        fprintf(stderr, "Failed to open log file, %s, exiting.\n", LOG_DEV);
        exit(1);
    }

    while ( (ii = read(con_fd, buf, 2048)) > 0 ) {
        ii = write(log_fd, buf, ii);
    }
    
    return 0;
}
