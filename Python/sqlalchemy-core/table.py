# -*- coding: utf-8 -*-
#
from sqlalchemy import create_engine
from sqlalchemy import Table, Column, Integer, String, MetaData, ForeignKey
from sqlalchemy.sql import select

metadata = MetaData()

users = Table('users', metadata,
              Column('id', Integer, primary_key=True),
              Column('username', String),
              Column('realname', String),
            )

addrs = Table('addresses', metadata,
              Column('id', Integer, primary_key=True),
              Column('user_id', None, ForeignKey('users.id')),
              Column('email', String, nullable=False),
            )

engine = create_engine('sqlite:///:memory:', echo=True)

metadata.create_all(engine)

ins = users.insert().values(username='jack', realname='Jack Hat')

conn = engine.connect()
conn.execute(ins)

ins = users.insert()
conn.execute(ins, username='hank', realname='Hank Hat')

conn.execute(users.insert(), [
    {'username': 'user1', 'realname': 'User One'},
    {'username': 'user2', 'realname': 'User Two'},
    {'username': 'user3', 'realname': 'User Three'},
])

conn.execute(addrs.insert(), [
    {'user_id': 1, 'email': 'jack@email.com'},
])

# --- Select ---

rs = conn.execute(select([users]))
for row in rs:
    print(row)

# --- Join ---

s = select([users.c.username, users.c.realname, addrs.c.email]).select_from(
    users.join(addrs)).where(addrs.c.email != None)
rs = conn.execute(s)
for row in rs:
    print(row)

