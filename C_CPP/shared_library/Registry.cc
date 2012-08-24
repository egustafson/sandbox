#include "Registry.hh"

// ////////////////////////////////////////////////////////

Registry* Registry::instance(NULL);

Registry::Registry() { }
Registry::~Registry() { }

Registry* Registry::get_instance() {

    if ( NULL == instance ) {
        instance = new Registry;
    }
    return instance;
}

void Registry::log(const string& name) {
    myEntries.push_back(name);
}

void Registry::print_entries() const {
    list<string>::const_iterator it;
    for ( it = myEntries.begin(); it != myEntries.end(); it++ ) {
        cout<<"Registered:  "<<(*it)<<endl;
    }
}

