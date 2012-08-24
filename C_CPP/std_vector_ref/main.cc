#include <vector>
#include <iostream>

using namespace std;

class Container {
public:
    Container() : d(0) { cout<<"Null Constructor"<<endl; }
    Container( const Container& right ) : d(right.d) { cout<<"Self Constructor"<<endl; }
    ~Container() { cout<<"Destructor"<<endl; }
    
    Container& operator=(const Container& right) 
    { 
        cout<<"Assignment"<<endl;
        if (&right != this) { 
            d = right.d; 
        } 
        return *this; 
    }
    int& data() { return d; }
    int data() const { return d; } 
private:
    int d;
};



int main() {

    vector<Container> vec;
    Container c;
    c.data() = 1;
    vec.insert(0, c);

    cout<<"----------"<<endl;

    const Container& cr = vec[0];

    cout<<"cr.data() = "<<cr.data()<<endl;
    
    cout<<"----------"<<endl;

}
