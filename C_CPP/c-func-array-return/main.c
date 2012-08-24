#include <stdio.h>

struct ret_arr {
    char  arr[5];
};


struct ret_arr my_func() {

    struct ret_arr array;
    int ii;

    printf("f(): &array           = 0x%08x\n", &array);
    printf("f(): &(array.arr)     = 0x%08x\n", &(array.arr));
    printf("f(): &(array.arr[0])  = 0x%08x\n", &(array.arr[0]));
    printf("f(): &(array.arr[1])  = 0x%08x\n", &(array.arr[1]));    
    printf("f(): &(array.arr[2])  = 0x%08x\n", &(array.arr[2]));    
    printf("f(): &(array.arr[3])  = 0x%08x\n", &(array.arr[3]));    
    printf("f(): &(array.arr[4])  = 0x%08x\n", &(array.arr[4]));    

    for ( ii = 0; ii < 5; ++ii ) {
        array.arr[ii] = ii;
    }

    return array;
}


int main() {

    struct ret_arr array;
    int ii;

    for ( ii = 0; ii < 5; ii++ ) {
        array.arr[ii] = 0;
    }


    array = my_func();

    printf("&array           = 0x%08x\n", &array);
    printf("&(array.arr)     = 0x%08x\n", &(array.arr));
    printf("&(array.arr[0])  = 0x%08x\n", &(array.arr[0]));
    printf("&(array.arr[1])  = 0x%08x\n", &(array.arr[1]));    
    printf("&(array.arr[2])  = 0x%08x\n", &(array.arr[2]));    
    printf("&(array.arr[3])  = 0x%08x\n", &(array.arr[3]));    
    printf("&(array.arr[4])  = 0x%08x\n", &(array.arr[4]));    

    for ( ii = 0; ii < 5; ++ii ) {
        printf("arr[%d] = %d;  ", ii, array.arr[ii]);
    }
    printf("\n");

    return 0;
}
