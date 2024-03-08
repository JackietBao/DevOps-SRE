#!/usr/bin/env python
#在Unix系统中用来在PATH中查找python的常用命令
import dns.resolver

domain = raw_input('Please input an domain: ')

A = dns.resolver.query(domain, 'A')
#通过 dns.resolver.query() 函数进行域名查询，查询类型为 'A'，即 IPv4 地址记录
for i in A.response.answer:
#遍历查询结果 A 中的每一个响应记录。在这个循环中，i 是每个 DNS 响应的一部分
    for j in i.items:
#在每个响应记录中，这行再次遍历每个响应记录中的项目
        print j.address
#打印每个项目的地址。在 A 记录中，我们关注的是地址字段，因此这行代码输出 IPv4 地址。