"""
#函数传参
def hello(a,b=0):
    #定义了一个名为 hello 的函数，该函数接受两个参数 a 和 b。
    # 其中，参数 a 是必须传递的，而参数 b 则是有默认值的，如果不提供该参数，则默认为 0
    print(f"这是在函数内引用变量{a}")
    #打印了一个字符串，这个字符串中包含了函数内部引用的变量 a 的值。
    # 这是一个调试语句，用于显示函数内部的变量值
    #f 表示这是一个格式化字符串。而 {a} 则是一个占位符，表示在这个位置插入变量 a 的值。
    # 所以 f"这是在函数内引用变量{a}" 的意思是，在这个字符串中，{a} 将被替换为变量 a 的值
    print(a)
    print(b)
    return f"总和: {a+b}"
    返回值在调用函数时被输出
print(hello(1,6))


def seq(n):
    result = []
    x = 0
    while x < n:
        result.append(x)
        x += 1
    return result
print(seq(9))
"""


"""
#接受任意数据参数
def func(*args, **kwargs):
    #*args 是一个元组，它接受任意数量的位置参数
    #**kwargs 是一个字典，它接受任意数量的关键字参数
    print(args)
    print(kwargs)

func(1,2,3,a=1,b=2,c=3)


s = lambda  a,b: a+b
#a 和 b 是函数的参数，而 a + b 是函数的返回值
print(s(6,6))

t = lambda x : "偶数" if x % 2 == 0 else "奇数"
print(t(7))

"""

"""
#作用域
a = 2
def func():
    b = 3
print(a)
print(b)


import builtins
#builtins模块包含了Python的内建函数、异常等核心组件
print(dir(builtins))
#dir(builtins)返回了builtins模块的所有属性和方法的名称列表
"""

"""
x = 123
def outer():
    #定义了一个名为 outer 的函数
    x = 456
    def inner():
        #outer 函数内部，又定义了一个名为 inner 的函数
        print(x)
        #inner 函数内部，打印了变量 x 的值。
        # 由于 inner 函数没有自己的 x 变量，所以它会按照 LEGB（Local -> Enclosing -> Global -> Built-in）规则查找，
        # 找到了 outer 函数内的 x 变量，因此打印的是 outer 函数内部的 x 的值，即 456
    return inner()
#outer 函数内部，返回了 inner 函数的调用结果。
# 由于 inner 函数在 outer 函数内部定义，因此它可以访问 outer 函数内部的变量，包括局部变量 x
s = outer()
#调用 outer 函数，并将返回的结果赋值给变量 s。
# 因为 inner 函数在 outer 函数内部被调用并返回了，所以 s 实际上被赋值为 inner 函数的调用结果，即打印 x 的值
#s()
#outer()()

"""


"""
def hello():
    #定义了一个名为 hello 的简单函数，该函数不接受任何参数
    print("我是函数")
dec = decorator(hello)
#将 hello 函数传递给装饰器函数 decorator，并将返回的结果赋值给变量 dec。这一步是在应用装饰器
dec()
#调用装饰后的函数。由于 decorator 函数返回了内部函数 f，所以实际上调用的是 f 函数。

"""

"""

# 给一个函数增加一个记录日志的功能
def decorator(func):
    #装饰器函数的定义，它接受一个函数作为参数
    def f(**name):
        #装饰器函数内部，定义了一个内部函数 f。这个内部函数接受任意数量的关键字参数
        print("记录一下开始。")
        func(**name)
        #调用了传入装饰器的原始函数 func，并将传入的关键字参数 name 解包后传递给它
        print("记录一下结束。")
    return f
#返回内部函数 f，它是一个装饰器函数的结果


@decorator #dec = decorator(hello)
def hello2(**name):
    print(name)
    print("我是一个求和函数，正在执行")
hello2(a=1,b=2)
"""

"""
def hello():
    print("hello world!")
hello()

"""



"""
def f(a,b):
    return a + b
print(f(1,2)) #按参数位置赋值
print(f(b=2,a=1))#按对应关系赋值

"""

"""
def seq(n):
    result = []
    x = 0
    while x < n:
        result.append(x)
        x += 1
    return result
print(seq(9))
"""


"""
def f(a,b=2):
    return a + b
print(f(1))
print(f(1,3))

"""

"""
def func(*seq):
    x = 0
    for n in seq:
        x += n
    return x

print(func(1,2,3))
"""

"""
def func(**computer):
    for k,v in computer.items():
        print("名称: %s\t价格:%s"%(k,v))
func(主机=5000,显示器=1000,鼠标=60,键盘=150)
"""

"""
def outer():
    x = 1
    def inner():
        print(x)
    return inner() #不加括号表示返回函数体
sf = outer()
"""

"""
#无参数
def hello():
    print("我是原函数")

    #当被调用时，它会打印一条消息 "我是原函数"

def decorator(func):
    #接受一个函数 func 作为参数。装饰器的作用是，接收一个函数，然后返回一个新的函数
    def f():
        print("原函数开始执行了")
        func()
        print("原函数执行结束了")
        #定义了一个新的函数 f()。这个新的函数 f 是装饰器的核心，它实际上是对原函数的一个包装
    return f #返回函数体
#最后返回了新函数 f 的函数体（但没有执行它）。
# 这样，装饰器提供了一种方式，通过返回新的函数来增强原有函数的功能。

#当调用 f() 时，它首先打印消息 "原函数开始执行了"，表明原函数的执行即将开始
#然后，它调用原始的 func() 函数（在这个例子中，是 hello 函数），这样就可以执行原函数定义的功能
#最后，它打印 "原函数执行结束了"，表明原函数的执行已经完成。

dec = decorator(hello) #装饰器传入函数
dec() #调用内部函数


#这一行代码演示了如何使用装饰器。它首先调用装饰器 decorator，
# 传递 hello 函数作为参数。装饰器返回新的函数 f，然后将这个新函数赋值给变量 dec
#最后，通过调用 dec()，实际上是调用了装饰器返回的新函数 f()。

"""








