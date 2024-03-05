# f'{表达式}' 引用变量
"""
name = 'LiuHui'
age = 16
print("name: %s, age: %d" %(name, age))
"""

# 字符串拼接
"""
s1 = "LiuHui"
s2 = "love sleep"
print(s1 + s2)
print(s1,s2)
"""

# 获取字符串长度 len()
"""
s = "hello liuhui"
print(len(s))
"""

# 字符串切片
"""
s = "hello liuhui"
print(s[4])
print(s[4:7])
print(s[-2])
print(s[0:-1])
"""

xxoo = "abcdef"
print("首字母大写: %s" % xxoo.capitalize())
print("字符l出现次数: %s" % xxoo.count('l'))
print("感叹号是否结尾: %s" % xxoo.endswith('!'))
print("w字符是否是开头: %s" % xxoo.startswith('w'))
print("w字符索引位置: %s" % xxoo.find('w')) # xxoo.index('W')
print("格式化字符串: Hello{0} world!".format(','))
print("是否都是小写: %s" % xxoo.islower())
print("是否都是大写: %s" % xxoo.isupper())
print("所有字母转为小写: %s" % xxoo.lower())
print("所有字母转为大写: %s" % xxoo.upper())
print("感叹号替换为句号: %s" % xxoo.replace('!','.'))
print("以空格分隔切分成列表: %s" % xxoo.split(' '))
print("切分为一个列表: %s" % xxoo.splitlines())
print("去除两边空格: %s" % xxoo.strip())
print("大小写互换: %s" % xxoo.swapcase())
print("----------")
print("a字符是否是开头: %s" % xxoo.startswith('a'))
s = "我你中国"
print("感叹号替换为句号: %s" % s.replace('中国','北京'))
s = "1,2,3"
print(s.split(','))
name = input("请输入你的名字: ").strip()
print(name)




