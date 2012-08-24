/* main.cpp          -*- C++ -*- */

#include "ClassDef.h"
#include <stdio.h>

int main() {

    A a("a");
    B b("b");
    B b2("b2");

    printf("Testing.\n");

    a.print();
    b.print();

    printf("a -> b\n");
    b.set(a);

    a.print();
    b.print();

    printf("b2 -> a\n");
    a.set(b2);

    a.print();
    b.print();
    
    printf("Finished.\n");
    exit(0);
}
