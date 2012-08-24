/* ClassDef.cpp          -*- C++ -*- */

#include "ClassDef.h"
#include <stdio.h>

A::A() 
    : value(NULL)
{ }

A::A(const A& val) 
    : value(NULL)
{
    value = new char[strlen(val.value)+3];
    strcpy(value, val.value);
}

A::A(const char* str)
    : value(NULL)
{
    value = new char[strlen(str)+3];
    strcpy(value, str);
}

A::~A() {

    if ( NULL != value ) {
        delete[] value;
    }
}

void A::set(const B& data) {

    if ( NULL != value ) {
        delete[] value;
    }
    value = new char[strlen(data.get())+3];
    strcpy(value, data.get());
}

char* A::get() const {

    return value;
}

void A::print() const {

    printf("type:  A;  value:  \"%s\"\n", value);
}


//////////

B::B() 
    : value(NULL)
{ }

B::B(const B& val) 
    : value(NULL)
{
    value = new char[strlen(val.value)+3];
    strcpy(value, val.value);
}

B::B(const char* str)
    : value(NULL)
{
    value = new char[strlen(str)+3];
    strcpy(value, str);
}

B::~B() {

    if ( NULL != value ) {
        delete[] value;
    }
}

void B::set(const A& data) {

    if ( NULL != value ) {
        delete[] value;
    }
    value = new char[strlen(data.get())+3];
    strcpy(value, data.get());
}

char* B::get() const {

    return value;
}

void B::print() const {

    printf("type:  B;  value:  \"%s\"\n", value);
}
