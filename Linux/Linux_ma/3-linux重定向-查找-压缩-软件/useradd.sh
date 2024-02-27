#!/usr/bin/bash

for i in {1..10}
do
	# 每次循环都会执行，所以每一次的密码都不一样
	sjmm=$(echo $RANDOM | md5sum  | cut -c 8-18)

	# 定义一个变量叫user，值是magedu_1..10
	user=magedu_$i
	# 添加用户
	useradd ${user}
	# 将随机密码给予当前的用户设定，并将密码存储到pass.txt文件
	echo "${sjmm}" | tee -a pass.txt |passwd --stdin $user
	
	# 输出当前的用户是：当前的密码是： 通过追加重定向写入到一个文件中；
	echo "User: $user  Passwd: ${sjmm}" >>user_pass.txt
done
