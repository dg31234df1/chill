# -*- coding: utf-8 -*-
from flask import Flask
from flask_sqlalchemy import SQLAlchemy

print ("Load paramters")
app = Flask(__name__)
app.config.from_object('engine.settings')

#创建数据库对象
print ("Create database instance")
db = SQLAlchemy(app)

from engine.model.IndexTable import IndexTable

from engine.controller import IndexManage
