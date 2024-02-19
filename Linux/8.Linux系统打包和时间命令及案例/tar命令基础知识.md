1、tar 打包压缩
1）【打包】 为什么要打包，压缩？
-zcvf 打包组合 zcf
z 压缩（gzip压缩）
c 创建
v 显示过程
f 针对文件
语法：
命令 选项
tar  zcf  【压缩包的名字.tar.gz】  【要打包的东西】
                  装东西的筐             苹果

2）【查看包里内容】
t 查看包里内容
tar tf 包名
[root@oldboy usr]# tar tf /tmp/local.tar.gz 

3）【解包】
组合zxvf，缩写xf
x 解压
cd /opt
tar zxvf etc.tar.gz  #解压到了当前目录
tar zxf etc.tar.gz 
tar xf etc.tar.gz 

-C 指定路径解压
[root@oldboy opt]# tar xf etc.tar.gz -C /home/oldboy/
[root@oldboy opt]# ls /home/oldboy/

-p 保持属性（zcfp）

-j 通过bzip2命令压缩或解压（非常少了）
tar jcvf 包名.tar.bz2   包名

两种解压方法：
[root@oldboy tmp]# tar jxvf a.tar.bz2 

[root@oldboy tmp]# tar xf a.tar.bz2 #gzip,bzip统一解压方法。

排除打包--exclude=
tar  zcvf /tmp/pai.tar.gz ./oldboy/ --exclude=file1
命令 参数   包名			目录     排除的文件

从文件中排除打包--exclude-from(-X)

建立排除的文件paichu.log，内容就是排除的文件名
cat >paichu.log<<EOF
file1
file5
EOF

[root@oldboy /]# cat paichu.log 
file1
file5
[root@oldboy /]# tar zcvfX /tmp/pai.tar.gz paichu.log ./oldboy/
./oldboy/
./oldboy/file2
./oldboy/file3
./oldboy/file4

tar zcvfX /tmp/pai.tar.gz       paichu.log         ./oldboy/
                             存储排除文件名的文件

-h  打包软链接
/etc/rc.local是软链接文件，指向真实路径/etc/rc.d/rc.local						 
默认打包的时候，只打包了软链接文件。
tar zcvf  /backup/rc.local_1.tar.gz /etc/rc.local

【而加上-h打包，可以打包软链接对应的真实文件】
[root@oldboy etc]# tar zcvfh  /backup/rc.local_1.tar.gz /etc/rc.local
tar: 从成员名中删除开头的“/”
/etc/rc.local
[root@oldboy etc]# cd /backup/
[root@oldboy backup]# tar xf rc.local_1.tar.gz 
[root@oldboy backup]# cat etc/rc.local 
#!/bin/bash
touch /var/lock/subsys/local

日期 时间命令 date,别和data搞混这是数据。
date -s "2030/5/14"          #修改日期
date -s "2030/5/14 23:45:12" #修改时间
clock -w #写到bios永久生效。
学习一下。。。
工作中服务器的时间是定时和互联网时间同步的。 
自己配置定时同步（设置搭建时间服务器） NTPD服务和chronyd

CentOS6中，默认使用ntpd时间服务
Centos7中，默认使用chrony时间服务

当下时间，特定格式时间显示
[root@oldboy ~]# date +%F
2030-05-15
[root@oldboy ~]# date +%Y
2030
[root@oldboy ~]# date +%m
05
[root@oldboy ~]# date +%d
15
[root@oldboy ~]# date +%Y-%m-%d
2030-05-15
[root@oldboy ~]# date +%H
00
[root@oldboy ~]# date +%M
07
[root@oldboy ~]# date +%S
31
[root@oldboy ~]# date +%H:%M:%S
00:07:51
[root@oldboy ~]# date +%Y-%m-%d\ %H:%M:%S
2030-05-15 00:08:16
[root@oldboy ~]# date +%Y-%m-%d %H:%M:%S  #空格要转义，否则报错。
date: 额外的操作数 "%H:%M:%S"
Try 'date --help' for more information.
[root@oldboy ~]# date +%F\ %T
2030-05-15 00:09:08





















































































