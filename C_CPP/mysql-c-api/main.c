#include <sys/time.h>
#include <stdio.h>
#include <mysql.h>

int main(int argc, char** argv) {
    
    MYSQL_RES* result;
    MYSQL_ROW  row;
    MYSQL*     connection;
    MYSQL      mysql;
    int        state;

    mysql_init(&mysql);

    connection = mysql_real_connect(&mysql, "sif", 0, 0, "test", 0, 0, 0);
    if( NULL == connection ) {
        printf(mysql_error(&mysql));
        return 1;
    }

    state = mysql_query(connection, "SELECT * from testac");

    if ( state != 0 ) {
        printf(mysql_error(connection));
        return 1;
    }

    result = mysql_store_result(connection);
    
    while ( ( row = mysql_fetch_row(result) ) != NULL ) {
        printf("id: %s, title: %s\n", 
               (row[0] ? row[0] : "NULL"),
               (row[1] ? row[1] : "NULL") );
    }

    mysql_free_result(result);

    mysql_close(connection);
        
    printf("Done\n");

    return 0;
}
