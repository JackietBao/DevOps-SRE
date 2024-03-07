#!/usr/bin/bash

back_dir=/backup

#1.判断/backup目录是否存在；不存在创建
if [ ! -d ${back_dir} ];then
	mkdir -p ${back_dir}
fi

#2.打包/etc/为一个tar包到/backup 目录   时间_主机名称_etc.tar.gz
cd / && tar czf ${back_dir}/$(date +%F_%s)_$(hostname)_etc.tar.gz etc

#3.使用find查询tar.gz的包；超过3天则删除；
find ${back_dir} -type f -name "*.tar.gz" -mtime +3 | xargs rm -f 


#4.放入计划任务中，每天运行
