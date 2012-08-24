/* tmp.c */
#include <signal.h>
#include <setjmp.h>
#include <stdio.h>

void timeout(int);

jmp_buf env;

main()
{
  int i;
  char buf[160];

  signal(SIGALRM, timeout);
  if (setjmp(env) == 0)
    {
      alarm(5);
      printf("Type a word; if you don't in 5 seconds I'll use \"WORD\": ");
      fgets(buf, sizeof(buf), stdin);
      alarm(0);
    }
  else
    {
      strcpy(buf, "WORD");
    }
  printf("\nThe word is %s\n", buf);
  exit(0);
}

/* ----- */

void timeout(sig)
int sig;
{
  signal(sig, SIG_IGN);
  signal(SIGALRM, timeout);
  longjmp(env, 1);
}
