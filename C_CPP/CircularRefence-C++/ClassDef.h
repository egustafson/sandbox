/* ClassDef.h          -*- C++ -*- */


#ifndef CLASSDEF_H
#define CLASSDEF_H

class A;
class B;

class A {
    char*  value;
public:
    A();
    A(const A& val);
    A(const char* str);
    ~A();

    void  set(const B& data);
    char* get() const;
    void  print() const;
};

class B {
    char*  value;
public:
    B();
    B(const B& val);
    B(const char* str);
    ~B();

    void  set(const A& data);
    char* get() const;
    void  print() const;
};

#endif // CLASSDEF_H
