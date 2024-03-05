#!/usr/bin/bash


# 获取软件包名称
web_site="http://nginx.org/packages/centos/7/x86_64/RPMS/"
pkg_name=$(curl -s ${web_site} | grep "<a" | awk -F '"' '{print $2}')
#首先使用curl命令下载指定网站的内容
#通过管道传递给grep命令来查找包含<a的行
#再使用awk命令以双引号作为分隔符，提取出链接地址
#提取出的链接地址将保存在变量pkg_name中

# 遍历包名称，依次赋予给i变量；
for i in ${pkg_name}
do
	wget -O /var/ftp/nginx/${i} ${web_site}/${i}
	#在每次循环中，使用wget命令下载指定链接地址的软件包
	#并将其保存到/var/ftp/nginx/目录下，文件名为${i}（即当前链接地址中的文件名部分）
done
