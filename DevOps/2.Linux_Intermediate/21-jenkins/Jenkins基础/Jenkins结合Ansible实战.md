# Jenkins结合Ansible实战

## Jenkins机器安装ansbile

```shell
[root@jenkins ~]# yum -y install epel-release ; yum -y install ansible
[root@jenkins ~]# ansible --version
```

![img](assets/Jenkins结合Ansible实战/1679016832101-f7861920-12df-4400-a30b-8203947a4d29.png)

## 配置ansible主机清单

```shell
[root@jenkins ~]# cat /etc/hosts
192.168.91.135 tomcat1
[root@jenkins ~]# cat /etc/ansible/hosts
[tomcat-server]
tomcat1
[root@jenkins ~]# ssh-keygen  #一路回车
[root@jenkins ~]# ssh-copy-id  tomcat1
[root@jenkins ~]# ansible  tomcat1 -m ping
```

![img](assets/Jenkins结合Ansible实战/1652604647061-95d77ff5-08db-41f1-a036-668471340993.png)

## jenkins安装ansible插件

![img](assets/Jenkins结合Ansible实战/1652603544389-7267f901-0756-4e88-b959-f3eed81c653a.png)

![img](assets/Jenkins结合Ansible实战/1652603568650-04b8061c-d6d4-4a09-9085-c31dc44dd169.png)

## jenkins配置ansible工具

![img](assets/Jenkins结合Ansible实战/1652603669517-940d69e3-a6f5-4eea-aef2-250f4ae62778.png)

## Jenkins使用ansible命令

![img](assets/Jenkins结合Ansible实战/1652603730238-be015a61-17ad-48c4-b9f8-5d2747a4ace1.png)

![img](assets/Jenkins结合Ansible实战/1652604195896-759d4b45-a307-4dc8-a5eb-a5e9caba8c21.png)

![img](assets/Jenkins结合Ansible实战/1652604809053-a306df5b-51c2-4494-a070-4df4a6433ade.png)

![img](assets/Jenkins结合Ansible实战/1652604830834-831ca24a-7523-47de-9928-7b56ccb3c504.png)

![img](assets/Jenkins结合Ansible实战/1652604727453-dcdc0fdb-0812-4449-b624-ed23259b9a8a.png)



## jenkins使用ansible-playbook剧本

```shell
[root@jenkins ~]# cat /etc/ansible/test.yaml 
---
- hosts: tomcat-server
  tasks:
  - name: ip
    shell: ip a
    register: result   # 显示结果为"result"
  - debug: var=result  # 将显示结果var赋值给result
```

![img](assets/Jenkins结合Ansible实战/1652605596723-393c498f-d5e3-4186-99dc-1a050ef5f35c.png)

![img](assets/Jenkins结合Ansible实战/1652607436031-2dcb7e30-c9d0-4180-8460-9987e1b12102.png)

![img](assets/Jenkins结合Ansible实战/1652607665526-d896ed6e-3380-43c2-afb2-b240a7ba9768.png)

## jenkins+ansible构建maven项目

### jenkins安装git

```shell
[root@jenkins ~]# yum -y install git
```

### 添加maven和jdk工具

![img](assets/Jenkins结合Ansible实战/1652610261361-8a0f6452-ffe9-428e-968b-cbbf7b1e2db0.png)

![img](assets/Jenkins结合Ansible实战/1652610273174-4e16bf3f-503a-4576-aa25-f4a085d28308.png)

### jenkins配置gitee公钥

```shell
[root@jenkins ~]# cat /root/.ssh/id_rsa.pub
```

![img](assets/Jenkins结合Ansible实战/1652609908349-dc52acf4-d8c2-4ace-8da3-80655ecff70d.png)

### 项目配置

![img](assets/Jenkins结合Ansible实战/1652610074494-3c9b8747-df51-41c1-ab62-87af9530d62f.png)

![img](assets/Jenkins结合Ansible实战/1652611355973-6b7c91f1-3e05-4fff-9e07-4d60f54adccf.png)

### git参数化构建

![img](assets/Jenkins结合Ansible实战/1652611383994-0a35ad2f-3cdb-4e79-87ef-a274b2e8cea2.png)

```shell
[root@jenkins ~]# cat /root/.ssh/id_rsa
```

![img](assets/Jenkins结合Ansible实战/1652609865581-be69c4bb-063f-4891-a237-37aeedb7d2b3.png)

gitee仓库地址：git@gitee.com:youngfit/easy-springmvc-maven.git

![img](assets/Jenkins结合Ansible实战/1652611795725-93970a10-b15a-4763-8fbe-2b2564df350e.png)

### 配置maven构建参数

![img](assets/Jenkins结合Ansible实战/1652611875771-8a6bc43a-9bd4-4016-af13-79901d34b40d.png)

### 配置ansible运行指令

![img](assets/Jenkins结合Ansible实战/1652611979490-53f61480-86e1-4928-a31b-88528b9e2e28.png)

![img](assets/Jenkins结合Ansible实战/1652612130399-769218e8-17af-4279-953e-a4e7917abb17.png)

![img](assets/Jenkins结合Ansible实战/1652612028320-22b7d1f0-2d47-4ddc-984f-e97bfb8e4c14.png)

![img](assets/Jenkins结合Ansible实战/1652612125644-6eba4ed2-d441-4c3c-8a0d-f87c15d95dca.png)

-->保存

### tomcat服务器配置

#### 安装jdk和tomcat

#### 配置环境变量

```shell
安装jdk8和tomcat
[root@tomcat local]# cat /etc/profile
TOMCAT_HOME=/data/application/tomcat
JAVA_HOME=/usr/local/java
MAVEN_HOME=/usr/local/java/maven
CLASSPATH=.:$JAVA_HOME/lib:$JAVA_HOME/jre/lib:$JAVA_HOME/lib/tools.jar
PATH=$PATH:$JAVA_HOME/bin:$MAVEN_HOME/bin
export PATH=$JAVA_HOME/bin:$JAVA_HOME/jre/bin:$PATH
```

#### 配置脚本

```shell
[root@tomcat local]# cat /opt/script/app-jenkins.sh
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

### 构建测试

![img](assets/Jenkins结合Ansible实战/1652612259599-3cc9d776-4b04-45e8-bff1-039e611e8a9f.png)

![img](assets/Jenkins结合Ansible实战/1652612288024-409c6f1a-1b92-41ac-9115-1754b8a3b24b.png)

### 验证

![img](assets/Jenkins结合Ansible实战/1652612332300-e4b305a3-214a-467f-b72d-ceecb03c64ad.png)

![img](assets/Jenkins结合Ansible实战/1652612358182-e6386625-b08f-4a44-b494-ec6bb5cab75a.png)

来一次版本回退试试

![img](assets/Jenkins结合Ansible实战/1652612541841-e3e2a74b-797d-4641-a24a-1819f77cecb7.png)

![img](assets/Jenkins结合Ansible实战/1652612529056-cbf855c2-df50-47cf-b35a-91c305ac3505.png)

没毛病