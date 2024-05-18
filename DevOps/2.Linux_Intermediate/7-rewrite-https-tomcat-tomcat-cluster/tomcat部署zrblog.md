二进制安装

```
yum install java-11-openjdk-devel -y
mkdir /app
tar xf jdk-8u60-linux-x64.tar.gz -C /app/
ln -s /app/jdk1.8.0_60/   /app/jdk

vim /etc/profile
export JAVA_HOME=/app/jdk
export PATH=$JAVA_HOME/bin:$JAVA_HOME/jre/bin:$PATH
export CLASSPATH=$CLASSPATH:$JAVA_HOME/lib:$JAVA_HOME/jre/lib:$JAVA_HOME/lib/tools.jar
source ~/.bash_profile
```

启动tomcat

```
tar xf tomcat-9.0.26.tar.gz -C /app/
/app/apache-tomcat-9.0.26/bin
./startup.sh
```

先为tomcat新增一个host虚拟主机

```
vim /app/apache-tomcat-9.0.26/conf/server.xml
...
    <Engine>
      <Host name="localhost"  appBase="/code" unpackWARs="true" autoDeploy="true">
        <Valve className="org.apache.catalina.valves.AccessLogValve" directory="logs"
               prefix="zrlog_access_log" suffix=".txt"
               pattern="%h %l %u %t &quot;%r&quot; %s %b" />
      </Host>
    </Engine>
...
```

安装mariadb

```
yum install mariadb mariadb-server -y
systemctl start mariadb
mysqladmin password '1'
mysql -u root -p1
```

使用远程数据库需要配置授权

```
create database zrlog charset utf8;
grant all privileges on *.* to tomcat@'%' identified by 'tomcat';
```

上传zrlog

```
[root@192 ROOT]# ls
admin  assets  error  favicon.ico  include  META-INF  WEB-INF
[root@192 ROOT]# pwd
/code/ROOT
```

启动tomcat

```
/app/apache-tomcat-9.0.26/bin
./startup.sh
```