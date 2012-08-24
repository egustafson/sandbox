// ////////////////////////////////////////////////////////////

#include <map>
#include <iostream>
#include <string>

using namespace std;

// ////////////////////////////////////////////////////////////

class Base {
public:
    virtual ~Base();
    virtual void doSomething() = 0;
protected:
    Base();
};


class WorkerA : Base {
public:
    static const std::string tag;
    virtual ~WorkerA();
    virtual void doSomething();
    static Base* factory();
protected:
    WorkerA();
private:
    static bool registered;
};


class WorkerB : Base {
public:
    static const std::string tag;
    virtual ~WorkerB();
    virtual void doSomething();
    static Base* factory();
protected:
    WorkerB();
private:
    static bool registered;
};


class Factory {
public:
    static Base* get(const std::string& name);
    static bool  reg(const std::string& name, Base* (*classFactory)() );
private:
    static std::map<std::string, Base* (*)()> factoryMap;
};


// ------------------------------------------------------------

std::map<std::string, Base* (*)() > Factory::factoryMap;

const std::string WorkerA::tag = "WorkerA";
const std::string WorkerB::tag = "WorkerB";

bool WorkerA::registered = Factory::reg( WorkerA::tag, WorkerA::factory );
bool WorkerB::registered = Factory::reg( WorkerB::tag, WorkerB::factory );

// ------------------------------------------------------------

bool Factory::reg(const std::string& name, Base* (*classFactory)() ) {

    factoryMap[name] = classFactory;
    cout<<"Registered '"<<name<<"' with id "<<classFactory<<"."<<endl;
    return true;
}

Base* Factory::get(const std::string& name) {
    Base* newObj = NULL;
    map<std::string, Base* (*)()>::iterator itor = factoryMap.find(name);
    if ( itor != factoryMap.end() ) {
        newObj = (*itor).second();
    }
    return newObj;
}

// ------------------------------------------------------------

Base::Base() { }
Base::~Base() { }

// ------------------------------------------------------------

WorkerA::WorkerA() { }
WorkerA::~WorkerA() { }

void WorkerA::doSomething() { cout<<"WorkerA working."<<endl; }

Base* WorkerA::factory() {
    return new WorkerA;
}

// ------------------------------------------------------------

WorkerB::WorkerB() { }
WorkerB::~WorkerB() { }

void WorkerB::doSomething() { cout<<"WorkerB working."<<endl; }

Base* WorkerB::factory() {
    return new WorkerB;
}

// ============================================================

int main() {

    Base* firstWorker  = Factory::get( WorkerA::tag );
    Base* secondWorker = Factory::get( WorkerB::tag );

    cout<<"First Worker (A):  ";
    firstWorker->doSomething();

    cout<<"Second Worker(B):  ";
    secondWorker->doSomething();

    return 0;
}
