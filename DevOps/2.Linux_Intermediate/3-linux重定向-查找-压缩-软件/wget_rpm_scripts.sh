#!/usr/bin/bash

get_zabbix_rpm_url=https://mirrors.aliyun.com/zabbix/zabbix/5.0/rhel/8/x86_64/
rpm_name=$(curl -s ${get_zabbix_rpm_url} | grep "^<a" | awk -F '"' '{print $2}')
rpm_dir=/var/ftp/ops

for name in ${rpm_name}
do
    if [ ! -d ${rpm_dir} ];then
        mkdir -p ${rpm_dir}
    fi

    wget -O ${rpm_dir}/${name} ${get_zabbix_rpm_url}${name}
done

