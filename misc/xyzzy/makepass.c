/* makepass.c */

#include <pwd.h>
#include <stdio.h>
#include <termios.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

char *get_passwd();
char *encrypt_passwd(char *);

main()
{
  char *pw_crypt;

  pw_crypt = get_passwd();
  printf("%s\n", pw_crypt);
  exit(0);
}

/* ======== get_passwd ======== */

char *get_passwd()
{
  char buff[20];
  char pass[20];
  char *encrypted;
  struct termios attr;
  static char passwd[20];

  printf("Password:  ");
  fflush(stdout);

  tcgetattr(STDIN_FILENO, &attr);
  attr.c_lflag &= ~ECHO;
  tcsetattr(STDIN_FILENO, TCSAFLUSH, &attr);

  fgets(buff, sizeof(buff), stdin);
  sscanf(buff, "%s", pass);

  attr.c_lflag |= ECHO;
  tcsetattr(STDIN_FILENO, TCSAFLUSH, &attr);
  printf("\n");

  encrypted = encrypt_passwd(pass);
  encrypted[14] = '\0';	          /* make sure string is null terminated */
  strcpy(passwd, encrypted);

  return(passwd);
}


/* ======== encrypt_passwd ======== */

char *encrypt_passwd(pass)
char *pass;
{
  char *encrypted;
  char key[2];

  srand((unsigned int)time(NULL));

  key[0] = (char)(97 + (rand() % 26));
  key[1] = (char)(97 + (rand() % 26));

  encrypted = crypt(pass, key);

  return(encrypted);
}
