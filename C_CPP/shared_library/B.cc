#include <string>

#include "Registry.hh"

// //////////

class B {
public:
    B(const string& myname);
    ~B();
private:
    static B instance;
}

B::instance("B");

B::B(const string& myname) {
    Registry::get_instance()->log(myname);
}

B::~B() { }

