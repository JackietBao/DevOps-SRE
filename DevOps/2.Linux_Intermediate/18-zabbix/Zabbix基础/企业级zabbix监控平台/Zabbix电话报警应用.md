# Zabbix电话报警应用

前提条件：

监控端，正在监控被监控端的http 80端口；设置触发器，我这里没有显示而已。

## 监控内容

![img](assets/Zabbix电话报警应用/1652483605602-99ccc9d8-d08b-44b0-aee2-9e23da301397.png)

![img](assets/Zabbix电话报警应用/1652483623563-3ab13063-553d-4435-af32-b0a44877d05e.png)

## 使用睿象云

https://www.aiops.com/ 睿象云

![img](assets/Zabbix电话报警应用/1652483769149-659a377f-649e-4e21-a653-9b9383b04700.png)

![img](assets/Zabbix电话报警应用/1652483906778-d29673d2-7f7a-49ce-8f4a-05ab8cec7add.png)

![img](assets/Zabbix电话报警应用/1652483958617-55d988c2-297a-49b9-8193-02d8ae6dd193.png)

![img](assets/Zabbix电话报警应用/1652484151624-41a4b908-38ca-4f70-b4af-1df7763dd14e.png)

应用名称：zabbix

AppKey：82e9d8885eb44fc7b1da61afc81d10a9

```shell
一、安装 Agent
1、切换到zabbix脚本目录 (注意查看自己的zabbix存放脚本的目录)：
cd /usr/local/zabbix-server/share/zabbix/alertscripts

2、获取Cloud Alert Agent包：
wget https://download.aiops.com/ca_agent/zabbix/ca_zabbix_release-4.0.1.tar.gz

3、解压、安装。
tar -xzf ca_zabbix_release-4.0.1.tar.gz
cd cloudalert/bin
bash install.sh {appKey}
注：1、在安装过程中根据安装提示，输入zabbix管理地址、管理员用户名、密码。

 2、zabbix管理地址正确示例：http://zabbix.server.com/zabbix 或是：https://zabbix.server.com/zabbix

4、修改运行zabbix服务权限与cloudalert探针目录权限
请保证运行zabbix服务的权限和cloudalert探针目录的权限保持一致，不一致会导致告警无法正常接入。

5、验证告警集成
产生新的zabbix告警(problem)，动作状态为“已送达”表示集成成功。
```

![img](assets/Zabbix电话报警应用/1652484315229-080938e9-a950-4123-9e60-5d8438bde8fe.png)

![img](assets/Zabbix电话报警应用/1652484346912-49d08957-26e2-4934-bc31-ad3ab61ad8c9.png)

安装完成之后，会生成zabbix新的用户和报警媒介：

![img](assets/Zabbix电话报警应用/1652486901501-3d687772-4e7f-48a0-8619-e92fb45fb1c2.png)

![img](assets/Zabbix电话报警应用/1652486923246-e509f072-77b9-4cbe-8223-83a966e3a259.png)

## 分配策略

分配策略：什么类型的报警发送到哪个用户

![img](assets/Zabbix电话报警应用/1652484493375-45713c3d-1c23-4acd-b75b-ff5b98c0babc.png)

![img](assets/Zabbix电话报警应用/1652485651434-3542d64c-0cca-4a9d-97ba-ca9991977b31.png)

## 通知策略

通知策略：采用什么方式进行通知

![img](assets/Zabbix电话报警应用/1652485898162-ee5ab03b-df1b-4e11-b659-996dab527207.png)

![img](assets/Zabbix电话报警应用/1652485960143-bcd0e577-bda6-4ab7-804c-e6ac41ddb326.png)

## Zabbix添加动作

关联触发器

发送到 Cloud Alert User 组 和 Cloud Alert User 用户

![img](assets/Zabbix电话报警应用/1652486745972-8f130b0b-6759-417d-9a73-855d8fc1c574.png)

![img](assets/Zabbix电话报警应用/1652486753358-54a20b97-80fc-4958-91ff-960064037830.png)

测试报警

```shell
[root@zabbix-agent1 ~]# systemctl stop httpd
```

![img](assets/Zabbix电话报警应用/1652486698929-729b4780-f576-46ea-827f-21dc75d79f42.png)

![img](assets/Zabbix电话报警应用/1652487434542-75b1f909-f68d-4e34-b7e1-bfdb8027ea6d.jpeg)