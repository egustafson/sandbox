#include <iostream>
#include <valarray>

#include "Timer.hh"

// //////////////////////////////////////////////////////////////////////

using namespace std;

#define ARR_SIZE 5000000
#define WIN_SIZE 100

void printArr(const valarray<float>& va);

inline void   getSlice( const valarray<float>& va, int start, int size, valarray<float>& sl );
inline float* getSlice( const float* ar, int start, int size );

// ----------------------------------------------------------------------

int main() {
    cout<<"Start..."<<endl;

    float* flar = new float[ARR_SIZE];

    valarray<float> flva(ARR_SIZE);
    cout<<"flva.size() = "<<flva.size()<<endl;

    float seed = 3.14159;       // mini-PI

    int ii;
    for ( ii = 0; ii < ARR_SIZE; ++ii ) {
        float fii = (float)ii;
        float num = sqrt( seed * fii);
        flva[ii] = num;
        flar[ii] = num;
    }

    Timer sw1;
    valarray<float> sl;

    sw1.start();
    for ( ii = 0; ii < (ARR_SIZE - WIN_SIZE); ++ii ) {
        getSlice( flva, ii, WIN_SIZE, sl );
        seed = sl[50];
    }
    sw1.stop();

    cout<<"valarray slice performance:"<<endl;
    sw1.print(cout);

    // ----------



    Timer sw2;
    float* sl2;
    
    sw2.start();
    for ( ii = 0; ii < (ARR_SIZE - WIN_SIZE); ++ii ) {
        sl2 = getSlice( flar, ii, WIN_SIZE );
        seed = sl2[50];
    }
    sw2.stop();

    cout<<"C array (ptr) slice performance:"<<endl;
    sw2.print(cout);

    cout<<"Finished."<<endl;
    return 0;
}

// ----------------------------------------------------------------------

void getSlice( const valarray<float>& va, int start, int size, valarray<float>& sl ) {

    sl = va[slice(start, size, 1)];
}

// ----------------------------------------------------------------------

inline float* getSlice( const float* ar, int start, int size ) {

    return const_cast<float*>( &(ar[start]) );
}

// ----------------------------------------------------------------------

void printArr(const valarray<float>& va) {

    int ii;
    for ( ii = 0; ii < va.size(); ++ii ) {
        cout<<"va["<<ii<<"] = "<<va[ii]<<endl;
    }
}



