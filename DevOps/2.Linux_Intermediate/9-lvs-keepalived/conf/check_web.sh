#!/bin/sh
nginxpid=$(pidof nginx | wc -l)
#1.判断Nginx是否存活,如果不存活则尝试启动Nginx
if [ $nginxpid -eq 0 ];then
    systemctl start nginx
    sleep 2
    #2.等待2秒后再次获取一次Nginx状态
	nginxpid=$(pidof nginx | wc -l)
    #3.再次进行判断, 如Nginx还不存活则停止Keepalived,让地址进行漂移,并退出脚本
    if [ $nginxpid -eq 0 ];then
        systemctl stop keepalived
	pkill keepalived
   fi
fi
