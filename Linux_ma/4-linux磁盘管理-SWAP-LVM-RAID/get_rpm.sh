#!/usr/bin/bash


# 获取软件包名称
web_site="https://mirrors.tuna.tsinghua.edu.cn/zabbix/zabbix/4.2/rhel/7/x86_64/"
pkg_name=$(curl -s ${web_site} | grep "<a" | awk -F '"' '{print $4}')


# 遍历包名称，依次赋予给i变量；
for i in ${pkg_name}
do
	wget -O /var/ftp/zabbix/${i} ${web_site}/${i}
done
