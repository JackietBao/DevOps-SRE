echo	$(mkpasswd	-l	10	-d	2	-c	2	-C	2	-s	4)	| tee	pass.txt |	passwd	--stdin	oldxu

#-l 10：指定密码的长度为10个字符
#-d 2：指定至少有2个数字。
#-c 2：指定至少有2个小写字母
#-C 2：指定至少有2个大写字母
#-s 4：指定至少有4个特殊字符

#tee 是一个常用的命令行工具，它的作用是从标准输入读取数据，然后将其复制到标准输出，并将数据写入一个或多个文件
#passwd --stdin 命令表示从标准输入读取密码

