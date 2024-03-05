#!/usr/bin/bash

for i in {1..10}
do
	# 每次循环都会执行，所以每一次的密码都不一样
	sjmm=$(echo $RANDOM | md5sum  | cut -c 8-18)
	#$RANDOM来生成一个随机数，然后通过md5sum计算其MD5哈希值
	#最后使用cut命令提取哈希值的一部分作为密码。
	#这样生成的密码长度为11个字符。
	#-c选项表示按字符进行切割

	# 定义一个变量叫user，值是magedu_1..10
	user=magedu_$i
	# 添加用户
	useradd ${user}
	# 将随机密码给予当前的用户设定，并将密码存储到pass.txt文件
	echo "${sjmm}" | tee -a pass.txt |passwd --stdin $user
	#先用echo命令将生成的随机密码打印出来
	#然后通过管道|将密码同时输出到tee命令和passwd命令
	#tee -a pass.txt将密码追加写入到pass.txt文件中
	#passwd --stdin $user则将密码作为标准输入传递给
	#passwd命令来为刚创建的用户设置密码
	
	# 输出当前的用户是：当前的密码是： 通过追加重定向写入到一个文件中；
	echo "User: $user  Passwd: ${sjmm}" >>user_pass.txt
done
#整个循环会重复执行10次，每次生成一个不同的用户名和密码
#并将它们写入到pass.txt和user_pass.txt文件中