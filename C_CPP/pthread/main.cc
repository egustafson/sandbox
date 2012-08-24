#include <iostream>
#include <pthread.h>

using namespace std;

void *runner(void * param);
int sum;

pthread_mutex_t mut;

main() {
    
    pthread_t tid;
    pthread_attr_t attr;

    int ii;

    pthread_attr_init(&attr);
    pthread_mutex_init(&mut, NULL);

    sum = 0; 
    cout<<"In main before creating the thread, sum="<<sum<<endl;

    pthread_create( &tid, &attr, runner, NULL );

    for ( ii = 0; ii < 10000; ii++ ) {
        sum++;
    }

    pthread_join( tid, NULL );

    cout<<"Main finished, sum="<<sum<<endl;
}


void* runner(void* param) {

    int ii;
    cout<<"In the thread, hello world!!"<<endl;

    for ( ii = 0; ii < 10000; ii++ ) {
        sum++;
    }

    pthread_exit(0);
}

