### 1.初始化配置（所有节点）

#关闭selinux和防火墙

```shell
[root@controller ~]# systemctl stop firewalld.service
[root@controller ~]# systemctl disable firewalld.service 
[root@controller ~]# sed -i 's/SELINUX=enforcing/SELINUXT=disabled/g' /etc/selinux/config
```

#修改网卡、静态ip、域名解析

```shell
[root@controller ~]# cat /etc/network/interfaces
# The provider network interface
auto INTERFACE_NAME
iface INTERFACE_NAME inet manual
up ip link set dev $IFACE up
down ip link set dev $IFACE down

[root@controller ~]# vim /etc/sysconfig/network-scripts/ifcfg-INTERFACE_NAME
DEVICE=INTERFACE_NAME
TYPE=Ethernet
ONBOOT="yes"
BOOTPROTO="none"

[root@controller ~]# cat /etc/hosts
# controller
192.168.71.177       controller

# compute1
192.168.71.175       compute1


```





### 2.配置时间同步服务chrony(所有节点)

```shell
[root@controller ~]# yum -y install chrony
[root@controller ~]# vim /etc/chrony.conf
server ntp3.aliyun.com iburst  #更换阿里云时间点（只留这一段）
allow all           #允许连接点
local stratum 10    #允许最大连接数

[root@controller ~]#systemctl restart chronyd
[root@controller ~]#chronyc sources
```

### 3.配置openstack版本源  （此版本为train版本）所有节点centos7版本

```shell
[root@controller ~]# yum install centos-release-openstack-train -y   (千万别用epel源，否则会冲突)
[root@controller ~]# yum install python-openstackclient openstack-selinux -y
```

### 4.安装数据库（controller节点）

```shell
[root@controller ~]# yum install mariadb mariadb-server python2-PyMySQL -y

[root@controller ~]# cat /etc/my.cnf.d/openstack.cnf 
[mysqld]
bind-address = 192.168.71.177（换成本机ip）

default-storage-engine = innodb
innodb_file_per_table = on
max_connections = 4096
collation-server = utf8_general_ci
character-set-server = utf8

[root@controller ~]# systemctl enable mariadb.service
[root@controller ~]# systemctl start mariadb.service
[root@controller ~]# mysql_secure_installation（初始数据库）
>如果新机器则回车，不是输入root密码：123
>Disallow root login remotely? (Press y|Y for Yes, any other key for No) : n
此选项为n，其余为y即可
```

### 5.安装消息队列

```shell
[root@controller ~]# yum install rabbitmq-server -y
[root@controller ~]# systemctl enable rabbitmq-server.service
[root@controller ~]# systemctl start rabbitmq-server.service
[root@controller ~]# rabbitmqctl add_user openstack openstack123(设置密码，此密码为openstack123)
[root@controller ~]# rabbitmqctl set_permissions openstack ".*" ".*" ".*"   （开放权限最大化）

[root@controller ~]# rabbitmqctl list_users #查看创建的用户
[root@controller ~]# rabbitmq-plugins list #查看需要启动的服务
[root@controller ~]# rabbitmq-plugins enable rabbitmq_management rabbitmq_management_agent   #启动图形化服务

访问：http://192.168.71.177:15672/  用户：guest 密码：guest
```

### 6.配置消息缓存

```shell
[root@controller ~]# yum install memcached python-memcached -y
[root@controller ~]# vim  /etc/sysconfig/memcached
PORT="11211"
USER="memcached"
MAXCONN="1024"
CACHESIZE="1024"
OPTIONS="-l 127.0.0.1,::1,controller"           #添加控制节点的域名或者ip

[root@controller ~]# systemctl enable memcached.service
[root@controller ~]# systemctl start memcached.service
```

### 7.etcd目前先不装

### 8.选择版本安装[Install OpenStack services](https://docs.openstack.org/install-guide/openstack-services.html)（选的版本为train）keystone>控制节点

```shell
[root@controller ~]# mysql -u root -p
[root@controller ~]# MariaDB [(none)]> CREATE DATABASE keystone;
[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'localhost' \
IDENTIFIED BY 'keystone123';
[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'%' \
IDENTIFIED BY 'keystone123';

#安装和配置组件：
[root@controller ~]# yum -y install openstack-keystone httpd mod_wsgi    >安装keystone服务
[root@controller ~]# vim /etc/keystone/keystone.conf

[database]
connection = mysql+pymysql://keystone:keystone123@controller/keystone

[token]
provider = fernet

#同步数据库
[root@controller ~]#  su -s /bin/sh -c "keystone-manage db_sync" keystone  

#创建令牌
[root@controller ~]# keystone-manage fernet_setup --keystone-user keystone --keystone-group keystone
[root@controller ~]# keystone-manage credential_setup --keystone-user keystone --keystone-group keystone

#引导身份服务:
[root@controller ~]# keystone-manage bootstrap --bootstrap-password admin \
  --bootstrap-admin-url http://controller:5000/v3/ \
  --bootstrap-internal-url http://controller:5000/v3/ \
  --bootstrap-public-url http://controller:5000/v3/ \
  --bootstrap-region-id RegionOne
  
  #controller此处可用域名或ip
```

### 9.配置keystone[¶](https://docs.openstack.org/keystone/train/install/keystone-install-rdo.html#configure-the-apache-http-server)

```shell
[root@controller ~]# vim /etc/httpd/conf/httpd.conf
ServerName controller:80      #在95行

创建到的链接/usr/share/keystone/wsgi-keystone.conf文件:
[root@controller ~]# ln -s /usr/share/keystone/wsgi-keystone.conf /etc/httpd/conf.d/

[root@controller ~]# systemctl enable httpd.service;systemctl start httpd.service

通过设置适当的环境变量来配置admin管理帐户:
[root@controller ~]# vim admin.sh
 #!/bin/bash
 export OS_USERNAME=admin
 export OS_PASSWORD=admin          #此处admin与引导身份服务步骤中的--bootstrap-password admin一致
 export OS_PROJECT_NAME=admin
 export OS_USER_DOMAIN_NAME=Default
 export OS_PROJECT_DOMAIN_NAME=Default
 export OS_AUTH_URL=http://controller:5000/v3
 export OS_IDENTITY_API_VERSION=3
 [root@controller ~]# source admin.sh
 [root@controller ~]# openstack endpoint list
```

### 10.创建域、项目、用户和角色

```shell
创建新域的方法是；
[root@controller ~]# openstack domain create --description "An Example Domain" example

创建service项目:
[root@controller ~]# openstack project create --domain default  --description "Service Project" service

创建myproject项目:
[root@controller ~]# openstack project create --domain default  --description "Demo Project" myproject

创建myuser用户:
[root@controller ~]# openstack user create --domain default  --password-prompt myuser    >密码为myuser

创建myrole角色:
[root@controller ~]# openstack role create myrole

添加myrole角色到myproject项目和myuser用户:
[root@controller ~]# openstack role add --project myproject --user myuser myrole

#取消设置临时OS_AUTH_URL和OS_PASSWORD环境变量:
[root@controller ~]# unset OS_AUTH_URL OS_PASSWORD

#作为admin用户，请求验证令牌（输入两次admin密码即可）:
[root@controller ~]# openstack --os-auth-url http://controller:5000/v3 --os-project-domain-name Default --os-user-domain-name Default --os-project-name admin --os-username admin token issue

#作为myuser在上一部分中创建的用户请求身份验证令牌:
[root@controller ~]# openstack --os-auth-url http://controller:5000/v3 --os-project-domain-name Default --os-user-domain-name Default --os-project-name myproject --os-username myuser token issue


#为创建客户端环境脚本admin和demo项目和用户
[root@controller ~]# vim admin.sh
 #!/bin/bash
 export OS_PROJECT_DOMAIN_NAME=Default
 export OS_USER_DOMAIN_NAME=Default
 export OS_PROJECT_NAME=admin
 export OS_USERNAME=admin
 export OS_PASSWORD=admin
 export OS_AUTH_URL=http://controller:5000/v3
 export OS_IDENTITY_API_VERSION=3
 export OS_IMAGE_API_VERSION=2


[root@controller ~]# vim myuser.sh
 #!/bin/bash
 export OS_PROJECT_DOMAIN_NAME=Default
 export OS_USER_DOMAIN_NAME=Default
 export OS_PROJECT_NAME=myproject
 export OS_USERNAME=myuser
 export OS_PASSWORD=myuser
 export OS_AUTH_URL=http://controller:5000/v3
 export OS_IDENTITY_API_VERSION=3
 export OS_IMAGE_API_VERSION=2
 
[root@controller ~]# source admin.sh 
[root@controller ~]# openstack token issue
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Field      | Value                                                                                                                                                                                   |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| expires    | 2023-05-31T03:27:27+0000                                                                                                                                                                |
| id         | gAAAAABkdrCPd5CrEjqAPgM3-W7hITXKyh2cwaRjD8J0-P_HqFSumkMauQqhNu-593pFZdcSUgVBeo79QRkL_mAZD-t_7DvJXm0ujwaHRpRXNuYBCRxwpFAdAihQSbLXAR-QN62J0hLAw37d-AgraSd9IDr3OV6O5FILRR7gyJazGVI7yU2gupU |
| project_id | 5a38e5e74b1744af9b4b2ea2e3f2648d                                                                                                                                                        |
| user_id    | 11cc825570aa48bb89c1c173c5cbffe0                                                                                                                                                        |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
------------------------------------------------------------+
[root@controller ~]# source myuser.sh 
[root@controller ~]# openstack token issue
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Field      | Value                                                                                                                                                                                   |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| expires    | 2023-05-31T03:28:21+0000                                                                                                                                                                |
| id         | gAAAAABkdrDFP-fMVpCgd0ap59TjZ7Qozz2C-t0H274MMNVYj29U6O0D7s_8BncDlD3-XudgI5shhtadNMHnQevKpj-B4HKqfyRFi0E1AoerIpaQ2JhF36r9Q-MpUaPvpBFxBqY7mYctVC_WH2h5Mbc8ySvhrHKLSLzxANpJcrJmkkQ8EkhzVx0 |
| project_id | 0f46da77dc614585b0d1a0d334669640                                                                                                                                                        |
| user_id    | 6de93efdc6a24e8aa947f5afb54a641a                                                                                                                                                        |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+

```

### 11.安装glance镜像服务 [glance installation for Train](https://docs.openstack.org/glance/train/install/) (Red Hat controller节点)

```shell
[root@controller ~]# mysql -u root -p123

[root@controller ~]# MariaDB [(none)]> CREATE DATABASE glance;

[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON glance.* TO 'glance'@'localhost' \
  IDENTIFIED BY 'glance123';  

[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON glance.* TO 'glance'@'%' \
  IDENTIFIED BY 'glance123';

#要创建服务凭据,创建glance用户:
[root@controller ~]# openstack user create --domain default --password-prompt glance  密码为glance

#添加admin角色到glance用户和service项目:
[root@controller ~]# openstack role add --project service --user glance admin

#创建glance服务实体:
[root@controller ~]# openstack service create --name glance --description "OpenStack Image" image
 
#创建影像服务API端点:
公网
[root@controller ~]# openstack endpoint create --region RegionOne image public http://controller:9292
内网
[root@controller ~]# openstack endpoint create --region RegionOne image internal http://controller:9292
环境网
[root@controller ~]# openstack endpoint create --region RegionOne image admin http://controller:9292

[root@controller ~]# openstack endpoint list
+----------------------------------+-----------+--------------+--------------+---------+-----------+----------------------------+
| ID                               | Region    | Service Name | Service Type | Enabled | Interface | URL                        |
+----------------------------------+-----------+--------------+--------------+---------+-----------+----------------------------+
| 23ecd515802240faaf7b9a38ee8e45b4 | RegionOne | glance       | image        | True    | internal  | http://controller:9292     |
| 56bfea953f4e40dda989a364bc3ff04f | RegionOne | glance       | image        | True    | admin     | http://controller:9292     |
| 6d4986a67851489b87cad2e3fb98678d | RegionOne | keystone     | identity     | True    | admin     | http://controller:5000/v3/ |
| 712323c75202486182efe76ee61fc620 | RegionOne | glance       | image        | True    | public    | http://controller:9292     |
| 88a87d1399a24ac086255d429ec8a205 | RegionOne | keystone     | identity     | True    | public    | http://controller:5000/v3/ |
| a41509b2a3734ee3ba30d98373e44fb3 | RegionOne | keystone     | identity     | True    | internal  | http://controller:5000/v3/ |
+----------------------------------+-----------+--------------+--------------+---------+-----------+----------------------------+


```



#### >安装glance服务：

```shell
[root@controller ~]# yum install openstack-glance -y
[root@controller ~]# vim /etc/glance/glance-api.conf
[database]
connection = mysql+pymysql://glance:glance123@controller/glance     #更换密码为：glance123  glance数据库用户密码
#替换GLANCE_DBPASS使用您为影像服务数据库选择的密码。

[keystone_authtoken]
www_authenticate_uri = http://controller:5000
auth_url = http://controller:5000
memcached_servers = controller:11211
auth_type = password
project_domain_name = Default
user_domain_name = Default
project_name = service
username = glance
password = glance      #openstack用户glance密码为：glance

[paste_deploy]
flavor = keystone     
#选择keystone认证方式，openstack配置文件不能有中文，注释的也不行

[glance_store]
stores = file,http
default_store = file
filesystem_store_datadir = /var/lib/glance/images/
#在……里[glance_store]部分，配置图像文件的本地文件系统存储和位置


#填充影像服务同步数据库:
[root@controller ~]#  su -s /bin/sh -c "glance-manage db_sync" glance

#启动服务：
[root@controller ~]# systemctl enable openstack-glance-api.service
[root@controller ~]# systemctl start openstack-glance-api.service

#glance日志：
[root@controller ~]# tail -f /var/log/glance/api.log

```



#### >验证网址：http://download.cirros-cloud.net  后缀为img

```shell
[root@controller ~]#  wget http://download.cirros-cloud.net/0.4.0/cirros-0.4.0-x86_64-disk.img

#上传镜像
[root@controller ~]#glance image-create --name "cirros" \
  --file cirros-0.4.0-x86_64-disk.img \
  --disk-format qcow2 --container-format bare \
  --visibility public


[root@controller ~]# openstack image list
+--------------------------------------+--------+--------+
| ID                                   | Name   | Status |
+--------------------------------------+--------+--------+
| c7e68626-710b-43c7-aafb-e56ce9fc81b9 | cirros | active |
+--------------------------------------+--------+--------+
```



### 12.安装placement     [placement installation for Train](https://docs.openstack.org/placement/train/install/)（统一资源，计算节点）

```shell
[root@controller ~]# mysql -u root -p

#创建授权库
[root@controller ~]#MariaDB [(none)]> CREATE DATABASE placement;

[root@controller ~]#MariaDB [(none)]> GRANT ALL PRIVILEGES ON placement.* TO 'placement'@'localhost' \
  IDENTIFIED BY 'placement123';
MariaDB [(none)]> GRANT ALL PRIVILEGES ON placement.* TO 'placement'@'%' \
  IDENTIFIED BY 'placement123';


[root@controller ~]# openstack user create --domain default --password-prompt placement #设置密码为123

#使用admin角色将Placement用户添加到服务项目:
[root@controller ~]# openstack role add --project service --user placement admin

#在服务目录中创建放置API条目:
[root@controller ~]# openstack service create --name placement \
  --description "Placement API" placement

#创建位置API服务端点:
[root@controller ~]# openstack endpoint create --region RegionOne \
  placement public http://controller:8778

[root@controller ~]# openstack endpoint create --region RegionOne \
  placement internal http://controller:8778

[root@controller ~]# openstack endpoint create --region RegionOne \
  placement admin http://controller:8778

[root@controller ~]#
```

#### >安装和配置组件

```shell
[root@controller ~]# yum install openstack-placement-api

[root@controller ~]# vim /etc/placement/placement.conf
[placement_database]
connection = mysql+pymysql://placement:placement123@controller/placement

[api]
auth_strategy = keystone

[keystone_authtoken]
auth_url = http://controller:5000/v3
memcached_servers = controller:11211
auth_type = password
project_domain_name = Default
user_domain_name = Default
project_name = service
username = placement
password = 123

[root@controller ~]# su -s /bin/sh -c "placement-manage db sync" placement

[root@controller ~]# systemctl restart httpd

#解决httpd版本bug问题  >=2.4版本的需要配置
[root@controller ~]#vim /etc/httpd/conf.d/00-nova-placement-api.conf
<Directory /usr/bin>
   <IfVersion >= 2.4>
      Require all granted
   </IfVersion>
   <IfVersion < 2.4>
      Order allow,deny
      Allow from all
   </IfVersion>
</Directory>

[root@controller ~]# systemctl restart httpd
```

#### >验证

```shell
[root@controller ~]# source admin.sh
[root@controller ~]# placement-status upgrade check
+----------------------------------+
| Upgrade Check Results            |
+----------------------------------+
| Check: Missing Root Provider IDs |
| Result: Success                  |
| Details: None                    |
+----------------------------------+
| Check: Incomplete Consumers      |
| Result: Success                  |
| Details: None                    |
+----------------------------------+

```



### 13安装nova服务

#### 控制节点安装https://docs.openstack.org/nova/train/install/controller-install-rdo.html

```shell
[root@controller ~]# mysql -u root -p123

#创库授权：
[root@controller ~]# MariaDB [(none)]> CREATE DATABASE nova_api;
[root@controller ~]# MariaDB [(none)]> CREATE DATABASE nova;
[root@controller ~]# MariaDB [(none)]> CREATE DATABASE nova_cell0;

[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON nova_api.* TO 'nova'@'%' \
  IDENTIFIED BY 'nova123';
[root@controller ~]#MariaDB [(none)]> GRANT ALL PRIVILEGES ON nova.* TO 'nova'@'%' \
  IDENTIFIED BY 'nova123';

[root@controller ~]#MariaDB [(none)]> GRANT ALL PRIVILEGES ON nova_cell0.* TO 'nova'@'%' \
  IDENTIFIED BY 'nova123';
  


#创建nova用户:
[root@controller ~]# openstack user create --domain default --password-prompt nova  密码为：nova

#添加admin角色到nova用户:
[root@controller ~]# openstack role add --project service --user nova admin

#创建nova服务实体:
[root@controller ~]# openstack service create --name nova \
  --description "OpenStack Compute" compute



#创建计算API服务端点:
[root@controller ~]# openstack endpoint create --region RegionOne \
  compute public http://controller:8774/v2.1
[root@controller ~]# openstack endpoint create --region RegionOne \
  compute internal http://controller:8774/v2.1
[root@controller ~]# openstack endpoint create --region RegionOne \
  compute admin http://controller:8774/v2.1

```

##### >安装和配置组件：

```shell
[root@controller ~]# yum install openstack-nova-api openstack-nova-conductor openstack-nova-novncproxy openstack-nova-scheduler -y
  
 openstack-nova-conductor：提供数据库连接
 openstack-nova-novncproxy：访问云主机
 openstack-nova-scheduler：调度统一资源
 
 
[root@controller ~]# vim /etc/nova/nova.conf
[DEFAULT]
enabled_apis = osapi_compute,metadata

[api_database]
connection = mysql+pymysql://nova:nova123@controller/nova_api

[database]
connection = mysql+pymysql://nova:nova123@controller/nova

[DEFAULT]
transport_url = rabbit://openstack:openstack123@controller:5672/    #此处为rabbit的密码为openstack123


[api]
auth_strategy = keystone

[keystone_authtoken]
www_authenticate_uri = http://controller:5000/
auth_url = http://controller:5000/
memcached_servers = controller:11211
auth_type = password
project_domain_name = Default
user_domain_name = Default
project_name = service
username = nova
password = nova

#[DEFAULT]
#my_ip = 192.168.71.177

[DEFAULT]
use_neutron = true
firewall_driver = nova.virt.firewall.NoopFirewallDriver

[vnc]
enabled = true
server_listen = 192.168.71.177
server_proxyclient_address = 192.168.71.177


[glance]
api_servers = http://controller:9292


[oslo_concurrency]
lock_path = /var/lib/nova/tmp


[placement]
region_name = RegionOne
project_domain_name = Default
project_name = service
auth_type = password
user_domain_name = Default
auth_url = http://controller:5000/v3
username = placement
password = 123


#同步数据库
[root@controller ~]# su -s /bin/sh -c "nova-manage api_db sync" nova
[root@controller ~]# su -s /bin/sh -c "nova-manage cell_v2 map_cell0" nova
[root@controller ~]# su -s /bin/sh -c "nova-manage cell_v2 create_cell --name=cell1 --verbose" nova
[root@controller ~]# su -s /bin/sh -c "nova-manage db sync" nova

#验证
[root@controller ~]# su -s /bin/sh -c "nova-manage cell_v2 list_cells" nova
+-------+--------------------------------------+------------------------------------------+-------------------------------------------------+----------+
|  名称 |                 UUID                 |              Transport URL               |                    数据库连接                   | Disabled |
+-------+--------------------------------------+------------------------------------------+-------------------------------------------------+----------+
| cell0 | 00000000-0000-0000-0000-000000000000 |                  none:/                  | mysql+pymysql://nova:****@controller/nova_cell0 |  False   |
| cell1 | 6289d2ac-75bd-41f6-ab2f-fd1e0e34d97e | rabbit://openstack:****@controller:5672/ |    mysql+pymysql://nova:****@controller/nova    |  False   |
+-------+--------------------------------------+------------------------------------------+-------------------------------------------------+----------+


#启动服务
[root@controller ~]# systemctl enable \
>     openstack-nova-api.service \
>     openstack-nova-scheduler.service \
>     openstack-nova-conductor.service \
>     openstack-nova-novncproxy.service

[root@controller ~]# systemctl start \
>     openstack-nova-api.service \
>     openstack-nova-scheduler.service \
>     openstack-nova-conductor.service \
>     openstack-nova-novncproxy.service

[root@controller ~]# tail -f /var/log/nova/nova*
```



#### 计算节点安装：

```shell
[root@compute1 ~]# yum install openstack-nova-compute -y

[root@compute1 ~]# vim /etc/nova/nova.conf

[DEFAULT]
enabled_apis = osapi_compute,metadata

[DEFAULT]
transport_url = rabbit://openstack:openstack123@controller

[api]
auth_strategy = keystone

[keystone_authtoken]
www_authenticate_uri = http://controller:5000/
auth_url = http://controller:5000/
memcached_servers = controller:11211
auth_type = password
project_domain_name = Default
user_domain_name = Default
project_name = service
username = nova
password = nova


#[DEFAULT]
#my_ip = 192.168.71.177    #控制节点ip

[DEFAULT]
use_neutron = true
firewall_driver = nova.virt.firewall.NoopFirewallDriver


[vnc]
enabled = true
server_listen = 0.0.0.0
server_proxyclient_address = 192.168.71.175
novncproxy_base_url = http://192.168.71.177:6080/vnc_auto.html

[glance]
api_servers = http://controller:9292

[oslo_concurrency]
lock_path = /var/lib/nova/tmp


[placement]
region_name = RegionOne
project_domain_name = Default
project_name = service
auth_type = password
user_domain_name = Default
auth_url = http://controller:5000/v3
username = placement
password = 123



#查看是否支持cpu虚拟化
[root@compute1 ~]# egrep -c '(vmx|svm)' /proc/cpuinfo


结果为0，则表示不支持，需要如下操作：
[root@compute1 ~]# vim /etc/nova/nova.conf

[libvirt]
virt_type = qemu


#启动服务
[root@compute1 ~]# systemctl enable libvirtd.service openstack-nova-compute.service
[root@compute1 ~]# systemctl start libvirtd.service openstack-nova-compute.service

[root@compute1 ~]# tail -f /var/log/nova/nova-compute.log
```



##### >控制节点验证：

```shell
[root@controller ~]# openstack compute service list --service nova-compute
+----+--------------+----------+------+---------+-------+----------------------------+
| ID | Binary       | Host     | Zone | Status  | State | Updated At                 |
+----+--------------+----------+------+---------+-------+----------------------------+
|  7 | nova-compute | compute1 | nova | enabled | up    | 2023-05-31T07:17:13.000000 |
+----+--------------+----------+------+---------+-------+----------------------------+


#主机发现：（每添加一台主机都要执行这条命令）
[root@controller ~]# su -s /bin/sh -c "nova-manage cell_v2 discover_hosts --verbose" nova
Found 2 cell mappings.
Skipping cell0 since it does not contain hosts.
Getting computes from cell 'cell1': 6289d2ac-75bd-41f6-ab2f-fd1e0e34d97e
Checking host mapping for compute host 'compute1': f7e75e97-a2a5-42a3-99b9-c23187240af2
Creating host mapping for compute host 'compute1': f7e75e97-a2a5-42a3-99b9-c23187240af2
Found 1 unmapped computes in cell: 6289d2ac-75bd-41f6-ab2f-fd1e0e34d97e


#配置节点控制主机发现
[root@controller ~]# vim /etc/nova/nova.conf
[scheduler]
discover_hosts_in_cells_interval = 300


#重启服务
[root@controller ~]# systemctl restart openstack-nova-api.service openstack-nova-scheduler.service openstack-nova-conductor.service openstack-nova-novncproxy.service



#验证：
[root@controller ~]#  openstack compute service list   列出服务组件以验证每个流程的成功启动和注册:
+----+----------------+------------+----------+---------+-------+----------------------------+
| ID | Binary         | Host       | Zone     | Status  | State | Updated At                 |
+----+----------------+------------+----------+---------+-------+----------------------------+
|  3 | nova-conductor | controller | internal | enabled | up    | 2023-05-31T07:33:06.000000 |
|  5 | nova-scheduler | controller | internal | enabled | up    | 2023-05-31T07:33:06.000000 |
|  7 | nova-compute   | compute1   | nova     | enabled | up    | 2023-05-31T07:33:13.000000 |
+----+----------------+------------+----------+---------+-------+----------------------------+

[root@controller ~]# openstack catalog list     列出身份服务中的API端点以验证与身份服务的连接:
+-----------+-----------+-----------------------------------------+
| Name      | Type      | Endpoints                               |
+-----------+-----------+-----------------------------------------+
| nova      | compute   | RegionOne                               |
|           |           |   public: http://controller:8774/v2.1   |
|           |           | RegionOne                               |
|           |           |   internal: http://controller:8774/v2.1 |
|           |           | RegionOne                               |
|           |           |   admin: http://controller:8774/v2.1    |
|           |           |                                         |
| keystone  | identity  | RegionOne                               |
|           |           |   admin: http://controller:5000/v3/     |
|           |           | RegionOne                               |
|           |           |   public: http://controller:5000/v3/    |
|           |           | RegionOne                               |
|           |           |   internal: http://controller:5000/v3/  |
|           |           |                                         |
| placement | placement | RegionOne                               |
|           |           |   public: http://controller:8778        |
|           |           | RegionOne                               |
|           |           |   admin: http://controller:8778         |
|           |           | RegionOne                               |
|           |           |   internal: http://controller:8778      |
|           |           |                                         |
| glance    | image     | RegionOne                               |
|           |           |   internal: http://controller:9292      |
|           |           | RegionOne                               |
|           |           |   admin: http://controller:9292         |
|           |           | RegionOne                               |
|           |           |   public: http://controller:9292        |
|           |           |                                         |
+-----------+-----------+-----------------------------------------+

[root@controller ~]# openstack image list  列出影像服务中的影像以验证与影像服务的连通性:

+--------------------------------------+--------+--------+
| ID                                   | Name   | Status |
+--------------------------------------+--------+--------+
| c7e68626-710b-43c7-aafb-e56ce9fc81b9 | cirros | active |
+--------------------------------------+--------+--------+

[root@controller ~]# nova-status upgrade check   检查单元和放置API是否成功工作，以及其他必要的先决条件是否到位:
+--------------------------------+
| Upgrade Check Results          |
+--------------------------------+
| Check: Cells v2                |
| Result: Success                |
| Details: None                  |
+--------------------------------+
| Check: Placement API           |
| Result: Success                |
| Details: None                  |
+--------------------------------+
| Check: Ironic Flavor Migration |
| Result: Success                |
| Details: None                  |
+--------------------------------+
| Check: Cinder API              |
| Result: Success                |
| Details: None                  |
+--------------------------------+


```





### 14.neutron网络组件部署：

##### >控制节点：

```shell
[root@controller ~]# mysql -u root -p123

#创建授权库
[root@controller ~]# MariaDB [(none)]> CREATE DATABASE neutron;
[root@controller ~]# MariaDB [(none)]> GRANT ALL PRIVILEGES ON neutron.* TO 'neutron'@'%' \
  IDENTIFIED BY 'neutron123';

#创建用户：
[root@controller ~]# openstack user create --domain default --password-prompt neutron   密码：neutron
[root@controller ~]# openstack role add --project service --user neutron admin
[root@controller ~]# openstack service create --name neutron --description "OpenStack Networking" network

[root@controller ~]# openstack endpoint create --region RegionOne network public http://controller:9696
[root@controller ~]# openstack endpoint create --region RegionOne network internal http://controller:9696
[root@controller ~]# openstack endpoint create --region RegionOne network admin http://controller:9696
```



##### >配置网络选项：

有两种网络模式根据实际选择：

选项1：https://docs.openstack.org/neutron/train/install/controller-install-option1-rdo.html

选项2：https://docs.openstack.org/neutron/train/install/controller-install-option2-rdo.html

选项1部署最简单的架构，该架构仅支持将实例连接到提供商(外部)网络。没有自助服务(私有)网络、路由器或浮动IP地址。只有`admin`或者其他特权用户可以管理提供商网络。

选项2用支持将实例附加到自助服务网络的第3层服务增强了选项1。这`demo`或者其他非特权用户可以管理自助服务网络，包括在自助服务网络和提供商网络之间提供连接的路由器。此外，浮动IP地址提供了从外部网络(如互联网)到使用自助服务网络的实例的连接。



我这里先使用的选项1

```shell
[root@controller ~]# yum -y install openstack-neutron openstack-neutron-ml2 openstack-neutron-linuxbridge ebtables

[root@controller ~]# vim /etc/neutron/neutron.conf
[database]
connection = mysql+pymysql://neutron:neutron123@controller/neutron


[DEFAULT]
core_plugin = ml2
service_plugins =


[DEFAULT]
transport_url = rabbit://openstack:openstack123@controller


[DEFAULT]
auth_strategy = keystone

[keystone_authtoken]
www_authenticate_uri = http://controller:5000
auth_url = http://controller:5000
memcached_servers = controller:11211
auth_type = password
project_domain_name = default
user_domain_name = default
project_name = service
username = neutron
password = neutron


[DEFAULT]
notify_nova_on_port_status_changes = true
notify_nova_on_port_data_changes = true

[nova]
auth_url = http://controller:5000
auth_type = password
project_domain_name = default
user_domain_name = default
region_name = RegionOne
project_name = service
username = nova
password = nova

#残缺文件地址：https://docs.openstack.org/ocata/config-reference/networking/samples/neutron.conf.html
[oslo_concurrency]
lock_path = /var/lib/neutron/tmp

```

ML2插件使用Linux桥接机制为实例构建第2层(桥接和交换)虚拟网络基础设施。

```shell
[root@controller ~]# vim /etc/neutron/plugins/ml2/ml2_conf.ini

#残缺文件地址将内容替换：https://docs.openstack.org/ocata/config-reference/networking/samples/ml2_conf.ini.html
[ml2]
type_drivers = flat,vlan    #启用平面和VLAN网络:
tenant_network_types =         #禁用自助服务网络:
mechanism_drivers = linuxbridge    #启用Linux桥机制:
extension_drivers = port_security    #启用端口安全扩展驱动程序:


[ml2_type_flat]
flat_networks = extnet        #将提供商虚拟网络配置为平面网络:

[securitygroup]
enable_ipset = true       #启用ipset以提高安全组规则的效率:
```



Linux桥接代理为实例构建第2层(桥接和交换)虚拟网络基础设施，并处理安全组。

```shell
[root@controller ~]# vim /etc/neutron/plugins/ml2/linuxbridge_agent.ini

#残缺文件地址将内容替换:https://docs.openstack.org/ocata/config-reference/networking/samples/linuxbridge_agent.ini.html

[linux_bridge]
physical_interface_mappings = extnet:ens33   #将提供商虚拟网络映射到提供商物理网络接口

[vxlan]
enable_vxlan = false

[securitygroup]
enable_security_group = true
firewall_driver = neutron.agent.linux.iptables_firewall.IptablesFirewallDriver

#配置内核参数
[root@controller ~]# vim /etc/sysctl.conf
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1

[root@controller ~]# modprobe br_netfilter      #加载内核模块
[root@controller ~]# sysctl -p
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
```

##### 配置DHCP代理[¶](https://docs.openstack.org/neutron/train/install/controller-install-option1-rdo.html#configure-the-dhcp-agent)

```shell
[root@controller ~]# vim /etc/neutron/dhcp_agent.ini
[DEFAULT]
interface_driver = linuxbridge
dhcp_driver = neutron.agent.linux.dhcp.Dnsmasq
enable_isolated_metadata = true
```

##### 配置元数据代理[¶](https://docs.openstack.org/neutron/train/install/controller-install-rdo.html#configure-the-metadata-agent)

```shell
[root@controller ~]# vim /etc/neutron/metadata_agent.ini
[DEFAULT]
nova_metadata_host = controller
metadata_proxy_shared_secret = METADATA_SECRET

[root@controller ~]# vim /etc/nova/nova.conf
[neutron]
# ...
auth_url = http://controller:5000
auth_type = password
project_domain_name = default
user_domain_name = default
region_name = RegionOne
project_name = service
username = neutron
password = neutron
service_metadata_proxy = true
metadata_proxy_shared_secret = METADATA_SECRET    #此值与上方对应，上方如更改此处记得更换 


[root@controller ~]#  ln -s /etc/neutron/plugins/ml2/ml2_conf.ini /etc/neutron/plugin.ini
#同步数据库
[root@controller ~]# su -s /bin/sh -c "neutron-db-manage --config-file /etc/neutron/neutron.conf \
  --config-file /etc/neutron/plugins/ml2/ml2_conf.ini upgrade head" neutron

#重新启动计算API服务: 
[root@controller ~]# systemctl restart openstack-nova-api.service

#启动网络服务，并将其配置为在系统启动时启动。
[root@controller ~]# systemctl enable neutron-server.service neutron-linuxbridge-agent.service neutron-dhcp-agent.service neutron-metadata-agent.service
[root@controller ~]# systemctl start neutron-server.service neutron-linuxbridge-agent.service neutron-dhcp-agent.service neutron-metadata-agent.service
```



##### >计算节点:

```shell
[root@compute1 ~]# yum install openstack-neutron-linuxbridge ebtables ipset -y
[root@compute1 ~]# vim /etc/neutron/neutron.conf
[DEFAULT]
transport_url = rabbit://openstack:openstack123@controller
auth_strategy = keystone

[keystone_authtoken]
www_authenticate_uri = http://controller:5000
auth_url = http://controller:5000
memcached_servers = controller:11211
auth_type = password
project_domain_name = default
user_domain_name = default
project_name = service
username = neutron
password = neutron

[oslo_concurrency]
lock_path = /var/lib/neutron/tmp
```

##### 配置网络选项[¶](https://docs.openstack.org/neutron/train/install/compute-install-rdo.html#configure-networking-options)

```shell
[root@compute1 ~]# vim /etc/neutron/plugins/ml2/linuxbridge_agent.ini
[linux_bridge]
physical_interface_mappings = extnet:ens33

[vxlan]
enable_vxlan = false

[securitygroup]
enable_security_group = true
firewall_driver = neutron.agent.linux.iptables_firewall.IptablesFirewallDriver

[root@compute1 ~]# vim /etc/sysctl.conf
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1

[root@compute1 ~]# modprobe br_netfilter

[root@compute1 ~]# sysctl -p
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1




```



##### 将计算服务配置为使用网络服务

```shell
[root@compute1 ~]# vim /etc/nova/nova.conf
[neutron]
auth_url = http://controller:5000
auth_type = password
project_domain_name = default
user_domain_name = default
region_name = RegionOne
project_name = service
username = neutron
password = neutron

[root@compute1 ~]# systemctl restart openstack-nova-compute.service


[root@compute1 ~]# systemctl enable neutron-linuxbridge-agent.service
[root@compute1 ~]# systemctl start neutron-linuxbridge-agent.service



#验证：
[root@controller ~]# openstack network agent list
+--------------------------------------+--------------------+------------+-------------------+-------+-------+---------------------------+
| ID                                   | Agent Type         | Host       | Availability Zone | Alive | State | Binary                    |
+--------------------------------------+--------------------+------------+-------------------+-------+-------+---------------------------+
| 2a188220-036f-4e92-88fe-934c52b3954f | Metadata agent     | controller | None              | :-)   | UP    | neutron-metadata-agent    |
| 33ffc3e3-ad82-47d0-8728-d841a93ada88 | DHCP agent         | controller | nova              | :-)   | UP    | neutron-dhcp-agent        |
| 99354cf8-758b-499d-bc87-d43380a2a510 | Linux bridge agent | compute1   | None              | :-)   | UP    | neutron-linuxbridge-agent |
| e6451885-43ea-4127-8ebd-e13b451e6f07 | Linux bridge agent | controller | None              | :-)   | UP    | neutron-linuxbridge-agent |
+--------------------------------------+--------------------+------------+-------------------+-------+-------+---------------------------+

```

![image-20230601181844442](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230601181844442.png)



#### 启动虚拟机进行尝试创建实例：

```shell
[root@controller ~]# openstack flavor create --id 0 --vcpus 1 --ram 64 --disk 1 m1.nano

[root@controller ~]#  ssh-keygen -q -N ""   一直回车即可

[root@controller ~]# openstack keypair create --public-key ~/.ssh/id_rsa.pub mykey

[root@controller ~]# openstack keypair list

[root@controller ~]# openstack security group rule create --proto icmp default

[root@controller ~]# openstack security group rule create --proto tcp --dst-port 22 default  允许安全外壳(SSH)访问
```

```shell
[root@controller ~]# openstack flavor list

[root@controller ~]# openstack image list       列出可用图像

创建网络：
[root@controller ~]#  openstack network create  --share --external --provider-physical-network extnet --provider-network-type flat provider

创建子网：
[root@controller ~]# openstack subnet create --network provider  --allocation-pool start=192.168.71.50,end=192.168.71.100 --dns-nameserver 114.114.114.114 --gateway 192.168.71.2  --subnet-range 192.168.71.0/24 provider  __子网名称

[root@controller ~]# openstack network list




[root@controller ~]# openstack server create --flavor m1.nano --image cirros --nic net-id=5f605576-c267-4c69-93a5-d5fd8672f15b --security-group default --key-name mykey vm-1

注释：类型  镜像名称  net网络id  默认安全组   vm-1：云主机名称

+-------------------------------------+-----------------------------------------------+
| Field                               | Value                                         |
+-------------------------------------+-----------------------------------------------+
| OS-DCF:diskConfig                   | MANUAL                                        |
| OS-EXT-AZ:availability_zone         |                                               |
| OS-EXT-SRV-ATTR:host                | None                                          |
| OS-EXT-SRV-ATTR:hypervisor_hostname | None                                          |
| OS-EXT-SRV-ATTR:instance_name       |                                               |
| OS-EXT-STS:power_state              | NOSTATE                                       |
| OS-EXT-STS:task_state               | scheduling                                    |
| OS-EXT-STS:vm_state                 | building                                      |
| OS-SRV-USG:launched_at              | None                                          |
| OS-SRV-USG:terminated_at            | None                                          |
| accessIPv4                          |                                               |
| accessIPv6                          |                                               |
| addresses                           |                                               |
| adminPass                           | MH9Zidiqv8tP                                  |
| config_drive                        |                                               |
| created                             | 2023-06-01T10:47:23Z                          |
| flavor                              | m1.nano (0)                                   |
| hostId                              |                                               |
| id                                  | d8e675d7-9c0d-4a06-b317-b09e451d74b5          |
| image                               | cirros (c7e68626-710b-43c7-aafb-e56ce9fc81b9) |
| key_name                            | mykey                                         |
| name                                | vm-1                                          |
| progress                            | 0                                             |
| project_id                          | 5a38e5e74b1744af9b4b2ea2e3f2648d              |
| properties                          |                                               |
| security_groups                     | name='7c322133-3f76-41f8-bd07-5ce50c0631db'   |
| status                              | BUILD                                         |
| updated                             | 2023-06-01T10:47:23Z                          |
| user_id                             | 11cc825570aa48bb89c1c173c5cbffe0              |
| volumes_attached                    |                                               |
+-------------------------------------+-----------------------------------------------+


检查实例的状态:
[root@controller ~]# openstack service list


获得虚拟网络计算(VNC)实例的会话URL，并从web浏览器访问它:
[root@controller ~]# openstack console url show vm-1
+-------+-------------------------------------------------------------------------------------------+
| Field | Value                                                                                     |
+-------+-------------------------------------------------------------------------------------------+
| type  | novnc                                                                                     |
| url   | http://controller:6080/vnc_auto.html?path=%3Ftoken%3Df4743b76-a778-4f7b-9321-cb67992ade4f |
+-------+-------------------------------------------------------------------------------------------+

换成本机ip进行访问：
http://192.168.71.177:6080/vnc_auto.html?path=%3Ftoken%3Df4743b76-a778-4f7b-9321-cb67992ade4f

```

这是一个bug，不代表成功奥！![image-20230601185421542](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230601185421542.png)

解决方案：（虚拟机都是在计算节点创建的）

```shell
[root@compute1 ~]# virsh capabilities

[root@compute1 ~]# vim /etc/nova/nova.conf 
hw_machine_type = x86_64=pc-i440fx-rhel7.2.0          #更改虚拟化类型
cpu_mode = host-passthrough                         #直接使用宿主机cpu，记得把汉字删掉


重启nova服务
[root@compute1 ~]# systemctl restart openstack-nova-*


在创建一个虚拟机：
[root@controller ~]# openstack server create --flavor m1.nano --image cirros --nic net-id=5f605576-c267-4c69-93a5-d5fd8672f15b --security-group default --key-name mykey vm-2

[root@controller ~]#  openstack console url show vm-2
+-------+-------------------------------------------------------------------------------------------+
| Field | Value                                                                                     |
+-------+-------------------------------------------------------------------------------------------+
| type  | novnc                                                                                     |
| url   | http://controller:6080/vnc_auto.html?path=%3Ftoken%3D105ad300-11a5-4826-af5f-681068507b8a |
+-------+-------------------------------------------------------------------------------------------+
http://192.168.71.177:6080/vnc_auto.html?path=%3Ftoken%3D105ad300-11a5-4826-af5f-681068507b8a
```

测试成功：用户：cirros    密码：gocubsgo

![image-20230601191018322](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230601191018322.png)

进行ping测试：

![image-20230601191540388](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230601191540388.png)

![image-20230601191552755](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230601191552755.png)

```shell
[root@controller ~]# ssh cirros@192.168.71.85
The authenticity of host '192.168.71.85 (192.168.71.85)' can't be established.
ECDSA key fingerprint is SHA256:sjFBOaZZuox/DsQedPr+p/cIEXCjc5PsIIZrgWV5xm4.
ECDSA key fingerprint is MD5:4b:96:b5:b0:68:87:f3:de:f4:2a:db:05:18:1b:93:53.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.71.85' (ECDSA) to the list of known hosts.
$ 
尝试远程连接也可以！
```





### 15.安装dashboard

```shell
[root@controller ~]# yum install openstack-dashboard -y
如若下载yum报错，可操作:
--2023-06-02 11:54:41-- wget http://mirrors.aliyun.com/repo/Centos-7.repo?spm=a2c6h.25603864.0.0.3d975969x9X4Sl
正在解析主机 mirrors.aliyun.com (mirrors.aliyun.com)... 失败：未知的名称或服务。
wget: 无法解析主机地址 “mirrors.aliyun.com”
[root@controller ~]# vim /etc/resolv.conf
nameserver 8.8.8.8 
nameserver 8.8.4.4
nameserver 223.5.5.5
nameserver 223.6.6.6
#（nameserver 223.5.5.5 和 nameserver 223.6.6.6选择其中一个添加即可）


[root@controller ~]# vim /etc/openstack-dashboard/local_settings
OPENSTACK_HOST = "controller"

#白名单
ALLOWED_HOSTS = ['one.example.com', 'two.example.com']
#注意，我将白名单改为了全部可以：ALLOWED_HOSTS = ['*']

#会话存储服务
SESSION_ENGINE = 'django.contrib.sessions.backends.cache'

CACHES = {
    'default': {
         'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
         'LOCATION': 'controller:11211',
    }
}



OPENSTACK_KEYSTONE_URL = "http://%s:5000/v3" % OPENSTACK_HOST


OPENSTACK_KEYSTONE_MULTIDOMAIN_SUPPORT = True    #启用对域的支持



#配置API版本:
OPENSTACK_API_VERSIONS = {
    "identity": 3,
    "image": 2,
    "volume": 3,
}



#安装ˌ使成形Default作为您通过仪表板创建的用户的默认域:
OPENSTACK_KEYSTONE_DEFAULT_DOMAIN = "Default"


#安装ˌ使成形user作为您通过仪表板创建的用户的默认角色:
OPENSTACK_KEYSTONE_DEFAULT_ROLE = "user"


#如果您选择了网络选项1，请禁用对第3层网络服务的支持:
OPENSTACK_NEUTRON_NETWORK = {
    ...
    'enable_router': False,
    'enable_quotas': False,
    'enable_distributed_router': False,
    'enable_ha_router': False,
    'enable_lb': False,
    'enable_firewall': False,
    'enable_vpn': False,
    'enable_fip_topology_check': False,
}


#配置时区:
TIME_ZONE = "Asia/Shanghai"

官方bug：！！！
WEBROOT = "/dashboard"

#将下面一行添加到如果不包括。
[root@controller ~]# vim /etc/httpd/conf.d/openstack-dashboard.conf
WSGIApplicationGroup %{GLOBAL}


[root@controller ~]# systemctl restart httpd.service memcached.service

[root@compute1 ~]# yum -y install bash-completion  tab键加强版
```

##### 访问：http://192.168.71.177/dashboard/auth/login/?next=/dashboard/

![image-20230602142741371](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230602142741371.png)

##### 测试实例：

![image-20230602145232649](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230602145232649.png)

![image-20230602152548437](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230602152548437.png)

配置下解析，否则控制台打不开，或者在控制,计算节点上nova.conf配置文件[vnc]模块更改：

![image-20230602153949801](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230602153949801.png)

测试成功：![image-20230602155229331](C:\Users\刁玉\AppData\Roaming\Typora\typora-user-images\image-20230602155229331.png)
