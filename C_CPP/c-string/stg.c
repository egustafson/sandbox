#include <stdio.h>

int main() {

    char root[4] = "root";

    int *str = (int*)root;

    printf("root -> %x\n", *str);
    return 0;
}
