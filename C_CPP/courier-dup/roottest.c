#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

void fatal(char* msg);

int main() {

    FILE*  root_data;
    int    pid;
    int    wait_stat;

    const char *args[2];
    char *env[2];

    fprintf(stderr, "roottest: UID(%d), EUID(%d)\n", getuid(), geteuid());

    args[0] = "./tester";
    args[1] = 0;
    env[0]  = "HOME=/home";
    env[1]  = 0;

    root_data = fopen("root.data", "r");
    if ( !root_data ) {
        fatal("Can't open 'root.data'\n");
    }
/*     root_data = stdin; */

    if ( (pid=fork()) < 0 ) {
        fatal("Can't fork()\n");
    }
    
    if ( pid ) { /* child */
        if ( fileno(root_data) != 0 ) {
            close(0);
            dup(fileno(root_data));
        }

        setuid(201);
        seteuid(201);
        
        execve("/home/egustafs/exper/courier-dup/tester", (char **)args, env);
        fatal("child failed to exec.\n");
    }

    /* Parent */
    while (wait(&wait_stat) != pid) /* do nothing */;
}

/* -------------------------------------------------- */
void fatal(char* msg) {
    fprintf(stderr, "FATAL:  %s", msg);
    exit(1);
}
