/* Actions to take: */
/*  */
/* 1. get a FILE* to an open file. */
/* 2. fork() */
/* 3. child: close(0) */
/* 4. child: dup(fileno(FILE*)) */
/* 5. exec a program that prints STDIN */

#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

void fatal(char* msg);

int main() {

    FILE*   data;
    int     pid;
    int     wait_stat;

    const char *args[4];
    char *env[2];

    fprintf(stderr, "tester: UID(%d), EUID(%d)\n", getuid(), geteuid());

    args[0] = "/bin/sh";
    args[1] = "-c";
    args[2] = "./child";
    args[3] = 0;
    env[0]  = "HOME=/home";
    env[1]  = 0;

/*     data = fopen("data", "r"); */
/*     if ( !data ) { */
/*         fatal("Can't open 'data'\n"); */
/*     } */
    data = stdin;

/*     if (fseek(data, 0L, SEEK_SET) == -1) { */
/*         fatal("Can't fseek()\n"); */
/*     } */

    if ( (pid=fork()) < 0 ) {
        fatal("Can't fork()\n");
    }
    
    if ( pid ) { /* child */
        if ( fileno(data) != 0 ) {
            close(0);
            dup(fileno(data));
        } else {
            fprintf(stderr,"stdin -> stdin\n");
        }
        
        execve("/bin/sh", (char **)args, env);
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
