#include <stdio.h>

int parse_octet(char* s, int* oct, char sep, int count, int min, int max, int base) {
// s   - pointer to the eth address to parse
// eth - int[6] to store the eth address in

    int ii;
    int valid;

    char* endptr;

    for ( ii = 0; ii < count; ++ii ) {
        oct[ii] = strtol(s, &endptr, base);

        valid  = ( ii >= count-1 || *endptr == sep );
        valid &= (oct[ii] >= min && oct[ii] <= max);
        if ( !valid ) {
            return 0;
        }
        s = endptr+1;
    }
    if ( endptr != NULL && *endptr != '\0' ) {
        s = endptr;
        while ( isspace(*s) ) ++s;
        if ( *s != '\0' ) {
            return 0;
        }
    }
    return 1;
}

// ////////////////////////////////////////////////////////////////////////////////

int checkip( char* s ) {
    
    int ip[4];

    if ( parse_octet(s, ip, '.', 4, 0, 255, 10) ) {
        printf("IP: %d.%d.%d.%d\n", ip[0], ip[1], ip[2], ip[3]);
    } else {
        printf("IP: '%s' - INVALID\n", s);
    }
}

int checketh( char* s ) {
    
    int eth[6];

    if ( parse_octet(s, eth, ':', 6, 0, 255, 16) ) {
        printf("eth: %02x:%02x:%02x:%02x:%02x:%02x\n", eth[0], eth[1], eth[2], eth[3], eth[4], eth[5]);
    } else {
        printf("eth: '%s' - INVALID\n", s);
    }
}

int main() {

    checkip("1.1.1.1");
    checkip("0.0.0.0");
    checkip("255.255.255.255");
    checkip("0.255.255.0");
    checkip("255.0.0.255");

    checkip("  1.1.1.1");
    checkip("  255.0.0.1");
    checkip("1.1.1.1  ");
    checkip("1.0.0.255  ");
    checkip(" 255.1.1.255  ");

    checkip("256.0.0.0");
    checkip("0.0.0.256");

    checkip("x.0.0.1");
    checkip("1.0.0.x");

    checkip("x 1.1.1.1");
    checkip("1.1.1.1 x");

    checkip("x1.1.1.1");
    checkip("1.1.1.1x");

    checketh("00:00:00:00:00:00");
    checketh("ff:ff:ff:ff:ff:ff");

    checketh("01:01:01:01:01:01");
    checketh("ff:00:00:00:00:ff");
    checketh("00:ff:ff:ff:ff:00");

    checketh("  01:01:01:01:01:01");
    checketh("01:01:01:01:01:01  ");
    checketh("  01:01:01:01:01:01  ");

    checketh("01:01:01:01:01:256");
    checketh("01:01:01:01:01:100");

    checketh("-1:01:01:01:01:01");
    checketh("100:01:01:01:01:01");
    checketh("256:01:01:01:01:01");

    checketh("x01:01:01:01:01:01");
    checketh("x 01:01:01:01:01:01");
    checketh("01:01:01:01:01:01x");
    checketh("01:01:01:01:01:01 x");
}
