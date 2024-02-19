# 静态IP设置

```shell
vim /etc/sysconfig/network-scripts/ifcfg-ens33       #静态IP配置文件
 
 
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"               #修改项，设置静态IP(原来是dhcp)
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
NAME="ens33"
UUID="3150ff72-d779-438c-a097-5d88b13f51a1"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="192.168.16.5"            #新增项，添加静态IP地址  
NETMASK="255.255.255.0"          #新增项，添加子网掩码
GATEWAY="192.168.16.2"           #新增项，添加网关地址
DNS1="8.8.8.8"                   #新增项，添加DNS服务器地址(谷歌的，网速慢但是安全)
DNS2="114.114.114.114"           #新增项，添加DSNS服务器地址(电信联通移动通用的DNS，快好稳安全)
=========================================================================================
systemctl  restart  network      #重启网卡
```

#### 如果一个DNS访问不了

#### 就会去你设置的另一个DNS访问

#### 所以我们可以设置多个DNS，无所谓的
