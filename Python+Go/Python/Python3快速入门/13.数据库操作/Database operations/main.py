import pymysql

conn = pymysql.connect(host='192.168.31.61',
                       port=3309,
                       user='root',
                       password='123456',
                       db='test',
                       charset='utf8',
                       #连接使用的字符集，通常是 'utf8' 或 'utf8mb4
                       cursorclass=pymysql.cursors.DictCursor
                       #指定了游标类，这里指定为 pymysql.cursors.DictCursor，表示返回的查询结果以字典的形式表示，这样可以通过列名访问结果集中的数据
                       )
cursor = conn.cursor()
#创建了一个游标对象，通过这个对象可以执行 SQL 查询和操作数据库

sql = "insert into user(username, password) values('aliang','123456')"
#将一个新用户的用户名和密码插入到名为 "user" 的表中
cursor.execute(sql)
#使用游标对象执行上述 SQL 插入语句，将新用户的信息插入到数据库中
conn.commit()
#提交事务，将之前的 SQL 插入操作提交到数据库中。
# 这一步是必要的，因为默认情况下 pymysql 的连接是自动提交模式，
# 但是在某些情况下（例如在事务中执行多个 SQL 操作），你可能需要手动提交以确保数据被正确保存
sql = "select * from user"
#义了一个 SQL 查询语句，从名为 "user" 的表中检索所有列的数据
cursor.execute(sql)
result = cursor.fetchall()
#从游标对象中获取所有查询结果，并将其存储在变量 result 中
#查询结果是一个列表，每个元素是一个字典，表示一行数据，字典的键是列名，值是对应的数据
print(result)
conn.close()
#关闭数据库连接，释放资源


sql = "insert into user(username, password) values(%s, %s)"
#{'张三':'123456','李四':'123456',...}
#[('张三','123456'),('李四','123456'),()...]
args = [('zhangsan','123456'),('lisi','123456'),('zhaowu','123456')]
cursor.executemany(sql, args)
#批量执行相同的 SQL 语句，但每次执行都使用不同的参数
#sql 是之前定义的 SQL 插入语句，args 是包含了多组参数的列表
#这行代码将会执行三次插入操作，每次使用不同的用户名和密码参数
conn.commit()
#提交事务，将之前的 SQL 操作提交到数据库中。
# 因为使用了 executemany() 方法，需要手动提交事务以确保所有的插入操作都被正确保存到数据库中


sql = "update user set password='456789' where username='aliang'"
sql = "delete from user where username='aliang'"
cursor.execute(sql)
conn.commit()
print(cursor.fetchone()) # 获取第一条记录
print(cursor.fetchmany(3)) # 获取指定多少条记录
print(cursor.fetchall())  # 获取所有记录


try:
    with conn.cursor() as cursor: # cursor = conn.cursor()
        sql = "select * from user"
        cursor.execute(sql)
        result = cursor.fetchall()
    for d in result:
        print(f"ID: {d['id']}, 用户名: {d['username']}, 密码: {d['password']}")
        #打印每行数据的 ID、用户名和密码字段
finally:
    conn.close()



























