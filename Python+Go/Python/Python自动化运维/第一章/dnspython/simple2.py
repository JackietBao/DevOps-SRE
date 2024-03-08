#!/usr/bin/env python
import dns.resolver

domain = raw_input('Please input an domain: ')

MX = dns.resolver.query(domain, 'MX')
for i in MX:
    print 'MX preference =', i.preference, 'mail exchanger =', i.exchange
#对于每个MX记录，打印出其优先级（preference）和邮件交换器（exchange）
#i.preference 表示优先级，i.exchange 表示邮件交换器的域名
