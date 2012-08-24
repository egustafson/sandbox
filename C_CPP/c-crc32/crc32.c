#include "crc32.h"

#define USE_STATIC_TABLE

#ifdef USE_STATIC_TABLE
# include "crc32_poly.h"
#else
# include <stdio.h>
  static unsigned long crc_table[256];
#endif

unsigned long crc32(uchar* data, int size) {

    int ii;
    unsigned long crc = 0xffffffff;

    for ( ii = 0; ii < size; ++ii ) {
        crc = ((crc>>8) &0x00ffffff) ^ crc_table[(crc^data[ii])&0xff];
    }

    return (crc ^ 0xffffffff);
}


// ////////// Unused Utilities //////////

#ifndef USE_STATIC_TABLE
// Create the polynomial table
static void crc32_poly() {

    unsigned long crc, poly;
    int ii, jj;

    poly = 0xedb88320L;

    for ( ii = 0; ii < 256; ++ii ) {

        crc = ii;
        for ( jj = 8; jj > 0; --jj ) {
            if ( crc & 1 ) {
                crc = (crc >> 1) ^ poly;
            } else {
                crc = crc >> 1;
            }
            crc_table[ii] = crc;
        }
    }

}

// Print the polynomial table in C syntax
void print_crc32_poly() {

    int ii;

    // 1. Initialize the values.
    crc32_poly();

    for ( ii = 0; ii < 256; ++ii ) {

        if ( 0 == (ii & 0x07) ) {
            printf("\n");
        }
        printf(" 0x%08x,", crc_table[ii]);
    }
    printf("\n");
}

#endif // ndef USE_STATIC_TABLE
