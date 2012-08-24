#include <errno.h>
#include <stdio.h>
#include <termios.h>

int main() {
    
    FILE* my_tty;
    char  ch;
    int   xx;
    
    my_tty = fopen("/dev/ttyqf", "r+");

    while ( (ch=fgetc(my_tty)) != EOF ) {
        putchar(ch);
    }
    putchar('\n');

/*     for( ch = 'A'; ch < ('Z'+1); ch++ ) { */
/*         fputc(ch, my_tty); */
/*     } */
/*     if ( 0 != (xx = tcsendbreak(fileno(my_tty), 1)) ) { */
/*         fprintf(stderr, "tcsendbreak() failed - returned %d\n", xx); */
/*         fprintf(stderr, "errno: %d\n", errno); */
/*     } */

    fclose(my_tty);
    return 0;
}
