#include <stdio.h>


int main (int argc, char* argv[]) {

    int nr;
    int ii;
    int subs[6];
    char buff[] = "010912282004.33";

    if ( argc != 2 ) {
        printf("Usage:  main MMDDhhmm[[CC]YY][.ss]\n\n");
        return 1;
    }

    subs[5] = 9999;

    nr = sscanf(argv[1], "%2d%2d%2d%2d%d.%2d",
                &(subs[0]),
                &(subs[1]),
                &(subs[2]),
                &(subs[3]),
                &(subs[4]),
                &(subs[5]) );

    printf("Matched %d substitutions.\n", nr);

    for ( ii = 0; ii < nr; ++ii ) {
        printf("%d\n", subs[ii]);
    }

    printf("subs[5] = %d\n", subs[5]);

    return 0;
}
