/* gopt.c */

#include <unistd.h>
#include <stdio.h>

void main(int argc, char** argv) {
    int c;

    while ((c = getopt(argc, argv, "ab:c:")) != EOF) {
        printf("%c, %s\n", c, optarg);
    }
    printf("%c\n", 'a');
}
