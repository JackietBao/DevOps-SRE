#!/usr/bin/env python
import dns.resolver

domain = raw_input('Please input an domain: ')
ns = dns.resolver.query(domain, 'NS')
for i in ns.response.answer:
#使用 for 循环遍历查询结果中的每个回答（answer）部分
     for j in i.items:
#在每个回答中，使用 for 循环遍历其中的每个项目（item）
          print j.to_text()
#打印每个项目的文本表示。这里的 j 是一个 DNS 记录对象，.to_text() 
#方法将该记录对象转换为其文本表示形式，并将其打印出来