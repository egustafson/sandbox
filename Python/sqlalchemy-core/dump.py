# -*- coding: utf-8 -*-
#
from sqlalchemy import create_engine
from sqlalchemy import Table, Column, Integer, String, MetaData
from sqlalchemy.sql import select

metadata = MetaData()

tab1 = Table('t1', metadata,
   Column('id', Integer, primary_key=True),
   Column('key', String),
   Column('val', String),
)

engine = create_engine('sqlite:///:memory:', echo=True)
metadata.create_all(engine)
conn = engine.connect()

conn.execute(tab1.insert(), [
    {'key': 'key-1', 'val': 'val-1'},
    {'key': 'key-2', 'val': 'val-2'},
    {'key': 'key-3', 'val': 'val-3'},
])

rs = conn.execute(select([tab1]))
for row in rs:
    print(row)
