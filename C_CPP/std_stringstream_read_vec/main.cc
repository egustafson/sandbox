#include <sstream>

using namespace std;


int main() {

    stringstream intvec;
    int ii;

    intvec.str(" 1 2 3 4 ");
    while ( intvec && !intvec.eof() ) {
        intvec>>ii>>ws;
        cout<<ii<<endl;
    }

    cout<<endl;

    intvec.clear();
    intvec.str(" 400 300 200 100");
    while ( intvec && !intvec.eof() ) {
        intvec>>ii>>ws;
        cout<<ii<<endl;
    }


    return 0;
}
