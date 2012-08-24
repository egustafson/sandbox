#include <stdio.h>

typedef unsigned long  datum;
typedef datum*         addr_ptr;

addr_ptr test_ram_block( volatile addr_ptr base_addr, unsigned long size );

int main() {

    const unsigned long mem_size = 16 * 1024 * 1024;
    char* mem_chunck = (char*)malloc( mem_size ); /* 1 Meg */

    test_ram_block( (unsigned long*)mem_chunck, mem_size );
    
    free(mem_chunck);
    return 0;
}


addr_ptr test_ram_block( volatile addr_ptr base_addr, unsigned long size ) {

    unsigned long offset;
    unsigned long n_words = size / sizeof(datum);

    datum pattern;
    datum antipattern;

    for ( pattern = 1, offset = 0; offset < n_words; pattern++, offset++ ) {
        base_addr[offset] = pattern;
    }

    for ( pattern = 1, offset = 0; offset < n_words; pattern++, offset++ ) {
        if ( base_addr[offset] != pattern) {
            return ((addr_ptr)&base_addr[offset]);
        }
        antipattern = ~pattern;
        base_addr[offset] = antipattern;
    }

    for ( pattern = 1, offset = 0; offset < n_words; pattern++, offset++ ) {
        antipattern = ~pattern;
        if ( base_addr[offset] != pattern) {
            return ((addr_ptr)&base_addr[offset]);
        }
        base_addr[offset] = 0;
    }
    return NULL;
}
