computer = ["主机","显示器","键盘","鼠标"]
# 查询
""""
print(computer.index("显示器"))# 获取索引
print(computer.count("键盘"))# 统计元素数量
computer.sort()# 正向排序，修改原有对象
print(computer)
computer.reverse()# 倒向排序，修改原有对象
print(computer)
"""

# 增加
"""
computer.append("音响")
print(computer)
computer.insert(1,"鼠标垫")
print(computer)
"""

# 修改
"""
computer[3] = '音响'
print(computer)
"""

# 删除
"""
computer.remove("鼠标")
print(computer)
computer.pop(2)
print(computer)
"""

# 切片
"""
print(computer[0])
print(computer[0:3])
print(computer[-1])
print(computer[0:-1])
"""

# 清空列表
print(len(computer))
computer.clear()
print(computer)
computer = []
print(computer)
del computer































