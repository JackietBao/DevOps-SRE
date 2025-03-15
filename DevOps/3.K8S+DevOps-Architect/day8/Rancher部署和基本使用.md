# Rancher部署和基本使用

#### 介绍

一个开源的企业级容器管理平台。通过Rancher，企业再也不必自己使用一系列的开源软件去从头搭建容器服务平台。Rancher提供了在生产环境中使用的管理Docker和Kubernetes的全栈化容器部署与管理平台。

帮助用户不需要深入了解kubernetes概念就可以使用*rancher* 

起步于美国硅谷，近年逐步发力中国市场。现已被suse收购

#### 安装

版本选型：

https://rancher.com/support-maintenance-terms/all-supported-versions/rancher-v2.5.2/

本文档使用最新的v2.5.2版本，v2.x的版本，安装配置过程都是相似的。

直接使用rancher官方镜像启动：

```powershell
$ docker run -d --privileged --name rancher --restart=unless-stopped -p 8080:80 -p 8443:443 -v /opt/rancher/:/var/lib/rancher/ rancher/rancher:v2.5.2
```

等待服务启动后，提供主机的`https://<host-ip>:8443`即可访问rancher管理界面，第一次访问需要重装管理员密码。



内部使用自家的k3s启动了内部集群，容器层面直接使用的containerd来管理镜像及容器。

```powershell
$ docker exec -ti rancher bash
# kubectl get no
# kubectl get po -A  
```



如果出现镜像无法正常拉取的情况，可以考虑使用离线的方式将所需的镜像导入到rancher容器内部，针对rancher 2.5.2版本，所需的镜像列表为：

```powershell
docker.io/rancher/pause:3.1
docker.io/rancher/coredns-coredns:1.6.9
docker.io/rancher/gitjob:v0.1.8
docker.io/rancher/rancher-operator:v0.1.1
docker.io/rancher/shell:v0.1.5
docker.io/rancher/rancher-webhook:v0.1.0-beta7
docker.io/rancher/fleet:v0.3.1
```

如果无法正常拉取镜像，可以采用下述方式：

```powershell
# 宿主机中拉取镜像
$ docker pull docker.io/rancher/pause:3.1
$ docker save -o rancher-pause.tar docker.io/rancher/pause:3.1
$ docker cp rancher-pause.tar rancher:/tmp/
$ docker exec -ti rancher bash
#/var/lib/rancher/k3s/data/19bb6a9b46ad0013de084cf1e0feb7927ff9e4e06624685ff87f003c208fded1/bin/ctr image import /tmp/rancher-pause.tar
```





#### 基本使用

##### 几个概念

- 集群

  rancher可以管理多个k8s集群，集群可以通过新建以及导入的方式纳入rancher的管控

  初始化会将内置k3s部署的集群接入，名为`local`

- 项目

  集群下的逻辑概念，一个集群可以包含多个项目，一个项目下可以包含多个命名空间。

  初始化会为接入的每个集群创建两个项目：

  - Default：对应集群的default命名空间
  - System：对应系统级别的命名空间，包含`kube-system`、`kube-public`、`cattle-system`、`ingress-nginx`等

- 命名空间

  对应k8s的命名空间概念，可以直接新建或者将命名空间移动到已有的项目中



##### 权限管理

rancher支持本地用户以及与LDAP账户对接，用户的权限是基于项目赋予的。

