#!/usr/bin/bash

VIP=172.16.1.100
DEV=lo:0

case $1 in
    start)
	echo "1" >/proc/sys/net/ipv4/conf/all/arp_ignore
	echo "1" >/proc/sys/net/ipv4/conf/default/arp_ignore
	echo "1" >/proc/sys/net/ipv4/conf/lo/arp_ignore

	echo "2" >/proc/sys/net/ipv4/conf/all/arp_announce
	echo "2" >/proc/sys/net/ipv4/conf/default/arp_announce
	echo "2" >/proc/sys/net/ipv4/conf/lo/arp_announce


	cat  >/etc/sysconfig/network-scripts/ifcfg-${DEV} <<-EOF
	DEVICE=lo:0
	IPADDR=${VIP}
	NETMASK=255.0.0.0
	ONBOOT=yes
	NAME=loopback
	EOF
	
	ifup ${DEV}	# 启动网卡
	systemctl start nginx
    ;;
    stop)
    echo "0" >/proc/sys/net/ipv4/conf/all/arp_ignore
    echo "0" >/proc/sys/net/ipv4/conf/default/arp_ignore
    echo "0" >/proc/sys/net/ipv4/conf/lo/arp_ignore
    
    echo "0" >/proc/sys/net/ipv4/conf/all/arp_announce
    echo "0" >/proc/sys/net/ipv4/conf/default/arp_announce
    echo "0" >/proc/sys/net/ipv4/conf/lo/arp_announce

        ifdown ${DEV}  # 停止网卡
        rm -f /etc/sysconfig/network-scripts/ifcfg-${DEV}
        systemctl stop nginx
        ;;
    *)
        echo "Usage: sh $0 { start | stop }"
esac
