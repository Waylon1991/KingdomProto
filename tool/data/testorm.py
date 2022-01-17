from sqlalchemy import Column, String, Integer, create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base

MyBase = declarative_base()


class User(MyBase):
    __tablename__ = 'user'

    id = Column('id', Integer(), primary_key=True)
    name = Column('name', String(255))
    addr = Column('addr', String(255))


engine = create_engine(
    'mysql+mysqlconnector://root:123456@localhost:3306/club',
    connect_args={'auth_plugin': 'mysql_native_password'})
DBSession = sessionmaker(bind=engine)

# 创建session对象:
session = DBSession()
# 创建新User对象:
new_user = User(id='6', name='Bob', addr='yuhang')
# 添加到session:
session.add(new_user)
# 提交即保存到数据库:
session.commit()
# 关闭session:
session.close()
