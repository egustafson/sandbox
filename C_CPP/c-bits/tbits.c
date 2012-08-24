#include <stdio.h>

#define UTIL 100

static int boot_order[4][8] = 
    {
        { 0, 1, 2, 3, 4, 5, 6, 7 }, // U/0 First
        { 2, 3, 0, 1, 6, 7, 4, 5 }, // A/0 First
        { 4, 5, 6, 7, 0, 1, 2, 3 }, // U/1 First
        { 6, 7, 4, 5, 2, 3, 0, 1 }, // A/1 First
    };


void tbit( int ksel, int bsel ) {

    int seq = ((bsel & 0x01)<<1) | (ksel==UTIL?0:1);
    printf("k = %d, b = %d --> %d\n", ksel, bsel, seq);

    int *prtnSeq = boot_order[seq];

    int ii;
    for ( ii = 0; ii < 8; ++ii ) {
        printf(" %d", prtnSeq[ii]);
    }
    printf("\n");
}


int main() {

    int k, b;

    k = 1;   b = 0; tbit(k, b);
    k = 100; b = 0; tbit(k, b);
    k = 1;   b = 1; tbit(k, b);
    k = 100; b = 1; tbit(k, b);
}
