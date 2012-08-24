/* prep.c */

/*  Prep takes a file saved by VMSNEWS and converts it to a standard
 * unix mail format file, which can subsequently be read using a normal
 * mail reader (i.e. elm).
 *
 * Currently prep inserts a static From line:
 *
 *  From egustafs Sat Jan 1 00:01:01 1994
 *
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define PROGRAM_NAME "prep"
#define LINE_LEN 80
#define FROM_LINE "From egustafs Sat Jan 1 00:01:01 1994\n"

struct list_elm *list_seek(const char *, struct list_elm *);
struct list_elm *list_rm_line(struct list_elm *);
void list_insert(char *);
void read_file(char *);
void write_file(char *);
void usage(void);

struct list_elm {
  char line[LINE_LEN];
  struct list_elm *next;
  struct list_elm *prev;
};

struct list_elm *list_head = NULL;	/* ptr to the head of the list */
struct list_elm *list_pos;	/* current position in list */
extern char **environ;		/* ptr to environment variables */


/* ---------- main() ---------- */

void main(argc, argv)
int argc;			/* num of arguments passed */
char **argv;			/* ptr to argument strings */
{
  struct list_elm *pos_ptr;	/* ptr into the list structure */

  if ( argc != 2 )
    usage();

  list_insert(FROM_LINE);
  read_file(argv[1]);

  pos_ptr = list_head;
  while (pos_ptr != NULL)
    {
      pos_ptr = list_seek("+---+", pos_ptr);
      if ( pos_ptr != NULL )
	{
	  pos_ptr = list_rm_line(pos_ptr);
	  strcpy(pos_ptr->line, FROM_LINE);
	}
    }

  write_file(argv[1]);

  printf("Exiting normally.\n\n");
  exit(0);
}


/* ---------- read_file() ---------- */

void read_file(filename)
char *filename;
{
  char line[100];		/* line buffer */
  int  line_count = 0;		/* number of lines read in */
  FILE *infile;			/* input file */
  
  if ((infile = fopen(filename, "r")) == NULL)
    {
      perror("Error: could not open file");
      usage();
    }
  
  while ( ~feof(infile) )
    {
      fgets(line, sizeof(line), infile);
      list_insert(line);
      line_count++;
    }
  fclose(infile);
  printf("%d lines read\n", line_count);
  return;
}


/* ---------- write_file() ---------- */

void write_file(filename)
char *filename;
{
  FILE *outfile;		/* output file */
  struct list_elm *pos_ptr;	/* pointer into the list */

  pos_ptr = list_head;
  if ((outfile = fopen(filename, "w")) == NULL)
    {
      perror("Error: could not open file");
      exit(1);
    }

  while ( pos_ptr != NULL )
    {
      fputs(pos_ptr->line, outfile);
      pos_ptr = list_rm_line(pos_ptr);
    }

  return;
}


/* ---------- list_seek() ---------- */

/*
 *  Seek forward to the line which has pattern in it.  Return
 * a pointer to that line.  Start searching at pos_ptr or the
 * beginning if start_pos is NULL.
 */

struct list_elm *list_seek(pattern, pos_ptr)
const char *pattern;
struct list_elm *pos_ptr;
{
  if (pos_ptr == NULL)
    pos_ptr = list_head;

  while ( (pos_ptr != NULL) && (strstr(pos_ptr->line, pattern) == NULL) )
    pos_ptr = pos_ptr->next;

  return(pos_ptr);
}


/* ---------- list_rm_line() ---------- */

/*
 *  Remove the line pointed to by pos_ptr, returns a pointer
 * to the next line.
 */

struct list_elm *list_rm_line(pos_ptr)
struct list_elm *pos_ptr;
{
  struct list_elm *tmp_ptr;

  if (pos_ptr == NULL)
    return NULL;

  if ((pos_ptr == list_head) && (pos_ptr->next == NULL))
    {
      tmp_ptr = NULL;
      /* do nothing, but don't execute any other option */
    }
  else if (pos_ptr == list_head)
    {
      pos_ptr->next->prev = NULL;
      list_head = pos_ptr->next;
    }
  else if (pos_ptr->next == NULL)
    {
      pos_ptr->prev->next = NULL;
    }
  else
    {
      pos_ptr->next->prev = pos_ptr->prev;
      pos_ptr->prev->next = pos_ptr->next;
    }

  tmp_ptr = pos_ptr->next;
  free(pos_ptr);
  return(tmp_ptr);
}


/* ---------- list_insert() ---------- */

/*
 *  Insert the string pointed to by line onto the tail end
 * of the list.
 */

void list_insert(line)
char *line;
{
  static struct list_elm *list_tail;
  struct list_elm *new_node;

  if ((new_node = (struct list_elm *)malloc(sizeof(struct list_elm))) == NULL)
    {
      perror("Out of memory");
      exit(1);
    }

  if ( list_head != NULL )
    {
      list_tail->next = new_node;
      new_node->prev = list_tail;
      new_node->next = NULL;
      list_tail = new_node;
    }
  else
    {
      list_head = new_node;
      list_tail = new_node;
      new_node->prev = NULL;
      new_node->next = NULL;
    }

  strcpy(new_node->line, line);

  return;
}

/* ---------- usage() ---------- */

void usage()
{
  fprintf(stderr, "Usage:  %s <filename>\n\n", PROGRAM_NAME);
  exit(1);
}
