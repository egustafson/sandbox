#include <iostream.h>
#define _DEBUG
#include "header.hh"

int main() {


    cout<<"Hello, world."<<endl;
    cout.width(4);
    cout<<'('<<4<<')'<<endl;
    cout<<'('<<5<<')'<<endl;
    LOG_DEBUG<<"Testing."<<endl;
    LOG_INFO<<"Testing."<<endl;
    LOG_WARN<<"Testing."<<endl;

}
