""""
age = int(input("小朋友，请输入你的年龄，让我来判断你是不是成年了："))
if age > 18:
    print("恭喜你，可以谈女朋友了。")
else:
    print("不好意思，你还不能谈女朋友！")
result = "恭喜你，可以谈女朋友了。" if age > 18 else "不好意思，你还不能谈女朋友！"
print(result)
"""

age = int(input("请输入你的年龄: "))
if age < 7:
    print("儿童")
elif age > 7 and age < 17:
    print("少年")
elif age >= 18 and age < 41:
    print("青年")
elif age >= 41 and age < 48:
    print("壮年")
else:
    print("老年")

