computer = ("主机","显示器","键盘","鼠标","显示器")
print(computer)

#特点
#1、不支持添加元素
#2、不支持删除元素
#3、不知处修改元素
#4、不支持查询元素

print(computer.index("显示器"))
print(computer.count("显示器"))
print(computer[0:-1])

#修改元素，先转为列表再修改，最后再转回元组
l = list(computer)
print(l)
l.append("音响")
print(l)
print(tuple(l))

#元组应用场景
"""
1、你希望这个列表数据是受到保护，不可修改
2、函数的返回值是用元组类型返回
"""