#!/usr/bin/python
# -*- coding: utf-8 -*-

import dns.resolver
import os
import httplib

iplist=[]    #用来存储解析域名得到的IP地址列表

appdomain="www.google.com.hk"    #存储需要检查的业务域名

def get_iplist(domain=""):    #域名解析函数，解析成功IP将追加到iplist
    #domain是一个参数名称
    try:
        A = dns.resolver.query(domain, 'A')    #解析A记录类型
    except Exception,e: #e通常用于获取异常的具体信息，比如错误消息等
        print "dns resolver error:"+str(e) #e是except语句捕获到的异常对象
        #调用str(e)，可以将异常对象转换成人类可读的字符串形式
        return
    for i in A.response.answer:
    #A是通过dns.resolver.query(domain, 'A')得到的结果对象，代表对指定域名进行A记录查询的结果
    #A.response是查询结果中包含DNS响应信息的部分
    #A.response.answer包含了DNS查询响应的答案部分，通常包括了一个或多个解析记录（如A记录、AAAA记录等）
    #，每个记录对应域名的一个IP地址
        for j in i.items:
        #每一个i代表一个解析记录，可能是一个A记录，包含了域名对应的一个或多个IP地址
        #i.items是这个解析记录中所有条目的集合，每个条目代表一个具体的解析结果，如一个IP地址
            iplist.append(j.address)    #追加到iplist
            #对于每个条目j，j.address是这个条目指向的IP地址
    return True

def checkip(ip):
    checkurl=ip+":80" #指定的 IP 地址后面拼接字符串 ":80" 来形成一个完整的URL地址
    getcontent=""
    httplib.socket.setdefaulttimeout(5)    #定义http连接超时时间(5秒)
    conn=httplib.HTTPConnection(checkurl)    #创建http连接对象

    try:
        conn.request("GET", "/",headers = {"Host": appdomain})  #发起URL请求，添加host主机头
        #指定了请求的资源路径，这里是根路径/，代表请求的是网站的首页
        #指定了要发送的HTTP头部
        r=conn.getresponse()
        getcontent =r.read(15)   #获取URL页面前15个字符，以便做可用性校验
    finally:
        if getcontent=="<!doctype html>":  #监控URL页的内容一般是事先定义好，比如“HTTP200”等
            print ip+" [OK]"
        else:
            print ip+" [Error]"    #此处可放告警程序，可以是邮件、短信通知

if __name__=="__main__":
#__name__:代表当前Python脚本的名称
#Python文件被直接运行时，__name__ 的值会被设置为 "__main__"
    if get_iplist(appdomain) and len(iplist)>0:    #条件：域名解析正确且至少要返回一个IP
        for ip in iplist:
            checkip(ip)
    else:
        print "dns resolver error."
