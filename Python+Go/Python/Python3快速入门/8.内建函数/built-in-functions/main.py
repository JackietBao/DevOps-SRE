#高阶函数
#需求：我给你一个列表，请你把列表中的偶数拿出来

"""

a = range(1,10)
b = []
for i in a:
    if i % 2 == 0:
        b.append(i)
print(b)
"""

#需求：处理列表中每一个元素，都乘以2


"""
def f(n):
    return n * 2
a = range(1, 11)
for i in a:
    print(f(i))

result = map(f, a)
print(list(result))

result = map(lambda n:n * 2, a)
print(list(result))
"""



"""
a = range(1, 11)
def f(n):
    if n % 2 == 0:
        return n
    #return n：如果 n 是偶数，那么将 n 返回
result = filter(f, a)
print(list(result))

result = filter(lambda n: n >5, a)
print(list(result))
"""


"""
#排序
n = [3,1,2,6,8]
s = ["b","c","a","x"]

print(sorted(n))
print(sorted(s))

d = {'鼠标': 50,'主机':5000,'键盘': 150,'显示器':1000}

def f(n):
    return n[1]
('主机', 5000)
result = sorted(d.items(),key=lambda n:n[1], reverse=True)
print(result)
"""

"""
l = [1,2,3,4,5,6]
print(list(reversed(l)))
"""


"""
s = "LiuHui"
for i in s:
    print(i + ".",end="")
##将当前字符 i 与一个点号 "." 进行拼接，并使用 end="" 来避免换行，即不在每次打印后自动换行。
# 这样，每个字符都会在同一行被打印出来，以点号 "." 分隔
print('.'.join(s))
#用点号 "." 来连接字符串中的每个字符。join()
# 方法将字符串中的每个字符用指定的分隔符连接起来，这里使用了点号 "." 作为分隔符

computer = ["主机","显示器","鼠标"]
print(''.join(computer))
"""


"""
x = [1,2,3]
y = [4,5,6]
print(list(zip(x,y)))
"""

a = 1
def f():
    b = 2
    print(locals())
f()
print(globals())
































