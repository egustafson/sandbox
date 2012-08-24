/* redial.c */

#include <termcap.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define SLEEP_TIME 1
#define SH_STRING "exec /usr/sbin/pppd -detach /dev/cua0 connect \"/usr/sbin/chat -v -f /home/egustafs/.chatrc2\""
/* #define SH_STRING "exec /usr/bin/sleep 2" */

void main()
{
    unsigned int counter = 0;
    char*  clearstring;
    char* termtype = getenv("TERM");
    if ( tgetent( NULL, termtype ) ) {
        clearstring = tgetstr("cl", NULL);
        printf("%s", clearstring);
    }

    printf("ReDial.\n");
    for (;;) {                   /* for ever */
        counter++;
        printf("Dialing[%7d]", counter);
        fflush(stdout);
        system( SH_STRING );
        printf(" -- connection lost.\n");
        sleep( SLEEP_TIME );
    }
}
