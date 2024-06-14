# 云服务器上线WordPress博客平台

## 1、部署LAMP架构

```shell
LAMP = Linux操作系统 + Apache网站服务 + MariaDB数据库 + PHP后端语言
```

## 2、安装Apache MariaDB PHP

```shell
[root@host-1 ~]# yum install httpd mariadb-server mariadb php php-mysql gd php-gd

httpd：就是apache的软件包名称
mariadb-server: 是Maridb服务安装包
maridb: 是连接数据库的客户端命令工具
php: php语言所依赖的基础环境
php-mysql: 能让php连接数据库
gd php-gd： 如果页面上有动态图，装上这两个软件包，能正常显示动态图效果
```

## 3、启动Apache

```shell
[root@host-1 ~]# systemctl start httpd   #启动apache
[root@host-1 ~]# systemctl status httpd  #查看apache的状态
```

![img](assets/云服务器上线WordPress博客平台/1687397307286-3c7eef8c-06fd-4aeb-97ed-7f2c5cecc881.png)

```shell
[root@host-1 ~]# systemctl enable httpd #将apache设置为开机自启
```

## 4、重启云服务器验证

```shell
[root@host-1 ~]# reboot   #重启当前机器
```

![img](assets/云服务器上线WordPress博客平台/1687397551673-21042848-769b-4429-bfbd-5af8aceab577.png)

## 5、启动MariaDB数据库

```shell
[root@host-1 ~]# systemctl start mariadb   #启动MariaDB数据库
[root@host-1 ~]# systemctl status mariadb   #查看MariaDB数据库的状态
```

![img](assets/云服务器上线WordPress博客平台/1687397686612-f306d14f-494e-4bb2-b34d-4ae2d38c0d01.png)

```shell
[root@host-1 ~]# systemctl enable mariadb   #将MariaDB数据库设置为开机自启
```

## 6、部署WordPress网站

### 6.1 准备数据库

```shell
[root@host-1 ~]# mysql
MariaDB [(none)]> create database wordpress;    #创建1个库，库名为wordpress
MariaDB [(none)]> \q   #退出
```

### 6.2 下载wordpress的源码包

```shell
[root@host-1 ~]# wget https://cn.wordpress.org/wordpress-4.9.4-zh_CN.tar.gz
```

![img](assets/云服务器上线WordPress博客平台/1687398109061-cba5e863-30e5-4724-b4bf-febba1ab2f51.png)

![img](assets/云服务器上线WordPress博客平台/1687398141697-bea3c3de-3795-4e4d-bc64-6848e0751ab7.png)

### 6.3 解压wordpress源码包

```shell
[root@host-1 ~]# tar xf wordpress-4.9.4-zh_CN.tar.gz
```

![img](assets/云服务器上线WordPress博客平台/1687398239309-519b9ccc-6e35-460a-aea2-c4aa28c70aac.png)

## 7、设置云服务器安全组策略

![img](assets/云服务器上线WordPress博客平台/1687398379140-8b2f6e00-775a-478a-a8eb-0c260b9908f0.png)

![img](assets/云服务器上线WordPress博客平台/1687398393971-b666420a-4dfd-4977-9aa9-366c84145e92.png)

![img](assets/云服务器上线WordPress博客平台/1687398423058-99cf97de-9d57-4558-a1f0-f094055bb2b7.png)

![img](assets/云服务器上线WordPress博客平台/1687398482926-cc1c78f5-0ba9-40f8-85ff-72bd2c0d3290.png)

## 8、拷贝博客系统资源到网站发布目录

```shell
[root@host-1 ~]# cp -rf wordpress/* /var/www/html/
```

## 9、再次访问安装

![img](assets/云服务器上线WordPress博客平台/1687398725882-cf5a64ec-7344-4309-b6a4-3a43fcae0310.png)

![img](assets/云服务器上线WordPress博客平台/1687398793359-32d0863e-c5d8-4b5c-b0c6-454e9857361b.png)

```shell
[root@host-1 ~]# vim /var/www/html/wp-config.php
按英文的冒号:
```

![img](assets/云服务器上线WordPress博客平台/1687415646152-4a40b0e2-805e-4c85-94b7-7b74572acff2.png)

敲回车

再按小写的i

![img](assets/云服务器上线WordPress博客平台/1687398865919-36c9bde5-3a9e-4449-8240-39d7f18ebbb7.png)

![img](assets/云服务器上线WordPress博客平台/1687398886354-db61eec3-a0fc-408c-8ff7-219fd03a20c5.png)

![img](assets/云服务器上线WordPress博客平台/1687398931954-ecd2bc21-bd11-41d3-aec5-08aec18d4856.png)

![img](assets/云服务器上线WordPress博客平台/1687398986234-0cd02cee-d0b2-4a36-8298-be69f4e2e041.png)

![img](assets/云服务器上线WordPress博客平台/1687399040252-da914145-9e94-4f87-a93a-74e303413ac9.png)

![img](assets/云服务器上线WordPress博客平台/1687399056670-5a5eff81-0cd8-427e-9070-8a29cc95aac1.png)

![img](assets/云服务器上线WordPress博客平台/1687399076917-b2e16307-91f7-4f4e-a576-9171504c9f18.png)

![img](assets/云服务器上线WordPress博客平台/1687399129769-70872676-a582-4c89-af9f-137c017aaa6b.png)