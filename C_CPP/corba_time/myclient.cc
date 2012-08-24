#include <iostream.h>
#include <iomanip.h>
#include "timeC.h"

int main(int argc, char* argv[]) {
    try {
        if ( argc != 2 ) {
            cerr<<"Usage: client IOR_string"<<endl;
            throw 0;
        }
        
        CORBA::ORB_var orb = CORBA::ORB_init(argc, argv);
        
        CORBA::Object_var obj = orb->string_to_object(argv[1]);
        if (CORBA::is_nil(obj)) {
            cerr<<"Nil Time reference"<<endl;
            throw 0;
        }
        
        Time_var tm = Time::_narrow(obj);
        if (CORBA::is_nil(tm)) {
            cerr<<"Argument is not a Time reference"<<endl;
            throw 0;
        }
        
        TimeOfDay tod = tm->get_gmt();
        cout<<"Time in Greenwich is "
            <<setw(2)<<setfill('0')<<tod.hour<<":"
            <<setw(2)<<setfill('0')<<tod.minute<<":"
            <<setw(2)<<setfill('0')<<tod.second<<endl;
    } catch (const CORBA::Exception &) {
        cerr<<"Uncaught CORBA exception"<<endl;
        return 1;
    } catch (...) {
        return 1;
    }
    return 0;
}
