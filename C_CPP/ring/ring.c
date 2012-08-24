/* ring.c */

#include <fcntl.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <utmp.h>

#define LONG_TIME 300000	/* half second worth of usec's */
#define SHORT_TIME 100000	/* quarter second worth of usec's */

#ifndef UTMP_FILE
#  define UTMP_FILE "/var/adm/utmp"
#endif

#define TRUE 1
#define FALSE 0
#define nonuser(ut) (ut.ut_type == 8)

#define BELL '\x7'


void ring(FILE *);


void main(argc, argv)
int argc;
char *argv[];
{
  FILE *terminal;
  char term_str[20] = "/dev/";
  int utmp_fileid;
  struct utmp utmp_entry;

  if (argc != 2)
    {
      printf("Usage: ring username\n\n");
      exit(0);
    }

  if ((utmp_fileid = open(UTMP_FILE, O_RDONLY)) == -1)
    {
      fprintf(stderr,"Can not open %s.\n", UTMP_FILE);
      exit(1);
    }

  while (read(utmp_fileid, &utmp_entry, sizeof(utmp_entry)) > 0)
    {
      if (nonuser(utmp_entry) == FALSE)
	if (strncmp(argv[1],utmp_entry.ut_name,sizeof(argv[1])) == 0)
	  break;
    }

  close(utmp_fileid);
  if (strncmp(argv[1],utmp_entry.ut_name,sizeof(argv[1])) != 0)
    {
      printf("%s is not logged in.\n", argv[1]);
      exit(1);
    }

  strcat(term_str, utmp_entry.ut_line);
  if ((terminal = fopen(term_str, "w")) == NULL)
    {
      printf("%s is not accepting messages on %s\n", argv[1], utmp_entry.ut_line);
      exit(1);
    }

  printf("Ringing %s on %s.\n", argv[1], utmp_entry.ut_line);
  ring(terminal);
  fclose(terminal);

  return;
}

/* ---------- ring() ---------- */

void ring(filevar)
FILE *filevar;
{

  putc(BELL,filevar);
  fflush(filevar);
  usleep(SHORT_TIME);
  putc(BELL,filevar);
  fflush(filevar);
  usleep(LONG_TIME);
  putc(BELL,filevar);
  fflush(filevar);

  return;
}
