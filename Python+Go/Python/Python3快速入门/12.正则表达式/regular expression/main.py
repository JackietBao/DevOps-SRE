"""
import re
s = "this is test string"
pattern = re.compile('this')
#使用字符串 'this' 作为模式，它表示要匹配的文本中包含 'this'
result = pattern.match(s)
#match() 方法尝试从字符串的开头进行匹配，如果匹配成功则返回一个 Match 对象，否则返回 None
print(result.group()) #匹配成功后，result对象会增加一个group()方法，可以用它来获取匹配结果


import re
s = "this is test string"

result = re.match('this', s)
print(result.group())


import re

s = "123\\abc"
result = re.match(r"123\\abc", s)
print(result)
"""






import re
s = "hello 666666"
result = re.match("hello 6+", s) # 贪婪匹配
print(result)
result = re.match("hello 6+?", s) # 非贪婪匹配
print(result)















