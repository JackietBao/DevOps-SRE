# 云服务器上线QQ农场平台

## 1、先还原云服务器操作系统

![img](assets/云服务器上线QQ农场平台/1687416154868-148d34be-f326-450c-93fa-fd6eaa532fcd.png)

![img](assets/云服务器上线QQ农场平台/1687416201567-41591a25-181f-4fb6-aa3f-9b4f9c29d203.png)

![img](assets/云服务器上线QQ农场平台/1687416239774-baa3feac-a760-4466-8712-6c678224f8de.png)

![img](assets/云服务器上线QQ农场平台/1687416267652-16603d60-07af-4c21-bb1e-65136bcce616.png)

![img](assets/云服务器上线QQ农场平台/1687416296222-f280ed22-07da-4974-a798-f6d1dc44dfbb.png)

![img](assets/云服务器上线QQ农场平台/1687416336910-895740d6-b6f8-46a4-941c-d407b615b20b.png)

![img](assets/云服务器上线QQ农场平台/1687416354899-0e4d13c6-862f-44b4-8d51-dd85a6dde356.png)

![img](assets/云服务器上线QQ农场平台/1687416389537-031999ab-2c79-4def-9d7a-9c6c09514126.png)



## 2、部署LNMP环境

```shell
LNMP = Linux操作系统  + Nginx网站服务  + MariaDB数据库  +  PHP后端服务
[root@host-1 ~]# yum -y install php php-fpm php-curl php-intl  php-mysql
软件包注释：
PHP-FPM是一个PHPFastCGI管理器
安装某些PHP源码时需求系统开启curl扩展
php-intl是ICU 库的一个包装器
php-mysql是连接数据库的模块
GD库是php处理图形的扩展库
```

## 3、安装MariaDB数据库

```shell
[root@host-1 ~]# yum -y install mariadb   mariadb-server
```

## 4、启动MariaDB数据库

```shell
[root@host-1 ~]# systemctl start mariadb
[root@host-1 ~]# systemctl status mariadb
```

![img](assets/云服务器上线QQ农场平台/1687416728306-ca7e755d-317b-4aa5-8284-393d1de59922.png)

```shell
[root@host-1 ~]# systemctl enable mariadb
[root@host-1 ~]# mysql
MariaDB [(none)]> create database farm;  #创建1个库，库名为farm
MariaDB [(none)]> quit									  #退出
```

## 5、安装Nginx网站服务

```shell
[root@host-1 ~]# yum -y install nginx
[root@host-1 ~]# systemctl start nginx
[root@host-1 ~]# systemctl status nginx
```

![img](assets/云服务器上线QQ农场平台/1687416824944-eb6f20e6-0ceb-41c3-ae97-8241481ddcaa.png)

```shell
[root@host-1 ~]# systemctl enable nginx
```

## 6、访问测试Nginx

![img](assets/云服务器上线QQ农场平台/1687416901964-422c8584-97f3-4a89-b42f-42cc19b218d3.png)

## 7、上线QQ农场项目

### 7.1 上传项目包

```shell
[root@host-1 ~]# yum -y install lrzsz  #下载上传工具
[root@host-1 ~]# rz
```

![img](assets/云服务器上线QQ农场平台/1687417045351-928e0ff1-51c0-4b48-b63e-319102eb0e02.png)

![img](assets/云服务器上线QQ农场平台/1687417099249-35f30760-4620-4dc0-9921-1dd49c07bf93.png)

![img](assets/云服务器上线QQ农场平台/1687417118190-e8c4b12f-aefd-483b-b609-2d91e3f68d7e.png)

### 7.2 设置网站发布目录

```shell
[root@host-1 ~]# mkdir /farm
```

### 7.3 解压项目包

```shell
[root@youngfit ~]# yum -y install unzip   #下载解压工具
[root@host-1 ~]# unzip farm-ucenter1.5.zip
```

![img](assets/云服务器上线QQ农场平台/1687417285686-9caef94f-cdaa-4365-85be-400a07a3cba1.png)

### 7.4 拷贝项目资源到默认发布目录/farm下

```shell
[root@host-1 ~]# cp -rf upload/* /farm/
```

### 7.5 修改Nginx配置

```shell
[root@host-1 ~]# rm -rf /etc/nginx/nginx.conf  #删除Nginx默认的主配置文件（因为这个文件内容过于繁琐）
[root@host-1 ~]# mv /etc/nginx/nginx.conf.default /etc/nginx/nginx.conf
[root@host-1 ~]# vim /etc/nginx/nginx.conf
```

![img](assets/云服务器上线QQ农场平台/1687417574701-52e217a1-3055-4059-b488-a52df3c056ea.png)

按小写的i进入编辑模式，修改为下图模样

![img](assets/云服务器上线QQ农场平台/1687417651506-c54906a6-db5f-4ec1-ac4e-1f8f3e2bf05a.png)

![img](assets/云服务器上线QQ农场平台/1687417754452-d63c565b-4e4f-49c2-8cc0-a0f7bd330824.png)

![img](assets/云服务器上线QQ农场平台/1687417786665-b3286207-e441-4004-999d-9c529bcbf2bb.png)

![img](assets/云服务器上线QQ农场平台/1687417828532-46210290-9a8e-4f01-aded-a91b2a2a3c28.png)

```shell
[root@host-1 ~]# systemctl restart nginx
```

![img](assets/云服务器上线QQ农场平台/1687417964408-5011a557-654c-4c8d-bd85-3bddff56ea1c.png)

```shell
[root@host-1 ~]# systemctl restart php-fpm
[root@host-1 ~]# chmod -R 777 /farm/
```

### 7.6 再次访问

![img](assets/云服务器上线QQ农场平台/1687418057554-28848884-922f-4e3f-a76b-cb7a6595130c.png)

![img](assets/云服务器上线QQ农场平台/1687418119811-292d41c2-df2c-4717-a418-2536fae46c57.png)

![img](assets/云服务器上线QQ农场平台/1687418141394-9c8270c8-4fee-423e-8bad-d0d16385cedb.png)

![img](assets/云服务器上线QQ农场平台/1687418159990-57c1a28b-fd2c-4263-91d2-c91756017922.png)