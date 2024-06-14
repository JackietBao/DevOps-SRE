### Jenkins构建CI/CD



**什么是CI/CD**：持续集成/持续发布---开发(git) -->git主库-->jenkins(git+jdk+tomcat+maven打包+测试）-->发布到tomcat服务器。



Jenkins是一个功能强大的应用程序，允许**持续集成和持续交付项目**，无论用的是什么平台。这是一个免费的源代码，可以处理任何类型的构建或持续集成。集成Jenkins可以用于一些测试和部署技术。Jenkins是一种软件。



### 为什么要 CI / CD 方法简介



软件开发的连续方法基于自动执行脚本，以最大限度地减少在开发应用程序时引入错误的可能性。从新代码的开发到部署，它们需要较少的人为干预甚至根本不需要干预。



它涉及在每次小迭代中不断构建，测试和部署代码更改，从而减少基于有缺陷或失败的先前版本开发新代码的机会。



这种方法有三种主要方法，每种方法都根据最适合您的策略进行应用。



**持续集成**(Continuous Integration, CI):  代码合并，构建，部署，测试都在一起，不断地执行这个过程，并对结果反馈。



**持续部署**(Continuous Deployment, CD):　部署到测试环境、预生产环境/灰度环境、生产环境。



**持续交付**(Continuous Delivery, CD):  将最终产品发布到生产环境、给互联网用户使用。



### 一、jenkins介绍



Jenkins是帮我们将代码进行统一的编译打包、还可以放到tomcat容器中进行发布。
我们通过配置，将以前：编译、打包、上传、部署到Tomcat中的过程交给Jenkins，Jenkins通过给定的代码地址URL（代码仓库地址），将代码拉取到其“宿主服务器”（Jenkins的安装位置），进行编译、打包和发布到Tomcat容器中。



##### 1、Jenkins概述



```plain
是一个开源的、提供友好操作界面的持续集成(CI)工具，主要用于持续、自动的构建的一些定时执行的任务。Jenkins用Java语言编写，可在Tomcat等流行的容器中运行，也可独立运行。
```



jenkins通常与版本管理工具(SCM)、构建工具结合使用；常用的版本控制工具有SVN、GIT。jenkins构建工具有Maven、Ant、Gradle。



##### 2、Jenkins目标



① 持续、自动地构建/测试软件项目。



② 监控软件开发流程，快速问题定位及处理，提高开发效率。



##### 3、Jenkins特性



```shell
① 开源的java语言开发持续集成工具，支持CI，CD。
② 易于安装部署配置：可通过yum安装,或下载war包以及通过docker容器等快速实现安装部署，可方便web界面配置管理。
③ 消息通知及测试报告：集成RSS/E-mail通过RSS发布构建结果或当构建完成时通过e-mail通知，生成JUnit/TestNG测试报告。
④ 分布式构建：支持Jenkins能够让多台计算机一起构建/测试。
⑤ 文件识别:Jenkins能够跟踪哪次构建生成哪些jar，哪次构建使用哪个版本的jar等。
⑥ 丰富的插件支持:支持扩展插件，你可以开发适合自己团队使用的工具，如git，svn，maven，docker等。
```



工作流程图:



![img](assets/Jenkins构建CICD/1569246908031.png)



```shell
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



##### 4、产品发布流程



产品设计成型 -> 开发人员开发代码 -> 测试人员测试功能 -> 运维人员发布上线



持续集成（Continuous integration，简称CI）



持续部署（continuous deployment）



持续交付（Continuous delivery）



![img](assets/Jenkins构建CICD/1583060962747.png)



### 二、部署应用Jenkins+Github+Tomcat实战



准备环境:



两台机器



git-server   ----https://github.com/bingyue/easy-springmvc-maven



jenkins-server   ---192.168.246.212---最好是3个G以上



java-server  -----192.168.246.210



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



![img](assets/Jenkins构建CICD/image-20210901223718598.png)



##### 3.访问登录



当您第一次访问一个新的 Jenkins 实例时，系统会要求您使用自动生成的密码将其解锁



1.浏览到`http://localhost:8080`（或您在安装时为 Jenkins 配置的任何端口）并等待**解锁 Jenkins**页面出现



![img](assets/Jenkins构建CICD/image-20210901223831661.png)



2.从 Jenkins 控制台日志输出中，复制自动生成的字母数字密码（在 2 组星号之间）。



![img](assets/Jenkins构建CICD/image-20210901223854511.png)



3.使用插件自定义 Jenkins  或者  推荐安装插件，这里飞哥使用的推荐安装插件。。。



![img](assets/Jenkins构建CICD/image-20210911094150230.png)



![img](assets/Jenkins构建CICD/image-20210911094330384.png)



##### 4.创建第一个管理员用户



![img](assets/Jenkins构建CICD/image-20210901224023748.png)



![img](assets/Jenkins构建CICD/image-20210901224037624.png)



![img](assets/Jenkins构建CICD/image-20210901224048113.png)



### War包安装



##### 1.下载安装包



百度搜索openjdk11、tomcat、maven、jenkins

openjdk11经过测试之后，不能运行当前版本jenkins。我这里换成了jdk1.8的

jdk官网：https://www.oracle.com/java/technologies/downloads/

jenkins官网：https://www.jenkins.io/

tomcat官网：https://tomcat.apache.org/

maven官网：https://maven.apache.org/



![img](assets/Jenkins构建CICD/image-20210911141241212.png)

![img](assets/Jenkins构建CICD/1652448511823-c97d9802-db51-449b-a7d6-60578acd622c.png)



![img](assets/Jenkins构建CICD/image-20210911141350166.png)



![img](assets/Jenkins构建CICD/image-20210911141450852.png)



![img](assets/Jenkins构建CICD/image-20210911141600694.png)



```shell
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



```shell
[root@jenkins ~]# vim /etc/profile
JAVA_HOME=/usr/local/java
MAVEN_HOME=/usr/local/java/maven
PATH=$PATH:$JAVA_HOME/bin:$MAVEN_HOME/bin
export PATH USER LOGNAME MAIL HOSTNAME HISTSIZE HISTCONTROL JAVA_HOME MAVEN_HOME

[root@jenkins ~]# source /etc/profile

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



![img](assets/Jenkins构建CICD/image-20210901230514289.png)



![img](assets/Jenkins构建CICD/image-20211010093156851.png)

![img](assets/Jenkins构建CICD/1670899108856-4a500590-09b9-4408-94db-c93b65170e82.png)



![img](assets/Jenkins构建CICD/1678412127199-c9a38c49-4a74-4a7f-a330-b88aa5d14902.png)

![img](assets/Jenkins构建CICD/1678412153070-90e5ae64-a266-450a-b26e-de4967709911.png)

![img](assets/Jenkins构建CICD/1678412171765-910e59ad-4c49-41b9-b1dd-c5d3c95461fa.png)

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



![img](assets/Jenkins构建CICD/image-20210902215714741.png)



![img](assets/Jenkins构建CICD/image-20210902215741814.png)



![img](assets/Jenkins构建CICD/image-20210902215826900.png)



![img](assets/Jenkins构建CICD/image-20210902215911264.png)



![img](assets/Jenkins构建CICD/image-20210902220025910.png)



![img](assets/Jenkins构建CICD/image-20210902220200043.png)

插件安装完成之后，最好重启jenkins，确保插件生效

![img](assets/Jenkins构建CICD/1678412593249-0d9afb98-0c28-4165-89bc-c582a1029c08.png)

##### 5.配置国内源

因为Jenkins下载，默认是国外地址，如果插件下载失败，我们就替换为国内地址;

官方下载插件慢 更新下载地址

Jenkins 安装时会默认从updates.jenkins-ci.org 拉取，我们需要换成国内源——清华大学开源软件镜像站;

https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/update-center.json

cd   {你的Jenkins工作目录}/updates  进入更新配置位置;

```shell
[root@jenkins-server1 updates]# pwd
/root/.jenkins/updates    #这是Jenkins默认的工作目录
[root@localhost updates]# vim  default.json      #修改配置文件
s/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /root/.jenkins/updates/default.json            #官方源替换清华源
s/http:\/\/www.google.com/https:\/\/www.baidu.com/g    #google替换成百度

或者直接进行一下操作（一步到位，不需要多步操作）
[root@localhost ~]# sed -i 's/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /root/.jenkins/updates/default.json && sed -i 's/http:\/\/www.google.com/https:\/\/www.baidu.com/g' /root/.jenkins/updates/default.json
```



之后，在网站后面加上restart进行jenkins重启。

http://192.168.153.147:8080/jenkins/restart



##### 5.邮箱配置(可选)



安装邮件插件，才能确保邮件发送成功。否则可能不会发送邮件



![img](assets/Jenkins构建CICD/image-20210911144642567.png)



![img](assets/Jenkins构建CICD/image-20210911144613368.png)



![img](assets/Jenkins构建CICD/image-20210911144701695.png)



![img](assets/Jenkins构建CICD/image-20210902221414177.png)



![img](assets/Jenkins构建CICD/1678412991631-a8d2c645-8a93-4f95-992e-eacc28defb12.png)

![img](assets/Jenkins构建CICD/1678413070292-b7db5d41-912f-45cd-b406-d8ca82c45922.png)

![img](assets/Jenkins构建CICD/1678413096005-1c994292-b2e0-461c-ba8a-057b1cfe5488.png)

![img](assets/Jenkins构建CICD/1678413112994-180042db-987a-4815-b6d1-775b667c47eb.png)



![img](assets/Jenkins构建CICD/image-20211012175738611.png)



![img](assets/Jenkins构建CICD/image-20211012175851300.png)



![img](assets/Jenkins构建CICD/image-20210902221346972.png)



可看到邮箱确实接收到了邮件，则配置成功；



当然邮箱能接收到邮件的前提是，邮箱要开启smtp服务



![img](assets/Jenkins构建CICD/image-20210902221434247.png)



##### 6.配置Jenkins私钥

GItlab-->Jenkins

```shell
[root@jenkins ~]# ssh-keygen
[root@jenkins ~]# cat /root/.ssh/id_rsa
```



![img](assets/Jenkins构建CICD/image-20210902223413244.png)



##### 7.添加后端服务器

```plain
公钥发送到后端服务器，才能实现免密；
[root@jenkins ~]# ssh-copy-id -i root@192.168.153.194
```



![img](assets/Jenkins构建CICD/image-20210902223816852.png)



##### 8.配置JDK和Maven



虽然Jenkins服务器上，已经安装了JDK和maven工具，但是，还需要在Jenkins服务中，进行配置；



这样Jenkins才能自动化的使用两个工具；



![img](assets/Jenkins构建CICD/image-20210902224113750.png)







![img](assets/Jenkins构建CICD/image-20210902224216097.png)



![img](assets/Jenkins构建CICD/image-20210902224325162.png)



![img](assets/Jenkins构建CICD/image-20210902224414234.png)



如果有多个jdk和maven需要配置的话，可以点击新增jdk或者新增maven；



##### 9.构建发布任务



![img](assets/Jenkins构建CICD/image-20210902224510552.png)



![img](assets/Jenkins构建CICD/image-20210902224539566.png)



![img](assets/Jenkins构建CICD/image-20210902224620999.png)



https://github.com/bingyue/easy-springmvc-maven

或者用以下链接

https://gitee.com/youngfit/easy-springmvc-maven.git



这里是我再github上，找的一个开源项目。能进行编译打包

如果jenkins服务器还未安装git客户端，请先安装

```shell
[root@jenkins-server ~]# yum -y install git
```



![img](assets/Jenkins构建CICD/image-20210902224715465.png)

![img](assets/Jenkins构建CICD/1686712407265-c1121f7d-0ae1-4707-9af6-e208ee6ca6e5.png)



![img](assets/Jenkins构建CICD/image-20210902224957863.png)



如果有多个后端服务器，可以点击 ADD server进行添加；



##### 10.java服务器添加脚本



```shell
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



![img](assets/Jenkins构建CICD/image-20210902225049643.png)



##### 12.启用邮箱



![img](assets/Jenkins构建CICD/image-20210902225124311.png)



##### 13.构建项目



![img](assets/Jenkins构建CICD/image-20210908135616162.png)



![img](assets/Jenkins构建CICD/image-20210908135410486.png)

##### 14.访问tomcat测试

##### 15.更新测试

随便找台机器，作为开发人员

```shell
# git clone https://gitee.com/youngfit/easy-springmvc-maven.git
[root@tomcat-server ~]# cd easy-springmvc-maven/
[root@tomcat-server easy-springmvc-maven]# vim src/main/webapp/index.jsp
```

![img](assets/Jenkins构建CICD/1670918518215-2552f3ca-ca0a-4a47-8b76-009acbafcbd2.png)

```shell
[root@tomcat-server easy-springmvc-maven]# git add *
[root@tomcat-server easy-springmvc-maven]# git config --global user.email "feigeyoungfit@163.com"
[root@tomcat-server easy-springmvc-maven]# git config --global user.name "feigeyoungfit"
[root@tomcat-server easy-springmvc-maven]# git commit -m "username & password"
[master 5e9f4fd] username & password
 1 file changed, 2 insertions(+), 2 deletions(-)
[root@tomcat-server easy-springmvc-maven]# git push origin master
```

![img](assets/Jenkins构建CICD/1670919912187-f18cd605-47a6-4f10-8a12-ca11021e0e23.png)

![img](assets/Jenkins构建CICD/1670919928266-aa1943a1-8b17-4a64-9e53-2a4fa91bbf92.png)



##### 16.准备开源项目

![img](assets/Jenkins构建CICD/1670919078484-e544e579-a23f-4806-9af8-c0625b6d369c.png)

![img](assets/Jenkins构建CICD/1670919234756-2fbc5a36-7d89-4fc0-a874-c4fe88dfa40b.png)

![img](assets/Jenkins构建CICD/1670919255822-92e587de-9e8e-4ed7-b6da-b9867306135c.png)

![img](assets/Jenkins构建CICD/1670919270148-4cbef903-4c46-45ed-aa52-5050b9dbb2ec.png)

![img](assets/Jenkins构建CICD/1670919311801-1ed69934-9bf2-4ac0-bf64-476be3affd3c.png)

随便找台机器，作为开发人员的开发环境，需要安装git，以及配置git邮箱、用户名；

```shell
[root@tomcat-server tmp]# git clone https://gitee.com/youngfit/testweb.git
[root@tomcat-server tmp]# cd testweb/
```

将easy-springmvc-maven-master.zip 源码包，上传到开发人员机器

![img](assets/Jenkins构建CICD/1670919554706-240cfaa3-2ac9-43f1-b76e-c5b493b9684c.png)

```shell
[root@tomcat-server testweb]# yum -y install unzip
[root@tomcat-server testweb]# unzip easy-springmvc-maven-master.zip
[root@tomcat-server testweb]# cp -r easy-springmvc-maven-master/* .
cp: overwrite ‘./README.md’? y
```

![img](assets/Jenkins构建CICD/1670919639327-a296021b-c0c2-4495-bfd4-26b4538475fb.png)

因为已经拷贝出来了。删除没用的压缩包、目录、文件

```shell
[root@tomcat-server testweb]# rm -rf easy-springmvc-maven-master easy-springmvc-maven-master.zip 
[root@tomcat-server testweb]# ls
pom.xml  README.en.md  README.md  src
[root@tomcat-server testweb]# rm -rf README.en.md 
[root@tomcat-server testweb]# ls
pom.xml  README.md  src
[root@tomcat-server testweb]# git add *
[root@tomcat-server testweb]# git commit -m "version1"
[root@tomcat-server testweb]# git push origin master
```

![img](assets/Jenkins构建CICD/1670919712737-dd65eae8-d3b5-4747-bce8-2b2ad0e6ae05.png)



![img](assets/Jenkins构建CICD/1670919730307-6ed78e76-c901-4f1d-9b7a-61473ef21966.png)

### 三、Jenkins+Gitlab+maven项目实战



Jenkins服务器去拉取代码。所以要下载git客户端



```shell
[root@jenkins-server ~]# yum -y install git
```



![img](assets/Jenkins构建CICD/1670983695896-0030f164-eea8-4d3d-9d74-02a5c0a6cf03.png)



开始构建maven项目

![img](assets/Jenkins构建CICD/1670983799423-7acbd9cb-8482-4943-98aa-3c423de2b0b4.png)

![img](assets/Jenkins构建CICD/1670983844940-cc88a9fe-c614-4e1e-bfbb-adba05ff9fc4.png)

![img](assets/Jenkins构建CICD/1670984239967-9e0a6b7a-4169-4bdc-93f0-b1639f4382a7.png)

git@192.168.91.168:root/cloudweb.git

![img](assets/Jenkins构建CICD/1670984295310-8508a7f4-86a5-49ec-aa34-40405f4ae71e.png)

![img](assets/Jenkins构建CICD/1670984350468-dc62707b-bb75-4782-8f8e-876a88b0ebee.png)

```plain
[root@jenkins ~]# cat /root/.ssh/id_rsa
```

![img](assets/Jenkins构建CICD/1670984467041-34a29369-80a9-4899-aa2b-ac9d87da0b2f.png)

```plain
[root@jenkins ~]# cat /root/.ssh/id_rsa.pub 
```

![img](assets/Jenkins构建CICD/1670984552841-87542d9c-7b8e-49d7-b937-e53f93ce43f8.png)

![img](assets/Jenkins构建CICD/1670984612609-59f1e9d9-c304-4979-a88e-f39c38fc6a87.png)

![img](assets/Jenkins构建CICD/1670984709739-27caf301-414a-4a96-a824-826753ece7ad.png)



![img](assets/Jenkins构建CICD/1670984942386-4b11447c-e012-4fb8-8cfc-5a6d19cb1eee.png)

![img](assets/Jenkins构建CICD/1670984967094-cb72f86f-bc21-492d-960a-0bac7994bf76.png)



访问测试

![img](assets/Jenkins构建CICD/1670985178071-7cfddb83-2705-49d9-bcd1-dd05f65f0e57.png)

开发更新代码

```plain
[root@git-client mnt]# git clone git@192.168.91.168:root/cloudweb.git
[root@git-client mnt]# cd cloudweb/
[root@git-client mnt]# vim src/main/webapp/index.jsp
```

![img](assets/Jenkins构建CICD/1670985229589-55325db8-cc65-4369-baf6-91c3e369e7e9.png)

```plain
[root@git-client mnt]# git add *
[root@git-client mnt]# git commit -m "用户名称 & 用户密码"
[master 31f2d4b] 用户名称 & 用户密码
 1 file changed, 2 insertions(+), 2 deletions(-)
[root@git-client mnt]# git push origin master
Counting objects: 11, done.
Compressing objects: 100% (5/5), done.
Writing objects: 100% (6/6), 559 bytes | 0 bytes/s, done.
Total 6 (delta 2), reused 0 (delta 0)
To git@192.168.91.168:root/cloudweb.git
   0121086..31f2d4b  master -> master
```

![img](assets/Jenkins构建CICD/1670985314800-5574461a-e76d-4629-81a1-75100608b81e.png)



![img](assets/Jenkins构建CICD/1670985343699-8f494317-ad9e-4d4d-bcb5-8e0253d356b3.png)

![img](assets/Jenkins构建CICD/1670985420814-7e21501c-6744-4092-9861-af5132a64863.png)







### 四、Jenkins git参数化构建



git参数化构建：开发人员推送代码之前，对此版本的代码，打一个标签(tag)。我们可以认作为是此套代码的版本号。后续可以方便我们进行版本之间的切换。尤其是刚上线一套代码有问题，可以运用jenkins立即进行版本回退/切换；



Gitlab仓库代码准备：

![img](assets/Jenkins构建CICD/1654067572140-a2b87e7c-1721-430e-b0d4-664bb7bbe168.png)

首先，需要安装插件"Git Parameter"。如图



![img](assets/Jenkins构建CICD/image-20220326112424871.png)



手动测试：



```shell
[root@git-client ~]# git clone git@192.168.91.168:root/cloudweb.git
[root@git-client ~]# cd cloudweb
[root@git-client cloudweb]# vim src/main/webapp/index.jsp
```



![img](assets/Jenkins构建CICD/image-20220326113630658.png)



```shell
[root@gitlab easy-springmvc-maven]# git add *
[root@gitlab easy-springmvc-maven]# git commit -m "修改用户为密码"
[root@gitlab easy-springmvc-maven]# git tag -a "v1.0" -m "修改用户为密码"
[root@gitlab easy-springmvc-maven]# git tag  #查看一下
v1.0
[root@gitlab easy-springmvc-maven]# git push origin v1.0
[root@gitlab easy-springmvc-maven]#
```



![img](assets/Jenkins构建CICD/1670987518538-79276bf9-3ac9-443a-9e89-06d8cadf420d.png)



##### 配置Jenkins参数化构建（tag方式）



![img](assets/Jenkins构建CICD/image-20220326113914745.png)



点击"高级"



![img](assets/Jenkins构建CICD/image-20220326113947739.png)



![img](assets/Jenkins构建CICD/image-20220326114125905.png)



![img](assets/Jenkins构建CICD/image-20220326114843521.png)

访问测试

![img](assets/Jenkins构建CICD/1670988020551-8b353d95-444b-4fd3-8037-d869e6a167c4.png)

开发人员再次更新代码，推送仓库

```shell
[root@git-client cloudweb]# vim src/main/webapp/index.jsp
```

![img](assets/Jenkins构建CICD/1670988048171-5da1ebd5-d6ad-4f68-ac47-6e268998ed00.png)

```shell
[root@git-client cloudweb]# git add *
[root@git-client cloudweb]# git commit -m "用户user & 密码pass"
[master ffa6acf] 用户user & 密码pass
 1 file changed, 2 insertions(+), 2 deletions(-)
[root@git-client cloudweb]# git tag -a "v1.1" -m "用户user & 密码pass"
[root@git-client cloudweb]# git tag
v1.0
v1.1
[root@git-client cloudweb]# git push origin v1.1
```

![img](assets/Jenkins构建CICD/1670988152608-80c3a16b-83e9-4363-b7f4-ba00d661a53c.png)

![img](assets/Jenkins构建CICD/1670988252805-47ea6916-a09e-44c6-a3a1-efd91dfdc5f4.png)

访问测试

尝试进行版本切换，再访问测试



##### 配置Jenkins参数化构建(commit修订号)

开发人员如果不会打标签，或者说他们不愿意配合打标签

![img](assets/Jenkins构建CICD/1668755336084-5813fea2-fece-4763-b914-ece7b2bafdd1.png)

![img](assets/Jenkins构建CICD/1668755355976-1e6f1c29-a626-428a-82d4-f47ba9013bae.png)

![img](assets/Jenkins构建CICD/1668755371041-b2cb07c0-e061-4993-a99e-dca651bb2920.png)



### 五、Jenkins多节点配置



在企业里面使用Jenkins自动部署+测试平台时，每天更新发布几个网站版本,很频繁,但是对于一些大型的企业来讲，Jenkins就需要同时处理很多的任务，这时候就需要借助Jenkins多个node或者我们所说的Jenkins分布式SLAVE，今天我们带大家来学习Jenkins多实例的配置；



添加Linux平台Jenkins SLAVE配置：



1. 由于Jenkins是Java程序，添加的SLAVE客户端服务器必须安装Java JDK环境；
2. 创建远程执行Jenkins任务的用户，一般为Jenkins用户，工作目录为/home/Jenkins;
3. Jenkins服务器免秘钥登录Slave服务器或者通过用户名和密码登录；



#### 1.添加从节点



![img](assets/Jenkins构建CICD/image-20211009204536504.png)



![img](assets/Jenkins构建CICD/image-20211009204609760.png)



![img](assets/Jenkins构建CICD/image-20211009204648831.png)



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



![img](assets/Jenkins构建CICD/image-20220328104326621.png)



![img](assets/Jenkins构建CICD/image-20211009221943512.png)



![img](assets/Jenkins构建CICD/image-20211009221815103.png)



#### 3.指定java命令路径



![img](assets/Jenkins构建CICD/image-20211009221903569.png)



![img](assets/Jenkins构建CICD/image-20211009222015840.png)



#### 4.测试从节点



项目指定到哪个节点运行。



![img](assets/Jenkins构建CICD/image-20220328104400072.png)



![img](assets/Jenkins构建CICD/image-20211009222752678.png)



![img](assets/Jenkins构建CICD/image-20211009224634465.png)