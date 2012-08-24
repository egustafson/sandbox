#!/usr/bin/env python
# ######################################################################
#
# An example of how to perform DB access transactionally using
# the Python DB-API through MySQLdb.
#
# This demo needs the following table created in the MySQL DB named
# below and accessible by the user/pass listed below:
#
#   CREATE TABLE demo ( id INTEGER PRIMARY KEY AUTO_INCREMENT,
#                       name VARCHAR(255)
#                     ) ENGINE=InnoDB ;
#
# Note:  ENGINE=InnoDB is necessary to ensure a table that supports
#   transactional semantics.
#

import MySQLdb


# ###################################
# Constants

DB_USER = 'test'
DB_PASS = 'test'
DB_HOST = 'localhost'
DB_PORT = 3306
DB_DB   = 'test'


# ######################################################################

def printRows(cursor, connection):
    connection.commit()  # Get the latest view
    cursor.execute('SELECT * FROM demo');
    result = cursor.fetchall()
    print("Table demo has %s rows." % (len(result)))
    

# ######################################################################
if __name__ == '__main__':

    conn1 = None
    conn2 = None
    try:
        conn1 = MySQLdb.connect( host=DB_HOST, port=DB_PORT, user=DB_USER, passwd=DB_PASS, db=DB_DB)
        conn2 = MySQLdb.connect( host=DB_HOST, port=DB_PORT, user=DB_USER, passwd=DB_PASS, db=DB_DB)
        #
        # Disable autocommit, transactions are desired
        #
        if hasattr(conn1, 'autocommit'):
            conn1.autocommit(0)
            conn2.autocommit(0)
        else:
            print("WARNING:  connection does not have 'autocommit' attribute, transactions may not be supported.")
    except:
        print("Problem connecting to the database.")


    cur1 = conn1.cursor()
    cur2 = conn2.cursor()

    printRows(cur1, conn1)
    
    cur2.execute("INSERT INTO demo (name) VALUES ('fred')")
    print("Inserted into table, but NO commit, yet.")

    printRows(cur1, conn1)

    conn2.rollback()
    print("Rolled back the insert.")

    conn2.commit()
    print("Committed insert, null effect if previously rolled back.")

    printRows(cur1, conn1)
    print("Demo complete.")
