#include <string>
#include <iostream>
#include <list>

using namespace std;

int main() {

    static const std::string AlphaNums  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-.";
    static const std::string WhiteSpace = " \n\t";

    list<string> lineList;
    
    lineList.push_back("Var1 = Value1");
    lineList.push_back(" Var2 = Value 2 ");
    lineList.push_back("Var3=Value 3");
    lineList.push_back(" # pure comment ");
    lineList.push_back(" Var-5 = A longer value with # a comment ");

    list<string>::iterator itor;

    for ( itor = lineList.begin(); itor != lineList.end(); ++itor ) {

        std::string varName, value;

        // ----------

        std::string::size_type idx;

        std::string pLineIn( (*itor) );

        idx = pLineIn.find('#');
        if ( idx != std::string::npos ) {
            // strip the comment
            pLineIn.erase( idx );
        }

        idx = pLineIn.find_first_of( AlphaNums );
        if ( idx == std::string::npos ) {
            cout<<"Comment only/blank line"<<endl;
            continue;
            // it was a comment only line
        }
        pLineIn.erase(0, idx);

        idx = pLineIn.find_first_not_of( AlphaNums );
        varName = pLineIn.substr(0, idx);

        idx = pLineIn.find('=');
        pLineIn.erase(0, idx+1);

        idx = pLineIn.find_first_not_of( WhiteSpace );
        pLineIn.erase(0, idx);

        idx = pLineIn.find_last_not_of( WhiteSpace );
        pLineIn.erase( idx+1 );

        value = pLineIn;

        // ----------

        cout<<"\'"<<varName<<"\' --> \'"<<value<<"\'"<<endl;
    }





    return 0;
}
