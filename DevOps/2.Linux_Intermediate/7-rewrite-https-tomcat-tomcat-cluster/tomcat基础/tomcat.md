## Tomcat 运维实战  JDK

### 1、JVM 虚拟机常识

**两个常识问题**

作为了解JVM 虚拟机的开始。我们很有必要弄明白以下两个问题。

JVM   JRE  JAVA类库

jvm:  java虚拟机

JRE：java运行环境

java类库：语言库

JVM是java virtual  Machine 的缩写，JVM是一种用于计算机设备的规范，它是一个虚构出来的计算机。使用java-version可以查看是否安装虚拟机



#### **1、什么是JAVA虚拟机**

```
所谓虚拟机，就是一台虚拟的计算机。他是一款软件，用来执行一系列虚拟计算机指令。大体上，虚拟机可以分为系统虚拟机和程序虚拟机。大名鼎鼎的VisualBox、VMware就属于系统虚拟机。他们完全是对物理计算机的仿真。提供了一个可以运行完整操作系统的软件平台。
程序虚拟机的典型代表就是Java虚拟机，它专门为执行单个计算机程序而设计，在Java虚拟机中执行的指令我们称为Java字节码指令。无论是系统虚拟机还是程序虚拟机，在上面运行的软件都限制于虚拟机提供的资源中。
```

#### **2、JAVA 如何做到跨平台**

```shell
同一个JAVA程序(JAVA字节码的集合)，通过JDK--JAVA虚拟机(JVM)运行于各大主流操作系统平台
比如Windows、CentOS、Ubuntu等。程序以虚拟机为中介，来实现跨平台.
java的跨平台原理：
运行过程：首先将Java文件编译成字节码（.class）文件【在所有平台生成的字节码文件都是相同的】；然后使用虚拟机jvm运行字节码文件【在不同的平台上运行字节码文件的java虚拟机是不同的，不同的平台有不同的版本】。
JDK是Java Development Kit的缩写，JDK是Java语言的软件开发工具包(SDK)。提供了java开发、编译、运行需要的文件和环境。
在JDK的安装目录下有一个jre目录，里面有两个文件夹bin和lib，在这里可以认为bin里的就是jvm，lib中则是jvm工作所需要的类库，而jvm和 lib合起来就称为jre。
JRE是Java Runtime Environment的缩写，JRE的内部有一个Java虚拟机（Java Virtual Machine）以及一些标准的类别函数库（Class Library），JRE是运行java字节码文件必备的工具
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735627453-fbb16858-55b9-4a7d-b22b-00b196c0d7e8.png#averageHue=%23e0e0e0&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=405&id=u5513bf51&name=image.png&originHeight=506&originWidth=637&originalType=binary&ratio=1&rotation=0&showTitle=false&size=131084&status=done&style=none&taskId=ua64fb5eb-a6f0-43b0-8d5b-9ce7bf4130b&title=&width=509.6)

#### 3、常用虚拟机参数

JVM 虚拟机提供了三种类型参数

##### 1、标准参数

```
标准参数中包括功能和输出的参数都是很稳定的，很可能在将来的JVM版本中不会改变。你可以用 java 命令（或者是用 java -help）检索出所有标准参数。java -
java-version:检查是否安装JDK

```

##### 2、X  类型参数

```
非标准化的参数，在将来的版本中可能会改变。所有的这类参数都以 -X 开始。
```

##### 3、XX  类型参数

```
在实际情况中 X 参数和 XX 参数并没有什么不同。X 参数的功能是十分稳定的。
用一句话来说明 XX 参数的语法。所有的 XX 参数都以"-XX:"开始，但是随后的语法不同，取决于参数的类型：
1）对于布尔类型的参数，我们有"+"或"-"，然后才设置 JVM 选项的实际名称。
   例如，-XX:+ 用于激活选项，而 -XX:- 用于注销选项。
   Example:
   开启GC日志的参数: -XX:+PrintGC
2) 对于需要非布尔值的参数，如 string 或者 integer，我们先写参数的名称，后面加上"="，最后赋值。
   例如: -XX:MaxPermSize=2048m
   例如：-XX:PermSize=2048m
```

#### 4、常用的JVM参数

##### 1、跟踪JAVA虚拟机的垃圾回收

GC日志：jvm垃圾回收，记录jvm的运行状态，OOM内存溢出的报错信息等。

- %t 将会被替代为时间字符串，格式为: YYYY-MM-DD_HH-MM-SS

开启GC日志:

```shell
-XX:+PrintGCDetails -XX:+PrintGCDateStamps -Xloggc:/data0/logs/gc-%t.log"
```

##### 2、配置JAVA虚拟机的堆空间

```
-Xms:初始堆大小
-Xmx:最大堆大小
实际生产环境中， 我们通常将初始化堆(-Xms) 和 最大堆(-Xmx) 设置为一样大。以避免程序频繁的申请堆空间。设置为物理内存的一半。
```

##### 3、配置JAVA虚拟机的永久区(方法区)

```
-XX:PermSize=2048m   	  内存永久保留区域  ：//所占用的内存是堆内存的一部分内存，不能超过堆内存
-XX:MaxPermSize=2048m   内存最大永久保留区域

JDK 1.8中 PermSize 和 MaxPermGen 已经无效。JDK 1.8 中已经不存在永久代的结论而以元空间代替。
java -jar  参数    文件路径
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735658983-076c894f-3a99-4ab0-9a18-103da6b53668.png#averageHue=%23f5f6f4&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=193&id=ua41998b3&name=image.png&originHeight=241&originWidth=799&originalType=binary&ratio=1&rotation=0&showTitle=false&size=137727&status=done&style=none&taskId=u3f5b64af-9a2d-4c71-b805-b0533aa12d1&title=&width=639.2)

nohup :后台运行，启动程序命令 &

Tomcat是Apache软件基金会（Apache Software Foundation）的Jakarta 项目中的一个核心项目，由Apache、Sun和其他一些公司及个人共同开发而成。

**Tomcat服务器是一个免费的开放源代码的Web应用服务器，属于轻量级应用服务器，在中小型系统和并发访问用户不是很多的场合下被普遍使用，是开发和调试JSP程序的首选。**

**Tomcat：JAVA容器，WEB容器，WEB中间件**

Tomcat,JBOSS，Weblogic  ---收费。

**apache和nginx  只能解析静态页面**

**web容器：uwsgi   php      tomcat   -----这些是解析动态页面的**

一个tomcat默认并发是200(官方)，可以修改，但实际用的时候也就150并发左右。

**tomcat端口:本身自己的端口:8005.还有一个端口是和其他应用通信的端口:8009。给浏览器(客户端)访问页面用的端口是8080。**

**https端口:443**

**使用方案**：

方案一：  Tomcat         //单独使用   ----基本不用
方案二：  Nginx+Tomcat *10      //反向代理和负载均衡                                    			
方案三：
Nginx
|
+--------------------------------------------------------+
|               |               | 		              |
Tomcat1 Tomcat2 Tomcat3         nginx服务器  ----解析静态页面

建议使用Nginx和Tomcat配合，Nginx处理静态，Tomcat处理动态程序
方案三中后端Tomcat可以运行在单独的主机，也可以是同一台主机上的多实例

**Tomcat**官网： [http://tomcat.apache.org](http://tomcat.apache.org/)

##### 1、Tomcat好帮手---JDK

JDK是 Java 语言的软件开发工具包，JDK是整个java开发的核心，JDK中包括完整的JRE（Java Runtime Environment），Java运行环境，包括了用于产品环境的各种库类，如基础类库rt.jar，以及给开发人员使用的补充库等。

**JDK下载面页：**

[http://www.oracle.com/technetwork/java/javase/downloads/index.html](http://www.oracle.com/technetwork/java/javase/downloads/index.html)

##### 2、安装Tomcat & JDK

安装时候选择tomcat软件版本要与程序开发使用的版本一致。jdk版本要进行与tomcat保持一致。

###### 1、系统环境说明

```shell
[root@java-tomcat1 ~]# cat /etc/redhat-release 
CentOS Linux release 7.4.1708 (Core) 
[root@java-tomcat1 ~]# uname -a 
Linux java-tomcat1 3.10.0-693.el7.x86_64 #1 SMP Tue Aug 22 21:09:27 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
[root@java-tomcat1 ~]# getenforce 
Disabled
[root@java-tomcat1 ~]# systemctl status firewalld
● firewalld.service - firewalld - dynamic firewall daemon
   Loaded: loaded (/usr/lib/systemd/system/firewalld.service; disabled; vendor preset: enabled)
   Active: inactive (dead)
     Docs: man:firewalld(1)
```

###### 2 、安装JDK

命令集：

```shell
上传jdk1.8到服务器。安装jdk
[root@java-tomcat1 ~]# tar xzf jdk-8u191-linux-x64.tar.gz -C /usr/local/
[root@java-tomcat1 ~]# cd /usr/local/
[root@java-tomcat1 local]# mv jdk1.8.0_191/ java
设置环境变量:
[root@java-tomcat1 local]# vim /etc/profile
export JAVA_HOME=/usr/local/java   #指定java安装目录
export PATH=$JAVA_HOME/bin:$JAVA_HOME/jre/bin:$PATH    #用于指定java系统查找命令的路径
export CLASSPATH=.:$JAVA_HOME/lib:$JAVA_HOME/jre/lib:$JAVA_HOME/lib/tools.jar  #类的路径，在编译运行java程序时，如果有调用到其他类的时候，在classpath中寻找需要的类。
检测JDK是否安装成功:
[root@java-tomcat1 local]# source /etc/profile
[root@java-tomcat1 local]# java -version
java version "1.8.0_191"
Java(TM) SE Runtime Environment (build 1.8.0_191-b12)
Java HotSpot(TM) 64-Bit Server VM (build 25.191-b12, mixed mode)
```

###### 3、安装Tomcat

命令集：

```shell
[root@java-tomcat1 ~]# mkdir /data/application -p
[root@java-tomcat1 ~]# cd /usr/src/
[root@java-tomcat1 ~]# yum -y install wget
[root@java-tomcat1 src]# wget http://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.46/bin/apache-tomcat-8.5.46.tar.gz
[root@java-tomcat1 src]# tar xzf apache-tomcat-8.5.46.tar.gz -C /data/application/
[root@java-tomcat1 src]# cd /data/application/
[root@java-tomcat1 application]# mv apache-tomcat-8.5.46/ tomcat
设置环境变量:
[root@java-tomcat1 application]# vim /etc/profile
export TOMCAT_HOME=/data/application/tomcat   #指定tomcat的安装目录
[root@java-tomcat1 application]# source  /etc/profile
查看tomcat是否安装成功:
[root@java-tomcat1 tomcat]# /data/application/tomcat/bin/version.sh
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
Server version: Apache Tomcat/8.5.42
Server built:   Jun 4 2019 20:29:04 UTC
Server number:  8.5.42.0
OS Name:        Linux
OS Version:     3.10.0-693.el7.x86_64
Architecture:   amd64
JVM Version:    1.8.0_191-b12i
JVM Vendor:     Oracle Corporation
```

#### 2、Tomcat目录介绍

##### 1、tomcat主目录介绍

```shell
[root@java-tomcat1 ~]# cd /data/application/tomcat/
[root@java-tomcat1 tomcat]# yum install -y tree   (树)
[root@java-tomcat1 tomcat]# tree -L 1
.
├── bin     #存放tomcat的管理脚本
├── BUILDING.txt
├── conf    #tomcat的配置文件
├── CONTRIBUTING.md
├── lib      #web应用调用的jar包存放路径
├── LICENSE
├── logs     #tomcat日志存放目录，catalin.out日志为只要输出日志
├── NOTICE
├── README.md
├── RELEASE-NOTES
├── RUNNING.txt
├── temp     #存放临时文件
├── webapps  #默认网站发布目录
└── work     #存放编译生产的.java与.class文件

7 directories, 7 files
```

**2、webapps目录介绍**

/usr/share/nginx/html        

index index.html

location /root

```shell
[root@java-tomcat1 tomcat]# cd webapps/
[root@java-tomcat1 webapps]# tree -L 1
.
├── docs  #tomcat的帮助文档
├── examples  #web应用实例
├── host-manager  #主机管理
├── manager    #管理
└── ROOT    #默认站点根目录

5 directories, 0 files
```

**3、Tomcat配置文件目录介绍（conf）**

```shell
[root@java-tomcat1 webapps]# cd ../conf/
[root@java-tomcat1 conf]# tree -L 1
.
├── Catalina
├── catalina.policy
├── catalina.properties
├── context.xml
├── logging.properties
├── logs
├── server.xml           # tomcat 主配置文件
├── server.xml.bak
├── server.xml.bak2
├── tomcat-users.xml    # tomcat 管理用户配置文件
├── tomcat-users.xsd
└── web.xml

2 directories, 10 files
```

**4、Tomcat的管理**

```shell
启动程序 #/data/application/tomcat/bin/startup.sh
关闭程序 #/data/application/tomcat/bin/shutdown.sh
```

启动停止

```shell
[root@java-tomcat1 conf]# cd ../bin/
[root@java-tomcat1 bin]# ./startup.sh 
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
Tomcat started.
[root@java-tomcat1 bin]# ./shutdown.sh 
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
```

注意：tomcat未启动的情况下使用shutdown脚本，会有大量的输出信息。

检查tomcat是否启动正常

```shell
[root@java-tomcat1 bin]# netstat -lntp  |grep java
tcp6       0      0 :::8080         :::*                   LISTEN      30560/java
tcp6       0      0 127.0.0.1:8005          :::*          LISTEN      30560/java
tcp6       0      0 :::8009                 :::*           LISTEN      30560/java
```

端口:

8005：这个端口负责监听关闭Tomcat的请求 shutdown:向以上端口发送的关闭服务器的命令字符串。

8009: 与其他服务通信接口，接受其他服务器转发过来的请求

8080: 建立http连接用。可以修改

**说明：**所有与java相关的，服务启动都是java命名的进程

启动完成浏览器进行访问

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735788144-0930855b-f219-456d-a9a5-26e4c074251d.png#averageHue=%23ddf0cf&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=241&id=uae8283f4&name=image.png&originHeight=301&originWidth=814&originalType=binary&ratio=1&rotation=0&showTitle=false&size=90538&status=done&style=none&taskId=u476833b1-ed76-4b7c-a7ec-cd98dda77c2&title=&width=651.2)

#### 

##### 查看日志

```shell
[root@java-tomcat1 bin]# tail -f /data/application/tomcat/logs/catalina.out 
org.apache.catalina.startup.HostConfig.deployDirectory Deployment of web application directory [/data/application/tomcat/webapps/host-manager] has finished in [21] ms
04-Jul-2019 22:40:00.026 INFO [localhost-startStop-1] org.apache.catalina.startup.HostConfig.deployDirectory Deploying web application directory [/data/application/tomcat/webapps/manager]
04-Jul-2019 22:40:00.042 INFO [localhost-startStop-1] org.apache.catalina.startup.HostConfig.deployDirectory Deployment of web application directory [/data/application/tomcat/webapps/manager] has finished in [16] ms
04-Jul-2019 22:40:00.048 INFO [main] org.apache.coyote.AbstractProtocol.start Starting ProtocolHandler ["http-nio-8080"]
04-Jul-2019 22:40:00.058 INFO [main] org.apache.coyote.AbstractProtocol.start Starting ProtocolHandler ["ajp-nio-8009"]
04-Jul-2019 22:40:00.062 INFO [main] org.apache.catalina.startup.Catalina.start Server startup in 479 ms
```

#### 4、Tomcat主配置文件详解

##### 1、server.xml组件类别

顶级组件：位于整个配置的顶层，如server。

容器类组件：可以包含其它组件的组件，如service、engine、host、context。

连接器组件：连接用户请求至tomcat，如connector。

```shell
<server>  #表示一个运行于JVM中的tomcat实例。
     <service> #服务。将connector关联至engine，因此一个service内部可以有多个connector，但只能有一个引擎engine。
     <connector /> #接收用户请求，类似于httpd的listen配置监听端口的
     <engine>  #核心容器组件，catalina引擎，负责通过connector接收用户请求，并处理请求，将请求转至对应的虚拟主机host。
     <host>   #类似于httpd中的虚拟主机，
     <context></context>  #配置context的主要目的指定对应对的webapp的根目录。其还能为webapp指定额外的属性，如部署方式等。
     </host>
     <host>
     <context></context>
     </host>
     </engine>
     </service>
</server>
```

##### 2、server.xml配置文件注释

```shell
<?xml version='1.0' encoding='utf-8'?>
<!--
<Server>元素代表整个容器,是Tomcat实例的顶层元素.它包含一个<Service>元素.并且它不能做为任何元素的子元素.
    port指定Tomcat监听shutdown命令端口
    shutdown指定终止Tomcat服务器运行时,发给Tomcat服务器的shutdown监听端口的字符串.该属性必须设置
-->
<Server port="8005" shutdown="SHUTDOWN">
  <Listener className="org.apache.catalina.startup.VersionLoggerListener" />
  <Listener className="org.apache.catalina.core.AprLifecycleListener" SSLEngine="on" />
  <Listener className="org.apache.catalina.core.JreMemoryLeakPreventionListener" />
  <Listener className="org.apache.catalina.mbeans.GlobalResourcesLifecycleListener" />
  <Listener className="org.apache.catalina.core.ThreadLocalLeakPreventionListener" />
  <GlobalNamingResources>
    <Resource name="UserDatabase" auth="Container"
              type="org.apache.catalina.UserDatabase"
              description="User database that can be updated and saved"
              factory="org.apache.catalina.users.MemoryUserDatabaseFactory"
              pathname="conf/tomcat-users.xml" />
  </GlobalNamingResources>
  <!--service服务组件-->
  <Service name="Catalina">
    <!-- Connector主要参数说明（见下面） -->
    <Connector port="8080" protocol="HTTP/1.1"
               connectionTimeout="20000"
               redirectPort="8443" />
    <Connector port="8009" protocol="AJP/1.3" redirectPort="8443" />
    <Engine name="Catalina" defaultHost="localhost">
      <Realm className="org.apache.catalina.realm.LockOutRealm">
        <Realm className="org.apache.catalina.realm.UserDatabaseRealm"
               resourceName="UserDatabase"/>
      </Realm>
      <!-- 详情常见（host参数详解）-->
      <Host name="localhost"  appBase="webapps"
            unpackWARs="true" autoDeploy="true">
        <!-- 详情见扩展（Context参数说明 ）-->
        <Context path="" docBase="" debug=""/>
        <Valve className="org.apache.catalina.valves.AccessLogValve" directory="logs"
               prefix="localhost_access_log" suffix=".txt"
               pattern="%h %l %u %t &quot;%r&quot; %s %b" />
      </Host>
    </Engine>
  </Service>
</Server>
```

##### 4、Connector主要参数说明

```shell
port:指定服务器端要创建的端口号，并在这个端口监听来自客户端的请求。
protocol：连接器使用的协议，支持HTTP和AJP。AJP（Apache Jserv Protocol）专用于tomcat与apache建立通信的， 在httpd反向代理用户请求至tomcat时使用（可见Nginx反向代理时不可用AJP协议）。
redirectPort：指定服务器正在处理http请求时收到了一个SSL传输请求后重定向的端口号
maxThreads：接收最大请求的并发数
connectionTimeout  指定超时的时间数(以毫秒为单位)
```

<Connector port="8080" protocol="HTTP/1.1"

               maxThreads="500"    ----默认是200
               connectionTimeout="20000"       ---------连接超时时间。单位毫秒
               redirectPort="8443" />

##### 5、host参数详解

            <Host name="localhost"  appBase="webapps"
            unpackWARs="true" autoDeploy="true">

```shell
host:表示一个虚拟主机。
name:指定主机名。
appBase:应用程序基本目录，即存放应用程序的目录.一般为appBase="webapps"，相对于CATALINA_HOME而言的，也可以写绝对路径。
unpackWARs:如果为true，则tomcat会自动将WAR文件解压，否则不解压，直接从WAR文件中运行应用程序。
autoDeploy:在tomcat启动时，是否自动部署。
```

#### 5、WEB站点部署

上线的代码有两种方式：

第一种方式是直接将程序目录放在webapps目录下面，这种方式大家已经明白了，就不多说了。

第二种方式是使用开发工具将程序打包成war包，然后上传到webapps目录下面。

**1、使用war包部署web站点**

```shell
[root@java-tomcat1 ~]# pwd
/root

下载jenkins的war包
[root@java-tomcat1 ~]# wget http://updates.jenkins-ci.org/download/war/2.129/jenkins.war
[root@java-tomcat1 ~]# ls
jenkins.war

[root@java-tomcat1 ~]# cd /data/application/tomcat   #进入tomcat目录
[root@java-tomcat1 tomcat]# cp -r webapps/ /opt/    #将原来的发布网站目录备份
[root@java-tomcat1 tomcat]# cd webapps/
[root@java-tomcat1 webapps]# ls
docs  examples  host-manager  manager  ROOT
[root@java-tomcat1 webapps]# rm -rf *    #清空发布网站里面的内容
[root@java-tomcat1 webapps]# cp /root/jenkins.war .   #将war包拷贝到当前目录
[root@java-tomcat1 webapps]# ../bin/startup.sh   #启动
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
Tomcat started.
[root@java-tomcat1 webapps]# ls
jenkins  jenkins.war

二、手动解压:
[root@java-tomcat1 webapps]# ../bin/shutdown.sh   #关闭tomcat
[root@java-tomcat1 ~]# cd /data/application/tomcat/webapps/
[root@java-tomcat1 webapps]# rm -rf *    
[root@java-tomcat1 webapps]# mkdir ROOT      #创建一个ROOT目录存放war包
[root@java-tomcat1 webapps]# ls
ROOT
[root@java-tomcat1 webapps]# cd ROOT/
[root@java-tomcat1 ROOT]# cp /root/jenkins.war .
[root@java-tomcat1 ROOT]# unzip jenkins.war
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735852251-693f61e3-08cd-417d-bc27-8c7b19890515.png#averageHue=%23312f2d&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=84&id=u1d29f87e&name=image.png&originHeight=105&originWidth=811&originalType=binary&ratio=1&rotation=0&showTitle=false&size=57707&status=done&style=none&taskId=uef075af3-e9b4-4460-90ea-1225fb5bc3b&title=&width=648.8)

浏览器访问：[http://192.168.1.7:8080/jenkins](http://192.168.1.7:8080/jenkins)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735862427-0083b47e-ecbe-47c4-aa9c-31f0410a3a89.png#averageHue=%23f4f4f4&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=374&id=ubc197f01&name=image.png&originHeight=468&originWidth=810&originalType=binary&ratio=1&rotation=0&showTitle=false&size=80459&status=done&style=none&taskId=u28e8a4ad-0803-445a-aa9c-d86fb387940&title=&width=648)

##### 2、自定义默认网站目录

1、修改默认发布目录:

```shell
[root@java-tomcat1 ~]# mkdir /data/application/webapp  #创建发布目录
[root@java-tomcat1 ~]# vim /data/application/tomcat/conf/server.xml
```

将原来的

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735887098-6eade42d-d25b-46e1-8d7b-abdb863fc861.png#averageHue=%2331201f&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=84&id=ua7d53f5b&name=image.png&originHeight=105&originWidth=708&originalType=binary&ratio=1&rotation=0&showTitle=false&size=16136&status=done&style=none&taskId=udbdfc1e4-4489-43af-b91e-63cd6f5deb6&title=&width=566.4)

修改为

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735891932-44b863f6-11fb-4cc8-8a21-f3923db24a3c.png#averageHue=%23463f3f&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=78&id=ub593ead9&name=image.png&originHeight=97&originWidth=820&originalType=binary&ratio=1&rotation=0&showTitle=false&size=16667&status=done&style=none&taskId=ub1b57b3b-ef58-486b-9659-1bbc10a385f&title=&width=656)

```shell
[root@java-tomcat1 ~]# cp /root/jenkins.war /data/application/webapp/
[root@java-tomcat1 ~]# /data/application/tomcat/bin/startup.sh
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
Tomcat started.
[root@java-tomcat1 ~]# ll /data/application/webapp/   #已经自动解压
jenkins/     jenkins.war
```

浏览器访问：[http://192.168.1.7:8080/jenkins](http://192.168.1.7:8080/jenkins)

**3、部署开源站点（jspgou商城）**

第一：安装配置数据库

```shell
MariaDB设置初始化密码及修改密码
方法一：
[root@localhost ~]# mysql -uroot
MariaDB [(none)]> use mysql;
MariaDB [mysql]> UPDATE mysql.user SET password = PASSWORD('newpassward') WHERE user = 'root';
MariaDB [mysql]> FLUSH PRIVILEGES;
方法二：
[root@localhost ~]# mysql -uroot
MariaDB [(none)]> use mysql;
MariaDB [mysql]> SET password=PASSWORD('newpassward');
MariaDB [mysql]> FLUSH PRIVILEGES;
方法三：
[root@localhost ~]# mysqladmin -u root password 'newpassword'
#如果root已经设置过密码，采用如下方法 
[root@localhost ~]# mysqladmin -u root -p 'oldpassword' password 'newpassword'
```



```shell
1.使用mariadb
[root@yangge ~]# yum -y install mariadb mariadb-server
[root@yangge ~]# systemctl start mariadb
[root@yangge ~]# mysql
create database jspgou default charset=utf8;	//在数据库中操作，创建数据库并指定字符集
flush privileges;		//(可选操作)
exit;
```

第二：jspgou商城上线

关掉tomcat

```shell
上传jspgou商城的代码
[root@java-tomcat1 ~]# unzip jspgouV6.1-ROOT.zip
[root@java-tomcat1 ~]# cp -r ROOT/ /data/application/tomcat/webapps/
[root@java-tomcat1 ~]#cd /data/application/tomcat/webapps/
[root@java-tomcat1 webapps]# ls
ROOT
```

```shell
将数据导入数据库:
[root@java-tomcat1 ~]# cd DB/
[root@java-tomcat1 DB]# ls
jspgou.sql
[root@java-tomcat1 DB]# mysql -uroot -p  jspgou < jspgou.sql

启动tomcat访问:
[root@java-tomcat1 ~]# /data/application/tomcat/bin/startup.sh
[root@java-tomcat1 ~]# netstat -lntp
```

访问:[http://192.168.1.7:8080/](http://192.168.1.7:8080/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735930868-d2140ca6-d837-4ea5-ab13-e42bdddd6dcb.png#averageHue=%23c8c3b9&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=242&id=u5a67db8b&name=image.png&originHeight=303&originWidth=812&originalType=binary&ratio=1&rotation=0&showTitle=false&size=191569&status=done&style=none&taskId=ua597c220-c784-4cc1-8831-924f1610640&title=&width=649.6)

做一遍：要求用mysql5.7（会遇到问题，百度解决），贴近企业环境

#### 4、Tomcat多实例配置

**多实例（多进程）**：同一个程序启动多次，分为两种情况:

第一种：一台机器跑多个站点；

第二种：多个机器跑一个站点多个实例，配合负载均衡;

##### 1、复制程序文件

```shell
[root@java-tomcat1 ~]# cd /data/application/
[root@java-tomcat1 application]# ls
tomcat
[root@java-tomcat1 application]# cp -r tomcat/ tomcat_2
[root@java-tomcat1 application]# ls
tomcat  tomcat_2
修改端口，以启动多实例。多实例之间端口不能一致
[root@java-tomcat1 application]# sed -i 's@8005@8011@;s/8080/8081/' tomcat/conf/server.xml
[root@java-tomcat1 application]# sed -i 's#8005#8012#;s#8080#8082#' tomcat_2/conf/server.xml
[root@java-tomcat1 application]# sed -i 's#8009#8019#' tomcat/conf/server.xml
[root@java-tomcat1 application]# sed -i 's#8009#8029#' tomcat_2/conf/server.xml
[root@java-tomcat1 application]# diff tomcat/conf/server.xml tomcat_2/conf/server.xml  #对比文件不同之处
22c22
< <Server port="8011" shutdown="SHUTDOWN">
---
> <Server port="8012" shutdown="SHUTDOWN">
67c67
<          Define a non-SSL/TLS HTTP/1.1 Connector on port 8081
---
>          Define a non-SSL/TLS HTTP/1.1 Connector on port 8082
69c69
<     <Connector port="8081" protocol="HTTP/1.1"
---
>     <Connector port="8082" protocol="HTTP/1.1"
75c75
<                port="8081" protocol="HTTP/1.1"
---
>                port="8082" protocol="HTTP/1.1"
115,116c115,116
<     <!-- Define an AJP 1.3 Connector on port 8019 -->
<     <Connector port="8019" protocol="AJP/1.3" redirectPort="8443" />
---
>     <!-- Define an AJP 1.3 Connector on port 8029 -->
>     <Connector port="8029" protocol="AJP/1.3" redirectPort="8443" />
```

启动tomcat多实例

```shell
[root@java-tomcat1 application]# cp -r /opt/webapps/ROOT/ tomcat/webapps/
[root@java-tomcat1 application]# cp -r /opt/webapps/ROOT/ tomcat_2/webapps/
[root@java-tomcat1 application]# echo 8081 >> tomcat/webapps/ROOT/index.jsp 
[root@java-tomcat1 application]# echo 8082 >> tomcat_2/webapps/ROOT/index.jsp
启动：
[root@java-tomcat1 application]# /data/application/tomcat/bin/startup.sh
[root@java-tomcat1 application]# /data/application/tomcat_2/bin/startup.sh
```

```shell
作业：通过设置环境变量，使启动变得简单
```

检查端口查看是否启动:

```shell
[root@java-tomcat1 application]# netstat -lntp | grep java 
tcp6       0      0 127.0.0.1:8011          :::*                    LISTEN      1729/java           
tcp6       0      0 127.0.0.1:8012          :::*                    LISTEN      1783/java           
tcp6       0      0 :::8081                 :::*                    LISTEN      1729/java           
tcp6       0      0 :::8082                 :::*                    LISTEN      1783/java           
tcp6       0      0 :::8019                 :::*                    LISTEN      1729/java           
tcp6       0      0 :::8029                 :::*                    LISTEN      1783/java
```

##### 2、在浏览器访问，进行测试

检查多实例的启动

[http://192.168.50.114:8081/](http://192.168.50.114:8081/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735979974-cb6f535e-6492-499b-8558-da9cd9f39cfe.png#averageHue=%23f4e3b9&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=549&id=u88b8fa91&name=image.png&originHeight=686&originWidth=812&originalType=binary&ratio=1&rotation=0&showTitle=false&size=211731&status=done&style=none&taskId=udad4599c-bcf0-4107-973c-7eebe51cbc1&title=&width=649.6)

[http://192.168.50.114:8082/](http://192.168.50.114:8082/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669735987813-a69104e9-76a1-4a90-92f1-4779c629c68d.png#averageHue=%23f3e3ba&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=514&id=u96408115&name=image.png&originHeight=643&originWidth=823&originalType=binary&ratio=1&rotation=0&showTitle=false&size=221034&status=done&style=none&taskId=uc47729db-581f-4f91-aaaf-489eecf9521&title=&width=658.4)

#### 7、tomcat反向代理集群

##### 1、负载均衡器说明

关闭防火墙和selinux

```shell
yum安装nginx
[root@nginx-proxy ~]# cd /etc/yum.repos.d/
[root@nginx-proxy yum.repos.d]# vim nginx.repo
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=0
enabled=1
[root@nginx-proxy yum.repos.d]# yum install yum-utils -y
[root@nginx-proxy yum.repos.d]# yum install nginx -y
```

**2、配置负载均衡器**

备份原配置文件并修改

```shell
[root@nginx-proxy ~]# cd /etc/nginx/conf.d/
[root@nginx-proxy conf.d]# cp default.conf default.conf.bak
[root@nginx-proxy conf.d]# mv default.conf tomcat.conf
[root@nginx-proxy conf.d]# vim tomcat.conf
upstream testweb {
	server 192.168.50.114:8081 weight=1 max_fails=1 fail_timeout=2s;
	server 192.168.50.114:8082 weight=1 max_fails=1 fail_timeout=2s;
}
server {
    listen       80;
    server_name  localhost;
    access_log  /var/log/nginx/proxy.access.log  main;

    location / {
       root /usr/share/nginx/html;
       index index.html index.htm;
       proxy_pass http://testweb;
       proxy_set_header Host $host:$server_port;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }       
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    } 
}
创建upstream配置文件:
[root@nginx-proxy conf.d]# vim upstream.conf
upstream testweb {
	server 192.168.50.114:8081 weight=1 max_fails=1 fail_timeout=2s;
	server 192.168.50.114:8082 weight=1 max_fails=1 fail_timeout=2s;
}
```

启动nginx

```shell
[root@nginx-proxy ~]# systemctl start nginx
```

##### 3、在浏览器上进行访问测试

[http://192.168.50.118/](http://192.168.50.118/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736058442-c53bb92d-1691-41dd-9bc5-140cf193864c.png#averageHue=%23f5e4ba&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=519&id=u06016e71&name=image.png&originHeight=649&originWidth=818&originalType=binary&ratio=1&rotation=0&showTitle=false&size=226708&status=done&style=none&taskId=u1887ddea-23e1-4d08-a00f-f4c3116ef68&title=&width=654.4)

[http://192.168.50.118/](http://192.168.50.118/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736065212-68f2ac49-6d89-4f50-bc64-ad71fee0771b.png#averageHue=%23f5e4ba&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=513&id=u690eabed&name=image.png&originHeight=641&originWidth=828&originalType=binary&ratio=1&rotation=0&showTitle=false&size=222256&status=done&style=none&taskId=uf5e57ca4-e587-4a48-a8b6-511275598e9&title=&width=662.4)

#### 1、日志格式配置

```shell
[root@java-tomcat1 ~]# cd /data/application/tomcat/conf/
[root@java-tomcat1 conf]# vim server.xml
<Valve className="org.apache.catalina.valves.AccessLogValve" directory="/data/www/logs"
               prefix="jenkins-" suffix="-access_log"
               pattern="%{X-Real-IP}i - %v %t &quot;%r&quot; - %s %b %T &quot;%{Referer}i&quot; &quot;%{User-Agent}i&quot; %a &quot;-&quot; &quot;-&quot;" />
[root@java-tomcat1 conf]# mkdir -p /data/www

日志参数解释：
    ％a - 远程IP地址
    ％A - 本地IP地址
    ％b - 发送的字节数，不包括HTTP头，或“ - ”如果没有发送字节
    ％B - 发送的字节数，不包括HTTP头
    ％h - 远程主机名
    ％H - 请求协议
    ％l (小写的L)- 远程逻辑从identd的用户名（总是返回' - '）
    ％m - 请求方法
    ％p - 本地端口
    ％q - 查询字符串（在前面加上一个“？”如果它存在，否则是一个空字符串
    ％r - 第一行的要求
    ％s - 响应的HTTP状态代码
    ％S - 用户会话ID
    ％t - 日期和时间，在通用日志格式
    ％u - 远程用户身份验证
    ％U - 请求的URL路径
    ％v - 本地服务器名
    ％D - 处理请求的时间（以毫秒为单位）
    ％T - 处理请求的时间（以秒为单位）
    ％I （大写的i） - 当前请求的线程名称
```

#### 2、JVM 参数优化

```shell
[root@java-tomcat1 conf]# cd ../bin/
[root@java-tomcat1 bin]# cp catalina.sh catalina.sh.bak
[root@java-tomcat1 bin]# vim catalina.sh
JAVA_OPTS="$JAVA_OPTS -Xms1024m -Xmx1024m -XX:PermSize=512m -XX:MaxPermSize=512m"  #jdk1.7
JAVA_OPTS="$JAVA_OPTS -Xms1024m -Xmx1024m -XX:MetaspaceSize=512m -XX:MaxMetaspaceSize=512m"   #jdk1.8
java -jar -Xms1024m -Xmx1024m -XX:MetaspaceSize=512m -XX:MaxMetaspaceSize=512m /root/666.jar

服务器配置变更。降本增效
2c4g-->1c2g
堆栈2G--->堆栈（忘了改）2G
```

#### 3、开启GC日志

```shell
# vim catalina.sh
JAVA_OPTS="$JAVA_OPTS -XX:+PrintGCDetails -XX:+PrintGCDateStamps -Xloggc:/data/logs/gc-%t.log"

可选参数:
-XX:+AggressiveOpts，加快编译。
-XX:+UseParallelGC，优化垃圾回收。
[root@java-tomcat1 bin]# mkdir /data/logs
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736080704-a95b7c0f-1c77-4317-b1a7-688bab9bfae1.png#averageHue=%23002a52&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=69&id=u7c132770&name=image.png&originHeight=86&originWidth=806&originalType=binary&ratio=1&rotation=0&showTitle=false&size=30166&status=done&style=none&taskId=u701bdde4-f222-4975-b01e-12bcb6cd1e4&title=&width=644.8)

#### 4、开启JMX端口便于监控

```shell
# vim catalina.sh
CATALINA_OPTS="$CATALINA_OPTS -Dcom.sun.management.jmxremote 
-Dcom.sun.management.jmxremote.port=10028 
-Dcom.sun.management.jmxremote.authenticate=false
-Dcom.sun.management.jmxremote.ssl=false 
-Djava.rmi.server.hostname=java69-matrix.zeus.lianjia.com"
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736090156-9685ba91-4ac6-43d8-84eb-dc942c372847.png#averageHue=%23012a51&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=107&id=u7eb80003&name=image.png&originHeight=134&originWidth=807&originalType=binary&ratio=1&rotation=0&showTitle=false&size=55768&status=done&style=none&taskId=u455aace4-8771-41cf-91f0-ace42037f98&title=&width=645.6)

#### 5、取消JVM 的默认DNS缓存时间

不缓存DNS记录，避免DNS解析更改后要重启JVM虚拟机

```shell
# catalina.sh  ---添加如下内容
CATALINA_OPTS="$CATALINA_OPTS -Dsun.net.inetaddr.ttl=0 -Dsun.net.inetaddr.negative.ttl=0
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736105297-dee4e169-05d7-42f7-94b7-dc077fb3b65f.png#averageHue=%23012951&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=122&id=u30bb1f57&name=image.png&originHeight=153&originWidth=807&originalType=binary&ratio=1&rotation=0&showTitle=false&size=67575&status=done&style=none&taskId=u79e77173-46ad-42a0-9a54-6556bb4e118&title=&width=645.6)

### 8、JVM 运维实用排障工具

**1、jps**

```shell
用来查看Java进程的具体状态, 包括进程ID，进程启动的路径及启动参数等等，与unix上的ps类似，只不过jps是用来显示java进程，可以把jps理解为ps的一个子集。
常用参数如下:
-q：忽略输出的类名、Jar名以及传递给main方法的参数，只输出pid
-m：输出传递给main方法的参数，如果是内嵌的JVM则输出为null
-l：输出完全的包名，应用主类名，jar的完全路径名
-v：输出传给jvm的参数

注意: 使用jps 时的运行账户要和JVM 虚拟机启动的账户一致。若启动JVM虚拟机是运行的账户为www，那使用jps指令时，也要使用www 用户去指定。 sudo -u www jps
```

Example

```shell
// 查看已经运行的JVM 进程的实际启动参数
[root@java-tomcat1 ~]# jps -v 
58154 Jps -Denv.class.path=.:/usr/local/java/lib:/usr/local/java/jre/lib:/usr/local/java/lib/tools.jar -Dapplication.home=/usr/local/java -Xms8m
58015 Bootstrap -Djava.util.logging.config.file=/data/application/tomcat/conf/logging.properties -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager -Djdk.tls.ephemeralDHKeySize=2048 -Djava.protocol.handler.pkgs=org.apache.catalina.webresources -Dorg.apache.catalina.security.SecurityListener.UMASK=0027 -Dignore.endorsed.dirs= -Dcatalina.base=/data/application/tomcat -Dcatalina.home=/data/application/tomcat -Djava.io.tmpdir=/data/application/tomcat/temp
```

#### 2、jstack

```
jstack用于打印出给定的java进程ID或core file或远程调试服务的Java堆栈信息。如果现在运行的java程序呈现hung的状态，jstack是非常有用的。此信息通常在运维的过程中被保存起来(保存故障现场)，以供RD们去分析故障。
常用参数如下:
jstack <pid>
jstack [-l] <pid> //长列表. 打印关于锁的附加信息
jstack [-F] <pid> //当’jstack [-l] pid’没有响应的时候强制打印栈信息
```

Example

```shell
// 打印JVM 的堆栈信息，以供问题排查
[root@mouse03 ~]# jstack -F 38360 > /tmp/jstack.log
```

#### 9、Tomcat安全优化

##### 1、telnet管理端口保护（强制）

| **类别**           | **配置内容及说明**                                           | **标准配置**                                      | **备注**                                                     |
| ------------------ | ------------------------------------------------------------ | ------------------------------------------------- | ------------------------------------------------------------ |
| telnet管理端口保护 | 1.修改默认的8005管理端口为不易猜测的端口（大于1024）；2.修改SHUTDOWN指令为其他字符串； | <Server port="**8527**" shutdown="**dangerous**"> | 1.以上配置项的配置内容只是建议配置，可以按照服务实际情况进行合理配置，但要求端口配置在**8000~8999**之间； |


##### 2、 ajp连接端口保护（推荐）

| **类别**         | **配置内容及说明**                                           | **标准配置**                                    | **备注**                                                     |
| ---------------- | ------------------------------------------------------------ | ----------------------------------------------- | ------------------------------------------------------------ |
| Ajp 连接端口保护 | 1.修改默认的ajp 8009端口为不易冲突的大于1024端口；2.通过iptables规则限制ajp端口访问的权限仅为线上机器； | <Connector port="**8528**"protocol="AJP/1.3" /> | 以上配置项的配置内容仅为建议配置，请按照服务实际情况进行合理配置，但要求端口配置在**8000~8999**之间；；保护此端口的目的在于防止线下的测试流量被mod_jk转发至线上tomcat服务器； |


##### 3、降权启动（强制）

| **类别** | **配置内容及说明**                                           | **标准配置** | **备注**                                                     |
| -------- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ |
| 降权启动 | 1.tomcat启动用户权限必须为非root权限，尽量降低tomcat启动用户的目录访问权限；2.如需直接对外使用80端口，可通过普通账号启动后，配置iptables规则进行转发； |              | 避免一旦tomcat 服务被入侵，黑客直接获取高级用户权限危害整个server的安全； |


```shell
[root@java-tomcat1 ~]# useradd tomcat
[root@java-tomcat1 ~]# chown tomcat.tomcat /data/application/tomcat/ -R
[root@java-tomcat1 ~]# su -c '/data/application/tomcat/bin/startup.sh' tomcat 
Using CATALINA_BASE:   /data/application/tomcat
Using CATALINA_HOME:   /data/application/tomcat
Using CATALINA_TMPDIR: /data/application/tomcat/temp
Using JRE_HOME:        /usr/local/java
Using CLASSPATH:       /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar
Tomcat started.
[root@java-tomcat1 ~]# ps -ef | grep tomcat 
tomcat     1065      1 64 20:33 ?        00:00:06 /usr/local/java/bin/java -Djava.util.logging.config.file=/data/applicationtomcat/conf/logging.properties -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager -Djdk.tls.ephemeralDHKeySize=2048 -Djava.protocol.handler.pkgs=org.apache.catalina.webresources -Dorg.apache.catalina.security.SecurityListener.UMASK=0027 -Dignore.endorsed.dirs= -classpath /data/application/tomcat/bin/bootstrap.jar:/data/application/tomcat/bin/tomcat-juli.jar -Dcatalina.base=/data/application/tomcat -Dcatalina.home=/data/application/tomcat -Djava.io.tmpdir=/data/application/tomcat/temp org.apache.catalina.startup.Bootstrap start
root       1112   1027  0 20:33 pts/0    00:00:00 grep --color=auto tomcat
```

##### 4、文件列表访问控制（强制）

![image.png](https://cdn.nlark.com/yuque/0/2022/png/25576613/1669736166111-bb4317b6-a096-4f4c-bad1-9d8af9fb410f.png#averageHue=%23f2eded&clientId=u4a23e747-7684-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=518&id=uee836fc5&name=image.png&originHeight=648&originWidth=1417&originalType=binary&ratio=1&rotation=0&showTitle=false&size=284161&status=done&style=none&taskId=ud5609173-fb9f-4465-a76f-1566909c0ba&title=&width=1133.6)

| **类别**         | **配置内容及说明**                                         | **标准配置**                                                 | **备注**                                             |
| ---------------- | ---------------------------------------------------------- | ------------------------------------------------------------ | ---------------------------------------------------- |
| 文件列表访问控制 | 1.conf/web.xml文件中default部分listings的配置必须为false； | <init-param><param-name>**listings**</param-name><param-value>**false**</param-value></init-param> | false为不列出目录文件，true为允许列出，默认为false； |

**5、起停脚本权限回收（推荐）**

| **类别**         | **配置内容及说明**                                           | **标准配置或操作**        | **备注**                             |
| ---------------- | ------------------------------------------------------------ | ------------------------- | ------------------------------------ |
| 起停脚本权限回收 | 去除其他用户对Tomcat的bin目录下shutdown.sh、startup.sh、catalina.sh的可执行权限； | chmod -R 744 tomcat/bin/* | 防止其他用户有起停线上Tomcat的权限； |


##### 6、 访问日志格式规范（推荐）

conf/server.xml

| **类别**         | **配置内容及说明**                                | **标准配置或操作**                                           | **备注**                                                     |
| ---------------- | ------------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 访问日志格式规范 | 开启Tomcat默认访问日志中的Referer和User-Agent记录 | <Valve className="org.apache.catalina.valves.AccessLogValve"                 directory="logs"  prefix="localhost_access_log." suffix=".txt"                 pattern="%h %l %u %t %r %s %b %{Referer}i %{User-Agent}i %D" resolveHosts="false"/> | 开启Referer和User-Agent是为了一旦出现安全问题能够更好的根据日志进行问题排查； |


#### 10、Tomcat性能优化

**上策：优化代码**

该项需要开发经验足够丰富，对开发人员要求较高

**中策：jvm优化机制** **垃圾回收机制** **把不需要的内存回收**

优化jvm--优化垃圾回收策略

优化catalina.sh配置文件。在catalina.sh配置文件中添加以下代码

```shell
# tomcat分配1G内存模板
JAVA_OPTS="-Djava.awt.headless=true -Dfile.encoding=UTF-8 -server -Xms1024m -Xmx1024m -XX:NewSize=512m -XX:MaxNewSize=512m -XX:PermSize=512m -XX:MaxPermSize=512m"     

# 重启服务
su -c '/home/tomcat/tomcat8_1/bin/shutdown.sh' tomcat
su -c '/home/tomcat/tomcat8_1/bin/startup.sh' tomcat
```

**下策：加足够大的内存**      灰度更新

该项的资金投入较大      

1c2g   1c4g

2c4g 2c8g 

4c16g  

**下下策：每天0点定时重启tomcat**

使用较为广泛