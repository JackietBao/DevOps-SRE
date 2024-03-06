#特点
"""
1、无序的
2、不重复
"""

#两种定义方法
computer = {"主机","显示器","鼠标","键盘"}
lst = ["主机","显示器","鼠标","键盘"]
computer2 = set(lst)
#set()函数将列表lst转换为一个集合
print(computer)
print(computer2)

#添加
computer.add("音响")
print(computer) #添加是无序的
#删除
computer.remove("键盘")
print(computer)
computer.pop()
#用于从集合中删除并返回一个元素的方法
#具体地说，它会从集合中移除并返回一个随机选择的元素
print(computer)
computer.discard("主机")
print(computer)

#去重
lst = ["主机","显示器","鼠标","键盘","显示器","键盘"]
print(list(set(lst)))
#set(lst)：利用 set() 函数将列表 lst 转换为一个集合。
#集合是一种无序且不重复的数据结构，所以在集合中每个元素都是唯一的
print("===========")
#关系测试

a = set([1, 2, 3, 4, 5, 6])
b = set([4, 5, 6, 7 ,8, 9])
print(a)
print(b)

#返回a集合中元素在b集合没有的
print(a - b)
#返回b集合中元素在a集合中没有的
print(b - a)
#返回交集，即两个集合中一样的元素
print(a & b)
#返回集合，即合并去重
print(a | b)
#判断是否不相等
print(a != b)
#判断是否相等
print(a == b)


















