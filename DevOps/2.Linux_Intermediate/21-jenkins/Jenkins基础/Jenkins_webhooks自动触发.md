# Jenkins webhooks自动触发

# Jenkins+Gitlab webhook触发自动构建项目

- 效果：只要Gitlab仓库代码更新，Jenkins自动拉取代码，自动完成构建任务。无需手动点击“立即构建”或者"参数化构建"
- 需求场景：

1、项目代码的更新迭代较多，运维很有可能不在场，每次点击比较麻烦

2、更新的可能不是代码，可能是一些资源（比如：静态文件等）

Jenkins版本：2.303.1 Gitlab版本：12.6.3

### 安装配置Gitlab

Yum安装即可 ， 过程 略

Gitlab平台root用户密码配置为12345678

##### 创建一个项目(私有仓库)

准备测试代码：

![img](assets/Jenkins_webhooks自动触发/1652339450086-a5163726-3591-49fb-ab72-99eaedb19147.png)

### 后端服务器准备

```shell
[root@docker-server ~]# tar -xvzf jdk-8u211-linux-x64.tar.gz  -C /usr/local/
[root@docker-server ~]# cd /usr/local/
[root@docker-server local]# mv jdk1.8.0_211/ java 
[root@docker-server local]# vim /etc/profile
#最文件最后面添加
JAVA_HOME=/usr/local/java
PATH=$JAVA_HOME/bin:$PATH

[root@docker-server local]# source /etc/profile
[root@docker-server local]# java -version
```

![img](assets/Jenkins_webhooks自动触发/1668657965540-c9e2d959-1187-43df-b9ef-ddd1bee4cede.png)



```shell
[root@docker-server ~]# tar -xvzf apache-tomcat-8.5.45.tar.gz -C /data/application/
[root@docker-server application]# mv apache-tomcat-8.5.45/ tomcat
[root@docker-server application]# ls
tomcat
[root@docker-server application]# rm -rf tomcat/webapps/*
```

### 安装配置Jenkins



------



安装过程略

1.安装jdk

2.安装Tomcat

3.安装Maven（可选，不确定是否编译）

4.配置环境变量

5.启动



------



##### 记得配置jdk和maven

![img](assets/Jenkins_webhooks自动触发/1652339450493-8b5749b0-cde5-4aab-b5f2-1a6ef40a9909.png)

![img](assets/Jenkins_webhooks自动触发/1652339450471-4e2975d3-907d-48a0-8ab3-7cede4a8e63e.png)

##### 安装Gitlab hooks plugins插件

因为要用gitlab hook自动拉取代码的功能，需要安装GItlab hooks插件，才具有自动构建的功能

去“插件管理”页面，“可选插件”，搜索“Gitlab Hook Plugin”，“Gitlab”，点击“直接安装即可”

**注意：安装过程和网络有关。网络必须顺畅。且能正常连同国外Jenkins网站，才能下载成功**

![img](assets/Jenkins_webhooks自动触发/1652339451257-03a9e617-4080-487c-9ef3-6eb31a83e03b.png)

##### 新建Gitlab webhook相关项目

![img](assets/Jenkins_webhooks自动触发/1652339451477-659d9eb1-cbeb-4103-aa9c-9148c97deefb.png)

![img](assets/Jenkins_webhooks自动触发/1652339451581-049775dc-c2d3-4896-aa6b-a9a9555035c4.png)

Jenkins具体配置

![img](assets/Jenkins_webhooks自动触发/1652339451724-29d99afa-cd2c-4804-afc9-30bc3fc183e6.png)

来到Gitlab的test1项目中，复制拉取地址

![img](assets/Jenkins_webhooks自动触发/1652339451783-1ad11f51-3850-4b7b-9172-aa55d79c1202.png)

粘贴到

![img](assets/Jenkins_webhooks自动触发/1652339452308-7ab21409-79fc-4f1b-a0c4-1b43dbf958cd.png)

出现一堆报红，正常！因为需要配置私钥和公钥

需要把Jenkins服务器的私钥，配置到test1项目中。把Jenkins服务器的公钥，配置到GItlab的服务里面。

这样拉取就可以免密了！

![img](assets/Jenkins_webhooks自动触发/1652339452330-0558eb84-97cd-4798-9960-6a7361e52a94.png)

```shell
[root@jenkins-server ~]# useradd jenkins 
[root@jenkins-server ~]# su - jenkins 
[jenkins@jenkins-server ~]$ ssh-keygen 
[jenkins@jenkins-server ~]$ cat .ssh/id_rsa #查看jenkins用户的私钥
```



![img](assets/Jenkins_webhooks自动触发/1652339452630-bcb3e967-f73e-491c-9560-6548e57cf3d7.png)

![img](assets/Jenkins_webhooks自动触发/1652339452704-d7fadc16-7915-4eb9-92f4-0fa260f2a263.png)

看到仍然报红，将jenkins服务器上面的jenkins用户的公钥添加到gitlab中

![img](assets/Jenkins_webhooks自动触发/1652339452746-df4fd441-7651-4747-afe0-ecd535e7b1f5.png)

![img](assets/Jenkins_webhooks自动触发/1652339453366-abb0080e-1cb3-4b37-9a9b-dd6ec542138a.png)

![img](assets/Jenkins_webhooks自动触发/1652339453499-a98b031a-9d85-483f-8ed2-d4d31a6d3ed7.png)

登录到jenkins服务器中

```shell
[jenkins@jenkins-server ~]$ cat .ssh/id_rsa.pub #查看jenkins用户的公钥
```



![img](assets/Jenkins_webhooks自动触发/1652339453533-430970e4-ad4f-4de0-be5a-222b39987336.png)

![img](assets/Jenkins_webhooks自动触发/1652339453540-f7318326-a77c-453b-bd24-2d91e9e77fd3.png)

![img](assets/Jenkins_webhooks自动触发/1652339454024-f8f61b4c-ec46-417f-bd59-a3a89927a1bc.png)

##### 构建触发器

![img](assets/Jenkins_webhooks自动触发/1652339454580-8f2e6f6a-b902-4997-80a7-720630d9cede.png)

![img](assets/Jenkins_webhooks自动触发/1652339454593-8550acbd-e0b9-45e1-9329-2ec9ab8053f0.png)

![img](assets/Jenkins_webhooks自动触发/1652339454878-bf29d7bb-25fa-4516-b561-c48c10b08a3e.png)

![img](assets/Jenkins_webhooks自动触发/1652339455074-c91f4733-2eb4-4b78-92e7-d6b59387f49b.png)

![img](assets/Jenkins_webhooks自动触发/1671003567422-b95fb9a9-7222-49c5-85df-bf3d2384a5b8.png)

**要记录下上边的URL和认证密钥，切换到gitlab，找到对应的git库点击setting --> webhook ,填写以下内容**

地址：http://192.168.182.130:8080/jenkins/project/test3

Secret token：aab74fc9bd2427f6989e7a2b8ab9b178

##### 配置GItlab

![img](assets/Jenkins_webhooks自动触发/1652339455002-72def8cc-e155-4be7-96a1-d48adeb01e86.png)

添加完成之后报错

![img](assets/Jenkins_webhooks自动触发/1652339455454-3c93b938-d71d-402d-95c2-dae1ad91ae79.png)

这是因为gitlab 10.6 版本以后为了安全，不允许向本地网络发送webhook请求，设置如下：

登录管理员账号

![img](assets/Jenkins_webhooks自动触发/1652339455584-c1be02e3-5c38-4738-a6b0-1ad3ab6fbd76.png)

![img](assets/Jenkins_webhooks自动触发/1652339455884-c2e843d0-4e5e-4f49-9569-24e0d09c2d69.png)

![img](assets/Jenkins_webhooks自动触发/1652339456434-7d2a65a8-2225-4c29-b488-b627e60e2551.png)

然后需要再次添加webhook，就会成功了。

![img](assets/Jenkins_webhooks自动触发/1652339456536-869ac280-ee9d-4e45-98bf-dc4ac782a015.png)

成功了，才会显示出来

![img](assets/Jenkins_webhooks自动触发/1652339456579-e153c650-644c-4d2f-9c03-ccdf1581aa00.png)

![img](assets/Jenkins_webhooks自动触发/1652339456667-b6f1ab95-d289-4ec7-a3ad-b5399ba20057.png)

**回到jenkins页面**

![img](assets/Jenkins_webhooks自动触发/1652339456889-488c1d9d-e4c1-4d63-b868-c8264ab8ee61.png)

注意：Jenkins需要配置git用户名 和 邮箱地址

```shell
[root@git-client ~]# su - jenkins
[root@git-client ~]# git config --global user.email "soho@163.com"
[root@git-client ~]# git config --global user.name "soho"
```

##### 开始测试

在任何一台测试都可以。我这里在gitlab机器上面测试： 

```shell
[root@git-client ~]# yum -y install git
[root@git-client ~]# git config --global user.email "fei@163.com"
[root@git-client ~]# git config --global user.name "fei"
[root@git-client ~]# ssh-keygen #生成秘钥 
[root@git-client ~]# cat .ssh/id_rsa.pub #查看生成的公钥添加到gitlab里面去
```

![img](assets/Jenkins_webhooks自动触发/1652339457522-e0dc8b3c-4cd9-4dec-81ff-9526fca3a1f4.png)

先克隆一下仓库 

```shell
[root@git-client ~]# git clone git@192.168.91.168:root/cloudweb.git
[root@git-client ~]# cd cloudweb
[root@gitlab-server cloudweb]# vi src/main/webapp/index.jsp
```

![img](assets/Jenkins_webhooks自动触发/1652339457611-c7cd98b1-81b3-43c5-b131-6f7150447207.png)

```shell
[root@git-client cloudweb]# git add *
[root@git-client cloudweb]# git commit -m "用户名 和 密码"
[root@git-client cloudweb]# git push origin main
```



**返回到jenkins页面查看是否自动发布**

![img](assets/Jenkins_webhooks自动触发/1652339457837-da51c7ad-d30b-47df-a527-33a19f6ec5f1.png)

![img](assets/Jenkins_webhooks自动触发/1652339457884-662e646b-17aa-4da0-b8d1-80a81bacf6c4.png)

```shell
[root@jenkins easy-springmvc-maven]# cd /root/.jenkins/workspace/test3
[root@jenkins test3]# ls
pom.xml  README.md  src  target
```

访问tomcat页面，验证：

![img](assets/Jenkins_webhooks自动触发/1652339457945-39707c2f-1640-4b42-b939-da12b60842f4.png)

##### 基于Git参数化自动构建项目

Gitlab仓库中有测试项目：github同步过来的

![img](assets/Jenkins_webhooks自动触发/1652339458793-75860ac5-6ced-48ad-88ba-255ae90736ce.png)

![img](assets/Jenkins_webhooks自动触发/1652339458960-29b3f694-a3e6-4287-be24-1836ead55d1f.png)

Jenkins配置：

修改原来的"test3"项目：

![img](assets/Jenkins_webhooks自动触发/1652339458921-bb9fd087-b3e9-42e1-8b26-6a5015b4497a.png)

![img](assets/Jenkins_webhooks自动触发/1671004781816-03fd5b23-070b-4ab1-8659-51181c036fb5.png)

![img](assets/Jenkins_webhooks自动触发/1652339458944-e8f8d880-5157-4679-a864-12b290e77111.png)

![img](assets/Jenkins_webhooks自动触发/1652339459581-8f071137-bb20-49df-81e1-14b017130b17.png)

![img](assets/Jenkins_webhooks自动触发/1652339459775-2f398091-f525-43dd-9bc3-9c3b614f242e.png)

测试：

推送代码，打tag。代码也会自动构建；