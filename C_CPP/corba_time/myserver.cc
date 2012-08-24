#include <time.h>
#include <iostream.h>
#include "timeS.h"

class Time_impl : public virtual POA_Time {
public:
    virtual TimeOfDay get_gmt() throw(CORBA::SystemException);
};

TimeOfDay Time_impl::get_gmt() throw(CORBA::SystemException) {
    time_t time_now = time(0);
    struct tm* time_p = gmtime(&time_now);
    
    TimeOfDay tod;
    tod.hour   = time_p->tm_hour;
    tod.minute = time_p->tm_min;
    tod.second = time_p->tm_sec;

    cerr<<"*"<<endl;

    return tod;
}

int main(int argc, char* argv[]) {
    try {
        CORBA::ORB_var orb = CORBA::ORB_init(argc, argv);

        CORBA::Object_var obj = orb->resolve_initial_references("RootPOA");
        PortableServer::POA_var poa = PortableServer::POA::_narrow(obj);

        PortableServer::POAManager_var mgr = poa->the_POAManager();
        mgr->activate();

        Time_impl time_servant;
        
        Time_var tm = time_servant._this();
        CORBA::String_var str = orb->object_to_string(tm);
        cout<<str<<endl;

        orb->run();
    } catch (const CORBA::Exception&) {
        cerr<<"Uncaught CORBA exception"<<endl;
        return 1;
    }
    return 0;
}
