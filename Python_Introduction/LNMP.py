# -*- coding:UTF-8 -*-
'''
部署LNMP业务服务环境
Created on 2024年2月22日
@author: GeekyJack
'''
 
from fabric.colors import *
from fabric.api import *
 
env.user = 'root'
#定义业务角色分组
env.roledefs = {
    'webservers':['192.168.209.121', '192.168.209.122'],
    'dbservers':['192.168.209.123']
}
 
env.passwords = {
    'root@192.168.209.121:22' : "密码1",
    'root@192.168.209.122:22' : "密码2",
    'root@192.168.209.123:22' : "密码3",
}
 
@roles('webservers') #webtask任务函数引用'webservers'角色修饰符
def webtask():  #部署nginx php php-fpm环境
    print yellow("Install nginx php php-fpm...")
    with settings(warn_only = True):
        run("yum -y install nginx")
        run("yum -y install php-fpm php-mysql php-mbstring php-xml php-mcrypt php-gd")
        run("chkconfig --levels 235 php-fpm on")
        run("chkconfig --levels 235 nginx on")
        
@roles('dbservers') #dbatsk任务函数引用dbservers角色修饰符
def dbtask():   #部署mysql环境
    print yellow("Install MySQL...")
    with settings(warn_only = True):
        run("yum -y install mysql mysql-server")
        run("chkconfig --levels 235 mysqld on")
        
@roles('webservers', 'dbservers')   #publictask任务函数同时引用两个角色修饰符
def publictask():   #部署公共类环境，如epel、ntp等
    print yellow("Install epel ntp...")
    with settings(warn_only = True):
        run("rpm -Uvh http://dl.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm")
        run("yum -y install ntp")
        
def deploy():
    execute(publictask)
    execute(webtask)
    execute(dbtask)