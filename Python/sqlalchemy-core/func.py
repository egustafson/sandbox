# -*- coding: utf-8 -*-
#
from sqlalchemy import create_engine
from sqlalchemy.sql import func

engine = create_engine('sqlite:///:memory:', echo=True)
conn = engine.connect()

fn = func.now()

rs = conn.execute(fn)

print(rs.fetchall())
