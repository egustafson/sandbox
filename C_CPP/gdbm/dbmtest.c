/* dbmtest.c */

#include<stdio.h>
#include<gdbm.h>
#include<sys/stat.h>

#define FILEPERM S_IRUSR|S_IWUSR|S_IRGRP|S_IROTH
#define FILENAME "database.gdbm"

typedef struct {
  char key[10];
  char data[50];
} data_rec;

void      insert(GDBM_FILE, data_rec *);
data_rec *retrive(GDBM_FILE, char *);

/* -------------------------------- */
void main(void)
{
  data_rec  new_record;
  data_rec  *record;
  GDBM_FILE dbm;
  
  strcpy(new_record.key, "b");
  strcpy(new_record.data, "Last Data");

  dbm = gdbm_open( FILENAME, 0, GDBM_WRCREAT, FILEPERM, NULL );

/*  insert(dbm, &new_record); */
  record = retrive(dbm, new_record.key);

  gdbm_close( dbm );

  printf("Key:  %s\n", record->key);  
  printf("Data: %s\n", record->data); 

  exit(0);
}

/* ---------------------------------------- */

data_rec *retrive( dbm_ptr, key )
GDBM_FILE dbm_ptr;
char      *key;
{
  datum  key_entry;
  datum  record_entry;
  
  key_entry.dptr = key;
  key_entry.dsize = 10;

  record_entry = gdbm_fetch( dbm_ptr, key_entry );

  if (record_entry.dptr == NULL)
    fprintf(stderr, "record_entry.dptr is NULL\n");

  return (data_rec *)record_entry.dptr;
}

/* ---------------------------------------- */

void insert( dbm_ptr, record_ptr )
GDBM_FILE dbm_ptr;
data_rec *record_ptr;
{
  datum  key_entry;
  datum  record_entry;
  int    ret;

  key_entry.dptr = record_ptr->key;
  key_entry.dsize = 10;
  record_entry.dptr = (char *)record_ptr;
  record_entry.dsize = sizeof(data_rec);

  ret = gdbm_store( dbm_ptr, key_entry, record_entry, GDBM_INSERT );
  if (ret != 0)
    fprintf(stderr, "gdbm_store returned:  %d\n", ret);

  return;
}
