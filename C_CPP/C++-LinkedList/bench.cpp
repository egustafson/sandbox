/* bench.cpp          -*- C++ -*- */

#include "intlist.h"

#include <stdio.h>


int main() {

    int     ii;
    IntList List;

    for ( ii = 0; ii < 10000; ii++ ) {
        List.insert(ii);
    }

    // List.print();

    printf("Benchmark done.\n");
    exit(0);
}
