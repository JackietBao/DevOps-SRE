# 类型转换四个函数 int() float() str() bool()
# int() 可以用来将其他的对象转换为整型
# 规则：
#   布尔值：True -> 1   False -> 0
#   浮点数：直接取整，省略小数点后的内容
#   字符串：合法的整数字符串，直接转换为对应的数字
#           如果不是一个合法的整数字符串，则报错 ValueError: invalid literal for int() with base 10: '11.5'
#   对于其他不可转换为整型的对象，直接抛出异常 ValueError
# float() 和 int()基本一致，不同的是它会将对象转换为浮点数
# str() 可以将对象转换为字符串
#  True -> 'True'
#  False -> 'False'
#  123 -> '123' 
#  。。。
#  bool() 可以将对象转换为布尔值，任何对象都可以转换为布尔值
#   规则：对于所有表示空性的对象都会转换为False，其余的转换为True
#       哪些表示的空性：0 、 None 、 '' 。。。

a = True

# 调用int()来将a转换为整型
# int()函数不会对原来的变量产生影响，他是对象转换为指定的类型并将其作为返回值返回
# 如果希望修改原来的变量，则需要对变量进行重新赋值
a = int(a)

a = False
a = int(a)

a = '123'
a = int(a)

a = 11.6
a = int(a)

a = '11.5'
# a = int(a)#int() 函数只能用于将字符串表示的整数转换为整数类型，而不能处理包含小数点的字符串

a = None
# a = int(a)#None 是一个特殊的对象，表示“没有值”或“空”。尝试将 None 转换为整数是不合法的操作，会导致 TypeError 错误

a = 1
a = float(a)

a = False
a = float(a)

a = 123
a = str(a)

a = None
a = bool(a)

print('a =',a)
print('a的类型是',type(a))
# b = 456
# print('hello'+str(b))