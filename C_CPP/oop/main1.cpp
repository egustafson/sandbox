/* main1.cpp          -*- C++ -*- */

// Experiment 1:  Can I do type checking through
// pointer assignments by overloading the assignment
// operator for the base class?

// Result:  Failed.  There appears to be no way to define an
//   assignment operator that works with pointers to classes.
//

#include <stdio.h>

class BC_C {
protected:
    int type;
public:
    BC_C();
    
    virtual int get_type() const;
//     friend void operator = (BC_C obj1, BC_C obj2);
};

class A_C : BC_C {
public:
    A_C();
};

class B_C : BC_C {
public:
    B_C();
};
  
////////// definition //////////

BC_C::BC_C() :
    type(0)
{
    printf("constructor BC_C()\n");
}

A_C::A_C() { 

    type = 1;
    printf("constructor A_C()\n");
}

B_C::B_C() { 

    type = 2;
    printf("constructor B_C()\n");
}

int BC_C::get_type() const {
    return type;
}

void operator = (BC_C& obj1, BC_C& obj2) {

    fprintf(stderr, "operator = ()\n");
//     if ( obj1 != obj2 ) {
//         delete obj1;
//     }
//     obj1 = obj2;
//     return obj2;
}


////////// main() //////////

int main() {

    A_C* objA1;
    A_C* objA2;

    objA1 = new A_C;

    objA2 = objA1;

    delete objA1;

    return 0;
}
