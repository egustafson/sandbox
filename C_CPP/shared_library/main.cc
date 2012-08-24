#include <iostream.h>
#include <string>

#include "Registry.hh"

// //////////

class A {
public:
    A(const string& myname);
    ~A();
private:
    static A instance;
}

A::instance("A");

A::A(const string& myname) {
    Registry::get_instance()->log(myname);
}

A::~A() { }

// //////////

int main() {

    Registry::get_instance()->print_entries();

    cout<<"Done."<<endl;
    return 0;
}
