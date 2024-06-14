# 堡垒机JumpServer-V2.20.2

# JumpserverV2.20.2



#### 安装登录



仅需两步快速安装 JumpServer：



准备一台 2核4G （最低）且可以访问互联网的 64 位 Linux 主机；
以 root 用户执行如下命令一键安装 JumpServer。



```plain
# curl -sSL https://github.com/jumpserver/jumpserver/releases/download/v2.20.2/quick_start.sh | bash
```



![img](assets/堡垒机JumpServer-V2.20.2/1683015641741-b2f1728e-cc7b-4eb5-8408-ab7c2ab7d771.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015642436-5f20a2db-25fc-4fec-8fe9-96a29f0f9e07.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015640439-aaf8fe92-d84e-440d-a1bc-40f9d9092d0a.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015640829-5f274509-ce9b-4bfa-87ba-1378fa772671.png)



#### 用户配置



**用户、系统用户、特权用户的关系**



- 用户管理里面的用户列表 是用来登录jumpserver平台的用户, 用户需要先登录jumpserver平台, 才能管理或者连接资产
- **系统用户**是JumpServer 登录资产时使用的账号，如 root `ssh root@host`，而不是使用该用户名登录资产（ssh admin[@host) ]() `;`
- **特权用户**是资产上已存在的, 并且拥有 高级权限 的系统用户， JumpServer 使用该用户来 `推送系统用户`、`获取资产硬件信息` 等;
  **普通用户** 可以在资产上预先存在，也可以由 特权用户 来自动创建。





![img](assets/堡垒机JumpServer-V2.20.2/1683015640825-d8559c26-e6ca-4076-8864-2dbae5ced525.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015642207-da6002bb-7a9b-478e-8e61-9abc75a42345.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015642316-e1008dc1-f98f-49d4-9d69-b080c57f6ff2.png)



#### 特权用户



**特权用户** 是资产已存在的, 并且拥有 高级权限 的系统用户， 如 root 或 拥有 `NOPASSWD: ALL` sudo 权限的用户。 JumpServer 使用该用户来 `推送系统用户`、`获取资产硬件信息` 等。



来到被控制资源节点上创建特权用户"tom"，密码设置为"tom"



```plain
[root@host-1 ~]# useradd tom
[root@host-1 ~]# passwd tom
[root@host-1 ~]# vim /etc/sudoers
```



![img](assets/堡垒机JumpServer-V2.20.2/1683015642427-85abc0e3-a4ba-4dfb-b624-7952a7643f1e.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015643658-38328e9d-ed1f-4ad0-bf85-672733fec6e1.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015643720-80a2dca6-5b53-45d6-b538-3db634916b3b.png)



#### 系统用户



![img](assets/堡垒机JumpServer-V2.20.2/1683015643829-1252af6e-003e-4cf3-9563-51a9a75214ee.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015644026-14dd987d-3524-42af-a512-77f674b8a592.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015644106-a0ddab45-db80-4d85-a211-1f862d8c3ef7.png)



#### 网域列表



网域功能是为了解决部分环境（如：混合云）无法直接连接而新增的功能，原理是通过网关服务器进行跳转登录。JMS => 网域网关 => 目标资产



![img](assets/堡垒机JumpServer-V2.20.2/1683015645090-3f8e2669-1fc5-4d25-84fc-3d817da7c38a.png)



#### 资产列表



![img](assets/堡垒机JumpServer-V2.20.2/1683015646175-81fff547-d7e4-4ba4-9d1a-80b582a33b41.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015645991-080781f3-b922-4726-aefc-3536cd854b58.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015646194-acd030ac-8e4f-4393-bff1-65da0e8ccf10.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015646364-fae2fe85-cb37-4c41-97fa-eaa86a79d978.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015646383-b63497d5-cd07-44ff-ace8-2df0ac8e8bec.png)



#### 资产授权



指的是将某些资产(虚拟机资源)赋予某个平台用户。让这个平台用户具有连接这些资产的权限；



![img](assets/堡垒机JumpServer-V2.20.2/1683015647196-07d257b8-1d84-4697-b5b0-d749c69f1004.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015647579-8eb49efb-65cb-415d-935b-0a6b792048ab.png)



#### 验证测试



##### web页面



退出admin,使用wangdana用户进行登录。查看是否可正常连接资产



![img](assets/堡垒机JumpServer-V2.20.2/1683015647866-4d8a602e-b9b3-4dcb-84a9-90dabc900350.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015648014-2b4237f2-00ed-43f8-921e-13088ffadaaf.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015648015-fd8fb45c-24d5-4ec3-93ef-68f3bee01368.png)



##### 通过 web 终端连接资产



随便找一台虚拟机



```plain
语法： ssh -p 2222 平台用户@堡垒机IP
# ssh -p 2222 wangdana@10.8.165.41
```



![img](assets/堡垒机JumpServer-V2.20.2/1683015648749-eb57b770-c3ed-4f43-9719-9aa40e06e947.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015648958-5fd013d7-be03-428f-ba9b-8eabf737b34e.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015648920-da92cd91-78d3-4cdc-96c9-259352a0aa9f.png)



#### 命令过滤



系统用户可以绑定一些命令过滤器，一个过滤器可以定义一些规则 当用户使用这个系统用户登录资产，然后执行一个命令 这个命令需要被绑定过滤器的所有规则匹配，高优先级先被匹配， 当一个规则匹配到了，如果规则的动作是允许，这个命令会被放行， 如果规则的动作是禁止，命令将会被禁止执行， 否则就匹配下一个规则，如果最后没有匹配到规则，则允许执行



![img](assets/堡垒机JumpServer-V2.20.2/1683015650401-61f4bb06-1118-44f6-bc1a-fd29c868c961.png)



![img](assets/堡垒机JumpServer-V2.20.2/1683015650380-a514fae7-2419-4f98-9162-319f508994f5.png)



将系统用户与命令过滤器绑定成功



切换为"wangdana"平台用户，用james登录host-1进行测试：



![img](assets/堡垒机JumpServer-V2.20.2/1683015650443-bf0e8d75-c667-4173-8276-08f77a27f57b.png)



#### 回放功能



点着找一找



# 二进制部署

# 部署准备

![img](assets/堡垒机JumpServer-V2.20.2/1656040367189-a3bd06b4-2e26-4b6d-ae1e-7e3edf78dad8.png)

```plain
#部署脚本
链接：https://pan.baidu.com/s/1qeLFg-d8-K4rKbAeKKywKQ 
提取码：cbc0 


后续执行脚本可能第一步出现download失败jumpserver-installer-v2.23.1.tar.gz
#下载jumpserver-installer-v2.23.1.tar.gz
链接：https://pan.baidu.com/s/1XVwsO98KDWQEGuU5NyTTWQ 
提取码：3zl2 

说明：将jumpserver-installer-v2.23.1.tar.gz解压到/opt下。执行脚本（脚本上传至服务器哪个目录，都可直接运行）

#内核版本修改
https://blog.csdn.net/alwaysbefine/article/details/108931626
```

# 部署

## 1、执行脚本及处理报错

上传脚本

\#赋予权限并执行

```plain
chmod +x ./quick_start.sh
bash ./quick_start.sh
```


\#download失败可单独上传上述tar.gz解压到/opt下再执行脚本

```plain
tar xf jumpserver-installer-v2.23.1.tar.gz -C /opt
```


再次执行sh脚本，正常（因为jumpserver-installer-v2.23.1已存在，脚本会跳过执行download）

![img](assets/堡垒机JumpServer-V2.20.2/1656040615311-455161fd-b94f-4213-96fa-e9aa5717ab1e.png)

## 2、更改内核

\#查看内核版本

```plain
[root@jumpserver opt]# cat /proc/version 
Linux version 3.10.0-693.el7.x86_64 (builder@kbuilder.dev.centos.org) (gcc version 4.8.5 20150623 (Red Hat 4.8.5-16) (GCC) ) #1 SMP Tue Aug 22 21:09:27 UTC 2017


发现为3.10，不符合环境要求，趁部署jumpserver，更改内核版本
https://blog.csdn.net/alwaysbefine/article/details/108931626
```

配置好后记得重启



**！！！等jumpserver部署好后再重启！！！**

## 3、安装完成后

一路n或者回车![img](assets/堡垒机JumpServer-V2.20.2/1656041757705-c3fa0f29-6249-4259-b75f-a98233120310.png)

\# 安装完成后配置文件 /opt/jumpserver/config/config.txt



```shell
cd /opt/jumpserver-installer-v2.23.1

# 启动
./jmsctl.sh start

# 停止
./jmsctl.sh down

# 卸载
./jmsctl.sh uninstall

# 帮助
./jmsctl.sh -h
```



**登录**

默认用户名/密码：admin/admin

![img](assets/堡垒机JumpServer-V2.20.2/1656043182702-99f8284c-150a-4bbf-a447-6c9fb2050473.png)**修改密码**![img](assets/堡垒机JumpServer-V2.20.2/1656043226359-dda85c29-776c-481a-910d-0448bede17d1.png)**重新登录**![img](assets/堡垒机JumpServer-V2.20.2/1656043270142-a30ade65-df07-420a-bc65-453bfbfb6155.png)**登录界面**![img](assets/堡垒机JumpServer-V2.20.2/1656043309135-ec830938-016c-4cea-9e62-717bd09a444e.png)

# 更多文档帮助

\#需要开翻墙

https://docs.jumpserver.org/zh/master/



# Jumpserver忘记管理员admin密码



![img](assets/堡垒机JumpServer-V2.20.2/1662029167448-e4482a86-eaf8-49e8-879c-9887ee6857ab.png)

```shell
[root@jumpserver ~]# docker exec -it jms_core /bin/bash
root@98567733dfb2:/opt/jumpserver# cd apps/
root@98567733dfb2:/opt/jumpserver/apps# python manage.py changepassword admin
```

![img](assets/堡垒机JumpServer-V2.20.2/1662029141705-17adac15-30e4-420a-91ca-cfcdc13173a5.png)

再次登录