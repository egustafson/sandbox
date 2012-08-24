#include <string>
#include <iostream>
#include <list>

using namespace std;

int main() {

    list<string> strList;
    
    strList.push_back("asdfg \\");
    strList.push_back("xyzzy \\");
    strList.push_back("abcde \\");
    strList.push_back("lmnop ");

    list<string>::iterator itor = strList.begin();
    bool finished = false;

    string line;

    while ( !finished && (itor != strList.end()) ) {

        line += (*itor);

        if ( line[line.size()-1] == '\\' ) {
            line.erase( line.size()-1 );
            // and append the next line;
        } else {
            finished = true;
        }
        ++itor;
    }

    cout<<"-->"<<line<<"<--"<<endl;

    
}
