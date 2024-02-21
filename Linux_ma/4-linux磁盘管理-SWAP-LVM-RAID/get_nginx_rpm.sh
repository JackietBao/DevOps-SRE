#!/usr/bin/bash


# 获取软件包名称
web_site="http://nginx.org/packages/centos/7/x86_64/RPMS/"
pkg_name=$(curl -s ${web_site} | grep "<a" | awk -F '"' '{print $2}')


# 遍历包名称，依次赋予给i变量；
for i in ${pkg_name}
do
	wget -O /var/ftp/nginx/${i} ${web_site}/${i}
done
