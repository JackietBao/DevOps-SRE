# 【百练成魔】Jenkins CI/CD

### 架构 Jenkins + Gitlab + Docker + Harbor + Maven + SonarQube

### 流程图

![image-20230914140232268](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914140232268.png)

### 服务器准备

虚拟机虚拟化出一台 4C 16G 80G  服务器 本次所有项目均使用Docker部署都在单节点运行，也可以拆解分配项目



##### 安装docker

关闭firewalld seliunx防火墙

```
systemctl  stop firewalld  && systemctl disable firewalld  && setenforce 0 
```

安装docker及docker-compose

安装必要系统工具

```
yum install -y yum-utils device-mapper-persistent-data lvm2
```

添加软件源信息

```
yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

安装docker

```
yum -y install docker-ce
```

##### 安装docker-compose

仓库地址

```
https://github.com/docker/compose/releases
```

下载compose // 由于github是国外的所以可能下载失败 需要多试几次，或使用主机开启魔法下载

```
wget https://github.com/docker/compose/releases/download/v2.20.2/docker-compose-linux-x86_64
```

```
chmod +x docker-compose-linux-x86_64  
mv docker-compose-linux-x86_64  /usr/bin/docker-compose
```

![image-20230914142622885](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914142622885.png)

启动docker 并设置开机自启 检查版本

```
systemctl  start docker 
systemctl  enable  docker 

docker --version 
docker-compose  --version 
```

![image-20230914142922896](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914142922896.png)

到这里docker，docker-compose安装完毕



##### 安装gitlab

使用docker-compose部署

拉取镜像

```
docker pull gitlab/gitlab-ce
```

创建gitlab目录，使用本目录作为gitlab挂载目录存放数据文件

```
mkdir /usr/local/gitlab
cd /usr/local/gitlab
```

准备docker-compose文件

```
vim docker-compose.yml
```

```
version: '3.1'
services:
  gitlab:
    image: 'gitlab/gitlab-ce:latest'
    container_name: gitlab
    restart: always
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://192.168.117.141:8929'   ##修改为宿主机IP
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - '8929:8929'
      - '2224:2224'
    volumes:
      - './config:/etc/gitlab'
      - './logs:/var/log/gitlab'
      - './data:/var/opt/gitlab'
```

启动容器

```
docker-compose up -d
```

查看是否启动

```
docker ps 
```

![image-20230914144447781](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914144447781.png)

访问IP：8929  若出现502则需要继续等待初始化一会

查看初始化root密码

```
docker exec -it gitlab cat /etc/gitlab/initial_root_password
```

![image-20230914144631490](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914144631490.png)



登录gitlab 

![image-20230914144733671](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914144733671.png)

修改密码 点击头像

![image-20230914144833298](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914144833298.png)

点击 preferences

![image-20230914144905745](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914144905745.png)

点击 password  修改完点击 Save password 保存

![image-20230914145022037](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145022037.png)

保存后会退出 重新登录就可以了 使用你设置的新密码

![image-20230914145147302](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145147302.png)

点击首页 Create a projeck

创建新项目

![image-20230914145248328](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145248328.png)

创建空白仓库

![image-20230914145329494](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145329494.png)

创建空仓库

![image-20230914145723865](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145723865.png)

###### 推送代码到仓库

安装git 用来拉取仓库

```
yum -y install git
```

仓库点击clone   复制 http拉取代码方式

![image-20230914145951398](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914145951398.png)



回到用户家目录 拉取代码

```
cd
git clone http://192.168.117.141:8929/root/devops-test.git  ##git clone 后面粘贴复制的仓库地址
```

配置git推送用户名及邮箱 ##随便写就可

```
git config --global user.name "Administrator"
git config --global user.email "admin@example.com"
```

进入目录并切换到main分支

```
cd /devops-test/
git checkout -b main
```

拉取我在gitee放的代码

```
git clone https://gitee.com/jd_688/hell.git -b main
```

移动gitee代码到 当前目录下 并删除gitee目录

```
mv hell/* ./ 
rm -rf hell/
```

![image-20230914150850065](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914150850065.png)

提交代码到gitlab仓库

```
git add * 
git commit -m "推送代码"
git push -u origin main
```

输入gitlab账号密码 即可推送成功 # 若密码输入错误可再次执行 git push -u origin main 推送命令 

![image-20230914151620728](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914151620728.png)

回到仓库点击刷新就可以看到推送上来的代码

![image-20230914151855994](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914151855994.png)



##### 安装Jenkins

创建jenkins目录及数据持久化目录

```
mkdir /usr/local/jenkins
cd /usr/local/jenkins
```

拉取Jenkins镜像

```
docker pull jenkins/jenkins
```

编写启动docker-compose文件

```
vim docker-compose.yml
```

```
version: "3.1"
services:
  jenkins:
    image: jenkins/jenkins
    container_name: jenkins
    restart: always
    ports:
      - 8080:8080
      - 50000:50000
    volumes:
      - ./data/:/var/jenkins_home/
```

启动Jenkins容器

```
docker-compose up -d 
```

![image-20230914152427321](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914152427321.png)

第一次启动会启动失败 是因为对数据库data没有读权限 我们加一下就可以了

```
chmod -R a+w data/
```

![image-20230914152744974](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914152744974.png)

再次启动jenkins容器

```
docker-compose up -d 
```

![image-20230914152824582](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914152824582.png)

先修改插件下载源，不然使用默认源会下载失败也很慢

```
cd /data 
vim vim hudson.model.UpdateCenter.xml
```

```
<?xml version='1.1' encoding='UTF-8'?>
<sites>
  <site>
    <id>default</id>
    <url>http://mirror.esuni.jp/jenkins/updates/update-center.json</url>
  </site>
</sites>
```

![image-20230914153257821](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153257821.png)

再次重启Jenkins容器，访问JenKins：IP:8080/restart   弹出界面点击yes

例如我的机器是

```
http://192.168.117.141:8080/restart 
```

![](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153730930.png)

![image-20230914153836849](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153836849.png)



这时就可以正常使用了 访问IP：8080 进入jenkins的操作页面

![image-20230914153004228](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153004228.png)

获取密码

```
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

![image-20230914153022876](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153022876.png)

输入密码点击继续

选择插件来安装

![image-20230914153354984](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914153354984.png)

搜索这个选择上  点击安装

```
Git Parameter
```

![image-20230914154156994](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914154156994.png)

等待安装完成，耗时会有点长，待会进入系统再去安装其他所用插件

安装完成后会自动跳转界面 设置新密码即可

![image-20230914154918786](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914154918786.png)

这个默认的即可 我们就是IP：8080访问 所以不需要修改

![image-20230914154951012](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914154951012.png)

保存开始 使用Jenkins

![image-20230914155014731](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155014731.png)

到这里我们就算是部署Jenkins完毕 ，下面开始配工程

![image-20230914155249314](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155249314.png)

创建后进入配置 选择源码管理 Git

![image-20230914155406756](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155406756.png)

###### 公开仓库http配置

回到gitlab仓库复制仓库地址

![image-20230914155511562](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155511562.png)

填入仓库地址看到没有报错  指定分支 填写我们创建的 main分支 点击保存即可

![image-20230914160450556](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914160450556.png)

###### 私有仓库http配置

公开的仓库直接跳过这一步

私有仓库直接填写地址 下面会出现红色报错 让你提供认证秘钥 我这里是公开仓库所以是没有显示的

配置密码的方式如下，配置秘钥这里就不写了

![image-20230914155704662](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155704662.png)

![image-20230914155943305](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914155943305.png)

添加选择就可以了  并修改指定配置为 main分支

![image-20230914160553225](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914160553225.png)

继续我们先使用jenkins拉取代码

点击构建 完成出现绿色对号 

![image-20230914160727689](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914160727689.png)

也可以点击#2  点击控制台输出 查看日志  我是构建过一次所以是2#  第一构建的话都是1#

![image-20230914160834071](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914160834071.png)

看到日志说 我们已经把代码拉取下来了 我们到服务器查看一下

这里目录/devops-test   就是jenkins创建工程的名字

```
cd /usr/local/jenkins/data/workspace/devops-test
ls
```

可以看到代码已经拉取到服务器了

![image-20230914160959550](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914160959550.png)

接下来使用maven打包代码为jar包

###### 配置jdk 和 maven

```
下载地址
jdk     https://d6.injdk.cn/openjdk/openjdk/17/openjdk-17.0.1_linux-x64_bin.tar.gz
maven   https://dlcdn.apache.org/maven/maven-3/3.9.4/binaries/apache-maven-3.9.4-bin.tar.gz
```

下载完上传到服务器/root目录

解压到Jenkins挂载目录让Jenkins使用

```
tar xf apache-maven-3.9.4-bin.tar.gz  -C /usr/local/jenkins/data/
tar xf openjdk-17.0.1_linux-x64_bin.tar.gz  -C /usr/local/jenkins/data/
cd /usr/local/jenkins/data/
mv apache-maven-3.9.4/ maven 
mv jdk-17.0.1/ jdk17
```

此时容器中/var/jenkins_home中 maven和jdk17 已经挂载进来了

![image-20230914162339062](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914162339062.png)



修改maven配置文件

```
vim /usr/local/jenkins/data/maven/conf/settings.xml 
```

改为以下内容

```
<?xml version="1.0" encoding="UTF-8"?>

<settings xmlns="http://maven.apache.org/SETTINGS/1.2.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.2.0 https://maven.apache.org/xsd/settings-1.2.0.xsd">
  <pluginGroups>
  </pluginGroups>
  <proxies>
  </proxies>
  <servers>
  </servers>

  <mirrors>
    <mirror>
    <id>aliyunmaven</id>
    <mirrorOf>*</mirrorOf>
    <name>阿里云公共仓库</name>
    <url>https://maven.aliyun.com/repository/public</url>
    </mirror>
  </mirrors>

  <profiles>
  </profiles>

</settings>

```



点击Dashboard 回到首页 点击Manage Jenkins

![image-20230914162521833](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914162521833.png)

点击Tookl 配置全局配置

![image-20230914162602601](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914162602601.png)

往下拉 看到maven安装

![image-20230914162813652](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914162813652.png)

![image-20230914162936212](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914162936212.png)

点击保存 回到主页 选择devops-test这个工程 

![image-20230914163037223](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914163037223.png)

配置 Build Steps 

调用顶层Maven目录 选择maven 目标输入下面内容  点击保存

```
clean  package
```

![image-20230914163954844](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914163954844.png)

再次构建 点击开始构建 首次编译需要下载插件所以需要时间会有点长

点击2# 点击控制台输出 查看日志

![image-20230914165238112](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914165238112.png)

Build Success表示打包完成 

查看jar包

```
cd /usr/local/jenkins/data/workspace/devops-test
ls target/
```

![image-20230914165332801](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914165332801.png)

使用Java 运行跑一下

```
/usr/local/jenkins/data/jdk17/bin/java -jar target/testinit-0.0.1-SNAPSHOT.jar
```

![image-20230914165502116](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914165502116.png)

浏览器访问IP：9901

![image-20230914165535426](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914165535426.png)

看完以后就可以 ctrl +c 取消运行jar包



###### 配置Git参数化构建

回到工程 配置

![image-20230914170414944](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914170414944.png)

看图配置 把构建操作×掉

![image-20230914170515995](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914170515995.png)

重新选择执行shell

![image-20230914171841976](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171841976.png)

```
cd /var/jenkins_home/workspace/${JOB_NAME}
git checkout $tag
/var/jenkins_home/maven/bin/mvn clean package
```

```
参数解释

${JOB_NAME} 为当前工程名字 也就是devops-test
$tag 就是上面配置tag标签
```

点击保存 

回到仓库目录 修改代码 作为标识

```
cd /root/devops-test
vim src/main/java/com/example/testinit/controller/HelloWord.java
```

![image-20230914171021570](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171021570.png)

修改完 推送代码

```
git add * 
git commit -m " create v1.0.0" 
git tag -a v1.0.0 -m "create v1.0.0" 
git push -u origin v1.0.0
```

![image-20230914171312113](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171312113.png)

刷新仓库 点击tag 可以看到v1.0.0已经推送上来了

![image-20230914171430042](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171430042.png)

![image-20230914171440437](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171440437.png)

回到Jenkins项目中  点击Build 可以看到tag版本已经出来了 选中v1.0.0 开始构建

![image-20230914171903819](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171903819.png)

查看工程的控制台输出

![image-20230914171952403](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914171952403.png)

再次查看jar包

```
cd /usr/local/jenkins/data/workspace/devops-test
ls target/
```

![image-20230914165332801](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914165332801.png)

使用Java 运行跑一下

```
/usr/local/jenkins/data/jdk17/bin/java -jar target/testinit-0.0.1-SNAPSHOT.jar
```

![image-20230914172125023](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914172125023.png)

浏览器访问IP：9901

![image-20230914172140853](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914172140853.png)

可以看到 V1.0.0已经出来了  ctrl+c 取消运行



##### 安装SonarQube

拉取镜像

```bash
docker pull postgres
docker pull sonarqube:8.9.3-community
```

创建SonarQube目录

```
mkdir /usr/local/sonarQube
cd /usr/local/sonarQube
```

编写docker-compose启动文件

```
vim docker-compose.yml
```

```
version: "3.1"
services:
  db:
    image: postgres
    container_name: db
    ports:
      - 5432:5432
    networks:
      - sonarnet
    environment:
      POSTGRES_USER: sonar
      POSTGRES_PASSWORD: sonar
  sonarqube:
    image: sonarqube:8.9.3-community
    container_name: sonarqube
    depends_on:
      - db
    ports:
      - "9000:9000"
    networks:
      - sonarnet
    environment:
      SONAR_JDBC_URL: jdbc:postgresql://db:5432/sonar
      SONAR_JDBC_USERNAME: sonar
      SONAR_JDBC_PASSWORD: sonar
networks:
  sonarnet:
    driver: bridge
```

修改系统文件最大打开数量

```
vim /etc/sysctl.conf
```

```
vm.max_map_count = 262144
```

刷新

```
sysctl  -p
```

![image-20230914173023364](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173023364.png)

启动sonar

```
docker-compose up -d 
```

![image-20230914173205235](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173205235.png)

启动需要一会时间

访问IP：9000 

![image-20230914173506212](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173506212.png)

登录进去需要重新修改密码

![image-20230914173538413](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173538413.png)

修改语言为中文   第三部 I understand the resk 允许后 chinses pack 后面会出现install 点击安装即可

![image-20230914173640592](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173640592.png)

安装完需要重启 弹出的页面 Restart Server 

![image-20230914173828388](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173828388.png)

再次登录进来就是中文

![image-20230914173916488](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914173916488.png)

###### 手动Sonar-scanner实现代码检测

下载 上传到/root

```
https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.6.2.2472.zip
```

解压

```
yum -y install unzip 
unzip sonar-scanner-cli-4.6.2.2472.zip
mv sonar-scanner-4.6.2.2472/ /usr/local/jenkins/data/sonar-scanner 
```

修改配置

```
vim /usr/local/jenkins/data/sonar-scanner/conf
```

取消注释 改为主机IP

![image-20230914174659763](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914174659763.png)

使用手动命令检测代码

配置Java环境 

```
vim /etc/profile
```

```
export JAVA_HOME=/usr/local/jenkins/data/jdk17
export PATH=$PATH:$JAVA_HOME/bin
export MAVEN_HOME=/usr/local/jenkins/data/maven
export PATH=$PATH:$MAVEN_HOME/bin
```

```
source /etc/profile
```

![image-20230914180110736](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180110736.png)

```
java -version 
mvn -version
```

![](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180148482.png)

![image-20230914180010245](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180010245.png)

获取令牌

![image-20230914175159961](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914175159961.png)

进入工程目录

```
cd /usr/local/jenkins/data/workspace/devops-test
```

```
/usr/local/jenkins/data/sonar-scanner/bin/sonar-scanner -Dsonar.sources=./ -Dsona.projectname=devops-test -Dsonar.projectKey=devops-test -Dsonar.java.binaries=target/ -Dsonar.login=fe2bb1043bd317199a3919c8d78d05fc2daa4d4c
```

```
参数解释
sonar.sources=./  #指定目录
sona.projectname=devops-test  #指定工程名
sonar.projectKey=devops-test  #指定sonar项目名
sonar.java.binaries=target/   #指定检测jar包存放目录
sonar.login=fe2bb1043bd317199a3919c8d78d05fc2daa4d4c  #指定sonar生成的秘钥
```

![image-20230914180235428](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180235428.png)

![image-20230914180250385](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180250385.png)

检查完成 回到

sonar 

![image-20230914180328408](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230914180328408.png)

这样就可以看到代码的检查详情  这个开发看就行 我们不需要看

###### Jenkins 调用Sonar-scaner自动检测

系统管理，插件管理  安装以下插件

```
SonarQube Scanner
```

![image-20230915090347074](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915090347074.png)

系统管理，全局工具管理

sonarqube scanner

取消自动安装输入容器中的sonarqube目录 点击保存

```
/var/jenkins_home/sonar-scanner
```

![image-20230915090639003](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915090639003.png)

系统管理，System 中 添加Sonarqube servers

![image-20230915092609129](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915092609129.png)

![image-20230915092812476](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915092812476.png)

来到工程项目中 构建中 增加构建步骤，选中 Execute SonarQube Scanner

![image-20230915091230064](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915091230064.png)

```
sonar.projectname=${JOB_NAME}
sonar.projectKey=${JOB_NAME}
sonar.source=./
sonar.java.binaries=target/
```

![image-20230915092934135](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915092934135.png)

保存后 开始构建

这时会报错因为我们手动检测过一次 会生成一个目录，这个用户对这个目录没有写权限我们去删掉即可

![image-20230915093101936](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915093101936.png) 

```
cd /usr/local/jenkins/data/workspace/devops-test
ls -a 
rm -rf .scannerwork/
```

![image-20230915093139811](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915093139811.png)

再次构建查看输出日志 已经检测完成

![image-20230915093345968](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915093345968.png)

来到Sonarqube的web界面中查看 这样就OK了

![image-20230915093428969](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915093428969.png)

##### 安装Harbor仓库

下载安装包

```
https://github.com/goharbor/harbor/releases/download/v2.3.4/harbor-offline-installer-v2.3.4.tgz
```

上传到/root

解压

```
tar -zxvf harbor-offline-installer-v2.3.4.tgz -C /usr/local/
```

```
cd /usr/local/harbor
cp harbor.yml.tmpl  harbor.yml
vim harbor.yml
```

修改为主机IP 端口我是使用7788，这个随意都行 注释https,因为没有证书

![image-20230915094019155](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915094019155.png)

保存后 执行安装脚本

```
./install
```

![image-20230915094130389](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915094130389.png)

![image-20230915094409535](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915094409535.png)

访问主页 IP：7788

```
用户admin
密码Harbor12345
```

![image-20230915094505455](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915094505455.png)

登录进来新建一个项目

![image-20230915094621166](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915094621166.png)

修改daemon.json，支持Docker仓库，并重启Docker

```
vim /etc/daemon.json
```

```
{
        "registry-mirrors": ["https://pee6w651.mirror.aliyuncs.com"],
        "insecure-registries": ["192.168.117.141:7788"]
}
```

![image-20230915095031378](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915095031378.png)

重启docker  由于sonar容器没有做随dokcer启动而启动 所以需要启动一下sonar

```
systemctl  restart docker
docker start $(docker ps -qa)
```

###### Jenkins容器使用宿主机Docker并编写构建脚本

设置宿主机docker.sock权限

```bash
chown root:root /var/run/docker.sock
chmod o+rw /var/run/docker.sock
```

挂载宿主机docker到容器

```
vim /usr/local/jenkins/docker-compose.yml 
```

```
version: "3.1"
services:
  jenkins:
    image: jenkins/jenkins
    container_name: jenkins
    ports:
      - 8080:8080
      - 50000:50000
    volumes:
      - ./data/:/var/jenkins_home/
      - /usr/bin/docker:/usr/bin/docker
      - /var/run/docker.sock:/var/run/docker.sock
      - /etc/docker/daemon.json:/etc/docker/daemon.json
```

![image-20230915095622890](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915095622890.png)

重启Jenkins 容器

```
cd /usr/local/jenkins/
docker-compose up -d 
```



![image-20230915095657281](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915095657281.png)

进入容器查看docker挂载情况，可以看到容器中也可以使用宿主机的docker

```
docker exec -it jenkins bash 
docker ps
```

![image-20230915095814234](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915095814234.png)

###### 制作自定义镜像

回到Jenkins工程中 添加构建操作，执行shell

```
cd /var/jenkins_home/workspace/${JOB_NAME}/docker
mv ../target/*.jar ./
docker build -t ${JOB_NAME}:$tag .
docker login -u admin -p Harbor12345 192.168.117.141:7788
docker tag ${JOB_NAME}:$tag 192.168.117.141:7788/devops/${JOB_NAME}:$tag
docker push 192.168.117.141:7788/devops/${JOB_NAME}:$tag
```

```
注意这里必须要改镜像名字 否则推送不上Harboer仓库
登录密码 要写自己的账号密码,登录地址也要写仓库IP：端口    若是80端口也必须加上IP：80
```

![image-20230915110229484](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110229484.png)

###### 编写Dockerfile

进入仓库目录创建dockerfile存放目录 

```
cd /root/devops-test/
mkdir docker 
vim docker/Dockerfile
```

![image-20230915100905812](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915100905812.png)

```
FROM java:openjdk-8u111
COPY *.jar /usr/local
WORKDIR /usr/local
CMD java -jar *.jar
```

![image-20230915101829279](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915101829279.png)

修改标识

```
vim src/main/java/com/example/testinit/controller/HelloWord.java
```

![image-20230915102808929](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915102808929.png)

推送代码到仓库

```
git add * 
git commit -m "create v1.0.1"
git tag -a v1.0.1 -m "create v1.0.1"
git push -u origin v1.0.1
```

![image-20230915103046912](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915103046912.png)



再次提交v1.0.2版本

```
vim src/main/java/com/example/testinit/controller/HelloWord.java
```

![image-20230915103214193](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915103214193.png)

推送

```
git add * 
git commit -m "create v1.0.2"
git tag -a v1.0.1 -m "create v1.0.2"
git push -u origin v1.0.2
```

可以看到gitlab已经有了两个新的分支

![image-20230915103701333](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915103701333.png)

构建v1.0.1    不要构建v1.0.0 这个没有推动dockerfile

![image-20230915103752311](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915103752311.png)

看输出台日志 表示已经推送到Harbor仓库

![image-20230915110337671](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110337671.png)



![image-20230915110446706](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110446706.png)

点击进来可以看到拉取命令及版本号

![image-20230915110518577](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110518577.png)

再次到Jenkins工程构建 v1.0.2版本

![image-20230915110609054](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110609054.png)

可以看到v1.0.2版本也推送进来了

![image-20230915110640335](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110640335.png)

sonar也有检测 因为代码是一样的 所以不会出现新的东西

![image-20230915110719410](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915110719410.png)

###### 编写启动脚本

```
vim devops_start.sh
```

```
harbor_url=$1
harbor_project_name=$2
project_name=$3
tag=$4
#定义拉取的镜像名
imageName=$harbor_url/$harbor_project_name/$project_name:$tag
#删除正在运行的工程容器
containerId=`docker ps -a | grep ${project_name} | awk '{print $1}'`
if [ "$containerId" != "" ] ; then
    docker stop $containerId
    docker rm $containerId
    echo "Delete Container Success"
fi
#删除当前工程名的镜像
imageId=`docker images | grep ${project_name} | awk '{print $3}'`

if [ "$imageId" != "" ] ; then
    docker rmi -f $imageId
    echo "Delete Image Success"
fi
#登录harboer
docker login -u admin -p Harbor12345 $harbor_url
#拉取当前工程构建版本的镜像
docker pull $imageName
#创建工程容器
docker run -d -p 9901:9901 --name $project_name $imageName

echo "Start Container Success"
echo $project_name
```

```
参数解释
$1 为harbor仓库地址
$2 为harbor项目名
$3 jenkins工程名
$4 当前所用gitlab仓库中项目tag版本
```

![image-20230915111352208](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915111352208.png)

全局可执行

```
mv devops_start.sh  /usr/bin/
chmod a+x /usr/bin/devops_start.sh
```

```
docker images
```

系统管理，插件管理

安装

```
Publish Over SSH
```

![image-20230915112438146](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915112438146.png)

##### 配置服务器

获取Jenkins容器中的key

```
docker exec  -it jenkins bash
ssh-keygen
一路回车
cat /var/jenkins_home/.ssh/id_rsa
```

![image-20230915112812350](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915112812350.png)

复制key

![image-20230915113100805](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915113100805.png)

发送秘钥到后端服务器

```
ssh-copy-id  root@192.168.117.141
yes
输出root密码
```

![image-20230915113414988](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915113414988.png)

系统管理，系统配置  Publish over SSH

![image-20230915113130148](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915113130148.png)



添加后端服务器

![image-20230915113551481](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915113551481.png)



配置jenkins中devops-test工程

增加构建后操作  Send build artifacts over SSH

![image-20230915130629641](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915130629641.png)

Exec command中填写

```
devops_start.sh 192.168.117.141:7788 devops  ${JOB_NAME} $tag
```

![image-20230915131945359](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915131945359.png)

点击保存

构建工程 build v1.0.1

 查看输出日志

![image-20230915132839541](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915132839541.png)

访问IP：9901

![image-20230915132856083](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915132856083.png)

Sonarqube 分析中也分析了一次

![image-20230915132927111](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915132927111.png)

harbor中也重新推送了新的镜像

![image-20230915133033501](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915133033501.png)

再次构建v1.0.2 

访问IP：9901 

![image-20230915133124674](assets/%E3%80%90%E7%99%BE%E7%BB%83%E6%88%90%E9%AD%94%E3%80%91Jenkins%20CICD/image-20230915133124674.png)

清理名字标签为<none>的镜像

```
docker image prune
```

完整的Jenkins 持续集成及持续发布就完成了