#include <stdio.h>


static inline unsigned long rol( unsigned long val ) {

    int hbit = val & 0x80000000;
    val <<= 1;
    if ( hbit ) {
        val |= 1;
    }
    return val;
}


int main() {

    unsigned long val;
    int ii;

    val = 0x80000000;
    for ( ii = 0; ii < 64; ++ii ) {
        
        val = rol(val);
        printf("%3d : 0x%08x\n", ii, val);
    }
}
