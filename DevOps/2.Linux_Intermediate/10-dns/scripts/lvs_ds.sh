#!/usr/bin/bash
VIP=172.16.1.100
RS1=172.16.1.5
RS2=172.16.1.6
PORT=80
DEV=eth1:1

case $1 in
    start)
	cat  >/etc/sysconfig/network-scripts/ifcfg-${DEV} <<-EOF
	TYPE=Ethernet
	BOOTPROTO=none
	DEFROUTE=yes
	NAME=${DEV}
	DEVICE=${DEV}
	ONBOOT=yes
	IPADDR=${VIP}
	PREFIX=24
	EOF

	# 启动网卡
	ifup ${DEV}

	# 配置LVS规则
	ipvsadm -C
	ipvsadm -A -t ${VIP}:${PORT} -s rr
	ipvsadm -a -t ${VIP}:${PORT} -r ${RS1} -g
	ipvsadm -a -t ${VIP}:${PORT} -r ${RS2} -g
	;;

	stop)
	    ifdown ${DEV}
	    rm -f /etc/sysconfig/network-scripts/ifcfg-${DEV}
	    ipvsadm -C
	   ;;
	*)
		echo "Usage: sh $0 { start | stop }"
	;;
esac
