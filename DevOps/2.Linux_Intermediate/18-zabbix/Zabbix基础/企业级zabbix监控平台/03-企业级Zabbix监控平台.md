# 03-企业级Zabbix监控平台

## Zabbix邮箱自动监控报警实战



邮件系统简要介绍：
电子邮件系统包括两个组件：



MUA(Mail User  Agent,邮件用户代理）和MTA(Mail Transport  Agent,邮件传送代理 postfix）;



MUA是邮件系统为用户提供的可以读写邮件的界面；而MTA是运行在底层，能够处理邮件的收发工作的程序;



mail和mailx即为负责查看、编写邮件和向MTA发送邮件的MUA。mailx是mail的功能加强版。



```shell
1.注册163邮箱
2.登陆网页邮箱设置客户端授权密码
```



#### 注册邮箱



注册地址：https://mail.163.com/



![img](assets/03-企业级Zabbix监控平台/1683012543319-dfe0bd57-09e0-46cc-a250-bf764bde8464.png)



#### 开启SMTP服务



![img](assets/03-企业级Zabbix监控平台/1683012544359-8b577e46-e7ea-4379-9eca-16df44489d30.png)



![img](assets/03-企业级Zabbix监控平台/1683012544415-dd7a8fe6-b34d-4901-bebf-57ba3bb24bc7.png)



![img](assets/03-企业级Zabbix监控平台/1683012543911-798d4008-1463-4741-b256-4e785646a4f3.png)



![img](assets/03-企业级Zabbix监控平台/1683012544252-fa21acf2-14c9-436d-8cee-2b36340409e7.png)



#### 监控端测试邮箱



**server服务器端：**



安装MUA软件：mailx



```shell
阿里云的yum源可以下载
[root@zabbix-server ~]# yum install mailx -y
[root@zabbix-server ~]# mailx -V 
12.5 7/5/10
注：使用新的方式--利用公网邮件服务器发送报警，需要关闭postfix服务
[root@zabbix-server ~]# systemctl stop postfix
```



配置公网邮箱信息：



```shell
[root@zabbix-server ~]# vim /etc/mail.rc  ---在最后添加如下：
```



```shell
set from=feigeyoungfit@163.com（邮箱地址） 
set smtp=smtp.163.com（smtp服务器） 
set smtp-auth-user=feigeyoungfit@163.com(用户名) 
set smtp-auth-password=TBELSWWIEOYTKPYF（这里是邮箱的授权密码） 
set smtp-auth=login
```



使用mailx发邮件的方式：



```shell
方式1：echo "正文内容" | mailx -s "邮件标题" 收件箱Email
方式2：mailx -s "邮件标题" 收件箱Email，回车按CTRL+D发送
参数：
-v ：显示发送的详细信息
```



手动发送邮件测试：



```shell
[root@zabbix-server ~]# mailx -v -s 'hello' 'feigeyoungfit@163.com'
123
123
EOT
```



手写邮件内容 （回车，然后ctrl+d正常结束)



手动使用mailx发送邮件测试结果：



![img](assets/03-企业级Zabbix监控平台/1683012544824-77227015-6be9-408a-a5a0-540284363f60.png)



zabbix添加邮件报警功能:



配置 zabbix 的邮件报警功能需要以下三个角色的参与。



```shell
1、 报警媒介(Media)
2、 触发器(Triggers)
3、 动作(Action)
```



```shell
报警媒介:
    指的是 zabbix 采用何种方式进行报警,目前 Zabbix 支持的示警媒体包括邮件、脚本、短信。
    
触发器:
    指的是当监控对象达到某个条件或条件集合的时候,触发 Zabbix 产生事件。
    
动作:
    指的是 Zabbix 产生对应事件后,它通过示警媒体发送报警。

接下来,我们配置一个邮件报警功能的范例。效果是当redis端口挂掉则触发报警,管理员将会收到一封 Zabbix 发出的报警邮件。
```



示警媒体的配置:



```shell
首先需要配置 Zabbix 的邮件功能。
点击 管理->报警媒介类型->创建媒体类型
```



1.然后在页面中填入你的报警媒介类型信息，例如下图所示:



```shell
注：脚本名称任意，存放于/usr/lib/zabbix/alertscripts
(生产上的测试服放这：/usr/local/zabbix/share/zabbix/alertscripts）
```



```plain
名称：sendmail.sh                   //名称任意
类型：脚本
脚本名称：sendmail.sh      
脚本参数：                          //一定要写，否则可能发送不成功
    {ALERT.SENDTO}              //照填，收件人变量
    {ALERT.SUBJECT}             //照填，邮件主题变量，变量值来源于‘动作’中的‘默认接收人’
    {ALERT.MESSAGE}           //照填，邮件正文变量，变量值来源于‘动作’中的‘默认信息’

配置完成后,不要忘记点击存档,保存你的配置。
```



![img](assets/03-企业级Zabbix监控平台/1683012545103-70723d68-4a8f-4ab6-a841-cfae427ff47a.png)



邮件内容模板：



发现问题：



![img](assets/03-企业级Zabbix监控平台/1683012545575-7be1bf25-c184-4cb6-a310-e4ab400dfadb.png)



邮件内容模板：



问题恢复：



![img](assets/03-企业级Zabbix监控平台/1683012545743-05ee0d3a-6e2e-4cf5-a099-72bd05be36d8.png)



```plain
默认信息：邮件的主题
主机: {HOST.NAME1}
时间: {EVENT.DATE} {EVENT.TIME}
级别: {TRIGGER.SEVERITY}
触发: {TRIGGER.NAME}
详情: {ITEM.NAME1}:{ITEM.KEY1}:{ITEM.VALUE1}
状态: {TRIGGER.STATUS}
项目：{TRIGGER.KEY1}
事件ID：{EVENT.ID}
```



![img](assets/03-企业级Zabbix监控平台/1683012545817-2f242b15-a341-4ba4-b62d-60b97c4e59a9.png)



2.修改zabbix服务端配置文件＆编写脚本：指定脚本的存储路径:



```shell
[root@zabbix-server ~]# vim /etc/zabbix/zabbix_server.conf
```



![img](assets/03-企业级Zabbix监控平台/1683012546427-15295304-2048-4771-8db2-a49b7247472e.png)



```shell
 AlertScriptsPath=/usr/lib/zabbix/alertscripts
```



编写邮件脚本:



脚本仍是发送邮件，只不过换了1种方式



```shell
[root@zabbix-server ~]# cd /usr/lib/zabbix/alertscripts/
[root@zabbix-server alertscripts]# vim sendmail.sh
#!/bin/bash 
#export.UTF-8
echo "$3" | sed s/'\r'//g | mailx -s "$2" $1
```



```shell
$1:接受者的邮箱地址：sendto
$2:邮件的主题：subject
$3:邮件内容：message
```



```plain
监控端测试：
[root@zabbix-server alertscripts]# echo 'feige' | mailx -s nihao feigeyoungfit@163.com
```



![img](assets/03-企业级Zabbix监控平台/1683012546608-9cdf987f-000d-4d41-9f97-454b80cdc92f.png)



修改权限：



```shell
[root@zabbix-server alertscripts]# chmod u+x sendmail.sh && chown zabbix.zabbix sendmail.sh
或者：
[root@zabbix-server alertscripts]# chmod +x sendmail.sh
```



修改admin用户的报警媒介：



用户默认是没有关联报警媒介的，设置后就可以接收报警消息了。



接下来,设置接受报警用户的电子邮件



点击：管理->报警媒介类型->Admin->报警媒介->添加



![img](assets/03-企业级Zabbix监控平台/1683012547238-d33c9d5e-38d4-46d1-8618-1d56abced7d8.png)



此时：



我们的node2节点上，有监控80端口的监控项，有关联到此监控项的触发器，还缺1个动作



#### 创建动作



![img](assets/03-企业级Zabbix监控平台/1683012547547-47dbbab5-07f8-404b-a5c2-cb2b33df8a03.png)



![img](assets/03-企业级Zabbix监控平台/1683012547292-76bb707a-71fa-4099-ba92-8ebdd406d8d2.png)



![img](assets/03-企业级Zabbix监控平台/1683012547716-cc4b5615-c405-470f-8690-f911c6ada043.png)



![img](assets/03-企业级Zabbix监控平台/1683012547876-a916f541-6a01-433e-ae8d-1ba98f6edb9b.png)



![img](assets/03-企业级Zabbix监控平台/1683012548903-b4c1fe4d-b6e8-441b-a791-22fcc4b33c78.png)



#### 测试



在被监控端操作：



```plain
[root@zabbix-agent-node2 ~]# systemctl stop httpd
```



可以看到，Zabbix检测到问题之后，会执行动作



![img](assets/03-企业级Zabbix监控平台/1683012549397-05ba8577-718e-4378-9d3f-05c66518079f.png)



报警邮件接收成功：



![img](assets/03-企业级Zabbix监控平台/1683012549948-482b93df-9954-4b44-be80-412b380da7fa.png)



```plain
[root@zabbix-agent-node2 ~]# systemctl start httpd
```



可看到，问题已恢复



![img](assets/03-企业级Zabbix监控平台/1683012549813-a8201e84-67e7-4c05-b2ae-fe95481216a1.png)



恢复邮件接收成功



![img](assets/03-企业级Zabbix监控平台/1683012550512-2eccdedd-0fd6-44d0-906d-bd6f261441be.png)



## Zabbix企业微信自动监控报警



这次我们来监控node2节点的redis服务，来呈现项目效果；



1、被监控节点node2



2、node2节点 监控项（监控redis6379端口）



3、node2节点触发器（关联到“redis6379端口”）



#### 被监控节点node2下载redis



```plain
[root@zabbix-agent-node2 ~]# yum -y install epel-release ; yum -y install redis
[root@zabbix-agent-node2 ~]# vim /etc/redis.conf
```



允许所有人来访问，这样zabbix监控端才能检测到redis 6379端口是否存活



![img](assets/03-企业级Zabbix监控平台/1683012550452-16fe7d74-c9c6-4c1a-8e86-b7f1e8f95afa.png)



```plain
[root@zabbix-agent-node2 ~]# systemctl start redis
```



#### 创建对应监控项



![img](assets/03-企业级Zabbix监控平台/1683012551308-f9b5d752-0285-40d5-8b2d-a66903971ea5.png)



#### 创建对应触发器



![img](assets/03-企业级Zabbix监控平台/1683012551522-91359ec3-22a2-4a10-aa2c-ef2ce599c4a9.png)



#### 测试触发器效果



被监控node2关闭redis服务



![img](assets/03-企业级Zabbix监控平台/1683012552055-5bec7a78-f47d-432a-bd76-ae266d56c497.png)



#### zabbix 微信报警【监控端】



手机端下载“企业微信”



![img](assets/03-企业级Zabbix监控平台/1683012552555-0cef9d49-213c-4c3f-a17b-41d9f407d9fc.png)



![img](assets/03-企业级Zabbix监控平台/1663845350802-8a39d58b-eff2-417a-9838-f4930455d0e2.png)



#### 企业登录后台【pc】



![img](assets/03-企业级Zabbix监控平台/1683012552611-b7886dba-dcc2-4b8e-a17e-caf347fdf6fd.png)



![img](assets/03-企业级Zabbix监控平台/1663845550005-7b9c24da-0aaf-4d5a-ae81-17d4bea7a13a.png)



![img](assets/03-企业级Zabbix监控平台/1663845593107-70b3bbca-002f-496b-8d53-662cc4075708.png)



![img](assets/03-企业级Zabbix监控平台/1683012552835-473380bb-53fd-4518-8bc5-046e87464f99.png)



![img](assets/03-企业级Zabbix监控平台/1683012553025-ef78ac89-add1-462b-b2a6-d36f58ed76a2.png)



![img](assets/03-企业级Zabbix监控平台/1683012553756-d7d8adb5-eae5-4316-a0f1-f9ca84bc414e.png)



![img](assets/03-企业级Zabbix监控平台/1683012553907-7007be96-8256-4ffe-9b55-8bb123b4ec19.png)



![img](assets/03-企业级Zabbix监控平台/1683012554216-7128bf54-65c7-4147-84f5-ef0459a1592a.png)



![img](assets/03-企业级Zabbix监控平台/1683012554336-c769238f-dea2-4a72-8d90-9d5481263330.png)



```plain
应用：报警机器人
AgentID 				1000021
Secret 					Fir-nxf0G2_TwO8woWXOWOPFWkaeECh3AfnfIJwubgk
企业 CorpID     wwd5348195e1cdd809
部门id 		   		1
```



#### 测试报警机器人



![img](assets/03-企业级Zabbix监控平台/1683012554413-a707d8b6-623f-43e4-b4ce-0892617e664e.png)



![img](assets/03-企业级Zabbix监控平台/1670464490927-30bb2032-ef39-40e8-9672-e2ce646ac715.png)



![img](assets/03-企业级Zabbix监控平台/1663846081634-8214ff3d-a321-4667-8097-36477018e1b1.png)

说明企业微信机器人可正常接受消息；

#### 企业微信可用ip问题

![img](assets/03-企业级Zabbix监控平台/1670464616073-fdf7f7ab-d13b-4ef7-ab33-e7a1493497b7.png)

![img](assets/03-企业级Zabbix监控平台/1670464644030-83092811-258d-47ac-a061-028d427358f7.png)

![img](assets/03-企业级Zabbix监控平台/1670464715397-435bf27c-988a-4292-b03c-9444b1ff3daf.png)



![img](assets/03-企业级Zabbix监控平台/1670406839546-394239cc-b55c-46ed-bc58-70131832b7a5.png)

![img](assets/03-企业级Zabbix监控平台/1670407898827-3b936ee8-c75b-40e2-85ca-a5c8d132631a.png)



![img](assets/03-企业级Zabbix监控平台/1670406948514-0cece3d0-f2e2-4f62-8d8f-8fe77020bfb5.png)

![img](assets/03-企业级Zabbix监控平台/1670407967536-48690029-d502-483b-b264-984f6b284440.png)

![img](assets/03-企业级Zabbix监控平台/1670407985816-c8449fc8-b3d1-4a46-8367-0ec763c9edfd.png)

![img](assets/03-企业级Zabbix监控平台/1670408637041-135606f8-3ca6-4db8-a95f-58d09f65c385.png)

![img](assets/03-企业级Zabbix监控平台/1670408593159-bff50cfc-3bfe-47e2-9585-ca6c06cf26e6.png)

![img](assets/03-企业级Zabbix监控平台/1670408876557-c7b2e873-244c-4ba1-8a00-c919ca517204.png)

![img](assets/03-企业级Zabbix监控平台/1670408890239-d25262d9-f4dc-4a37-a51e-4a4b80ff9cf0.png)



#### Python报警监控脚本



```plain
需要修改脚本中的四个信息：
self.__corpid = 'wwd5348195e1cdd809' //公司的corpid
self.__secret = '2QvlfpUxh4k-JeIuxVNmkh2N7ijfkCs1lzb4Tkgr6xQ' //应用的secret
'toparty':1, //部门id
'agentid':"1000017", //应用id
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



![img](assets/03-企业级Zabbix监控平台/1683012555100-3573a808-2d8b-4fba-80a6-544322e7e88b.png)



```plain
监控脚本测试：
注：ShiYuFei为上图的企业成员微信名 test为标题 双引号内为内容，标题和内容均可自定义
[root@zabbix-server alertscripts]# python wechat.py ShiYuFei test-zabbi "飞哥"
{u'invaliduser': u'', u'errcode': 0, u'errmsg': u'ok'}
```

![img](assets/03-企业级Zabbix监控平台/1663846422931-c70cd3f9-1933-470a-88cd-74bd09f06728.png)

![img](assets/03-企业级Zabbix监控平台/1663846813239-b5259ee2-1849-441c-a524-ffdc5e56c4a5.jpeg)



如果上述方式不能用

企业微信创建群

![img](assets/03-企业级Zabbix监控平台/1679488133631-e767da14-fab6-471d-b5e0-92e879ac452e.png)

```shell
#!/usr/bin/env python3
#coding:utf-8
#zabbix钉钉报警
import requests,json,sys,os,datetime
#webhook="https://oapi.dingtalk.com/robot/send?access_token=1e913cb1af600fdfb4fdf94d277dd5538ac2ea05e69e97d9051c2c30ec91208c"
# 这里替换为企业微信群机器人的webhook
webhook="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=bc616385-4c0f-4001-8e06-d5f56eab56f4"
user=sys.argv[1]
text=sys.argv[3]
data={
    "msgtype": "text",
    "text": {
        "content": text
    },
    "at": {
        "atMobiles": [
            user
        ],
        "isAtAll": False
    }
}
headers = {'Content-Type': 'application/json'}
x=requests.post(url=webhook,data=json.dumps(data),headers=headers)
if os.path.exists("/var/log/zabbix/dingding.log"):
    f=open("/var/log/zabbix/dingding.log","a+")
else:
    f=open("/var/log/zabbix/dingding.log","w+")
f.write("\n"+"--"*30)
if x.json()["errcode"] == 0:
    f.write("\n"+str(datetime.datetime.now())+"    "+str(user)+"    "+"发送成功"+"\n"+str(text))
    f.close()
else:
    f.write("\n"+str(datetime.datetime.now()) + "    " + str(user) + "    " + "发送失败" + "\n" + str(text))
    f.close()
[root@zabbix-server alertscripts]# ./dingding.py 123 123 123
```

![img](assets/03-企业级Zabbix监控平台/1679488195417-bc64f722-5fca-41d3-815c-17e4c30793b2.png)

#### zabbix 创建告警媒介



创建报警媒介



注：(以哪种方式发送报警信息,短信，脚本等等)



![img](assets/03-企业级Zabbix监控平台/1683012555234-92983139-c9b1-4214-b0e7-b42bf4360806.png)



设置报警消息模板和恢复消息模板

```plain
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



![img](assets/03-企业级Zabbix监控平台/1663847019377-49b6da0a-bdbd-431d-88b1-73e7fe749d02.png)



5.7 zabbix 添加告警用户



这里我们使用普通用户，来看如何设置，主要注意权限的问题



```plain
创建用户组：web-group，包含feige、alice...等用户
创建用户：feige、alice...等用户，并关联到微信告警
希望的结果是：发微信到组web-group，组中的用户feige、alice都能收到
```



#### 创建用户组



![img](assets/03-企业级Zabbix监控平台/1683012555561-cfa5420f-7e3e-49c9-bd75-f9c08c048865.png)



确保web-group用户组，对youngfit主机群组有读写的权限，这样web-group组



![img](assets/03-企业级Zabbix监控平台/1683012555576-edbd65b3-9678-4a72-a31e-c75c19adc3dd.png)



#### 创建用户



![img](assets/03-企业级Zabbix监控平台/1683012556402-0df78ebd-e8da-4d25-9a1d-1fff8a5d6dcc.png)



![img](assets/03-企业级Zabbix监控平台/1683012556494-a0aee4cd-8edc-48f2-a91d-b21bb2e017a7.png)



#### Zabbix 添加报警动作



![img](assets/03-企业级Zabbix监控平台/1683012556904-9fba65cd-d9fb-4ada-811b-03167e770408.png)



![img](assets/03-企业级Zabbix监控平台/1683012557019-ccb90a72-7fe3-4e61-9f48-64640375ecc1.png)



![img](assets/03-企业级Zabbix监控平台/1683012557024-98f8fbc4-177d-44e4-b637-4acd47590ea1.png)



#### 微信测试报警



被监控端操作



```plain
[root@zabbix-agent-node2 ~]# systemctl stop redis
```



![img](assets/03-企业级Zabbix监控平台/1683012557940-6371a5ed-d4c5-4f92-8daa-245eb0d9c59b.png)



```plain
[root@zabbix-agent-node2 ~]# systemctl start redis
```



![img](assets/03-企业级Zabbix监控平台/1683012558013-188806e4-0cac-4d0b-a88d-05593668721a.png)



## Zabbix 钉钉自动监控报警

下载钉钉PC端程序：https://page.dingtalk.com/wow/z/dingtalk/default/dddownload-index?from=zebra:offline

下载钉钉APP移动端程序：手机应用商店查找

#### 创建钉钉报警机器人

先组合你的小伙伴，创建钉钉群。

![img](assets/03-企业级Zabbix监控平台/1663813269875-c6fe81be-1b11-49ef-8410-f343ef6f57b5.png)

![img](assets/03-企业级Zabbix监控平台/1663758192141-c97a571c-60ae-44f9-9de1-a0feb3de7fdb.png)



![img](assets/03-企业级Zabbix监控平台/1663758169830-5ffa864c-7e43-42dd-9deb-813619992372.png)



![img](assets/03-企业级Zabbix监控平台/1683012558109-dc9856ac-2b05-494d-9793-386f7f5e6201.png)



![img](assets/03-企业级Zabbix监控平台/1683012558155-08367a1c-a030-44e9-9c46-b5fb73ec4799.png)



![img](assets/03-企业级Zabbix监控平台/1683012558389-2ea4db62-f2b0-415f-a407-91b05b38de45.png)



![img](assets/03-企业级Zabbix监控平台/1683012559295-a722718d-d882-4611-b095-34ce2c3fce1f.png)



![img](assets/03-企业级Zabbix监控平台/1683012559441-9cf6d103-fceb-4154-8cad-b84af47dc149.png)

或者用IP地址段，这里我没用

![img](assets/03-企业级Zabbix监控平台/1683012559422-a5516918-e451-41ca-bd36-4bfde35bb2ff.png)



把webhook的地址记录下来

![img](assets/03-企业级Zabbix监控平台/1663813511299-62938394-4405-46e9-b3ce-11176a0d9a84.png)

```shell
机器人的webhook身份地址：
https://oapi.dingtalk.com/robot/send?access_token=d3f1f891b45215fe63fdc9caddebf236746fd1d3642cfff2e79318e533933d43
```

#### 安装Python3环境

```shell
[root@zabbix-server ~]# yum -y install zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel gdbm-devel db4-devel libpcap-devel xz-devel libffi-devel gcc make

[root@zabbix-server ~]# wget https://www.python.org/ftp/python/3.7.4/Python-3.7.4.tgz
```

![img](assets/03-企业级Zabbix监控平台/1663813792726-529c4715-bd44-4589-9668-41bc441a88e6.png)

```shell
[root@zabbix-server ~]# mkdir -p /usr/local/python3

[root@zabbix-server ~]# tar xvzf Python-3.7.4.tgz
[root@zabbix-server ~]# cd Python-3.7.4

[root@python Python-3.7.4]# ./configure --prefix=/usr/local/python3/
[root@python Python-3.7.4]# make && make install
[root@python Python-3.7.4]# ln -s /usr/local/python3/bin/python3.7 /usr/bin/python3

[root@python Python-3.7.4]# ln -s /usr/local/python3/bin/pip3.7 /usr/bin/pip3
[root@python Python-3.7.4]# python3 -V
[root@python Python-3.7.4]# pip3 -V
```

![img](assets/03-企业级Zabbix监控平台/1663814169912-fbda5517-05e8-457c-813c-0eedc173b700.png)

#### 创建报警脚本



```shell
[root@zabbix-server ~]# cd /usr/lib/zabbix/alertscripts
[root@zabbix-server alertscripts]# vim dingding.py
#!/usr/bin/env python3
#coding:utf-8
#zabbix钉钉报警
import requests,json,sys,os,datetime
webhook="https://oapi.dingtalk.com/robot/send?access_token=d3f1f891b45215fe63fdc9caddebf236746fd1d3642cfff2e79318e533933d43"
user=sys.argv[1]
text=sys.argv[3]
data={
    "msgtype": "text",
    "text": {
        "content": text
    },
    "at": {
        "atMobiles": [
            user
        ],
        "isAtAll": False
    }
}
headers = {'Content-Type': 'application/json'}
x=requests.post(url=webhook,data=json.dumps(data),headers=headers)
if os.path.exists("/var/log/zabbix/dingding.log"):
    f=open("/var/log/zabbix/dingding.log","a+")
else:
    f=open("/var/log/zabbix/dingding.log","w+")
f.write("\n"+"--"*30)
if x.json()["errcode"] == 0:
    f.write("\n"+str(datetime.datetime.now())+"    "+str(user)+"    "+"发送成功"+"\n"+str(text))
    f.close()
else:
    f.write("\n"+str(datetime.datetime.now()) + "    " + str(user) + "    " + "发送失败" + "\n" + str(text))
    f.close()
[root@zabbix-server alertscripts]# chmod +x dingding.py
```

安装python-pip和requests库



```shell
[root@zabbix-server alertscripts]# yum -y install epel-release
[root@zabbix-server alertscripts]# yum -y install python-pip
[root@zabbix-server alertscripts]# pip3 install requests
```



![img](assets/03-企业级Zabbix监控平台/1679472302555-0d6f702a-1b07-4641-9ced-486fbfa8b4ee.png)

```plain
[root@zabbix-server ~]# touch /var/log/zabbix/dingding.log
[root@zabbix-server ~]# chown zabbix.zabbix /var/log/zabbix/dingding.log
```

#### 测试脚本是否可行



```shell
[root@zabbix-server alertscripts]# ./dingding.py test  test  "问题"
```



![img](assets/03-企业级Zabbix监控平台/1663758100288-3d800a61-ecbc-4f37-954e-71f53930dc4d.png)



pc端接收成功：



![img](assets/03-企业级Zabbix监控平台/1663758074344-51cb49f5-ede0-452e-974c-69c055e7f53f.png)



移动端App接收成功：



![img](assets/03-企业级Zabbix监控平台/1663814540968-de77e269-563c-49d6-ba62-610a0618a4b5.jpeg)



#### 创建新的报警媒介



![img](assets/03-企业级Zabbix监控平台/1663756295596-2babdf1c-3887-4285-bb83-d97ca26e9ca8.png)



```plain
主机: {HOST.NAME1}
时间: {EVENT.DATE} {EVENT.TIME}
级别: {TRIGGER.SEVERITY}
触发: {TRIGGER.NAME}
详情: {ITEM.NAME1}:{ITEM.KEY1}:{ITEM.VALUE1}
状态: {TRIGGER.STATUS}
项目：{TRIGGER.KEY1}
事件ID：{EVENT.ID}
```



![img](assets/03-企业级Zabbix监控平台/1663814820058-f1341b41-30bd-44e1-9ffb-21ed5f824972.png)



![img](assets/03-企业级Zabbix监控平台/1663814856413-477ddc47-bd3c-49f4-9d0a-09eea5d5f160.png)

![img](assets/03-企业级Zabbix监控平台/1663814874430-2719c9e6-ce91-4821-be7d-48e3b7976f1c.png)

#### 关联媒介



将feige用户关联到此媒介：

![img](assets/03-企业级Zabbix监控平台/1663815400494-1dd719fa-651c-4567-903b-72cbafdc49a7.png)

![img](assets/03-企业级Zabbix监控平台/1663815437566-162713d9-e6b3-4bb1-8b91-25797ca5fb13.png)

![img](assets/03-企业级Zabbix监控平台/1663815462301-e8cac75b-b2b5-43e5-a3f3-fa959bab9cfa.png)



#### 创建监控项

监听的仍然是82端口号

![img](assets/03-企业级Zabbix监控平台/1663815079994-3c2296f5-60d3-42b9-ad12-b0a0999fe81f.png)

然后点击下面的"添加"

#### 创建触发器

![img](assets/03-企业级Zabbix监控平台/1663815230337-4a6ee8b7-486b-4d7c-b315-c891a3531b98.png)

#### 设置动作



这里直接更改一下动作里面的媒介，看效果更快捷



![img](assets/03-企业级Zabbix监控平台/1663815311689-afc121a0-9656-466b-a781-a93ff2a16955.png)



![img](assets/03-企业级Zabbix监控平台/1683012559570-53028140-0e33-477b-894a-1ab0ef6875e9.png)



![img](assets/03-企业级Zabbix监控平台/1683012559625-93d0c46d-f452-4566-904e-7be74e8b820a.png)



#### 钉钉报警测试



```plain
[root@zabbix-agent-node2 ~]# systemctl stop nginx
```



PC端接收成功：



![img](assets/03-企业级Zabbix监控平台/1683012560301-7457e3ad-d2ba-445e-b19a-c3864603dc30.png)



移动端APP接收成功：



![img](assets/03-企业级Zabbix监控平台/1683012560596-a5f7d257-c322-4e01-b463-c301b1f8ec7a.png)



```plain
[root@zabbix-agent-node2 ~]# systemctl start nginx
```



![img](assets/03-企业级Zabbix监控平台/1683012560736-c405ae23-2709-4b1b-945e-aeb881d03526.png)



![img](assets/03-企业级Zabbix监控平台/1683012561602-0f6adeff-5ca3-488c-99c5-9a195e0fe79d.png)