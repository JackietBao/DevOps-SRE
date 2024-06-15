# gitlab搭建流程

## 一、关闭防火墙与内核

```
systemctl stop firewalld
setenforce 0
```



## 二、配置yum源并下载依赖

```
[root@gitlab]# cd /etc/yum.repos.d/
[root@gitlab]# vim gitlab-ce.repo
[gitlab-ce]
name=Gitlab CE Repository
baseurl=https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/yum/el$releasever
gpgcheck=0
enabled=1
####下载依赖
[root@gitlab]# yum install -y postfix curl policycoreutils-python openssh-server
[root@gitlab]# systemctl enable sshd
[root@gitlab]# systemctl start sshd
#####修改postfix的配置文件
[root@gitlab]# vim  /etc/postfix/main.cf
inet_interfaces = all
#inet_interfaces = $myhostname
#inet_interfaces = $myhostname, localhost
#inet_interfaces = localhost
[root@gitlab]# systemctl enable postfix
[root@gitlab]# systemctl start postfix
####安装最新版本gitlab
[root@gitlab]# yum install -y gitlab-ce
```

## 三、修改gitlab配置文件

```
[root@gitlab]# vim /etc/gitlab/gitlab.rb
##设置访问地址与时区
external_url 'http://ip+端口（或者域名）'

gitlab_rails['time_zone'] = 'Asia/Shanghai'
###打开数据路径（去掉注释，可更改路径）
 git_data_dirs({
   "default" => {
     "path" => "/mnt/nfs-01/git-data"
    }
 })
###开启ssh服务
 gitlab_rails['gitlab_shell_ssh_port'] = 22
```

## 四、初始化gitlab

```
[root@gitlab]# gitlab-ctl reconfigure
```

## 五、启动gitlab

```
[root@gitlab]# gitlab-ctl start
```

