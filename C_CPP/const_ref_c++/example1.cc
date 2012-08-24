// Question:
//
// Does the copy constructor (or assignment operator) get called if a
// stack allocated (local var) is returned for a function with return
// type of const& (const reference)?
//

#include <stream.h>

class Foo {
public:
    Foo();
    Foo( const Foo& in_foo );
    Foo( int in_value );
    ~Foo();
    
    const Foo& operator=( const Foo& in_foo );

    int get_value() const { return foo_value; }

private:
    int foo_value;
};


Foo::Foo() 
    : foo_value(-1)
{
    cout<<"Foo::Foo() - called, this = "<<this<<endl;
}

Foo::Foo( const Foo& in_foo )
    : foo_value( in_foo.foo_value )
{
    cout<<"Foo( const Foo& in_foo ) - called, this = "<<this<<endl;
}

Foo::Foo( int in_value ) 
    : foo_value( in_value )
{
    cout<<"Foo::Foo( int in_value ) - called, this = "<<this<<endl;
}

Foo::~Foo() {
    cout<<"Foo::~Foo() - called, this = "<<this<<endl;
}


const Foo& Foo::operator=( const Foo& in_foo ) {
    if ( &in_foo != this ) {
        foo_value = in_foo.foo_value;
    }
    cout<<"Foo::operator= - called, this = "<<this<<endl;
    return *this;
}

// ======================================================================


const Foo& some_func() {

    static Foo returnFoo(3);

    cout<<"some_func() created a Foo of value:  "<<returnFoo.get_value()<<endl;
    return returnFoo;
}

// ==========

int main() {

    Foo myFoo;

    cout<<"main() - myFoo:  "<<myFoo.get_value()<<endl;
    cout<<"main() - myFoo is at:  "<<&myFoo<<endl;

    myFoo = some_func();

    cout<<"main() - myFoo:  "<<myFoo.get_value()<<endl;
    cout<<"main() - myFoo is at:  "<<&myFoo<<endl;

    return 0;
}
