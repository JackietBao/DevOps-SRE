# Mysql多源主从复制

## 实验环境

| 操作系统   | 主机地址        | 主机名        |
| ---------- | --------------- | ------------- |
| CentOS 7.4 | 192.168.153.220 | mysql-master1 |
| CentOS 7.4 | 192.168.153.215 | mysql-master2 |
| CentOS 7.4 | 192.168.153.148 | mysql-slave1  |
| CentOS 7.4 | 192.168.153.219 | mysql-slave2  |

MySQL在5.7.2上面添加了多源复制（Multi-Source）功能，意味着一个从库可以连多个主库，从而同时进行同步，但是如果是同一个表的话，会存在主键和唯一索引冲突的风险，需要提前做好规划

个人认为MySQL能实现多源复制的关键有两点

\1.      MySQL5.7中的从库进行同步的SQL_THREAD，IO_THREAD可以并发执行，这使得使多数据源的binlog同时同步变为可能。我们只需要对每一个Master执行Change Master 语句，只需要在每个语句最后使用For Channel来进行区分。

master1:

change master to master2 channel 'master';



master2:

change master to master1 channel 'master1';



slave1:

change master to master1 channel 'master1';

change master to master2 channel 'master';

\2.      MySQL5.7中添加了channel（通道）来判别不同的数据源，这样slave可以非常简单的进行多数据源的配置与区分。

3.	由于复制的原理没有改变，在没有开启GTID的时候Master的版本可以是MySQL5.5、5.6、5.7。并且从库需要[master-info-repository](http://dev.mysql.com/doc/refman/5.6/en/replication-options-slave.html#option_mysqld_relay-log-info-repository)、[relay-log-info-repository](http://dev.mysql.com/doc/refman/5.6/en/replication-options-slave.html#option_mysqld_relay-log-info-repository)设置为table，否则会报错：

```shell
ERROR 3077 (HY000): To have multiple channels, repository cannot be of type FILE; Please check the repository configuration and convert them to TABLE.
```

## 实战操作

前提：本次使用的binlog日志方式。手动指定binlog日志名称以及位置的方式；

## 配置host解析

```shell
# cat /etc/hosts
192.168.153.220 mysql-master1
192.168.153.215 mysql-master2
192.168.153.148 mysql-slave1
192.168.153.219 mysql-slave2
```

## 安装Mysql 5.7

4台机器均安装，这里采用yum安装

Mysql官网地址：https://www.mysql.com/

![img](assets/Mysql多源主从复制/1639307709532-b5afb877-9ebf-4829-9541-ebc4f7eb301d.png)

```shell
# wget https://dev.mysql.com/get/mysql80-community-release-el7-4.noarch.rpm
# rpm -ivh mysql80-community-release-el7-4.noarch.rpm
# vi /etc/yum.repos.d/mysql-community.repo   //配置yum源，下载mysql5.7
```

![img](assets/Mysql多源主从复制/1639307815671-d5fea98b-6b72-4b97-9edd-5035ece6d1ae.png)

```shell
# yum -y install mysql-community-server
[root@mysq-master1 ~]# mysql --version   #一定要确定4台机器下载的Mysql版本一致
mysql  Ver 14.14 Distrib 5.7.36, for Linux (x86_64) using  EditLine wrapper
```

修改4台mysql的配置文件

打开/etc/my.cnf，在[mysqld]标签下添加内容：

```shell
server-id=1
log-bin=mylog
master-info-repository=TABLE
relay-log-info-repository=TABLE
```

![img](assets/Mysql多源主从复制/1642166719990-398d77da-9485-4161-bd8c-4a1ac134f434.png)![img](assets/Mysql多源主从复制/1642166738148-f5cbdda7-9151-4b97-9f35-68495d69a8f6.png)![img](assets/Mysql多源主从复制/1642166755219-2e3352c2-5aa7-4b3f-842c-96dad3e8a876.png)![img](assets/Mysql多源主从复制/1642166777246-612efa04-b5db-418d-b562-b7935d6fc4bf.png)

## 启动配置Mysql

mysql-master1节点操作

```shell
[root@mysq-master1 ~]# systemctl start mysqld
[root@mysq-master1 ~]# grep password /var/log/mysqld.log
```

![img](assets/Mysql多源主从复制/1639308306577-a466f992-25b4-4f23-91ea-b19a6a539598.png)

```shell
[root@mysq-master1 ~]# mysqladmin -uroot -p'WdJ_H?xVk8-k' password 'Youngfit@123456'
[root@mysq-master1 ~]# mysql -uroot -p'Youngfit@123456'
mysql> grant replication slave,replication client on *.* to 'rep'@'192.168.153.%' identified by 'Rep@123456';  --创建同步使用的用户，注意网段

mysql> flush privileges;
```

mysql-master2操作如同mysql-master1一样步骤

## 双主配置

查看mysql-master2的binlog日志名称和位置：

![img](assets/Mysql多源主从复制/1642166901447-e06dd14d-c683-4216-910c-8a73ff1c9bde.png)

在mysql-master1上指定操作：

```shell
mysql> change master to						--设置mysql-mastre2为自己的主节点
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002',
master_log_pos=856916 for channel 'mysql-master2';
Query OK, 0 rows affected, 2 warnings (0.00 sec)

mysql> start slave;
mysql> show slave status\G
```

查看mysql-master1的binlog日志名称和位置：

![img](assets/Mysql多源主从复制/1642166979149-0d9ed087-4183-4801-b8d0-8277510813f3.png)

在mysql-master2上指定操作：

```shell
mysql> change master to						--设置mysql-mastre1为自己的主节点
master_host='mysql-master1',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000003',   --我这里不一样，根据你自己的名称和位置去写
master_log_pos=659 for channel 'mysql-master1';

mysql> start slave;
mysql> show slave status\G
```

2台分别查看同步的io线程和sql线程是否为YES

```shell
mysql> show slave status\G
```

![img](assets/Mysql多源主从复制/1639309660142-a8b457a2-0d3b-448a-b49a-563c2c5cf8c1.png)

```shell
mysql> show slave status\G
```

![img](assets/Mysql多源主从复制/1639309694467-7121b912-bf67-4964-8cef-f2be19a5a378.png)

经常操作：

```shell
# 查看单个channel的状态
mysql> show slave status for channel 'mysql-master2'\G
# 停止单个channel同步
mysql>  stop slave for channel 'mysql-master2';
# 开启单个channel同步
mysql>  start slave for channel 'mysql-master2';
# 重置单个channel
mysql>  reset slave for channel 'mysql-master2';
# 查看所有channel
mysql> show slave status\G
# 开启所有channel
mysql> start slave;
```

 [监控](https://dev.mysql.com/doc/refman/5.7/en/replication-multi-source-monitoring.html)：系统库performance_schema增加了一些replication的监控表

```shell
mysql> use performance_schema
mysql> show tables like 'replicat%';
+-------------------------------------------+
| Tables_in_performance_schema (replicat%)  |
+-------------------------------------------+
| replication_applier_configuration         |###查看各个channel是否配置了复制延迟
| replication_applier_status                |###查看各个channel是否复制正常（service_state）以及事务重连的次数
| replication_applier_status_by_coordinator |###查看各个channel是否复制正常，以及复制错误的code、message和时间
| replication_applier_status_by_worker      |###查看各个channel是否复制正常，以及并行复制work号，复制错误的code、SQL和时间
| replication_connection_configuration      |###查看各个channel的连接配置信息：host、port、user、auto_position等
| replication_connection_status             |###查看各个channel的连接信息
| replication_group_member_stats            |###
| replication_group_members                 |###
+-------------------------------------------+
```

## 双从配置

mysql-slave1

```shell
[root@mysql-slave1 ~]# systemctl start mysqld
[root@mysql-slave1 ~]# grep password /var/log/mysqld.log 
[root@mysql-slave1 ~]# mysqladmin -uroot -p'oUm?lgeoA8yY' password 'Youngfit@123456'
[root@mysql-slave1 ~]# mysql -uroot -p'Youngfit@123456'
mysql> change master to						--设置mysql-mastre1为自己的主节点
master_host='mysql-master1',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000003',   --我这里不一样，根据你自己的名称和位置去写
master_log_pos=659 for channel 'mysql-master1';
Query OK, 0 rows affected, 2 warnings (0.00 sec)
    
mysql> change master to						--设置mysql-mastre2为自己的主节点
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002',   --我这里不一样，根据你自己的名称和位置去写
master_log_pos=856916 for channel 'mysql-master2';
Query OK, 0 rows affected, 2 warnings (0.00 sec)

mysql> start slave for channel 'mysql-master1';
Query OK, 0 rows affected (0.00 sec)

mysql> start slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)

mysql> show slave status\G
```

![img](assets/Mysql多源主从复制/1639310029238-97e50314-f95d-4600-9304-cc2b87864626.png)

![img](assets/Mysql多源主从复制/1639310051303-2efc12e7-259e-4c54-877a-f95bccdd4c4d.png)

mysql-slave2同样操作

```shell
[root@mysql-slave2 ~]# systemctl start mysqld
[root@mysql-slave2 ~]# grep password /var/log/mysqld.log 
[root@mysql-slave2 ~]# mysqladmin -uroot -p'oUm?lgeoA8yY' password 'Youngfit@123456'
[root@mysql-slave2 ~]# mysql-uroot -p'Youngfit@123456'
mysql> change master to						--设置mysql-mastre1为自己的主节点
master_host='mysql-master1',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000003',   --我这里不一样，根据你自己的名称和位置去写
master_log_pos=659 for channel 'mysql-master1';
Query OK, 0 rows affected, 2 warnings (0.00 sec)
    
mysql> change master to						--设置mysql-mastre2为自己的主节点
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002',   --我这里不一样，根据你自己的名称和位置去写
master_log_pos=856916 for channel 'mysql-master2';
Query OK, 0 rows affected, 2 warnings (0.00 sec)

mysql> start slave for channel 'mysql-master1';
Query OK, 0 rows affected (0.00 sec)

mysql> start slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)

mysql> show slave status\G
```

![img](assets/Mysql多源主从复制/1639310176398-cff256e1-7cc5-42c5-adee-d6571c2311c3.png)

![img](assets/Mysql多源主从复制/1639310196008-079c2321-c7c2-46fe-add0-0578404b13be.png)

## 验证集群

在mysql-master1上创建1个库，查看其他节点是否同步

```shell
mysql> create database testbase1;
Query OK, 1 row affected (0.00 sec)
```

![img](assets/Mysql多源主从复制/1639310286380-569af0f6-e47f-47ee-a0f2-9fc6a795f2dd.png)

![img](assets/Mysql多源主从复制/1639310306545-e5f3c76a-4005-4ea9-b879-7d631855a848.png)

![img](assets/Mysql多源主从复制/1639310315812-378bd8ea-e615-4167-95be-9441777e4173.png)

以上看出没有问题，再来试试mysql-master2插入数据，在mysql-master2节点操作

```shell
mysql> create table testbase1.table1(id int,age int);
```

![img](assets/Mysql多源主从复制/1639310423508-f1f65f37-09a2-4fca-868a-a443ea5f3f4a.png)

![img](assets/Mysql多源主从复制/1639310438675-1f6e221d-9ce3-4fd6-bafe-85801991d0ac.png)

![img](assets/Mysql多源主从复制/1639310450030-f2b1c83e-5d78-4aff-aa42-6cfc8a5afec5.png)

## 高可用性验证

停止mysql-master2节点

```shell
[root@mysq-master2 ~]# systemctl stop mysqld
```

能看到所有从节点，连接mysql-master2的io线程异常

![img](assets/Mysql多源主从复制/1639310643812-193c31f2-11cc-4aa4-869c-bb372be573eb.png)

![img](assets/Mysql多源主从复制/1639310684019-02277324-0586-4d52-a288-bc63cc41e280.png)

但是连接mysql-master1的仍为正常

![img](assets/Mysql多源主从复制/1639310726360-81b03dfe-bb1b-4c72-a452-b09a4866e77b.png)

![img](assets/Mysql多源主从复制/1639310740517-1a0fe344-a24e-4423-ac27-16d4e6e78d7c.png)

那么下面测试一下在mysql-master1上继续创建数据，看2个从节点是否还正常同步

```shell
mysql> create table  testbase1.table2(name char(3),home varchar(20));
Query OK, 0 rows affected (0.01 sec)
```

![img](assets/Mysql多源主从复制/1639310800671-b0333125-8517-4f48-909a-7d12e2f4079e.png)![img](assets/Mysql多源主从复制/1639310811913-859d7183-75c0-4292-bcaa-0afa31fb02fc.png)

## 修复问题

修复流程：

1.停止mysql-master2节点的数据库，清空数据，

2.mysql-master1将数据全部导出，发送给mysql-master2

3.mysql-master2导入数据，重新change master to

这种修复操作，一般在Mysql没有访问的时候，像我们公司，是晚上11点以后，基本不会有请求。如果不放心，可以直接先让防火墙屏蔽掉数据库端口3306，这样所有请求都无法进入，但是要对mysql-master1,mysql-slave1,mysql-slave2放行。

在mysql-master2操作

```shell
[root@mysq-master2 ~]# systemctl stop mysqld  #停止数据库
[root@mysq-master2 ~]# rm -rf /var/lib/mysql/*  #删除所有原来的数据，当然要确保其他节点正常，最好是再对所有数据进行一次备份，以防万一
[root@mysq-master2 ~]# rm -rf /var/log/mysqld.log
[root@mysq-master2 ~]# systemctl start mysqld
[root@mysq-master2 ~]# grep password /var/log/mysqld.log 
2021-12-12T12:17:13.471634Z 1 [Note] A temporary password is generated for root@localhost: #pHeA4KFauP?
[root@mysq-master2 ~]# mysqladmin -uroot -p'#pHeA4KFauP?' password 'Youngfit@123456'
```

在mysql-master1备份数据，并发送给mysql-master2

```shell
[root@mysq-master1 ~]# mysqldump -uroot -p'Youngfit@123456' -A --set-gtid-purged=OFF > alldata.sql
[root@mysq-master1 ~]# scp alldata.sql mysql-master2:/root/
```

在mysql-master2导入数据，保证跟集群数据一致

```shell
[root@mysq-master2 ~]# mysql -uroot -p'Youngfit@123456' < alldata.sql
[root@mysq-master2 ~]# mysql -uroot -p'Youngfit@123456'
mysql> grant replication slave,replication client on *.* to 'rep'@'192.168.153.%' identified by 'Rep@123456';  --创建同步使用的用户，注意网段
Query OK, 0 rows affected, 1 warning (0.00 sec)

mysql> flush privileges;
Query OK, 0 rows affected (0.00 sec)
```

下面进行恢复：

查看mysql-master1的binlog日志名称和位置：

![img](assets/Mysql多源主从复制/1642164843489-47fd4a5f-d116-4baf-8eb1-2bf2e2fc1600.png)

在Mysql-master2操作：

```shell
mysql> change master to                                                
master_host='mysql-master1',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000003',   #注意对应名称和日志位置
master_log_pos=498 for channel 'mysql-master1';

mysql> start slave for channel 'mysql-master1';
Query OK, 0 rows affected (0.00 sec)

mysql> show slave status for channel 'mysql-master1'\G
```

![img](assets/Mysql多源主从复制/1642165076758-0a8c3b80-e5a0-4e0a-8669-c2c0872138cb.png)

mysql-master1未恢复正常：

![img](assets/Mysql多源主从复制/1642166374376-c32a303e-9a42-4cea-a225-3bcd0b2b9674.png)

所以，mysql-master1需要重新指定：

先查看mysql-master2的binlog日志名称和位置：

![img](assets/Mysql多源主从复制/1642166444987-7d495a68-ff9a-4a31-817f-4895669fd8e9.png)

在mysql-master1上重新指定操作：

```shell
mysql> stop slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)

mysql> reset slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)

mysql> change master to                                                
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002',
master_log_pos=880 for channel 'mysql-master2';

mysql> start slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)

mysql> show slave status for channel 'mysql-master2';
```

![img](assets/Mysql多源主从复制/1642166481889-09b25e4a-13ee-43cb-802b-5dd5da2da510.png)

但是mysql-slave1还未恢复正常：

![img](assets/Mysql多源主从复制/1642165152221-7e3117ab-591f-4b14-857a-2681c02dfbe4.png)

在mysql-slave1上操作：

```shell
mysql> stop slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.01 sec)

mysql> reset slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)
```

看下mysql-master2的日志名称和位置：

![img](assets/Mysql多源主从复制/1642165298438-1b274990-e5cd-4e97-9ceb-8dba08cc2127.png)

在mysql-slave1上操作：

```shell
mysql> change master to                                                
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002',  -- 注意对应binlog日志名称和位置
master_log_pos=856916 for channel 'mysql-master2';

mysql> start slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)
```

![img](assets/Mysql多源主从复制/1642165445652-cb1a5990-2827-4779-a382-dcb360f00368.png)

在mysql-slave2上查看，也未恢复正常

![img](assets/Mysql多源主从复制/1642165486124-acff8185-8776-497a-82e8-c0691f3621fa.png)

在mysql-slave2上操作：

```shell
mysql> stop slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.01 sec)

mysql> reset slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.00 sec)
```

查看mysql-master2的binlog日志名称和位置：

![img](assets/Mysql多源主从复制/1642165596383-fa535129-2f91-420b-96a4-46435c6aa748.png)

在mysql-slave2上操作：

```shell
mysql> change master to
master_host='mysql-master2',
master_user='rep',
master_password='Rep@123456',
master_log_file='mylog.000002', -- 注意对应binlog日志名称和位置
master_log_pos=856916 for channel 'mysql-master2';

mysql> start slave for channel 'mysql-master2';
Query OK, 0 rows affected (0.01 sec)

mysql> show slave status for channel 'mysql-master2'\G
```

![img](assets/Mysql多源主从复制/1642165722475-33914506-1a08-4804-9dc9-aeb16eebecf8.png)

## 再次验证

在mysql-master1上新建数据，查看其他节点是否同步

```shell
mysql> create table testbase1.table2(id int,age int);
Query OK, 0 rows affected (0.01 sec)
```

其他3个节点查看验证：

![img](assets/Mysql多源主从复制/1642167460923-2bbfe0fb-9280-4100-9cb0-5a1c1a4e91f2.png)

![img](assets/Mysql多源主从复制/1642167471272-e15c60a7-20ea-4c6b-8ba7-18fdea1817b0.png)![img](assets/Mysql多源主从复制/1642167481386-571adb9c-0188-4495-b2dd-42cbfc6d27a7.png)

666！