#include <stdio.h>

/* #define MEM_BASE 0xa000000 */
#define MEM_SIZE 1024*1024*64

inline unsigned long rol( unsigned long val ) {
    int hbit = val & 0x80000000;
    val <<= 1;
    if ( hbit ) {
        val |= 1;
    }
    return val;
}

void runpattern( unsigned long base, unsigned long size, unsigned long pattern ) {

    unsigned long  pp;
    unsigned long  ii;
    unsigned long  pat;
    unsigned long* mbase = (unsigned long*)base;

    unsigned long bound = size / sizeof(mbase);

    for ( pp = 0; pp < 32; ++pp ) {

        pat = pattern;
        for ( ii = 0; ii < pp; ++ii ) {
            pat = rol(pat);
        }
        printf("Pass %2d: 0x%08x\n", pp, pat);

        /* Set memory */
        for ( ii = 0; ii < bound; ++ii ) {
            mbase[ii] = pat;
            pat = rol(pat);
        }

        /* Test value stuck */
        pat = pattern;
        for ( ii = 0; ii < pp; ++ii ) {
            pat = rol(pat);
        }
        for ( ii = 0; ii < bound; ++ii ) {
            if ( mbase[ii] != pat ) {
                printf("Pattern mismatch at 0x%08x, expect 0x%08x, got 0x%08x\n", ii, pat, mbase[ii]);
            }
            pat = rol(pat);
        }
    }
}

int main() {

    unsigned long  mbase;
    unsigned long  msize;


    mbase = malloc( MEM_SIZE );
    msize = MEM_SIZE;

    if ( !mbase ) {
        fprintf(stderr, "Error: could not allocated %d bytes\n", MEM_SIZE);
        exit(1);
    }

/*     runpattern( mbase, msize, 0x00000001 ); */
    runpattern( mbase, msize, 0xfffffffe );
}
