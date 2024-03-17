f = open('computer.txt', encoding="utf8", mode='w+')
#'w' 表示写入模式。在写入模式下打开文件会清空文件的内容，如果文件不存在则会创建文件。
#'+' 表示可读写模式。它允许对文件进行读取和写入操作。如果文件不存在则会创建文件

#写入模式打开的（'w+'），这意味着如果文件不存在，则会被创建；如果文件已存在，那么它的内容将会被清空。
#print(f.read())
#print(f.read(4))
#读取下一行内容
#print(f.readline())
f.write("\n键盘")
#在文件中写入一个新行并添加文字“键盘”
f.flush()
#将缓冲区的数据写入文件，确保数据被立即写入而不是等待缓冲区满或文件关闭时才写入

computer = ["主机","显示器","键盘","鼠标"]
#创建一个列表变量computer，其中包含四个字符串元素
for i in computer:
    f.write("\n" + i)
f.flush()
f.close()
#关闭文件，结束文件操作

computer = ["主机", "显示器", "键盘", "鼠标","音响"]
with open('computer.txt', encoding="utf8", mode='w+') as f:
#with确保在程序块执行完毕后自动关闭文件，不论是否出现异常。在此，文件句柄被赋值给变量f
#as f将打开的文件对象（或上下文管理器）赋值给变量f，以便在with语句块中使用。
    for i in computer:
        f.write("\n" + i )
    f.flush()













