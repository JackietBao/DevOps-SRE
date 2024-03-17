"""
import os
print(os.name)
#返回操作系统类型

import sys
arg_list = sys.argv
#sys.argv 是一个列表，其中包含了命令行调用脚本时传递的所有参数，其中第一个参数是脚本的名称
arg_len = len(arg_list)-1
#计算了除去脚本名称之外的参数数量，并将结果存储在 arg_len 变量中
if arg_len == 2:
    print("参数输入正确")
else:
    print("位置参数数量不对，示例：%s -a -b" %arg_list[0])
    #%arg_list[0] 将会用脚本的名称替换它


import  sys

print("hello world")
sys.exit(0)  # 0正常退出，非0异常退出
sys.exit(1)
print("hello world2")


import platform
print(platform.platform())
print(platform.uname())
print(platform.system())
print(platform.version())


import subprocess
cmd="ls"
result = subprocess.run(cmd, shell=True, stdout=subprocess.PIPE,stderr=subprocess.PIPE)
#subprocess.run() 函数来执行命令。参数 cmd 指定了要执行的命令
#shell=True 表示在系统的 shell 中执行命令（例如，在 Unix-like 系统上使用 Bash）
#stdout=subprocess.PIPE 表示将命令的标准输出重定向到一个管道
#stderr=subprocess.PIPE 表示将命令的标准错误输出重定向到另一个管道
#subprocess.run() 函数将等待命令执行完成，并返回一个 CompletedProcess 对象，其中包含了执行结果的信息
print(result)
print(result.returncode)
#打印执行命令后的返回码。返回码为 0 通常表示命令执行成功，而非 0 的返回码通常表示出现了错误
print(result.stdout.decode())
#将命令的标准输出解码为字符串并打印出来。由于之前使用了 stdout=subprocess.PIPE，
# 命令的输出被捕获并存储在 result.stdout 中。.decode() 方法用于将字节流解码为字符串
print(result.stderr.decode(encoding='gbk'))  # 使用 GBK 解码
#命令的标准错误输出解码为字符串并打印出来。
#由于使用了 stderr=subprocess.PIPE，命令的错误输出被捕获并存储在 result.stderr 中。.decode() 方法用于将字节流解码为字符串







import pickle

computer = {"主机": 5000, "显示器": 1000, "鼠标": 60, "键盘": 150}

with open("data.pkl", "wb") as f:
    #with 上下文管理器打开一个名为 data.pkl 的文件，以便将 computer 字典对象序列化并保存到文件中。
    # 打开文件的模式为二进制写入模式（"wb"）
    pickle.dump(computer, f)
    #函数将 computer 字典对象序列化并写入到文件 f 中
    #把字典对象转换为字节流，并写入到文件中，以便后续反序列化时使用

with open("data.pkl", "rb") as f:
    #再次使用 with 上下文管理器打开同一个文件 data.pkl
    #以二进制读取模式（"rb"）打开，以便从文件中读取已序列化的对象
    d = pickle.load(f)
    #从文件 f 中反序列化已保存的对象。这将读取文件中的字节流，并将其转换为原始的 Python 对象

print(d)



import json

computer = {'主机':5000,"显示器":1000,"鼠标":60,"键盘":150}
json_obj = json.dumps(computer)
#函数将 Python 对象 computer 编码为 JSON 格式的字符串
#
print(type(json_obj))
#打印变量 json_obj 的类型
print(computer)
print(json_obj)
#打印经过 JSON 编码后的 JSON 字符串 json_obj。

data = json.loads(json_obj)
#json.loads() 函数将 JSON 字符串 json_obj 解码为 Python 对象
print(data)

import time

print(time.time())  # 当前时间戳
print(time.strftime("%Y-%m-%d %H:%M:%S"))  # 指定格式时间，例如2020-11-08 16:06:25
print(time.localtime())  # struct时间类型

now = time.time()  # 时间戳 -> struct -> strftime
struct_time = time.localtime(now)
print(time.strftime("%Y-%m-%d %H:%M:%S", struct_time))

d1 = time.time()
time.sleep(2)
d2 = time.time()
print(d2 - d1)
#计算了执行 time.sleep(2) 前后的时间差，time.sleep() 函数会让程序暂停指定的秒数。
# 在这里，程序暂停了 2 秒钟，然后再次获取了当前时间戳。两次时间戳的差值即为程序暂停的时间长度，单位是秒，因此会打印出 2


from datetime import date, datetime, timedelta
#这一行从 datetime 模块中导入了 date、datetime 和 timedelta 类。
# date 类用于处理日期，datetime 类用于处理日期和时间，而 timedelta 类用于表示时间间隔
now = datetime.now()  # 获取当前时间 2020-11-08 16:19:12.883672
print(date.strftime(now, "%Y-%m-%d %H:%M:%S"))  # 指定格式时间
print(date.today())  # 2020-11-08


from datetime import date, datetime, timedelta
import time
date_array = datetime.fromtimestamp(time.time())
#将当前的时间戳转换为一个 datetime 对象，并将其赋值给 date_array 变量
print(date_array)  # 2020-11-08 16:19:12.883672

from datetime import date, datetime, timedelta
yesterday = date.today() - timedelta(days=3)  # 获取昨天日期
print(yesterday)
tomorrow = date.today() + timedelta(days=1)  # 获取明天日期
print(tomorrow)


from urllib import request

res = request.urlopen("http://google.com")

print(res.getcode())
#返回的是这个响应的状态码
print(res.geturl())
#获取 HTTP 请求的实际 URL
print(res.headers)
#HTTP 响应的头部信息
print(res.read())
#读取 HTTP 响应的内容
print(res.readline())
#从 HTTP 响应中逐行读取内容
from urllib import request
url = ["http://www.baidu.com","http://www.jd.com"]
for i in url:
     res = request.urlopen(i)
     #打开当前迭代到的网址 i，并将返回的响应对象赋值给变量 res
     if res.getcode() == 200 or res.getcode == 301:
         print(f"{i} 访问成功")
     else:
         print(f"{i} 访问失败")
"""


from urllib import request
url = "http://www.ctnrs.com"
user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
#存储了一个模拟浏览器的用户代理字符串。用户代理字符串告诉服务器正在发送请求的客户端的类型和版本信息
header = {"User-Agent": user_agent}
#定义了一个字典 header，包含了 HTTP 请求的头部信息。这里设置了一个键值对，键是 "User-Agent"，值是之前定义的用户代理字符串
req = request.Request(url, headers=header)
#接受一个 URL 和一个可选的 headers 参数，用于设置请求的头部信息。
# 在这里，我们将之前定义的 URL 和头部信息传递给了这个函数，以创建一个带有自定义头部信息的请求对象
res = request.urlopen(req)
#使用 urllib.request.urlopen() 函数发送 HTTP 请求，并将返回的响应对象赋值给变量 res。
#这个函数接受一个 HTTP 请求对象作为参数，并返回一个响应对象，其中包含了服务器对请求的响应信息
print(res.getcode())























































