# Zabbix结合Grafana图形展示

## Grafana简介

Grafana 是 Graphite 和 InfluxDB 仪表盘和图形编辑器。Grafana 是开源的，功能齐全的度量仪表盘和图形编辑器，支持 Graphite，InfluxDB 和 OpenTSDB。

Grafana 主要特性：灵活丰富的图形化选项；可以混合多种风格；支持白天和夜间模式；多个数据源；Graphite 和 InfluxDB 查询编辑器等等。

因为zabbix本身自带的图形比较少，不能满足我们的需求。所以，我们可以安装grafana来配合zabbix出图，让数据更加直观、形象地体现出来。

Grafana官网：https://grafana.com/

## Granfana下载

官网下载方法：https://grafana.com/grafana/download?edition=oss

选择开源版本，我这里和zabbix-server安装在一台虚拟机

```shell
[root@zabbix-server ~]# wget https://dl.grafana.com/oss/release/grafana-8.3.4-1.x86_64.rpm
[root@zabbix-server ~]# yum -y install grafana-8.3.4-1.x86_64.rpm
[root@zabbix-server ~]# systemctl start grafana-server
# 日志位置
[root@zabbix-server ~]# tail -f /var/log/grafana/grafana.log
```

![img](assets/Zabbix结合Grafana图形展示/1652529796856-d2becb64-f061-4ed0-ac9b-f29193b6f856.png)Granfana官网文档：https://grafana.com/docs/grafana/latest/getting-started/getting-started/

首次登录 Grafana：

1. 打开您的网络浏览器并转到 http://localhost:3000/。3000除非您配置了不同的端口，否则Grafana 侦听的默认 HTTP 端口是。
2. 在登录页面上，输入admin用户名和密码（默认admin）。
3. 单击**登录**。如果登录成功，您将看到更改密码的提示。
4. 在提示上单击“**确定**”，然后更改您的密码。
5. ![img](assets/Zabbix结合Grafana图形展示/1652532733850-b16d47d7-7b7d-4f4c-bdf2-cdb7d89ddfdf.png)

![img](assets/Zabbix结合Grafana图形展示/1652529874824-9d3febca-1d5e-4e59-923c-370b5fd954fa.png)

## 下载连接zabbix的插件

https://grafana.com/grafana/plugins/?utm_source=new-data-source&search=zabbix

![img](assets/Zabbix结合Grafana图形展示/1652530142128-dd6fa0a6-18aa-4cb1-8429-9d60bb973a42.png)

![img](assets/Zabbix结合Grafana图形展示/1652530688040-415a8324-66c2-40f5-92ea-ec551e56d5a3.png)

```shell
[root@grafana-server ~]# grafana-cli -h

# 是从github上下载，下载失败。多次尝试。不行尝试下载到本地，上传虚拟机。手动解压
# 我就是手动的
第一种方法：
[root@grafana-server ~]# grafana-cli plugins list-remote
[root@grafana-server ~]# grafana-cli plugins list-remote|grep -i zabbix
[root@grafana-server ~]# grafana-cli plugins install alexanderzobnin-zabbix-app
```

![img](assets/Zabbix结合Grafana图形展示/1652530724874-bd58ace0-4e05-4891-bc2c-c24863e6d0b7.png)

```shell
# 第二种方法
[root@grafana-server ~]# cp alexanderzobnin-zabbix-app-4.2.8.zip  /var/lib/grafana/plugins/
[root@grafana-server plugins]# unzip alexanderzobnin-zabbix-app-4.2.8.zip
[root@grafana-server ~]# systemctl restart grafana-server
```

## Grafana开启Zabbix插件

![img](assets/Zabbix结合Grafana图形展示/1652533046966-7055a640-e907-439e-8d7a-12b8991f1e10.png)

![img](assets/Zabbix结合Grafana图形展示/1652533092841-557b0f4e-2413-4357-b543-24a734ec5111.png)

![img](assets/Zabbix结合Grafana图形展示/1652533123598-ca03b39e-2caf-4965-9b9c-90861bec0391.png)

## 配置zabbix数据源

![img](assets/Zabbix结合Grafana图形展示/1652533158624-55f6c40c-e877-4eb2-95fb-7137b1f2a6d4.png)

![img](assets/Zabbix结合Grafana图形展示/1652533183224-8c65f0dd-9b94-4a54-9425-682c562e2e26.png)

往下拉，自己安装的插件zabbix在最下面

http://192.168.91.134/zabbix/api_jsonrpc.php

![img](assets/Zabbix结合Grafana图形展示/1652533316260-497e4700-c895-425a-92ff-dcc05a303dfd.png)

![img](assets/Zabbix结合Grafana图形展示/1652533347685-2629a3db-213d-44f2-bed9-c7e14242118b.png)

![img](assets/Zabbix结合Grafana图形展示/1652533376335-b906739b-9c0b-4486-ac95-ec8ae76455d8.png)

```shell
Name：自定义一个名称
Type： 下拉框中选择zabbix   #如果没有安装zabbix插件，此处则没有zabbix选项
URL：http://zabbix-serverIP/zabbix/api_jsonrpc.php   #zabbixAPI接口地址
Access: 默认Server(Default)，表示grafana直接到Server取数据，brower表示让每个客户端单独获取，一般不用。
Username：填写你zabbix的用户名
Password： 填写你zabbix用户名的密码
Trends：勾选
```

## Grafana可视化Zabbix数据

![img](assets/Zabbix结合Grafana图形展示/1652535256903-84144503-4e38-49db-8429-4b5a83dfcc75.png)

![img](assets/Zabbix结合Grafana图形展示/1652535266634-3cd58ed7-7d33-4607-9572-75e523d3abe0.png)

![img](assets/Zabbix结合Grafana图形展示/1652535341784-e87a13e3-b651-4132-b011-9b79ce11ad3d.png)

如没有以上选项，重启grafana-server，再次测试

![img](assets/Zabbix结合Grafana图形展示/1652535392721-7c3c4fc2-3682-433c-8f69-141ecebe8f79.png)

![img](assets/Zabbix结合Grafana图形展示/1652535501798-4df8f363-1d56-432f-8f19-0c094967931a.png)

![img](assets/Zabbix结合Grafana图形展示/1652535580574-22af8c41-e14f-4cd2-94c8-9047bd8d8ad2.png)

![img](assets/Zabbix结合Grafana图形展示/1652535604025-2020eb1f-871e-41c3-879d-14029f9833d1.png)

![img](assets/Zabbix结合Grafana图形展示/1652535610332-7789831b-38c5-4622-8bba-89eb40fcd7a0.png)