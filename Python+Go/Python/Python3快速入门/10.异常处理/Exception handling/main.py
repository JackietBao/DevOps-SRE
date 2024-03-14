n1 = '2'
n2 = 2
try:
    #try 块中的代码引发了一个异常，程序控制流转移到 except 块
    n = n1 + n2
except  Exception as e:
    #except 块中，异常被捕获，Exception as e 语句将异常对象赋给了变量 e
    print("错误：%s" %e)
    #打印了异常的消息，这里是类型错误的消息
    #%s 表示字符串格式化
    #%e 用于表示异常对象
    print("发生错误执行的代码")

'''
1、兼容类的处理
2、明确的要求
3、上面表达式有已知类的错误，用except备选处理
'''


try:
    f = open('computer.txt')
    #程序尝试打开名为 computer.txt 的文件，并将文件对象赋给变量 f
    f.read()
    #调用 f.read() 读取文件内容
except Exception as e:
    #如果在 try 块中的任何代码引发了异常，控制流将转移到 except 块
    print("文件不存在")