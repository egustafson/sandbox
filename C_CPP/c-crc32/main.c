#include <stdio.h>
#include <stdlib.h>

#include "crc32.h"

#define DSIZE 256

int main() {

    int ii;
    uchar  data[DSIZE];
    unsigned long crc;
    
    for ( ii = 0; ii < DSIZE; ++ii ) {
        data[ii] = 0; //(uchar)(0xff & rand());
    }

    crc = crc32( data, sizeof(data) );

    printf("CRC-32:  0x%08x\n", crc);

    return 0;
}
