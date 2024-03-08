#比较操作符
n1 = 8
n2 = 6

print(n1>n2)

#逻辑操作符
a = 3
b = 4
c = 3

print(a < b)
print(c > b)
#都为真才为真
print(a < b and c > b)
# 一个为真则为真，都为假才为假
print(a < b or c > b)
# 取反
print(not c > b)

#成员操作符
print("==========")
computer = ["主机","显示器","鼠标","键盘"]
d = {"主机": 5000}
print("主机" in computer)

if "主机" in d:
    print("在里面")
if "音响" not in d:
    print("不在里面")
print(d)

#身份操作符
print("==========")
a = []
b = []
print(id(a))
print(id(b))
print(a is b)
#由于 Python 中的列表是可变对象，即使两个列表具有相同的内容，
#它们在内存中的地址也可能不同。在这种情况下，a is b 会返回 False，
#因为它们指向不同的内存位置





