# 02-企业级Zabbix监控平台

## 1、Zabbix可视化



##### 1、 简介



随着大量数据流入Zabbix，对用户来说，如果他们可以通过可视化图表查看数据而不仅仅是看到数字，那会更容易了解发生了什么事情。



图表可以使您一目了然地掌握数据流，关联问题，发现某些事情何时开始，或者表现其可能会出现问题



##### 2、 自定义图形（Graphs）



自定义图形中可以集中展示多个时间序列的数据流。支持“**线状图**(normal)”、“**堆叠面积图**(stacked)”、“**饼图**(pie)” 和“**分离型饼图**(exploded)”四种不同形式的图形。
　　具体的设置过程如下：
　　进入 配置 ---> 主机 ---> node1 ---> 图形，选择右上角创建图形：



![img](assets/02-企业级Zabbix监控平台/1683012454735-fe7e5ee7-ac73-4dd0-b221-bbbc6990a9a6.png)



查看图形：



![img](assets/02-企业级Zabbix监控平台/1683012455040-2028843b-3c9e-4940-97d8-a35f06075dc1.png)







**2个类似的监控项，也可关联到1个图形当中**



比如：创建2个类似的监控项



![img](assets/02-企业级Zabbix监控平台/1683012455029-1f94d596-bcbd-466e-a84d-d5934f477652.png)



![img](assets/02-企业级Zabbix监控平台/1683012454807-4818dfb8-6beb-4a67-b8d2-4dca39e592e8.png)



再克隆



![img](assets/02-企业级Zabbix监控平台/1683012455033-a5997872-113b-4900-865e-146d76903a23.png)



再次创建图形（包含以上2个监控项）：



![img](assets/02-企业级Zabbix监控平台/1683012456187-27d0bc37-f55e-4536-8a1a-6bc9e99af25a.png)



去查看图形：



![img](assets/02-企业级Zabbix监控平台/1683012456448-6fefd804-6812-4907-84fd-0a42960c292e.png)



##### 3、聚合图形



在Zabbix的聚合图形页面上，你可以把从各种图形聚合到一起，以便在单个屏幕上快速查看。构建聚合图形非常简单只管



![img](assets/02-企业级Zabbix监控平台/1683012457869-ecd12214-3690-4c62-9ae4-482cf6ddfe5d.png)



![img](assets/02-企业级Zabbix监控平台/1683012457868-d07fabe2-4510-40bd-a3d7-96394b48afa5.png)



![img](assets/02-企业级Zabbix监控平台/1683012458407-e90cd457-df83-46ac-8522-2ce447af2b07.png)



![img](assets/02-企业级Zabbix监控平台/1683012458443-91a85d25-b4f7-46f5-8078-dd039395f5b9.png)



效果图：



![img](assets/02-企业级Zabbix监控平台/1683012458598-e6785865-328d-4ca1-8f7c-5862ce984aa7.png)



##### 4、幻灯片



![img](assets/02-企业级Zabbix监控平台/1683012458967-c321c5f4-8b9d-45fb-ab8a-4b05a96704e8.png)



![img](assets/02-企业级Zabbix监控平台/1683012459287-e8f1992b-2c49-4ec9-b5c7-c6ab004131c6.png)



##### 5、解决字体乱码问题



在windows的“控制面板”中，找到1款字体；



![img](assets/02-企业级Zabbix监控平台/1683012459691-515f3395-e709-4077-899f-57a4caa110c8.png)



![img](assets/02-企业级Zabbix监控平台/1683012459674-dad9ca88-6cee-4726-b219-294689cf2ba5.png)



![img](assets/02-企业级Zabbix监控平台/1683012460376-5ba065a9-ca47-4d66-a682-cd8652d5abfc.png)



上传到Zabbix监控端



![img](assets/02-企业级Zabbix监控平台/1683012460381-a7fc65be-1d9d-4f48-91e5-8bbbe1509f0e.png)



```shell
[root@zabbix-server ~]# cd /usr/share/zabbix/assets/fonts
[root@zabbix-server fonts]# ls
graphfont.ttf
[root@zabbix-server fonts]# mv graphfont.ttf graphfont.ttf.bak
[root@zabbix-server fonts]# ls
graphfont.ttf.bak    
[root@zabbix-server fonts]# mv /root/simkai.ttf .
[root@zabbix-server fonts]# ls
graphfont.ttf.bak  simhei.ttf
[root@zabbix-server fonts]# mv simkai.ttf graphfont.ttf
[root@zabbix-server fonts]# chmod 777 graphfont.ttf
```



![img](assets/02-企业级Zabbix监控平台/1683012460511-896eaef6-7a0f-4f73-a0e2-10723db68fc0.png)



![img](assets/02-企业级Zabbix监控平台/1683012460928-7f7b176d-a252-41c4-877c-65078e7ed958.png)



##### 6、模板



模板是可以方便地应用于多个主机的一组实体。 实体可以是：



- 监控项
- 触发器
- 图形
- 应用集
- 聚合图形
- 自动发现规则
- web场景



当模板链接到主机时，模板的所有实体（监控项，触发器，图形，…）都将添加到主机。模板直接分配给每个单独的主机（而不是主机组）。



模板通常用于为特定服务或应用程序（如Apache，MySQL，PostgreSQL，Postfix …）分组实体，然后应用于运行这些服务的主机。



使用模板的另一个好处是当所有主机都需要更改时。只需要在模板上更改某些内容将会将更改应用到所有链接的主机。



因此，使用模板是减少工作量并简化Zabbix配置的好方法。



![img](assets/02-企业级Zabbix监控平台/1683012460997-45972208-a185-43df-acc5-11ddc1849c7c.png)



![img](assets/02-企业级Zabbix监控平台/1683012461451-dda73862-a16a-41fe-a74f-a9586343037d.png)



![img](assets/02-企业级Zabbix监控平台/1683012461863-d45568ab-96c2-454a-9434-b251a1701517.png)



![img](assets/02-企业级Zabbix监控平台/1683012461870-b474c26d-4154-4407-a9ca-70f16a8f4d57.png)



将 模板 关联到 主机



![img](assets/02-企业级Zabbix监控平台/1683012462271-bf8fe4ac-97e3-4b7c-b986-8c0ab01f8edb.png)



![img](assets/02-企业级Zabbix监控平台/1683012462468-75e24ac2-c15f-4b5d-903a-d93ad52e3bd9.png)



再次去查看主机的监控项：



![img](assets/02-企业级Zabbix监控平台/1683012463475-589abb79-340b-46ac-a525-ca5cda4130ba.png)



去  最新数据 中也可看到：



![img](assets/02-企业级Zabbix监控平台/1670311060080-e687d87c-58f0-4887-b1cc-2972afb9f0ed.png)



注意：1、一个主机可以链接多个模板.
　　   2、如果我们有多个主机，同时这些主机也在一个主机组里，这样的话，我们只需要在这个主机组里添加模板。



移除模板链接



一定要点击 清除链接并清理；这样才会把之前保留的数据清除；减缓数据库压力；



![img](assets/02-企业级Zabbix监控平台/1683012463511-0371235c-9f4c-4281-95ee-1ffae594cbf1.png)



## 2、用户参数（自定义监控）



##### 1、介绍和用法



① 介绍



自定义用户参数，也就是自定义监控项的键值



有时，你可能想要运行一个代理检查，而不是Zabbix的预定义



你可以**编写一个命令**来**检索需要的数据**，并将其包含在agent代理配置文件("UserParameter"配置参数)的**用户参数**中



```shell
用法格式 syntax
UserParameter=<key>,<command>
一个用户参数也包含一个键
在配置监控项时，key是必需的
注意:需要重新启动 agentd 服务
```



##### 2、用法展示



（1）修改agent 端的配置，设置用户参数



①自己需要查找的参数的命令



```shell
[root@zabbix-agent-none1 ~]# free | awk '/^Mem/{print $3}'
```



![img](assets/02-企业级Zabbix监控平台/1683012463555-ac1576a2-e63b-41e4-a544-519320b0b061.png)



② 修改agent配置文件，把查找参数的命令设为用户参数



```shell
[root@zabbix-agent-none1 ~]# cd /etc/zabbix/zabbix_agentd.d/
[root@zabbix-agent-none1 zabbix_agentd.d]# vim memory_usage.conf
UserParameter=memory.used,free | awk '/^Mem/{print $3}'
```



③ 重启agent 服务



```shell
[root@zabbix-agent-none1 zabbix_agentd.d]# systemctl restart zabbix-agent.service
```



（2）在zabbix-server 端，查询



```shell
[root@zabbix-server ~]# zabbix_get -s 192.168.153.178 -k memory.used -p 10050
```



（3）在监控上，设置一个item监控项，使用这个用户参数



配置-->主机-->none1-->监控项-->创建监控项



![img](assets/02-企业级Zabbix监控平台/1683012464134-133817a0-a3ce-457f-9ee2-18f7afeeb05b.png)



（4）查询 最新数据 中的 图形



![img](assets/02-企业级Zabbix监控平台/1683012464229-1bf79436-5750-47f4-93b3-fbaa35c159d1.png)

（5）创建3个监控项

已用内容，空闲内存，可用内存

```shell
[root@zabbix-agent-none1 zabbix_agentd.d]# cat memory_usage.conf 
UserParameter=memoryused,free | awk '/^Mem/{print $3}'
UserParameter=memoryfree,free | awk '/^Mem/{print $4}'
UserParameter=memoryAvia,free | awk '/^Mem/{print $NF}'
```

![img](assets/02-企业级Zabbix监控平台/1670313812779-964695b4-64eb-4125-984a-adb6a3c04c0e.png)

## 3、用法升级（自定义监控）



#### （1）修改agent 端的配置，设置用户参数



① 命令行查询参数的命令



![img](assets/02-企业级Zabbix监控平台/1683012464641-90442700-3230-4a24-bbe8-5a7f9ab762c3.png)



② 修改配置文件，把查找参数的命令设为用户参数



```shell
[root@zabbix-agent-none1 zabbix_agentd.d]# ls
memory_usage.conf
[root@zabbix-agent-none1 zabbix_agentd.d]# vim memory_usage.conf  ----继续添加
UserParameter=memory.stats[*],cat /proc/meminfo | awk '/^$1/{print $$2}'     --添加到文件中注意去掉反斜杠
[root@zabbix-agent-none1 zabbix_agentd.d]# systemctl restart zabbix-agent.service
```



```shell
注意：$$2：表示不是前边调位置参数的$1，而是awk 的参数$2
注意：$1是调用前边的[*]，位置参数，第一个参数
```



#### （2）在zabbix-server 端，查询使用这个用户参数的key



```shell
传参:
[root@zabbix-server fonts]# zabbix_get -s 192.168.153.178 -p 10050 -k "memory.stats[MemTotal]"
999696
[root@zabbix-server fonts]# zabbix_get -s 192.168.153.178 -p 10050 -k "memory.stats[MemFree]"
698736
[root@zabbix-server fonts]# zabbix_get -s 192.168.153.178 -p 10050 -k "memory.stats[MemAvailable]"
164
```



![img](assets/02-企业级Zabbix监控平台/1670314718159-1dc8b195-620b-4ae9-a861-347e9aee2178.png)



#### （3）在监控上，设置一个item监控项，使用这个用户参数



① 添加Memory Total 的item监控项，使用**memory.stats[MemTotal]** 的用户参数



![img](assets/02-企业级Zabbix监控平台/1683012464633-a01655f9-311c-4fba-bdbd-da8002b6375d.png)



② clone 克隆Memory Total 创建Memory Free 的监控项，使用**memory.stats[MemFree]** 用户参数



![img](assets/02-企业级Zabbix监控平台/1683012465355-714b32a1-7f4c-4036-bf7a-650ed92df91e.png)



#### （4）上面2个监控项的graph 图形



① memory total



![img](assets/02-企业级Zabbix监控平台/1683012465791-a192af2f-9ddb-4a8c-beae-9fff6d1b43de.png)



② memory free



![img](assets/02-企业级Zabbix监控平台/1683012465859-b1bb223c-4ca1-4a8d-887a-a4ac2ec3dbd6.png)



实战案例1：



## 4、使用用户参数监控php-fpm服务的状态



在agent 端：192.168.153.178



#### （1）下载，设置php-fpm 、nginx



```plain
[root@zabbix-agent-node1 ~]# yum -y install php-fpm
[root@zabbix-agent-node1 ~]# vim /etc/php-fpm.d/www.conf     #修改如下
user = nginx
group = nginx
pm.status_path = /php-fpm-status    #php-fpm 的状态监测页面 ，#打开注释并修改
ping.path = /ping      #ping 接口，存活状态是否ok   #打开注释
ping.response = pong    #响应内容pong  #打开注释
[root@zabbix-agent-none1 ~]# yum -y install nginx
[root@zabbix-agent-none1 ~]# systemctl start php-fpm
```



#### （2）设置nginx ，设置代理php，和php-fpm的状态页面匹配



```plain
[root@zabbix-agent-none1 ~]# vim /etc/nginx/nginx.conf
server {
        listen       80 default_server;
        server_name  localhost;
        root         /usr/share/nginx/html;

        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

        location / {
        }
        
        location ~ \.php$ {
            fastcgi_pass   127.0.0.1:9000;
            fastcgi_index  index.php;
            fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
            include        fastcgi_params;
        }
        location ~* /(php-fpm-status|ping) {
            fastcgi_pass   127.0.0.1:9000;
            fastcgi_index  index.php;
            fastcgi_param  SCRIPT_FILENAME  $fastcgi_script_name;
            include        fastcgi_params;

            access_log off;   #访问这个页面就不用记录日志了
       }
}
```



```plain
[root@zabbix-agent-node1 ~]# nginx -t 
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
[root@zabbix-agent-node1 ~]# systemctl start nginx
```



#### （3）在agent 端，设置用户参数



① 查询 curl 192.168.153.178/php-fpm-status



![img](assets/02-企业级Zabbix监控平台/1670377427004-ef3ebcfe-d8b2-4a8a-8de0-0a4522b963b6.png)



```plain
pool – fpm池子名称，大多数为www
process manager – 进程管理方式,值：static, dynamic or ondemand. dynamic
start time – 启动日期,如果reload了php-fpm，时间会更新
start since – 运行时长
accepted conn – 当前池子接受的请求数
listen queue – 请求等待队列，如果这个值不为0，那么要增加FPM的进程数量
max listen queue – 请求等待队列最高的数量
listen queue len – socket等待队列长度
idle processes – 空闲进程数量
active processes – 活跃进程数量
total processes – 总进程数量
max active processes – 最大的活跃进程数量（FPM启动开始算）
max children reached - 大道进程最大数量限制的次数，如果这个数量不为0，那说明你的最大进程数量太小了，请改大一点。
slow requests – 启用了php-fpm slow-log，缓慢请求的数量
```



② 设置



```shell
[root@zabbix-agent-node1 ~]# cd /etc/zabbix/zabbix_agentd.d/
[root@zabbix-agent-node1 zabbix_agentd.d]# vim php_status.conf     ---添加
UserParameter=php-fpm.stats[*],curl -s http://192.168.246.226/php-fpm-status | awk '/^$1/{print $$NF}'
#设置用户参数为php-fpm.stats[*]，$1为第一个参数；$$NF为awk中的参数
```



③ 重启服务



```plain
[root@zabbix-agent-node1 zabbix_agentd.d]# systemctl restart zabbix-agent
```



#### （4）在zabbix-server 端，查询使用这个用户参数的key



```plain
[root@zabbix-server ~]# zabbix_get -s 192.168.153.178 -p 10050 -k "php-fpm.stats[idle]"
4
[root@zabbix-server ~]# zabbix_get -s 192.168.153.178 -p 10050 -k "php-fpm.stats[max active]"
1
[root@zabbix-server ~]# zabbix_get -s 192.168.153.178 -p 10050 -k "php-fpm.stats[total processes]"
5
[root@zabbix-server ~]# zabbix_get -s 192.168.153.178 -p 10050 -k "php-fpm.stats[active]"
1
```



![img](assets/02-企业级Zabbix监控平台/1683012466031-fc368564-bcef-452f-a28a-a6f55384dc48.png)



#### （5）创建一个模板，在模板上创建4个item监控项，使用定义的用户参数



① 创建一个模板



![img](assets/02-企业级Zabbix监控平台/1683012466508-55d21086-a9f3-4964-a96e-58f4ca24cf3b.png)



② 在模板上配置items 监控项，使用刚定义的用户参数



fpm.stats[total processes]



![img](assets/02-企业级Zabbix监控平台/1683012466792-b8cd2834-263a-4757-be34-5b70d70ba017.png)



③ 再clone克隆一个items监控项



**fpm.stats[active processes]**



![img](assets/02-企业级Zabbix监控平台/1683012467445-2e444ac4-2d9b-4e48-88d8-dbf5c94ef9b5.png)



#### （6）host主机链接模板



![img](assets/02-企业级Zabbix监控平台/1683012467499-a55aa4de-d8ba-461d-86c4-e7c5486ad756.png)



（7）查看graph 图形



① php-fpm total processes



![img](assets/02-企业级Zabbix监控平台/1683012467615-31e039b2-e017-42b3-82cc-e2bed90901fb.png)



② php-fpm active processes



![img](assets/02-企业级Zabbix监控平台/1683012467748-a8f41998-9d12-429d-87b4-fb84d2b0e199.png)



（8）模板导出和导入，可以给别人使用。或者使用别人的



![img](assets/02-企业级Zabbix监控平台/1683012468014-59ab1a39-d087-4826-8958-aa3fb6afd493.png)



自己定义用户参数的文件，也不要忘记导出



```plain
/etc/zabbix/zabbix_agentd.d/php_status.conf
```



## 5、Network discovery 网络发现(自动发现)



### 1、介绍



（1）介绍



```shell
网络发现：zabbix server扫描指定网络范围内的主机；
网络发现是zabbix 最具特色的功能之一，它能够根据用户事先定义的规则自动添加监控的主机或服务等;
优点：
加快Zabbix部署
简化管理
在快速变化的环境中使用Zabbix，而不需要过度管理
```



（2）发现方式：



```shell
ip地址范围；
可用服务（ftp, ssh, http, ...）
zabbix_agent的响应；
snmp_agent的响应；
```



（3）网络发现通常包含两个阶段：discovery发现 和 actions动作



```shell
① discovery：
Zabbix定期扫描网络发现规则中定义的IP范围；
② actions：网络发现中的事件可以触发action，从而自动执行指定的操作，把discvery events当作前提条件；
添加/删除主机
启用/禁用host
向组中添加主机
移除组中的主机
从模板链接主机或取消链接
```



=======================================================



### 2、配置网络发现Network discovery



（1）利用第二台用于可被扫描发现的主机----192.168.153.143



① 安装agent 端的包



```shell
[root@zabbix-agent-node2 ~]#  rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm
[root@zabbix-agent-node2 ~]# yum -y install zabbix-agent zabbix-sender
[root@zabbix-agent-node2 ~]# systemctl stop firewalld
[root@zabbix-agent-node2 ~]# setenforce 0
```



② 设置agent 配置，可以把之前设置好的none1的配置传过来



```shell
[root@zabbix-agent-node2 ~]# vim /etc/zabbix/zabbix_agentd.conf
```



```shell
Server=192.168.153.147

ServerActive=192.168.153.147

Hostname=zabbix-agent-none2 #只需修改hostname
```



```shell
[root@zabbix-agent-none2 ~]# visudo       #修改sudo的配置,添加如下信息
```



```shell
zabbix ALL=(ALL) NOPASSWD: ALL
```



![img](assets/02-企业级Zabbix监控平台/1683012469060-e6fcacb7-7f3a-42f5-b49d-b89efafb03d2.png)



④ 开启服务



```shell
[root@zabbix-agent-none2 ~]# systemctl start zabbix-agent
```



（2）设置自动发现规则discovery



配置--自动发现--创建自动发现规则



![img](assets/02-企业级Zabbix监控平台/1683012469161-f7f12f27-f25f-496d-8c5c-396cabc0b9e6.png)



注释：在zabbix-server端



![img](assets/02-企业级Zabbix监控平台/1683012469299-3f4e4310-962e-44c8-9aa3-2a9dd4e03c96.png)



或者，换个键值:



![img](assets/02-企业级Zabbix监控平台/1683012469358-9318bf60-8e7e-40a5-9a66-4c2b1a9df3ab.png)



② 更新间隔：1h就好，不要扫描太过频繁，扫描整个网段，太费资源；这里为了实验，设为1m



（3）自动发现成功



![img](assets/02-企业级Zabbix监控平台/1683012469939-3805836a-05ce-44b3-a655-44ac9f6307bb.png)



### 3、设置自动发现discovery 的动作action



a) 创建



![img](assets/02-企业级Zabbix监控平台/1683012469985-a8410d46-47f8-4a2b-a6cd-c3334ce8e7b7.png)



① 设置A条件，自动发现规则=test net



② 设置B条件，自动发现状态=up



![img](assets/02-企业级Zabbix监控平台/1683012470934-4b94c836-7831-432a-9939-91d585adde8f.png)



③ 要做什么操作



添加主机到监控



自动链接Template OS Linux 到此host



![img](assets/02-企业级Zabbix监控平台/1683012470928-79e27c8f-4b64-4f0b-8fc2-ff44e940a32f.png)



![img](assets/02-企业级Zabbix监控平台/1683012470925-98dc28f5-d588-47ac-8b6e-de0a55171fb8.png)



d) 启用动作，查看效果



确实已经生效，添加主机成功，模板链接成功



![img](assets/02-企业级Zabbix监控平台/1683012471741-eb3c1404-3f71-4ec0-879d-2c3751013b94.png)



（5）如果自己需要添加的主机已经扫描添加完成，就可以关闭网络扫描了，因为太耗资源



![img](assets/02-企业级Zabbix监控平台/1683012471575-efe8cba6-6bac-42d7-978d-86de1da0bc5b.png)



## 6、web页面监控



### 1、介绍



（1）介绍



要使用Web监控，您需要定义web场景。Web场景包括一个或多个HTTP请求或“步骤”。Zabbix server根据预定义的命令周期性的执行这些步骤。如果主机是通过代理监控的话，这些步骤将由代理执行。



从 Zabbix2.2 开始，Web 场景和监控项，触发器等一样，是依附在主机/模板上的。这意味着 web 场景也可以创建到一个模板里，然后应用于多个主机。



任何web场景会收集下列数据：



- 整个场景中所有步骤的平均下载速度
- 失败的步骤数量
- 最近的错误信息



对于web场景的所有步骤，都会收集下列数据：



- 每秒下载速度
- 响应时间
- 响应码



### 2、创建设置web场景



配置--主机--none1--web场景--创建web场景



（1）创建



![img](assets/02-企业级Zabbix监控平台/1683012471985-57ece872-c9f8-4431-9f49-146233c9df54.png)



（2）配置web 监测



① 点击步骤，设置web page web页面



![img](assets/02-企业级Zabbix监控平台/1683012472161-b3d15ac5-9227-4571-b1cc-a79977f6d1d9.png)



a) 设置名为home page，URL为http://192.168.153.178/index.html 的web页面



![img](assets/02-企业级Zabbix监控平台/1683012472208-48459888-ab00-4ff2-924c-feda32fad77b.png)



![img](assets/02-企业级Zabbix监控平台/1683012472692-a3fe1984-6aaa-4263-96da-6c35860a2730.png)



b) 设置名为fpm status，URL为http://192.168.153.178/php-fpm-status 的web页面



![img](assets/02-企业级Zabbix监控平台/1683012473369-9256230a-67fd-48cd-94c1-d8bb909d78df.png)



![img](assets/02-企业级Zabbix监控平台/1683012473418-2cf7e98e-abf1-4ac4-85af-bc8a02897f91.png)



c) 设置2个web页面成功



![img](assets/02-企业级Zabbix监控平台/1683012473642-5c03fd17-09ed-42d6-a7b5-433b9d616f57.png)



### 3、查看测试



![img](assets/02-企业级Zabbix监控平台/1683012474280-82615817-1f79-40ca-91c6-306585994191.png)

```shell
web场景名称：node1 web test
这个web场景中包含检测2个页面

Download speed for scenario "node1 web test".
整个web场景的下载速度（也就是2个页面共同下载的速度）
		
Download speed for step "home page" of scenario "node1 web test".
web场景中页面(home page)的下载速度
		
Download speed for step "php-fpm-status" of scenario "node1 web test".
web场景中页面(php-fpm-status)的下载速度
		
Failed step of scenario "node1 web test".
web场景访问失败的步骤数量

Last error message of scenario "node1 web test".
web场景最近一次访问的错误信息

Response code for step "home page" of scenario "node1 web test".
web场景中页面(home page)的状态码
		
Response code for step "php-fpm-status" of scenario "node1 web test".
web场景中页面(php-fpm-status)的状态码

Response time for step "home page" of scenario "node1 web test".
web场景中页面(home page)的响应速度

Response time for step "php-fpm-status" of scenario "node1 web test"
web场景中页面(php-fpm-status)的响应速度
```



## 7、实现分布式 zabbix proxy 监控



### 1、实验前准备



-  ntpdate 192.168.153.147 同步时间 
-  关闭防火墙，selinux 
-  proxy主机设置主机名 hostnamectl set-hostname zbxproxy.youngfit.com 
-  vim /etc/hosts 每个机器都设置hosts，以解析主机名；DNS也行 



### 2、环境配置（4台主机）

| 机器名称              | IP配置          | 服务角色  |
| --------------------- | --------------- | --------- |
| zabbix-server         | 192.168.153.147 | 监控      |
| node1                 | 192.168.153.178 | 被监控端  |
| node2                 | 192.168.153.143 | 被监控端  |
| zbxproxy.youngfit.com | 192.168.153.148 | 代理proxy |



zabbix-server 直接监控一台主机 node1



zabbix-server 通过代理 proxy 监控 node2



所有机器配置解析



```shell
监控端
[root@zabbix-server ~]# hostnamectl set-hostname zabbix-server
[root@zabbix-server ~]# cat /etc/hosts
192.168.153.147 zabbix-server
192.168.153.148 zbxproxy.youngfit.com
192.168.153.143 zabbix-agent-node2

代理端
[root@192 ~]# hostnamectl set-hostname zbxproxy.youngfit.com
[root@zbxproxy ~]# cat /etc/hosts
192.168.153.147 zabbix-server
192.168.153.148 zbxproxy.youngfit.com
192.168.153.143 zabbix-agent-node2

被监控端
[root@node1 ~]# hostnamectl set-hostname zabbix-agent-node2
[root@node1 ~]# cat /etc/hosts
192.168.153.147 zabbix-server
192.168.153.148 zbxproxy.youngfit.com
192.168.153.143 zabbix-agent-node2
```



### 3、在 zbxproxy.youngfit.com 上配置 mysql



##### 1、安装配置 mysql5.7



```shell
[root@zbxproxy ~]# wget https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm
[root@zbxproxy ~]# rpm -ivh mysql80-community-release-el7-3.noarch.rpm
[root@zbxproxy ~]# vim /etc/yum.repos.d/mysql-community.repo
```



![img](assets/02-企业级Zabbix监控平台/1683012474192-a04ae4f9-2226-4bc1-b1a0-50257e1eab62.png)



```shell
[root@zbxproxy ~]# yum -y install mysql-server mysql
```



##### 2、启动数据库服务



```plain
[root@zbxproxy ~]#  systemctl start mysqld
```



![img](assets/02-企业级Zabbix监控平台/1683012474330-0a74839d-fa14-4be0-89a6-c10d6a48e932.png)



##### 3、顺便更改密码



```shell
[root@zbxproxy ~]# mysqladmin -uroot -p'wTp0z:pkkweI' password 'FeiGe@2021'
```



##### 4、创建数据库 和 授权用户



```shell
mysql> create database zbxproxydb character set 'utf8';
mysql> grant all on zbxproxydb.* to 'zbxproxyuser'@'192.168.153.%' identified by 'ZBXproxy@2023';
mysql> grant all on zbxproxydb.* to 'zbxproxyuser'@'localhost' identified by 'ZBXproxy@2021';
mysql> flush privileges;
```



### 4、在 zbxproxy.youngfit.com 上下载zabbix 相关的包，主要是代理proxy的包



```shell
[root@zbxproxy ~]# rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm
[root@zbxproxy ~]# yum -y install zabbix-proxy-mysql zabbix-get zabbix-agent zabbix-sender
```



1、导入数据



zabbix-proxy-mysql 包里带有，导入数据的文件



```plain
[root@zbxproxy ~]# rpm -ql zabbix-proxy-mysql
```



![img](assets/02-企业级Zabbix监控平台/1683012474862-6e64190a-cb40-4d28-b7ef-7c5917b2ddb1.png)



```shell
[root@zbxproxy ~]# cp /usr/share/doc/zabbix-proxy-mysql-5.0.12/schema.sql.gz  .
[root@zbxproxy ~]# gzip -d schema.sql.gz 
[root@zbxproxy ~]# mysql -uroot -p'FeiGe@2021' zbxproxydb < schema.sql
```



2、查看数据已经生成



![img](assets/02-企业级Zabbix监控平台/1683012474930-a19dce6a-7fff-4b3b-95e7-b7987374dcb4.png)



### 5、配置 proxy 端



```plain
[root@zbxproxy ~]# vim /etc/zabbix/zabbix_proxy.conf
```



![img](assets/02-企业级Zabbix监控平台/1683012475538-2335a8a8-2bcf-42c8-941e-da82ba413a51.png)



```plain
Server=192.168.153.147        # server 的IP
ServerPort=10051             # server 的端口

Hostname=zbxproxy.youngfit.com  # 主机名
ListenPort=10051             # proxy自己的监听端口
EnableRemoteCommands=1       # 允许远程命令
LogRemoteCommands=1          # 记录远程命令的日志

# 数据的配置
DBHost=localhost
DBName=zbxproxydb
DBUser=zbxproxyuser
DBPassword=ZBXproxy@2023

ConfigFrequency=30      # 多长时间，去服务端拖一次有自己监控的操作配置；为了实验更快的生效，这里设置30秒，默认3600s
DataSenderFrequency=1   # 每一秒向server 端发一次数据，发送频度
```



2、开启服务



```shell
[root@zbxproxy ~]# systemctl start zabbix-proxy
```



### 6、配置node2端允许proxy代理监控



如未安装zabbix-agent，先安装



```shell
[root@zabbix-agent-node2 ~]# yum install zabbix-agent zabbix-sender -y
```



```shell
[root@zabbix-agent-node2 ~]# vim /etc/zabbix/zabbix_agentd.conf
Server=192.168.153.148,192.168.153.147      #写proxy的地址即可，监控端可选
ServerActive=192.168.153.148,192.168.153.147
Hostname=zabbix-agent-node2
[root@zabbix-agent-node2 ~]# systemctl restart zabbix-agent #启动服务
```



### 7、把代理加入监控 server 创建配置 agent 代理



1、创建 agent 代理



![img](assets/02-企业级Zabbix监控平台/1683012475788-c9cf79ba-10ed-4b30-8adc-0daa9d4fdfff.png)



![img](assets/02-企业级Zabbix监控平台/1683012475804-80ceee18-a88a-44ee-8cdc-f0a6817e6d41.png)



### 8、创建node2 主机并采用代理监控



![img](assets/02-企业级Zabbix监控平台/1683012476212-0d422d9b-3281-4a42-9a5a-223be8d9ffd0.png)



### 9、创建item监控项



##### 1、随便创一个监控项 CPU Switches



![img](assets/02-企业级Zabbix监控平台/1683012476330-24a5ee4c-f644-412b-b41c-ef83456c8ba5.png)



![img](assets/02-企业级Zabbix监控平台/1683012476936-59dfbb89-3c30-4655-a573-628990eabc37.png)



**下面可以在被监控端下载个Apache,进行监控，查看是否生效**



被监控端下载Apache



```shell
[root@zabbix-agent-node2 ~]# yum -y install httpd
[root@zabbix-agent-node2 ~]# systemctl start httpd
```



在监控端页面创建监控tcp 80端口的监控项



![img](assets/02-企业级Zabbix监控平台/1683012477074-2cddaac8-f1fd-45b9-ac10-1b1fb3a68928.png)



在zabbix-server端测试，查看是否能获取的key值



```shell
[root@zbxproxy ~]# zabbix_get  -s 192.168.153.143 -k net.tcp.listen[80] -p 10050
1
```



![img](assets/02-企业级Zabbix监控平台/1683012477741-e02d24f3-371a-4992-a4e3-98d22ce6d633.png)



创建触发器



![img](assets/02-企业级Zabbix监控平台/1683012477893-5d49a865-9812-46b3-ab1a-4fa02ab2d7d4.png)



```plain
测试触发器效果
[root@zabbix-agent-node2 ~]# systemctl stop httpd
```



![img](assets/02-企业级Zabbix监控平台/1683012478055-5deec7a4-200a-4a9f-bcf8-9a4b703e57e0.png)



可见，触发器也能生效



但是现在还未学习发送报警消息



等学会了设置自动发送报警消息



再回来，测试是否可以正常报警



我这边已经经过实验，可以报警，报警图片如下图：



![img](assets/02-企业级Zabbix监控平台/1683012478237-62aa7bf8-9f36-4597-8b65-0f26b90e815c.png)