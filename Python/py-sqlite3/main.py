#!/usr/bin/env python

import sqlite3

def pr_tables(c):

    c.execute('''CREATE TABLE IF NOT EXISTS demo
                 (id INTEGER PRIMARY KEY AUTOINCREMENT, v TEXT)''')

    for r in c.execute('select * from sqlite_master'):
        print(r)
    c.close()



if __name__ == '__main__':

    conn = sqlite3.connect('example.db')
    pr_tables(conn.cursor())

    c = conn.cursor()

#     c.executemany('INSERT INTO demo (v) VALUES (?)', ('a','b','c','d'))
#     conn.commit()

    for r in c.execute('SELECT * FROM demo'):
        print(r)

    c.close()

    conn.close()
    print("done.")

