#include <pthread.h>
#include <stdio.h>

#include <iostream>

void* do_one_thing(void *);
void* do_another_thing(void *);
void do_wrap_up(int, int);

int r1 = 0;
int r2 = 0;

using namespace std;

int main() {
    pthread_t  thread1;
    pthread_t  thread2;

    pthread_create( &thread1, NULL, do_one_thing, (void*)&r1 );
    pthread_create( &thread2, NULL, do_another_thing, (void*)&r2 );

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    do_wrap_up(r1, r2);

    return 0;
}

void* do_one_thing(void *p) {

    int i, j, x;
    int* pnum_times = (int*)p;
    
    for ( i = 0; i < 4; ++i ) {
        cout<<"doing one thing"<<endl;
        for ( j = 0; j < 10000; ++j ) {
            x = x + i;
        }
        (*pnum_times)++;
    }
}


void* do_another_thing(void *p) {

    int i, j, x;
    int* pnum_times = (int*)p;
    
    for ( i = 0; i < 4; ++i ) {
        cout<<"doing another thing"<<endl;
        for ( j = 0; j < 10000; ++j ) {
            x = x + i;
        }
        (*pnum_times)++;
    }
}

void do_wrap_up(int one_times, int another_times) {

    int total;
    
    total = one_times + another_times;
    printf("wrap up: one thing %d, another %d, total %d\n", 
           one_times, another_times, total);
}
