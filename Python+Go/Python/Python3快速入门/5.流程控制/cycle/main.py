""""
lst = [1, 2, 3, 4]
#列表 lst 中的每个元素
for i in lst:
    print("正在种第%s种子" %i)
    #%s 是 Python 中用于字符串格式化的占位符
    #包括整数、浮点数、字符串等。 %s 表示将一个字符串插入到指定位
"""

"""
s = "abcde"
l = ['a', 'b', 'c']
#定义了一个列表 l，其中包含三个字符串元素 'a'、'b' 和 'c'
n = 1
for i in l:
    print(f"现在正在处理第{n}个元素，当前元素值: {i}")
    #这条消息使用了 f-string 格式化来插入当前循环中的元素编号和元素值。
    # {n} 和 {i} 将分别被循环中的 n 和 i 的值替换，n 用于表示当前元素的编号，而 i 则表示当前元素的值
    n += 1
"""


"""
computer = {"主机":5000,"显示器":1000,"鼠标":60,"键盘":150}
for i in computer.items():
    print(f"名称：{i[0]}, 价格：{i[1]}")
"""

#多重赋值
"""
computer = {"主机":5000,"显示器":1000,"鼠标":60,"键盘":150}
k, v = ("主机", 5000)
for k, v in computer.items():
    print(f"名称：{k}, 价格：{v}")

"""

"""
l1 = [1, 2, 3, 4, 5, 6]
l2 = [4, 5, 6, 7, 8, 9]
for i in l1:
    for x in l2:
        if i == x:
            print(i)
"""

"""
#range函数
lst = []
for i in range(1,100):
    lst.append(i)
    print(lst)
# 应用场景
# 1、可以用它指定循环次数
# 2、生成序列，保存到一个列表
"""

#while语句

"""
n = 0
while n <= 10:
    print(f"重复工作执行第{n}次。")
    n += 1
"""

"""
#死循环
import time
#导入了Python标准库中的time模块，该模块包含了处理时间相关的函数和类
while True:
#True是一个布尔值，表示真。因此，while True表示“无论什么情况下”，
# 它将一直执行下去，除非在循环内部使用break语句跳出循环
    print("正在访问测试网站是否可连通...")
    time.sleep(3)
#使程序在此处暂停3秒钟。time.sleep()函数让程序休眠指定的秒数，
# 以秒为单位。在这里，它让程序每次迭代之间等待3秒钟，然后再次进入下一个循环

"""


"""
#满足条件跳出循环
for i in range(1,6):
    if i == 3:
        print("循环到3了，我做了一件事")
        continue
        print("循环到3了，我退出了")
        break
    else:
        print(i)
print("==========")
"""

"""
n = 0
while n <= 10:
    if n == 3:
        print("循环到3了，我做了一件事")
        n += 1
        break
    else:
        print(n)
        n += 1
"""

#案例：用户登录，三次错误机会
count = 0
total = 3
while True:
    if count < 3:
        name = input("请输入你的名字：").strip()
        #.strip()是一个字符串方法，用于去除字符串两端的空格
        # （包括空格、制表符、换行符等）并返回去除空格后的新字符串
        if name == "LiuHui":
            print("登录成功")
            break
        elif 0 == len(name):
            #检查用户输入的用户名是否为空（长度为0），如果是，则输出"输入不能为空"
            print("输入不能为空！")
            count += 1
            total -= 1
            print(f"你现在还剩余{total}次重试机会")
        else:
            print("用户名错误，请重新输入！")
            count += 1
            total -= 1
            print(f"你现在还剩余{total}次重试机会")
    else:
        print("超出错误次数，退出！")
        break























































