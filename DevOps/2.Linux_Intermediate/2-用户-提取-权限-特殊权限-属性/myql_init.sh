#!/usr/bin/bash

mysql_tmp_file=/tmp/mysql.init
User=JackietBao

#1.初始化MySQL服务
touch "${mysql_tmp_file}"

#2.模拟用户删除文件
useradd jackietBao
su - "${User}" -c "rm -rf '${mysql_tmp_file}' > /dev/null"

#3.检查是否删除成功
if [ $? -eq 0 ]; then
    echo "${mysql_tmp_file} 文件被 ${User} 用户删除成功，该目录不是sbit, mysql初始化失败"
else
    echo "${mysql_tmp_file} 文件被 ${User} 用户删除失败，该目录是sbit,mysql初始化成功"
fi