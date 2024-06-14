# Zabbix微信监控报警

# 实训项目：基于微信实现自动化监控报警

笔记本：千锋教育Linux云计算实训项目

Author： Youngfit

Email：ashiyufei[@youngfit.com ]() 

更新时间：2022/9/19



监控环境准备

部署zabbix-server【监控端】

部署zabbix-agent【被监控端】

实现对web服务器监控【监控端】

zabbix 企业微信报警【监控端】



#### 一、监控环境准备

###### 1.1重置云主机

关闭云主机，可能需要些时间

等待云主机启动

#### 二、部署zabbix-server【监控端】

###### 2.1 安装zabbix-server

```shell
[root@youngfit ~]# rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm
```

**替换国内阿里云yum源**
由于用zabbix官网网络Yum源下载速度比较慢，我这里替换为阿里云源

```shell
[root@zabbix-server ~]# sed -i 's#http://repo.zabbix.com#https://mirrors.aliyun.com/zabbix#' /etc/yum.repos.d/zabbix.repo
```

![img](assets/Zabbix微信监控报警/1663327636730-31e2c072-8ab0-48d3-b63d-219eeb87c1a4.png)

**2. 安装 Zabbix 服务器和代理**

```shell
[root@youngfit ~]# yum -y install zabbix-server-mysql  zabbix-agent
```

###### 2.2 安装zabbix前端

```shell
[root@youngfit ~]# yum install centos-release-scl -y
这里下载的是个zabbix相关的yum源，才能继续下载其他的软件包
```

编辑文件 /etc/yum.repos.d/zabbix.repo 并启用 zabbix-frontend 存储库。

```shell
[root@youngfit ~]# vim /etc/yum.repos.d/zabbix.repo
```

![img](assets/Zabbix微信监控报警/1663556200394-a3036189-f4b5-47a3-bf18-eb96864b2a4d.png)

安装 Zabbix 前端包

```shell
 [root@youngfit ~]# yum -y install zabbix-web-mysql-scl zabbix-apache-conf-scl
```

###### 2.3 准备数据库及授权

```plain
[root@youngfit ~]# yum -y install mariadb-server
[root@youngfit ~]# systemctl start mariadb
[root@youngfit ~]# mysql

> create database zabbix character set utf8 collate utf8_bin;

> grant all privileges on zabbix.* to zabbix@localhost identified by 'zabbix';

> flush privileges;

> \q
Bye
```



###### 2.4 导入数据库文件



```plain
[root@youngfit ~]# zcat /usr/share/doc/zabbix-server-mysql*/create.sql.gz | mysql -uzabbix -pzabbix zabbix
```

验证数据是否导入成功

```plain
[root@youngfit ~]# mysql
MariaDB [(none)]> use zabbix;
MariaDB [zabbix]> show tables;
```

![img](assets/Zabbix微信监控报警/1663658087866-5e9dd4fd-2676-4581-9d05-5d523e5cb031.png)

```plain
MariaDB [zabbix]> quit
Bye
```

###### 2.5 修改zabbix配置文件



```plain
[root@youngfit ~]# vim /etc/zabbix/zabbix_server.conf
DBHost=localhost
DBName=zabbix
DBUser=zabbix
DBPassword=zabbix
注意删除前面的空格
```



###### 2.6 启动zabbix-server



```plain
[root@youngfit ~]# systemctl restart zabbix-server zabbix-agent  //重启zabbix服务和agent
[root@youngfit ~]# systemctl enable zabbix-server zabbix-agent  //设置为开机自启动
```



###### 2.7 修改Apache配置文件，修改时区

编辑文件/etc/opt/rh/rh-php72/php-fpm.d/zabbix.conf，取消注释并为您设置正确的时区。

```plain
[root@youngfit ~]# vim /etc/opt/rh/rh-php72/php-fpm.d/zabbix.conf #最后一行
php_value date.timezone Asia/Shanghai
```

![img](assets/Zabbix微信监控报警/1663573935677-435604e9-407c-4339-80b4-99ec207d4d1a.png)

###### 2.8 启动 Zabbix 服务器和代理进程



```plain
[root@youngfit ~]# systemctl restart zabbix-server zabbix-agent httpd rh-php72-php-fpm
[root@youngfit ~]# systemctl enable zabbix-server zabbix-agent httpd rh-php72-php-fpm
```



###### 2.8 安装zabbix-server



```plain
打开浏览器用域名访问：
http://youngfit.online/zabbix
如果没有备案，选用公网ip的方式访问：
http://39.97.124.193/zabbix
```



由于老师的服务器没有备案，我们就用ip的方式访问

![img](assets/Zabbix微信监控报警/1663574065307-8943cd98-1afe-4c72-abb7-70f0b6f1aa5d.png)

![img](assets/Zabbix微信监控报警/1663574079325-2cb5d973-453c-47c7-be2f-ed3292a6f13e.png)

![img](assets/Zabbix微信监控报警/1663574102541-5ab4a447-d27d-4e6a-a044-7cfd9327fbde.png)

![img](assets/Zabbix微信监控报警/1663574135355-944502e1-c604-4d0b-958d-d2726349a286.png)

![img](assets/Zabbix微信监控报警/1663574151343-1f4408e2-3b48-41a0-beec-ef347730cfa4.png)

![img](assets/Zabbix微信监控报警/1663574162006-034acf54-036d-4ffc-9eda-0025f6357ecd.png)



安装成功之后，用户名为：Admin 密码为：zabbix

![img](assets/Zabbix微信监控报警/1663574189135-cc02ff06-589f-405b-899d-db326064a13b.png)

![img](assets/Zabbix微信监控报警/1663574290925-f42e5821-b3ff-4e77-9970-903716799dba.png)

页面显示为英文，可以设置为中文

Admin用户，可以设置任何用户的语言

设置自己的语言如下图

![img](assets/Zabbix微信监控报警/1663574374215-1e0cece4-0d27-4f34-8031-94b4f404ce08.png)

![img](assets/Zabbix微信监控报警/1663574404622-9753508d-25fe-4823-af64-7331b71137cb.png)

![img](assets/Zabbix微信监控报警/1663574417364-f299a104-a347-4b8c-9cd7-e1e930e44631.png)

#### 三、部署zabbix-agent【被监控端】



由于我们只有一台云主机，我们被就拿同一台云主机作为被监控端



###### 3.1 安装zabbix-agent



```plain
  [root@youngfit   ~]# yum -y install zabbix-agent
```



###### 3.2 zabbix-agent配置



```plain
[root@youngfit  ~]# vim /etc/zabbix/zabbix_agentd.conf
Server=47.75.81.162 #zabbix-server IP 即监控端的IP
ServerActive=47.75.81.162 #zabbix-server IP 即监控端的IP
Hostname=web1
```



###### 3.3 启动zabbix-agent



```plain
[root@web1 ~]# systemctl restart zabbix-agent
[root@web1  ~]# systemctl enable zabbix-agent
```

###### 3.4 开放被监控端口

安全组策略，打开10051和10050



###### 3.5 准备网站测试环境



```plain
如果是新环境，则做如下操作：
[root@jspgou ~]# yum -y install nginx
[root@jspgou ~]# vim /etc/nginx/nginx.conf
```

![img](assets/Zabbix微信监控报警/1663727261666-1aa318f1-2f14-43c0-a84b-70aa49751990.png)

```plain
[root@jspgou ~]# systemctl start nginx

#查看82端口是否被占用
[root@jspgou ~]# netstat -tnlp
```

![img](assets/Zabbix微信监控报警/1663727379768-c40c208b-3ba3-460a-8553-083325031e35.png)

目标：82端口存活，意味网站存活；82端关闭，意味着网站故障，要有报警消息；

![img](assets/Zabbix微信监控报警/1663727486242-f09b8270-ec9a-4bce-b328-02af12bcb97b.png)

#### 四、 实现对web服务器监控【监控端】

###### 4.1 添加【主机组】

![img](assets/Zabbix微信监控报警/1663575863429-b99f3189-2071-4a72-9280-6af8eec41899.png)

![img](assets/Zabbix微信监控报警/1663575876026-59956e71-8cbd-4fae-b31d-d19e18f2fdc2.png)



###### 4.2 添加主机

![img](assets/Zabbix微信监控报警/1663575942598-bb3b7c7b-4b8f-45da-b70b-fd2032557081.png)

![img](assets/Zabbix微信监控报警/1663576002507-4830e56b-81f2-46a4-b82c-76d15a4f3e1d.png)



###### 4.3 监听nginx的82端口



添加监控项

![img](assets/Zabbix微信监控报警/1663576159049-f0742643-db24-49dd-add5-044cc767c8cc.png)

![img](assets/Zabbix微信监控报警/1663576184409-46f58574-c9dd-474e-a711-dfa2ab796e0e.png)

![img](assets/Zabbix微信监控报警/1663727969911-e3cea69e-62a9-45c4-be78-62e92e228b1a.png)

上步做完，往下拉，再点击添加



创建触发器

![img](assets/Zabbix微信监控报警/1663576287253-fec264b1-a107-4b36-84c6-efe5151dfec5.png)

![img](assets/Zabbix微信监控报警/1663576296292-3c69e500-e323-4fb7-87ed-3545531a9d2f.png)

![img](assets/Zabbix微信监控报警/1663728122382-485ad7d7-1814-42da-bfab-a0e9f871b5d6.png)

![img](assets/Zabbix微信监控报警/1663728142962-3ab78f13-4fd8-48fd-b8de-9e845adcdbc9.png)

![img](assets/Zabbix微信监控报警/1663728168333-d1e3a66e-c16a-4b44-a9d8-d314108fbc9d.png)

![img](assets/Zabbix微信监控报警/1663728199865-abc2470d-dec1-4baf-83ef-5f3585d57860.png)

###### 4.4 测试触发器效果



```plain
被监控端停止网站服务器
[root@youngfit ~]# systemctl stop nginx
```

![img](assets/Zabbix微信监控报警/1663576436326-dae1597f-21a1-4407-90b9-1cfe30848191.png)





```shell
被监控端启动网站服务器
[root@youngfit ~]# systemctl start nginx
```

![img](assets/Zabbix微信监控报警/1663576473268-54d01c86-02db-4151-8e9c-b33bb498b0d5.png)

问题已恢复



#### 五、 zabbix 微信报警【监控端】



###### 5.1 下载企业微信【手机APP】

创建自己的企业，自定义即可



###### 5.2 企业登录后台【pc】

https://work.weixin.qq.com/

![img](assets/Zabbix微信监控报警/1663742887044-f04c2e08-4f72-4778-b7b3-b28361247bcc.png)

![img](assets/Zabbix微信监控报警/1663577218771-f2f4233d-8b2f-4afd-bd44-f7171eb72418.png)



###### 5.3 创建报警机器人

![img](assets/Zabbix微信监控报警/1663577359302-ddf7278b-cb19-4d22-a9e3-146d0178cac0.png)

![img](assets/Zabbix微信监控报警/1663577371354-54cd0e33-8090-4999-8505-3e881aa63dba.png)

![img](assets/Zabbix微信监控报警/1663577535458-0361755b-0de1-46bb-8b35-4943a27f8ee3.png)



点击“我的企业”，拉到最下面，查看企业ID

点击“应用管理”，机器人里面，查看Secret认证id

点击“应用管理，查看AgentId

点击“通讯录”，点击公司名称后面的

![img](assets/Zabbix微信监控报警/1663743645041-d88f61c1-aa45-4635-b9dd-3cf697708ce5.png)

```plain
应用：报警机器人，以下信息提前记录下来；
AgentID: 		  1000020
Secret: 			Zj8CzUUhqgxL4BkSveeCwFk5rOsjSuxsxW7FpwIA2ds
企业 CorpID:      wwd5348195e1cdd809
部门id: 		   1
```





###### 5.4 python报警监控脚本



```plain
需要修改脚本中的四个信息：
```



```plain
[root@zabbix-server ~]# vim /usr/lib/zabbix/alertscripts/wechat.py
#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib,urllib2,json
import sys
reload(sys)
sys.setdefaultencoding( "utf-8" )


class WeChat(object):
        __token_id = ''
        # init attribute
        def __init__(self,url):
                self.__url = url.rstrip('/')
                self.__corpid = 'wwd5348195e1cdd809'
                self.__secret = 'QyASUhjgIRaHHOHKdPKrPpGm8bxlNUrXLqU3toK6E4Q'


        # Get TokenID
        def authID(self):
                params = {'corpid':self.__corpid, 'corpsecret':self.__secret}
                data = urllib.urlencode(params)


                content = self.getToken(data)


                try:
                        self.__token_id = content['access_token']
                        # print content['access_token']
                except KeyError:
                        raise KeyError


        # Establish a connection
        def getToken(self,data,url_prefix='/'):
                url = self.__url + url_prefix + 'gettoken?'
                try:
                        response = urllib2.Request(url + data)
                except KeyError:
                        raise KeyError
                result = urllib2.urlopen(response)
                content = json.loads(result.read())
                return content


        # Get sendmessage url
        def postData(self,data,url_prefix='/'):
                url = self.__url + url_prefix + 'message/send?access_token=%s' % self.__token_id
                request = urllib2.Request(url,data)
                try:
                        result = urllib2.urlopen(request)
                except urllib2.HTTPError as e:
                        if hasattr(e,'reason'):
                                print 'reason',e.reason
                        elif hasattr(e,'code'):
                                print 'code',e.code
                        return 0
                else:
                        content = json.loads(result.read())
                        result.close()
                return content

        # send message
        def sendMessage(self,touser,message):
                self.authID()
                data = json.dumps({
                        'touser':touser,
                        'toparty':1,
                        'msgtype':"text",
                        'agentid':"1000017",
                        'text':{
                                'content':message
                        },
                        'safe':"0"
                },ensure_ascii=False)


                response = self.postData(data)
                print response

if __name__ == '__main__':
        a = WeChat('https://qyapi.weixin.qq.com/cgi-bin')
        a.sendMessage(sys.argv[1],sys.argv[3])
```



```plain
监控脚本测试：
给予脚本执行权限：
[root@youngfit ~]# chmod +x /usr/lib/zabbix/alertscripts/wechat.py

注：Youngfit为上图的企业微信名 test为标题 yufei为内容，标题和内容均可自定义
[root@youngfit ~]# /usr/lib/zabbix/alertscripts/wechat.py Youngfit test yufei
{u'errcode': 60020, u'errmsg': u'not allow to access from your ip, hint: [1663579212640293053095122], from ip: 124.70.32.161, more info at https://open.work.weixin.qq.com/devtool/query?e=60020'}

以上结果，属于发送失败，需要在企业微信报警机器人中，添加可信ip。将上面ip添加进去
```

![img](assets/Zabbix微信监控报警/1663744117523-b714695f-548f-4971-98d1-eabea16230ee.png)

![img](assets/Zabbix微信监控报警/1663744175886-51f33bfd-7cc2-4602-b06e-5b4cd3c483e6.png)



![img](assets/Zabbix微信监控报警/1663579735650-e141177f-523a-4dfa-96c4-6b1261cc131b.png)

再次测试

![img](assets/Zabbix微信监控报警/1663579786552-f1f60b91-8d92-4f2f-9c4f-5c3d5e23b9b9.png)

![img](assets/Zabbix微信监控报警/1663579812793-f4b84264-59eb-48c2-a494-5162bb0ebc5f.jpeg)

###### 5.6 zabbix 创建告警媒介



注：(以哪种方式发送报警信息,短信，脚本等等)

![img](assets/Zabbix微信监控报警/1663639875851-8e4fb516-2810-4225-9431-136b6f380b9f.png)

![img](assets/Zabbix微信监控报警/1663639968834-3b8bfe4d-f2d2-4d65-875f-23d679672692.png)

```plain
3个参数：
{ALERT.SENDTO}
{ALERT.SUBJECT}
{ALERT.MESSAGE}
ZABBIX告警通知
告警状态：【{TRIGGER.STATUS}】
告警主机：【{HOST.NAME}】
主机地址：【{HOST.IP}】
告警时间：【{EVENT.DATE} {EVENT.TIME}】
告警等级：【{TRIGGER.SEVERITY}】
告警名称：【{TRIGGER.NAME}】
当前状态：【{ITEM.NAME}：{ITEM.KEY}={ITEM.VALUE}】
事件代码：【{EVENT.ID}】
ZABBIX告警恢复
告警状态：【{TRIGGER.STATUS}】
告警主机：【{HOST.NAME}】
主机地址：【{HOST.IP}】
恢复时间：【{EVENT.RECOVERY.DATE} {EVENT.RECOVERY.TIME}】
告警等级：【{TRIGGER.SEVERITY}】
告警名称：【{TRIGGER.NAME}】
当前状态：【{ITEM.NAME}：{ITEM.KEY}={ITEM.VALUE}】
事件代码：【{EVENT.ID}】
```

![img](assets/Zabbix微信监控报警/1663744782062-ac16f2ea-e0b5-43c1-90aa-ec0e1872d67e.png)

![img](assets/Zabbix微信监控报警/1663744820099-0f77ec85-2f97-4bff-9c22-988fb292e6a5.png)

![img](assets/Zabbix微信监控报警/1663744861957-222a6a8f-e536-49a8-8d64-803d15afd450.png)

![img](assets/Zabbix微信监控报警/1663744881002-c5954338-941a-4f41-9d98-93aec9174e25.png)

测试一下媒介是否可用

![img](assets/Zabbix微信监控报警/1663640223102-c392a978-e22b-4174-951d-21acafabe8ab.png)

###### 5.7 zabbix 添加告警用户



```plain
创建用户组：web-group，包含yangge、alice...等用户
创建用户：yangge、alice...等用户，并关联到微信告警
希望的结果是：发微信到组web-group，组中的用户yangge、alice都能收到
```



创建用户组：web-group

![img](assets/Zabbix微信监控报警/1663640268567-85bbf173-0519-4e4d-953d-bc3948f057f6.png)

设置用户组web-group对主机群组youngfit-webserver拥有读写(发送消息)权限

![img](assets/Zabbix微信监控报警/1663746965420-1e224e7c-216b-446f-85ec-aedaf427c428.png)

![img](assets/Zabbix微信监控报警/1663747005394-b60ef2c5-5a71-4ea4-b389-efeeaf777c63.png)

![img](assets/Zabbix微信监控报警/1663747035142-1e70f75e-0826-4ad9-9cfe-d0a25ae66baa.png)

![img](assets/Zabbix微信监控报警/1663747055204-6fb44f73-6639-4b11-b072-5d137e1f2706.png)



创建用户 feige并关联告警媒介

![img](assets/Zabbix微信监控报警/1663640336018-944fd2d6-cca0-4987-936d-c016166791f0.png)

![img](assets/Zabbix微信监控报警/1663747168155-58d54b97-9cf8-4a63-a2f5-e74c014bd511.png)

![img](assets/Zabbix微信监控报警/1663640388291-387a0706-f572-4ed6-8d3c-54a5b86c4f78.png)

![img](assets/Zabbix微信监控报警/1663747211572-1cfa7653-e123-4f59-bb09-41a495a62f50.png)

###### 5.8 Zabbix 添加告警动作

![img](assets/Zabbix微信监控报警/1663640426750-6f55db0a-cceb-49ba-9c55-61099a5b96d1.png)

![img](assets/Zabbix微信监控报警/1663747299838-85838394-d67e-43f2-976a-f09de2e3138e.png)

![img](assets/Zabbix微信监控报警/1663640465074-2c456e52-fc3b-4727-bc38-e4e472a95cde.png)

![img](assets/Zabbix微信监控报警/1663747376262-6195e76a-4d82-459b-90eb-ae63a9016cc9.png)

![img](assets/Zabbix微信监控报警/1663640516776-90a43df4-2c57-452a-8781-e09dcce95270.png)



![img](assets/Zabbix微信监控报警/1663747403536-1e5dbc7e-2b68-4a34-8781-c27ea3c289f2.png)



###### 5.9 测试微信报警



被监控端操作



```plain
[root@youngfit ~]# systemctl stop nginx
```



![img](assets/Zabbix微信监控报警/1663747820600-63730911-aa6e-404e-bc05-77ea674ce9c3.jpeg)

```plain
[root@youngfit ~]# systemctl start nginx
```

![img](assets/Zabbix微信监控报警/1663747855851-20ffdd4a-027f-4bd4-8fdf-635b07d02212.png)