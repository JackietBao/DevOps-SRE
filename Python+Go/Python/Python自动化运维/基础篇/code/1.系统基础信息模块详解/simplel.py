#!/usr/bin/env python
#这是一个 shebang，用于告诉操作系统以哪个解释器来执行这个脚本。
#在这里，它告诉操作系统使用环境中的 Python 解释器来执行这个脚本

from IPy import IP
#导入了 IPy 模块中的 IP 类

ip_s = raw_input('Please input an IP or net-range: ')
#使用 raw_input 函数获取用户输入的 IP 地址或 IP 网段
ips = IP(ip_s)
#将用户输入的 IP 地址或 IP 网段传递给 IP 类来创建一个 IP 对象。


if len(ips) > 1:
#检查如果输入的是一个 IP 网段而不是单个 IP 地址
    print('net: %s' % ips.net())
    #打印 IP 网段的网络地址
    print('netmask: %s' % ips.netmask())
    #打印 IP 网段的子网掩码。
    print('broadcast: %s' % ips.broadcast())
    #打印 IP 网段的广播地址。
    print('reverse address: %s' % ips.reverseNames()[0])
    #打印 IP 地址或 IP 网段的反向地址（域名）
    #ips.reverseNames()[0] 表示获取该列表中的第一个元素，
    #即获取到的反向地址列表中的第一个反向地址（域名）
    print('subnet: %s' % len(ips))
    #打印 IP 网段中的子网数量
else:
    print('reverse address: %s' % ips.reverseNames()[0])
    #打印 IP 地址或 IP 网段的反向地址（域名）

print('hexadecimal: %s' % ips.strHex())
#打印 IP 地址或 IP 网段的十六进制表示
print('binary ip: %s' % ips.strBin())
#打印 IP 地址或 IP 网段的二进制表示
print('iptype: %s' % ips.iptype())
#打印 IP 地址或 IP 网段的类型（IPv4 或 IPv6）