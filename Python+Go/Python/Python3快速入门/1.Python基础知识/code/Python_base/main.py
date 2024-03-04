# 算术运算符
'''
print(6+6)
print(12-8)
print(6*6)
print(6/6)
print(type("6"))
print(6 ** 3)
'''

# 变量

'''
name = "aliang"
print("你叫什么名字：")
print(name)


name, age = "aliang", 30
print(name)
print(age)

calc = 6 / 8
name = "aliang"
age = 30
print("你叫什么名字：%s，年龄: %d" %(name, age))

# calc = 6
# calc += 6 # calc = 6 + 6
# print(calc)
calc = 6
calc *= 6  # calc = 6 * 6
print(calc)
'''

# 转义符
'''
name = 'aliang'
age = 30
print("姓名: %s, \\n年龄: %d" %(name, age))

print("hello \nworld!")
print("hello \tworld!")
print("hello \rworld!")
print("hello \world!")
print("hello \
      world")

'''

# 获取用户输入
'''
age = input("小朋友，今年多大啦？请输入你的年龄：")
print("呦，都%s岁了！可以谈恋爱了。" %age)
'''

# 简单的计算器

print("选择算术运算符：")
# 方式1
# print("1、加")
# print("2、减")
# print("3、乘")
# print("4、除")
# 方法2
# print("1、加\n2、减\n3、乘\n4、除")
# 推荐方式
'''
print("""
    1、加
    2、减
    3、乘
    4、除
""")
choice = input("请输入编号：")
n1 = int(input("请输入第一个数字："))
n2 = int(input("请输入第二个数字："))
if choice == '1':
    print(n1+n2)
elif choice == '2':
    print(n1-n2)
elif choice == '3':
    print(n1*n2)
elif choice == '4':
    print(n1/n2)
else:
    print("输入的编号不对！")
'''

