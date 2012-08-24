#include <stdio.h>

int main() {

    char teststr[] = "<30>Log message.";

    int num;
    char buff[1024];

    int result = sscanf(teststr, "<%u>%[^\0]", &num, buff );

    printf("Result = %d\n", result);

    printf("Num    = %d\n", num);
    printf("Buffer = %s\n", buff);

    return 0;
}
