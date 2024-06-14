# jenkins

### Jenkins构建CI/CD

什么是CI/CD：持续集成/持续发布---开发(git) -->git主库-->jenkins(git+jdk+tomcat+maven打包+测试）-->发布到tomcat服务器

Jenkins是一个功能强大的应用程序，允许持续集成和持续交付项目，无论用的是什么平台。这是一个免费的源代码，可以处理任何类型的构建或持续集成。集成Jenkins可以用于一些测试和部署技术。Jenkins是一种软件.

官网地址：jenkins.io

### 为什么要 CI / CD 方法简介

软件开发的连续方法基于自动执行脚本，以最大限度地减少在开发应用程序时引入错误的可能性。从新代码的开发到部署，它们需要较少的人为干预甚至根本不需要干预。

它涉及在每次小迭代中不断构建，测试和部署代码更改，从而减少基于有缺陷或失败的先前版本开发新代码的机会。

这种方法有三种主要方法，每种方法都根据最适合您的策略进行应用。

持续集成(Continuous Integration, CI):  代码合并，构建，部署，测试都在一起，不断地执行这个过程，并对结果反馈。

持续部署(Continuous Deployment, CD): 部署到测试环境、预生产环境/灰度环境、生产环境。

持续交付(Continuous Delivery, CD):  将最终产品发布到生产环境、给用户使用。

### 一、jenkins介绍

Jenkins是帮我们将代码进行统一的编译打包、还可以放到tomcat容器中进行发布。我们通过配置，将以前：编译、打包、上传、部署到Tomcat中的过程交由Jenkins，Jenkins通过给定的代码地址URL（代码仓库地址），将代码拉取到其“宿主服务器”（Jenkins的安装位置），进行编译、打包和发布到Tomcat容器中。

##### 1、Jenkins概述

是一个开源的、提供友好操作界面的持续集成(CI)工具，主要用于持续、自动的构建的一些定时执行的任务。Jenkins用Java语言编写，可在Tomcat等流行的容器中运行，也可独立运行。

jenkins通常与版本管理工具、构建工具结合使用；常用的版本控制工具有SVN、gitlab。jenkins构建工具有Maven、Ant、Gradle。

##### 2、Jenkins目标

① 持续、自动地构建/测试软件项目。

② 监控软件开发流程，快速问题定位及处理，提高开发效率。

##### 3、Jenkins特性

```plain
① 开源的java语言开发持续集成工具，支持CI，CD。
② 易于安装部署配置：可通过yum安装,或下载war包以及通过docker容器等快速实现安装部署，可方便web界面配置管理。
③ 消息通知及测试报告：集成RSS/E-mail通过RSS发布构建结果或当构建完成时通过e-mail通知，生成JUnit/TestNG测试报告。
④ 分布式构建：支持Jenkins能够让多台计算机一起构建/测试。
⑤ 文件识别:Jenkins能够跟踪哪次构建生成哪些jar，哪次构建使用哪个版本的jar等。
⑥ 丰富的插件支持:支持扩展插件，你可以开发适合自己团队使用的工具，如git，ssh，svn，maven，docker等。
```

工作流程图:

![img](assets/Jenkins/1683013580490-fe20a5cb-765f-454f-beaa-b10fb2a14cd2.png)

```plain
测试环境中：
1.开发者会将代码上传到版本库中。
2.jenkins通过配置版本库的连接地址，获取到源代码。
3.jenkins获取到源代码之后通过参数化构建(或者触发器)开始编译打包。
4.jenkins通过调用maven（Ant或者Gradle）命令实现编译打包过程。
5.生成的war包通过ssh插件上传到远程tomcat服务器中通过shell脚本自动发布项目。

生产环境：
测试环境将项目测试没问题后，将项目推送到线上正式环境。
1.可以选择手动。
2.也可以通过调用脚本推送过去。
```

#### 4、产品发布流程

产品设计成型 -> 开发人员开发代码 -> 测试人员测试功能 -> 运维人员发布上线

持续集成（Continuous integration，简称CI）

持续部署（continuous deployment）

持续交付（Continuous delivery）

![img](assets/Jenkins/1683013580492-8d01f6b3-8fcf-42eb-9791-d2b3198247f6.png)

### 二、部署应用Jenkins+Github+Tomcat实战

准备环境:

两台机器

git-server    ----https://github.com/bingyue/easy-springmvc-maven

jenkins-server    ---192.168.246.212---最好是3个G以上

java-server   -----192.168.246.210

https://github.com/bingyue/easy-springmvc-maven

### Jenkins2.303.1版本安装

### Yum安装

##### 1.配置 Jenkins的yum源

```plain
# wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
# rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key
# yum install epel-release java-11-openjdk-devel -y
# yum install jenkins -y
# systemctl daemon-reload
```

##### 2.启动Jenkins

```plain
# systemctl start jenkins
# systemctl status jenkins
[root@jenkins yum.repos.d]# tail -f /var/log/jenkins/jenkins.log
```

![img](assets/Jenkins/1683013580276-c934d8f8-4e84-4f2a-8326-3adcec1df9e6.png)

##### 3.访问登录

当您第一次访问一个新的 Jenkins 实例时，系统会要求您使用自动生成的密码将其解锁

1.浏览到http://localhost:8080（或您在安装时为 Jenkins 配置的任何端口）并等待解锁 Jenkins页面出现

![img](assets/Jenkins/1683013580274-2dd61155-7887-4c2d-949e-aa0e16097987.png)

2.从 Jenkins 控制台日志输出中，复制自动生成的字母数字密码（在 2 组星号之间）。

![img](assets/Jenkins/1683013579994-bfdb0c72-fff4-4056-8cc6-9101519ac5e1.png)

3.使用插件自定义 Jenkins  或者  推荐安装插件，荐安装插件。。。

![img](assets/Jenkins/1683013581217-703d6e6e-8842-41ef-b2a0-b27f2e2c9645.png)

![img](assets/Jenkins/1683013581970-b90b138c-0daf-478a-bb04-54af7f4b0adc.png)

##### 4.创建第一个管理员用户

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332924438-dc5d5155-6730-4e09-a480-a931d1f06c15.png#averageHue=%23fcfafa&crop=0&crop=0&crop=1&crop=1&from=url&height=231&id=gfHve&originHeight=923&originWidth=1354&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=&width=339)

![img](assets/Jenkins/1683013582311-b2db59e1-49ec-47da-bb27-de1e7a026f9a.png)

![img](assets/Jenkins/1683013582807-7f3a0010-4685-4320-bc09-cee925d903d1.png)

### War包安装

##### 1.下载安装包

百度搜索openjdk11、tomcat、maven、jenkins

![img](assets/Jenkins/1683013582981-ef63ca9c-ff49-469f-98a9-40ec542d887f.png)

![img](assets/Jenkins/1683013583226-5b653705-ecab-4ab2-af9d-c4772e240ba1.png)

![img](assets/Jenkins/1683013585196-5de1d590-7b43-4cb0-bf16-6b2eaffb87dc.png)

![img](assets/Jenkins/1683013584944-e157b6d6-d91a-4aad-87ef-cd3804dee6de.png)

```plain
[root@jenkins ~]# wget https://download.java.net/openjdk/jdk11/ri/openjdk-11+28_linux-x64_bin.tar.gz
[root@jenkins ~]# wget https://get.jenkins.io/war/2.303/jenkins.war
[root@jenkins ~]# wget https://downloads.apache.org/maven/maven-3/3.8.2/binaries/apache-maven-3.8.2-bin.tar.gz
[root@jenkins ~]# wget https://dlcdn.apache.org/tomcat/tomcat-8/v8.5.70/bin/apache-tomcat-8.5.70.tar.gz
还有openjdk11
[root@jenkins ~]# cd /usr/local
[root@jenkins local]# tar -xvzf apache-maven-3.8.2-bin.tar.gz
[root@jenkins local]# tar -xvzf apache-tomcat-8.5.70.tar.gz
[root@jenkins local]# tar -xvzf openjdk-11+28_linux-x64_bin.tar.gz
[root@jenkins local]# mv jdk-11/ java
[root@jenkins local]# mv apache-tomcat-8.5.70 tomcat
[root@jenkins local]# rm -rf tomcat/webapps/*
[root@jenkins local]# mv apache-maven-3.8.2 java/maven
[root@jenkins ~]# cp jenkins.war  /usr/local/tomcat/webapps/
```

##### 2.配置环境变量

```plain
[root@jenkins ~]# vim /etc/profile
JAVA_HOME=/usr/local/java
MAVEN_HOME=/usr/local/java/maven
PATH=$PATH:$JAVA_HOME/bin:$MAVEN_HOME/bin
export PATH USER LOGNAME MAIL HOSTNAME HISTSIZE HISTCONTROL JAVA_HOME MAVEN_HOME

[root@jenkins ~]# java -version
openjdk version "11.0.12" 2021-07-20 LTS
OpenJDK Runtime Environment 18.9 (build 11.0.12+7-LTS)
OpenJDK 64-Bit Server VM 18.9 (build 11.0.12+7-LTS, mixed mode, sharing)

[root@jenkins ~]# mvn -v
Apache Maven 3.8.2 (ea98e05a04480131370aa0c110b8c54cf726c06f)
Maven home: /usr/local/java/maven
Java version: 11, vendor: Oracle Corporation, runtime: /usr/local/java
Default locale: en_US, platform encoding: UTF-8
OS name: "linux", version: "3.10.0-693.el7.x86_64", arch: "amd64", family: "unix"

[root@jenkins ~]# /usr/local/tomcat/bin/startup.sh
```

补充：如果启动访问报错

请更换jdk版本为1.8的，修改环境变量配置，重新启动即可；

##### 3.访问登录

![img](assets/Jenkins/1683013584750-f04b162d-e9c7-4014-8e9b-23fd92f608af.png)

![img](assets/Jenkins/1683013584840-4242d275-db1c-4040-b1a4-886bbfac6020.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332924438-dc5d5155-6730-4e09-a480-a931d1f06c15.png#averageHue=%23fcfafa&clientId=uc546fde2-80e0-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=462&id=u4433142a&name=image.png&originHeight=923&originWidth=1354&originalType=binary&ratio=1&rotation=0&showTitle=false&size=43751&status=done&style=none&taskId=u7ea6aebf-f84c-4a84-b521-671d5d400e7&title=&width=677)

##### 4.插件安装

```plain
安装插件:
所需的插件:
• Maven插件 Maven Integration plugin
• 发布插件 Deploy to container Plugin
需要安装插件如下：
=====================================================================================
安装插件
Deploy to container    ---支持自动化代码部署到tomcat容器
GIT plugin  可能已经安装,可在已安装列表中查询出来
Maven Integration   jenkins利用Maven编译，打包所需插件
Publish Over SSH  通过ssh连接
ssh  插件
安装过程:
系统管理--->插件管理---->可选插件--->过滤Deploy to container---->勾选--->直接安装
```

![img](assets/Jenkins/1683013585356-e878a23a-9b0d-47f7-8c99-831a2134f584.png)

![img](assets/Jenkins/1683013586832-b407cf08-81cc-4739-9661-b496a2a83d5a.png)

![img](assets/Jenkins/1683013587576-10b5e95c-479f-43f3-8654-132b9030e1b2.png)

![img](assets/Jenkins/1683013587586-07e4ab96-e374-4de6-9125-c583c117f905.png)

![img](assets/Jenkins/1683013587413-fad8e576-f50d-431c-b427-b06a7b3c5e10.png)

![img](assets/Jenkins/1683013587656-cb73dcec-3d27-4756-8493-1d5e2bc67e4c.png)

##### 5.配置国内源

因为Jenkins下载，默认是国外地址，如果插件下载失败，我们就替换为国内地址

官方下载插件慢 更新下载地址

Jenkins 安装时会默认从updates.jenkins-ci.org 拉取，我们需要换成国内源——清华大学开源软件镜像站。

https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/update-center.json

cd    {你的Jenkins工作目录}/updates  进入更新配置位置

```plain
[root@jenkins-server1 updates]# pwd
/root/.jenkins/updates    #这是Jenkins默认的工作目录
[root@localhost updates]# vim  default.json      #修改配置文件
s/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /root/.jenkins/updates/default.json            #官方源替换清华源
s/http:\/\/www.google.com/https:\/\/www.baidu.com/g    #google替换成百度

或者直接进行一下操作（一步到位，不需要多步操作）sed -i 's@原始信息@新内容@g'
[root@localhost ~]# sed -i 's/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /root/.jenkins/updates/default.json && sed -i 's/http:\/\/www.google.com/https:\/\/www.baidu.com/g' /root/.jenkins/updates/default.json
```

之后，在网站后面加上restart进行jenkins重启。

http://192.168.153.147:8080/jenkins/restart

##### 5.邮箱配置(可选)

安装邮件插件，才能确保邮件发送成功。否则可能不会发送邮件

![img](assets/Jenkins/1683013588981-0eb3e4da-e02c-4bdf-b634-7b8e27656550.png)

![img](assets/Jenkins/1683013588974-aeb2238a-c272-41fb-9f07-d5cfb75e7f95.png)

![img](assets/Jenkins/1683013589865-e62d8a83-1f72-4f2b-9ca4-b38bab12255d.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672331291677-c334ee9f-45c3-4ab5-8648-a506abdb5c2d.png#averageHue=%23fefefe&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=306&id=ucb4338e3&name=image.png&originHeight=612&originWidth=2650&originalType=binary&ratio=1&rotation=0&showTitle=false&size=58207&status=done&style=none&taskId=ud33c5c3a-7fc1-4a4a-86d3-dc166f83e4a&title=&width=1325)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332659245-b2fd19c4-5e4d-43a8-89c1-dc68dbb7919b.png#averageHue=%23fdfbfa&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=293&id=uaaf15446&name=image.png&originHeight=1173&originWidth=2362&originalType=binary&ratio=1&rotation=0&showTitle=false&size=81487&status=done&style=none&taskId=u4f1bff0f-6329-4115-aa9a-82a0a429ebe&title=&width=591)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332123231-ed182f5a-982b-481a-8471-60db40f5f297.png#averageHue=%23fefefe&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=568&id=uca9456e8&name=image.png&originHeight=1136&originWidth=2530&originalType=binary&ratio=1&rotation=0&showTitle=false&size=91250&status=done&style=none&taskId=uc1bf973d-c646-4094-a933-ee59481ddb5&title=&width=1265)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332253363-cc823e2f-d891-42e5-9657-64a2234d107c.png#averageHue=%23fefdfd&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=613&id=u23b5d54d&name=image.png&originHeight=1225&originWidth=2576&originalType=binary&ratio=1&rotation=0&showTitle=false&size=108247&status=done&style=none&taskId=u8c7ad7fd-07e9-4aef-a3f7-3c313741e20&title=&width=1288)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672332411650-51421c85-7679-4e9c-afb9-7d80fb51444e.png#averageHue=%23fefcfc&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=298&id=ue41a2beb&name=image.png&originHeight=596&originWidth=2557&originalType=binary&ratio=1&rotation=0&showTitle=false&size=52258&status=done&style=none&taskId=u64cbf40e-1570-49e2-860e-f4c956ab653&title=&width=1278.5)

可看到邮箱确实接收到了邮件，则配置成功；KYXDIMUFCNDXTBPJ

当然邮箱能接收到邮件的前提是，邮箱要开启smtp服务

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672331243616-939a19be-f50c-492b-bbfd-7b7820fb74f1.png#averageHue=%23f0e2c2&clientId=ub5198fbe-3902-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=443&id=ucade9d70&name=image.png&originHeight=886&originWidth=2584&originalType=binary&ratio=1&rotation=0&showTitle=false&size=319372&status=done&style=none&taskId=u63812e7f-7173-4055-9f9d-77a1aca54fa&title=&width=1292)

##### 6.配置Jenkins私钥

```plain
[root@jenkins ~]# ssh-keygen
[root@jenkins ~]# cat /root/.ssh/id_rsa
```

![img](assets/Jenkins/1683013589458-095e73b7-076c-4664-be48-1b3868d512d6.png)

##### 7.添加后端服务器

```plain
公钥发送到后端服务器，才能实现免密；
[root@jenkins ~]# ssh-copy-id -i root@192.168.153.194
```

![img](assets/Jenkins/1683013589626-f48e3382-effd-44eb-8b7d-d106b5ceb952.png)

##### 8.配置JDK和Maven

虽然Jenkins服务器上，已经安装了JDK和maven工具，但是，还需要在Jenkins服务中，进行配置；

这样Jenkins才能自动化的使用两个工具；

部署多个jdk环境，应对不同代码的构建 111111111111111111111111111111

![img](assets/Jenkins/1683013590958-60ca76f9-fa76-44f3-a641-c63ce29bc9da.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/25576613/1672325715936-6d336083-5709-41a2-b077-6ec634227500.png#averageHue=%23fefdfc&clientId=ue996132e-bc3a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=137&id=u4186fc3f&name=image.png&originHeight=274&originWidth=2114&originalType=binary&ratio=1&rotation=0&showTitle=false&size=13233&status=done&style=none&taskId=uf97bf382-8424-4b17-a36b-02ebec8a7be&title=&width=1057)

![img](assets/Jenkins/1683013590996-4224f131-b212-4e30-bf42-ed58d212487e.png)

![img](assets/Jenkins/1683013590904-a8310e96-147a-4ce2-ac77-19f580167fd8.png)

![img](assets/Jenkins/1683013591070-fc404de4-15e0-46d6-b307-bc80f41b48f2.png)

如果有多个jdk和maven需要配置的话，可以点击新增jdk或者新增maven；

##### 9.构建发布任务

![img](assets/Jenkins/1683013591450-dcc36672-8edc-445d-be23-e74aa6202c15.png)

![img](assets/Jenkins/1683013592404-d9c77601-234a-4294-846a-8f80b726e4d8.png)

![img](assets/Jenkins/1683013592769-552ebc91-229f-41b0-bce4-d36daa8a0c47.png)

https://github.com/bingyue/easy-springmvc-maven

这里是我再github上，找的一个测试的小项目。能进行编译打包

![img](assets/Jenkins/1683013592901-8eec63ba-a9a0-46f9-9a31-0e5452314503.png)

![img](assets/Jenkins/1683013593088-940fb7d0-5242-4964-a7cf-b4fb7872e075.png)

如果有多个后端服务器，可以点击 ADD server进行添加；

##### 10.java服务器添加脚本

```plain
[root@java-server ~]# mkdir -p /data/application
上传jdk
[root@java-server ~]# tar xzf jdk-8u191-linux-x64.tar.gz -C /usr/local/
[root@java-server ~]# cd /usr/local/
[root@java-server local]# mv jdk1.8.0_191/ java
下载tomcat
[root@java-server ~]# wget http://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.42/bin/apache-tomcat-8.5.42.tar.gz
[root@java-server ~]# tar xzf apache-tomcat-8.5.42.tar.gz -C /data/application/
[root@java-server ~]# cd /data/application/
[root@java-server application]# mv apache-tomcat-8.5.42/ tomcat
设置环境变量
[root@java-server ~]# vim /etc/profile
export JAVA_HOME=/usr/local/java
export PATH=$JAVA_HOME/bin:$JAVA_HOME/jre/bin:$PATH
export CLASSPATH=.:$JAVA_HOME/lib:$JAVA_HOME/jre/lib:$JAVA_HOME/lib/tools.jar
export TOMCAT_HOME=/data/application/tomcat
[root@java-server ~]# source /etc/profile
测试:
[root@java-server ~]# java -version 
删除tomcat默认发布目录下面的内容:
[root@java-server ~]# rm -rf /data/application/tomcat/webapps/*
[root@java-server ~]# cd /data/application/tomcat/webapps/
[root@java-server webapps]# ls
创建目录和脚本:
[root@java-server ~]# mkdir /opt/script  #创建脚本目录
[root@java-server ~]# vim app-jenkins.sh   #创建脚本
脚本内容在后面：
[root@java-server ~]# chmod +x app-jenkins.sh  #添加执行权限
[root@java-server ~]# mv app-jenkins.sh /opt/script/
脚本内容:
[root@java-server script]# cat app-jenkins.sh 
#!/usr/bin/bash
#本脚本适用于jenkins持续集成，实现备份war包到代码更新上线！使用时请注意全局变量。
#================
#Defining variables
export JAVA_HOME=/usr/local/java
webapp_path="/data/application/tomcat/webapps"
tomcat_run="/data/application/tomcat/bin"
updata_path="/data/update/`date +%F-%T`"
backup_path="/data/backup/`date +%F-%T`"
tomcat_pid=`ps -ef | grep tomcat | grep -v grep | awk '{print $2}'`
files_dir="easy-springmvc-maven"
files="easy-springmvc-maven.war"
job_path="/root/upload"

#Preparation environment
echo "Creating related directory"
mkdir -p $updata_path
mkdir -p $backup_path

echo "Move the uploaded war package to the update directory"
mv $job_path/$files $updata_path

echo "========================================================="
cd /opt
echo "Backing up java project"
if [ -f $webapp_path/$files ];then
    tar czf $backup_path/`date +%F-%H`.tar.gz $webapp_path
    if [ $? -ne 0 ];then
        echo "打包失败，自动退出"
        exit 1
    else
        echo "Checking if tomcat is started"
        if [ -n "$tomcat_pid" ];then
            kill -9 $tomcat_pid
            if [ $? -ne 0 ];then
                echo "tomcat关闭失败，将会自动退出"
                exit 2
            fi
        fi
        cd $webapp_path
        rm -rf $files && rm -rf $files_dir
        cp $updata_path/$files $webapp_path
        cd /opt
        $tomcat_run/startup.sh
        sleep 5
        echo "显示tomcat的pid"
        echo "`ps -ef | grep tomcat | grep -v grep | awk '{print $2}'`"
        echo "tomcat startup"
        echo "请手动查看tomcat日志。脚本将会自动退出"
    fi
else
    echo "Checking if tomcat is started"
        if [ -n "$tomcat_pid" ];then
            kill -9 $tomcat_pid
                if [ $? -ne 0 ];then
                    echo "tomcat关闭失败，将会自动退出"
                        exit 2
                fi
        fi
    cp $updata_path/$files $webapp_path
    $tomcat_run/startup.sh
        sleep 5
        echo "显示tomcat的pid"
        echo "`ps -ef | grep tomcat | grep -v grep | awk '{print $2}'`"
        echo "tomcat startup"
        echo "请手动查看tomcat日志。脚本将会自动退出"
fi
```

##### 11.调用maven打包命令

![img](assets/Jenkins/1683013592993-26170951-828b-4b66-8cd7-9b0f5fb8f5da.png)

##### 12.启用邮箱

![img](assets/Jenkins/1683013593865-4553519a-588e-4230-b5e7-649d2b6ad610.png)

##### 13.构建项目

![img](assets/Jenkins/1683013594141-b9a49cbf-7a03-4c5b-89b8-d142e0387021.png)

![img](assets/Jenkins/1683013594481-8ed7e11f-5d4f-4272-b420-8c58f369c166.png)

### 三、Jenkins+Gitlab

Jenkins服务器去拉取代码。所以要下载git客户端

[root@jenkins-server ~]# yum -y install git

![img](assets/Jenkins/1683013594552-01dad18c-5781-4a85-807c-44d87e92ee34.png)

开始一个简单的项目

![img](assets/Jenkins/1683013595094-cafd1129-52ef-4e26-aaa2-9111e81a3e56.png)

![img](assets/Jenkins/1683013597200-101f274f-eafb-4c41-a5c0-e4fefb37402f.png)

![img](assets/Jenkins/1683013596551-da6bb2dc-5ce6-48d0-a403-31d064568c03.png)

![img](assets/Jenkins/1683013597009-2e5423c3-a07d-4fe4-aeee-d039602043a2.png)

![img](assets/Jenkins/1683013597199-99b013b0-2cdd-4582-b29a-04c6d5a1eb6e.png)

![img](assets/Jenkins/1683013597259-4f705a77-9f12-43d2-a49c-91f11b2a94bd.png)

![img](assets/Jenkins/1683013597950-defbe969-75f5-48f9-bd0a-1d2ece11f10e.png)

![img](assets/Jenkins/1683013598135-f35303e1-4aa1-475d-bfec-f8fd170aa015.png)

[root@jenkins ~]# cat /root/.ssh/id_rsa     //查看Jenkins服务器的私钥

![img](assets/Jenkins/1683013598906-d532f4b9-a784-4e5f-9bfc-6d7044a6b7fa.png)

![img](assets/Jenkins/1683013598950-7f66cb55-ebe7-4aca-86c5-e31a182a9193.png)

![img](assets/Jenkins/1683013599004-2d835545-55ab-4437-8838-17518f1113ae.png)

![img](assets/Jenkins/1683013599001-4a46d5e1-5627-406b-b0f1-59266d502f2e.png)

Jenkins端配置好之后，还应在gitlab端配置Jenkins服务器的公钥

在jenkins服务器上查看公钥

![img](assets/Jenkins/1683013599422-ee8c053c-3c13-4955-8649-e81c45431a56.png)

复制粘贴到gitlab

![img](assets/Jenkins/1683013600112-047c8d06-40e9-4e23-8161-dc623424b03e.png)

然后去构建项目。自动拉取代码

![img](assets/Jenkins/1683013600651-8d75a9b7-6c28-488e-8fbe-bff331ec9c1c.png)

注意看拉取到了哪个目录下

![img](assets/Jenkins/1683013600713-6a9b5fb8-cdc2-4cd4-bb9f-664466113015.png)

在Jenkins服务器上认证

在这个目录下能找到自己拉取git的项目；证明项目成功完成

```plain
[root@jenkins ~]#  ls /root/.jenkins/workspace/demo
beifen.sh
```

### 四、Jenkins git参数化构建

git参数化构建：开发人员推送代码之前，对此版本的代码·，打一个标签(tag)。我们可以认作为是此套代码的版本号。后续可以方便我们进行版本之间的切换。尤其是刚上线一套代码有问题，可以运用jenkins立即进行版本回退/切换；

首先，需要安装插件"Git Parameter"。如图

![img](assets/Jenkins/1683013602984-e8c887c8-6ec7-4799-a16b-7c17ec2e2479.png)

手动测试：

```plain
[root@gitlab ~]# git clone git@192.168.182.128:root/easy-springmvc-maven.git
[root@gitlab ~]# cd easy-springmvc-maven
[root@gitlab easy-springmvc-maven]# cd src/main/webapp/
[root@gitlab easy-springmvc-maven]# cat src/main/webapp/index.jsp
```

![img](assets/Jenkins/1683013602530-cf3f9147-b1b0-4f34-b724-bc8c3900d423.png)

```plain
[root@gitlab easy-springmvc-maven]# git add *
[root@gitlab easy-springmvc-maven]# git commit -m "修改用户为密码"
[root@gitlab easy-springmvc-maven]# git tag -a "v1.0" -m "修改用户为密码"
[root@gitlab easy-springmvc-maven]# git tag  #查看一下
v1.0
[root@gitlab easy-springmvc-maven]# git push origin v1.0
[root@gitlab easy-springmvc-maven]#
```

![img](assets/Jenkins/1683013602980-6a6d6883-fcc5-4d5a-9227-be1a2f85a776.png)

##### 配置Jenkins参数化构建

![img](assets/Jenkins/1683013602979-be6e31b1-167d-44d0-9784-34662a4f4202.png)

点击"高级"

![img](assets/Jenkins/1683013603146-bf1b4b87-2fb5-4e11-8446-1882b29b50bf.png)

![img](assets/Jenkins/1683013605357-aff0e6e8-ea8f-4bdb-a6ed-6e5b77497024.png)

![img](assets/Jenkins/1683013605213-c9a6c42e-3e2a-4188-96bd-2508365ef8ab.png)

### 五、Jenkins多节点配置

在企业里面使用Jenkins自动部署+测试平台时，每天更新发布几个网站版本,很频繁,但是对于一些大型的企业来讲，Jenkins就需要同时处理很多的任务，这时候就需要借助Jenkins多个node或者我们所说的Jenkins分布式SLAVE，今天我们带大家来学习Jenkins多实例的配置；

添加Linux平台Jenkins SLAVE配置：

1. 由于Jenkins是Java程序，添加的SLAVE客户端服务器必须安装Java JDK环境；
2. 创建远程执行Jenkins任务的用户，一般为Jenkins用户，工作目录为/home/Jenkins;
3. Jenkins服务器免秘钥登录Slave服务器或者通过用户名和密码登录；

Jenkins-master Jenkins-slave1

#### 1.添加从节点

![img](assets/Jenkins/1683013606227-970b13f1-22db-49ff-9a74-3a2036a3c938.png)

![img](assets/Jenkins/1683013605352-3bddf16f-8b2a-4cc2-970b-f8add7cd7d28.png)

![img](assets/Jenkins/1683013605351-ece16ad0-7c03-4384-b457-2d16734e3fc2.png)

#### 2.参数详解

```plain
名字：节点的名字
描述：说明这个节点的用途等
of executors:并发构建数量
远程工作目录：用于存放jenkins的工作空间的
标签：分配job会以标签的名称去分配
用法：节点的使用策略
启动方法：windows的话就不要给自己添堵了，选择 Java web start
```

![img](assets/Jenkins/1683013606260-ba4c7c60-1a76-4441-acc3-78cea9781a7d.png)

![img](assets/Jenkins/1683013606425-cd74d879-cc42-4103-8307-f00e0831e796.png)

![img](assets/Jenkins/1683013606474-ea0ae6fe-d6e5-44c3-8ec8-29d07326e1f5.png)

#### 3.指定java命令路径

![img](assets/Jenkins/1683013606995-c4cfa082-4140-467b-b8e9-3b86b0d539b6.png)

![img](assets/Jenkins/1683013607802-fbdfc0bc-d682-4fb2-a426-8a3a3dbc8278.png)

#### 4.测试从节点

项目指定到哪个节点运行。

![img](assets/Jenkins/1683013608068-2b38ab5a-a013-4108-930d-b67f6dfd1a84.png)

![img](assets/Jenkins/1683013608291-82ab9850-06ad-4a13-9a6d-291c07b2e476.png)

![img](assets/Jenkins/image-20211009224634465.png)