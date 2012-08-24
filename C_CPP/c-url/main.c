/* ----- from parseurl.c ----- */

typedef struct {
	char *type;
    char *user;
    char *pass;
	char *server;
	char *port;
	char *file;
} url_t;

void parseurl(url_t *url, char *in);

/* --------------------------- */

#include <stdio.h>

/* ----- */

int main() {

    char  url_string[] = "ftp://user:pa%5b%3Fss@host:123/path/elem";
    url_t url_struct;

    printf("----------------------------------------\n");
    printf("url:     %s\n", url_string);

    parseurl(&url_struct, url_string);

    printf("----------------------------------------\n");
    printf("type:    %s\n", url_struct.type);
    printf("user:    %s\n", url_struct.user);
    printf("pass:    %s\n", url_struct.pass);
    printf("server:  %s\n", url_struct.server);
    printf("port:    %s\n", url_struct.port);
    printf("path:    %s\n", url_struct.file);
    printf("----------------------------------------\n");

    return 0;
}

