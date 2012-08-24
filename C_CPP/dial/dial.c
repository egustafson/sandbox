/* dial.c */

#include <errno.h>
#include <fcntl.h>
#include <setjmp.h>
#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <termios.h>
#include <unistd.h>

/* #define MODEM "/dev/cua0" */
#define MODEM "/dev/tty0"
#define MODEM_SPEED B38400
#define MODEM_WAIT_PERIOD 5
#define MODEM_RESET_STRING "atz"
#define MODEM_DIAL_STRING "atds"

#define DEBUG

/* -- Prototypes -- */
void main();
int setup_modem();
void modem_send_command( int, char *);
void modem_receive_response( int );
void timeout(int);
/* -- end prototypes -- */

/* -- Global Var's -- */
jmp_buf env;
char message[100];

void main()
{
  int fildes;

  fildes = setup_modem();
  modem_send_command( fildes, MODEM_RESET_STRING );  
  modem_receive_response( fildes );
  if ( strcmp(message, "OK\n") != 0 )
    {
      fprintf( stderr, "Modem did not return OK on reset.\n" );
      exit(1);
    }

  do
    {
      modem_send_command( fildes, MODEM_DIAL_STRING );
      modem_receive_response( fildes );
    }
  while ( strcmp(message, "BUSY\n") == 0 );

  printf("modem returned:  %s\n", message);

  close(fildes);
  exit(0);
}


/* ---------- setup_modem() ---------- */

int setup_modem()
{
  struct termios attr;
  int fildes;
  
  if ( (fildes = open( MODEM, O_RDWR )) == -1 )
    {
      perror("device open failed");
      exit(1);
    }

  tcgetattr( fildes, &attr );
  cfsetispeed( &attr, MODEM_SPEED );
  cfsetospeed( &attr, MODEM_SPEED );
  tcflush( fildes, TCIOFLUSH );
  tcsetattr( fildes, TCSANOW, &attr );

  return fildes;
}


/* ---------- modem_send_command() ---------- */

void modem_send_command( fildes, command )
int fildes;
char *command;
{
  strcpy( message, command );
  strcat( message, "\n" );

  if ( write( fildes, message, strlen(message) ) == -1 )
    perror(NULL);

#ifdef DEBUG
  fprintf( stderr, "SEND: %s", message );
#endif /* DEBUG */

/*  sleep(1); */

  modem_receive_response( fildes );

  return;
}


/* ---------- modem_receive_response() ---------- */

void modem_receive_response( fildes )
int fildes;
{
  FILE *modem;

  message[0] = '\0';
  modem = fdopen(fildes, "rw");
  signal(SIGALRM, timeout);

/*  if (setjmp(env) == 0) */
  if (1)
    {
/*      alarm(MODEM_WAIT_PERIOD); */
/*    read( fildes, buff, sizeof(buff)-1 ); */
      fgets( message, sizeof(message), modem );
	  
#ifdef DEBUG
      fprintf( stderr, "RECEIVED: %s", message );
#endif /* DEBUG */
	  
      if ( message[strlen(message)-1] != '\n' ) 
	strcat(message, "\n");
      alarm(0);
      signal(SIGALRM, SIG_DFL);
      fclose(modem);
    }
  else
    {
      message[0] = '\0';
      signal(SIGALRM, SIG_DFL);
    }
  return;
}


/* ---------- timeout() ---------- */

void timeout(sig)
int sig;
{
  signal(sig, SIG_IGN);
  signal(SIGALRM, timeout);
  longjmp(env, 1);
}



