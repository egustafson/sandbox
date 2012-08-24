#include <string.h>


static char *skip_char(char *str, char *del) {
	int i;
	for (i=0; str[i] != '\0'; i++) {
		int j;
		int f=1;
		for (j=0; del[j] != '\0'; j++) if (str[i] == del[j]) f=0;
		if (f) return str+i;
	}
	return str+i;
}

static char *find_char(char *str, char *del) {
	int i;
	int j;
	for (i=0; str[i] != '\0'; i++)
		for (j=0; del[j] != '\0'; j++)
			if (str[i] == del[j]) return str+i;
	return str+i;
}

static void kill_char(char *str, char *del) {
	for (str=str; *str != '\0'; str++) {
		int i;
		for (i=0; del[i] != '\0'; i++) {
			if (del[i] == *str) {
				*str = '\0';
				return;
			}
		}
	}
}

static void split_line(char *s, char *d, int *n, char **v)
{
	int m = 0;
	char *t;
	while (1) {
		s = skip_char(s, d);
		if (*s != '\0') {
			v[m++] = s;
			t = find_char(s, d);
			if (*t == '\0') goto end;
			else {
				s = t + 1;
				*t = '\0';
			}
		} else goto end;
	}
	end:
	*n = m;
}

static char *fix_not_null(char *str) {
	if (str && *str != '\0') {
		*str = '\0';
		return str+1;
	}
	return str;
}

static void decode_url(char *str) {

    int ii;
    char dec_char;
    char tmp[1024];


    strncpy(tmp, str, 1024);
    
    ii = 0;
    while ( tmp[ii] != '\0' ) {
        if ( '%' == tmp[ii] ) {
            if ( tmp[ii+1] < 'A' ) {
                dec_char = tmp[ii+1] - '0';
            } else if ( tmp[ii+1] < 'a' ) {
                dec_char = tmp[ii+1] - 'A' + 10;
            } else {
                dec_char = tmp[ii+1] - 'a' + 10;
            }
            dec_char = dec_char << 4;

            if ( tmp[ii+2] < 'A' ) {
                dec_char += tmp[ii+2] - '0';
            } else if ( tmp[ii+2] < 'a' ) {
                dec_char += tmp[ii+2] - 'A' + 10;
            } else {
                dec_char += tmp[ii+2] - 'a' + 10;
            }

            ii  = ii + 2;
            *str = dec_char;
        } else {
            *str = tmp[ii];
        }
        str++;
        ii++;
    } 
    *str = '\0';
}

/*
  http://www.apple.com:80/dir/index.html
  http://www.apple.com:80/index.html
  ^      ^            ^   ^
  |      |            |   |
  |      |            |   |
  |      |            |   +--- file
  |      |            +--- port
  |      +--- server
  +--- type

  file:/file
  file:/dir/file

NOTE:

  should also support:
  http://username:password@ftp.apple.com:80/dir/index.htm
*/

static char common_type(char *s) {
	if (!strncmp(s, "http", 4)) return 1;
	if (!strncmp(s, "ftp", 3)) return 1;
	if (!strncmp(s, "file", 4)) return 1;
	return 0;
}

typedef struct {
	char *type;
    char *user;
    char *pass;
	char *server;
	char *port;
	char *file;
} url_t;

void parseurl(url_t *url, char *in) {
	// 0) establish starting values
	url->type = "";
    url->user = "";
    url->pass = "";
	url->server = "";
	url->port = "";
	url->file = "";

	// 1) skip any leading whitespace

	url->type = skip_char(in, " \t");

	if (!common_type(url->type)) {
		url->server = url->type;
		url->type = "file";
	} else {
		// 2) find ':' that separates type and server
		url->server = find_char(url->type, ":");

		// 3) replace ':' with '\0' and advance server
		url->server = fix_not_null(url->server);

		// 4) skip leading '/'
		url->server = skip_char(url->server, "/");
	}

    if (strcmp(url->type, "ftp") == 0) {
        // - Check for username/password
        if ( find_char(url->server, "@") < find_char(url->server, "/") ) {
            // - there's at least a username in there.
            url->user = url->server;
            url->server = find_char(url->user, "@");
            url->server = fix_not_null(url->server);
            
            // - Extract the possible password field
            url->pass = find_char(url->user, ":");
            url->pass = fix_not_null(url->pass);

            // - decode any encoded characers
            decode_url(url->user);
            decode_url(url->pass);
        }
    }

	if (strcmp(url->type, "file") == 0) {
		url->file = url->server;
		url->server = "";
	} else {
		// 5) find '/' that separates server and file
		url->file = find_char(url->server, "/");

		// 6) replace '/' with '\0' and advance file
		url->file = fix_not_null(url->file);

		// 7) kill whitespace to cleanup the end of the filename
		kill_char(url->file, " \t\r\n");

		// 8) go back to find port by looking for ':' in server
		url->port = find_char(url->server, ":");

		// 9) replace ':' with '\0' and advance port
		url->port = fix_not_null(url->port);

#if 0
		// 10) go back to find dir
		for (i=strlen(url->file)-1; i>=0; i--) {
			if (url->file[i] == '/') {
				url->file[i] = '\0';
				url->dir = url->file;
				url->file = &url->dir[i+1];
				break;
			}
		}
#endif
	}
}
