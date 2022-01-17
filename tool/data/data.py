import mysql.connector

conn = mysql.connector.connect(
    host='localhost',
    user='root',
    password='123456',
    database='club',
    auth_plugin='mysql_native_password'
    )
cursor = conn.cursor()

cursor.execute('select * from login_user_infos')
values = cursor.fetchall()
print(values)
