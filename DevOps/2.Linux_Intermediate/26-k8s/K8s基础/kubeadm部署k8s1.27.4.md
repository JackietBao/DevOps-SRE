# **kubeadm部署k8s1.27.4**

## 1.编辑本地解析且修改主机名

![img](assets/kubeadm部署k8s1.27.4/1688450403484-e4d34a61-c2ae-4a6b-a99a-bd645c0d5c72.png)

master节点：

```plain
hostnamectl set-hostname k8s-master
```

node节点：

```plain
hostnamectl set-hostname k8s-node1
```

master节点产生ssh密钥拷贝给node节点实现免密登录

```plain
ssh-keygen
ssh-copy-id k8s-node1
```

## 2.开启路由转发功能

```plain
##添加配置文件
[root@k8s-master ~]# cat > /etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness = 0
EOF
##加载br_netfilter模块 && 查看是否加载
[root@k8s-master ~]# modprobe br_netfilter && lsmod | grep br_netfilter
br_netfilter           22256  0 
bridge                151336  1 br_netfilter
##加载⽹桥过滤及内核转发配置⽂件
[root@k8s-master ~]# sysctl -p /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness = 0
```

## 3.配置阿里云yum源

```plain
[root@k8s-master ~] cd /etc/yum.repos.d/
[root@k8s-node yum.repos.d]# mv CentOS-Base.repo CentOS-Base.repo.bak
[root@k8s-node yum.repos.d]# curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
[root@k8s-node yum.repos.d]# yum clean all && yum makecache
```

## 4.配置ipvs功能

```plain
[root@k8s-master ~]# yum -y install ipset ipvsadm
##编辑需要添加的模块儿
[root@k8s-master ~]# cat > /etc/sysconfig/modules/ipvs.modules <<EOF
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack
EOF
##给予执行权限
[root@k8s-master ~]# chmod 755 /etc/sysconfig/modules/ipvs.modules
##运行模块儿脚本
[root@k8s-master ~]# bash /etc/sysconfig/modules/ipvs.modules
##检验是否成功
[root@k8s-master ~]# lsmod | grep -e ip_vs -e nf_conntrack
ip_vs_sh               12688  0 
ip_vs_wrr              12697  0 
ip_vs_rr               12600  0 
ip_vs                 145458  6 ip_vs_rr,ip_vs_sh,ip_vs_wrr
nf_conntrack_netlink    36396  0 
nfnetlink              14519  2 nf_conntrack_netlink
nf_conntrack_ipv4      15053  2 
nf_defrag_ipv4         12729  1 nf_conntrack_ipv4
nf_conntrack          139264  7 ip_vs,nf_nat,nf_nat_ipv4,xt_conntrack,nf_nat_masquerade_ipv4,nf_conntrack_netlink,nf_conntrack_ipv4
libcrc32c              12644  4 xfs,ip_vs,nf_nat,nf_conntrack
```

## 5.关闭防火墙、内核以及swap分区

```plain
[root@k8s-master ~]# systemctl stop firewalld
[root@k8s-master ~]# setenforce 0
setenforce: SELinux is disabled
[root@k8s-master ~]# swapoff -a（临时关闭）
##永久关闭需要在/etc/fstab文件中注释掉
#/dev/mapper/centos-swap swap                    swap    defaults        0 0
```

## 6.安装Docker

```plain
# curl https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo -o /etc/yum.repos.d/docker.repo
# yum -y install docker-ce
```

配置docker加速器

```plain
[root@k8s-master ~]# mkdir /etc/docker
[root@k8s-master ~]# cat /etc/docker/daemon.json 
{
        "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn",
                        "https://docker.m.daocloud.io",
                        "http://hub-mirrors.c.163.com"],
        "max-concurrent-downloads": 10,
        "log-driver": "json-file",
        "log-level": "warn",
        "data-root": "/var/lib/docker"

}
[root@k8s-master ~]# systemctl restart docker
```

配置Cgroup驱动程序

（在 Linux 上，控制组（CGroup）⽤于限制分配给进程的资源，官⽅建议配置容器运⾏时和 kubelet 使⽤ systemd（systemd是Linux系统第⼀个初始进程）作为容器的控制组（CGroup）， 以此使系统更为稳定 。）

```plain
##在/etc/docker/daemon.json增加一条内容
[root@k8s-master ~]# cat /etc/docker/daemon.json
{
        "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn",
                        "https://docker.m.daocloud.io",
                        "http://hub-mirrors.c.163.com"],
        "max-concurrent-downloads": 10,
        "log-driver": "json-file",
        "log-level": "warn",
        "data-root": "/var/lib/docker",
				"exec-opts": ["native.cgroupdriver=systemd"]

}
```

启动docker并设置开机自启

## 7.安装cri-docker

作为接替 Docker 运行时的 Containerd 在早在 Kubernetes1.7 时就能直接与 Kubelet 集成使用，只是大部分时候我们因熟悉 Docker，在部署集群时采用了默认的 dockershim。在V1.24起的版本的 kubelet 就彻底移除了dockershim，改为默认使用Containerd了，当然也使用 cri-dockerd适配器来将Docker Engine与 Kubernetes 集成。可以参考官方文档：

 https://kubernetes.io/zh-cn/docs/setup/production-environment/container-runtimes/#docker

CRI-Dockerd cri-dockerd是什么

https://blog.51cto.com/u_16099281/6469310

```plain
##下载压缩包
[root@k8s-master ~]# wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.4/cri-dockerd-0.3.4.amd64.tgz
##解压
[root@k8s-master ~]# tar zxvf cri-dockerd-0.3.4.amd64.tgz
##拷贝二进制命令文件
[root@k8s-master ~]# cp cri-dockerd/* /usr/bin/
##配置systemctl管理
[root@k8s-master ~]# vim /usr/lib/systemd/system/cri-docker.service
[Unit]
Description=CRI Interface for Docker Application Container Engine
Documentation=https://docs.mirantis.com
After=network-online.target firewalld.service docker.service
Wants=network-online.target
Requires=cri-docker.socket

[Service]
Type=notify
ExecStart=/usr/bin/cri-dockerd --network-plugin=cni --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.7
ExecReload=/bin/kill -s HUP $MAINPID
TimeoutSec=0
RestartSec=2
Restart=always
StartLimitBurst=3
StartLimitInterval=60s
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
TasksMax=infinity
Delegate=yes
KillMode=process

[Install]
WantedBy=multi-user.target
[root@k8s-master ~]# vim /usr/lib/systemd/system/cri-docker.socket
[Unit]
Description=CRI Docker Socket for the API
PartOf=cri-docker.service

[Socket]
ListenStream=%t/cri-dockerd.sock
SocketMode=0660
SocketUser=root
SocketGroup=docker

[Install]
WantedBy=sockets.target

[root@k8s-master ~]# systemctl daemon-reload
[root@k8s-master ~]# systemctl enable --now cri-docker
Created symlink from /etc/systemd/system/multi-user.target.wants/cri-docker.service to /usr/lib/systemd/system/cri-docker.service.
[root@k8s-master ~]# systemctl status cri-docker
```

安装kubelet、kubeadm、kubectl

（有个小坑，之前更换阿里源里面没有这些包，需要再添加一个）

```plain
##添加源
[root@k8s-master ~]# cd /etc/yum.repos.d/
[root@k8s-master yum.repos.d]# vim kubernetes.repo 
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
##下载
[root@k8s-master ~]# yum -y install kubelet-1.27.4-0 kubeadm-1.27.4-0 kubectl-1.27.4-0 --disableexcludes=kubernetes
##设置kubelet开机自启
[root@k8s-master ~]# systemctl enable --now kubelet
Created symlink from /etc/systemd/system/multi-user.target.wants/kubelet.service to /usr/lib/systemd/system/kubelet.service.
##初始化之前kubelet无法启动，可以查看它的状态，下面这种情况代表正在等待指令
[root@k8s-master ~]# systemctl is-active kubelet
activating
```

## 8.kubeadm初始化

--image-repository registry.aliyuncs.com/google_containers：使用阿里云镜像仓库
--kubernetes-version=v1.27.3：指定k8s的版本
--pod-network-cidr=10.10.20.0/24：指定pod的网段

--cri-socket unix:///var/run/cri-dockerd.sock：指定容器运行时的Socket文件路径，原本默认是dockershim.sock，但现在改成cri-docker.sock

```plain
[root@k8s-master ~]# kubeadm init \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.27.4 \
--pod-network-cidr=10.10.20.0/24 \
--cri-socket unix:///var/run/cri-dockerd.sock
```

初始化成功后根据回显创建文件，并记录下节点加入集群的命令

![img](assets/kubeadm部署k8s1.27.4/1688459002399-9cf2fd77-30ff-4951-ada3-ece57b75f7ba.png)

master创建文件

```plain
[root@k8s-master ~]# mkdir -p $HOME/.kube
[root@k8s-master ~]# sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
[root@k8s-master ~]# sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

若是出了问题可以使用  kubeadm reset 清除所有配置和数据重新初始化

## 9.加入集群

node节点执行

```plain
kubeadm join 10.8.166.240:6443 --token uq6h3m.jgrn7qffjk4e7ihi --discovery-token-ca-cert-hash sha256:ce882c56d3f55023c9859e5cdf80cab130d6613d91ab5e395d8d85506c198ccb --cri-socket unix:///var/run/cri-dockerd.sock
```

![img](assets/kubeadm部署k8s1.27.4/1690365365136-dd877426-f012-4016-bad4-a56169b84e95.png)

master节点查看

```plain
[root@k8s-master ~]# kubectl get node
NAME         STATUS     ROLES           AGE    VERSION
k8s-master   NotReady   control-plane   144m   v1.27.4
k8s-node1    NotReady   <none>          50s    v1.27.4
k8s-node2    NotReady   <none>          4s     v1.27.4
```

## 10.配置网络插件

我这边用的是calico网络插件

## 1.master下载kube-flannel.yml文件

```plain
curl -O https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

![img](assets/kubeadm部署k8s1.27.4/1688462029701-503c4720-d90f-4326-8ea8-6a16ff7586d3.png)

改为初始化时指定的网段

![img](assets/kubeadm部署k8s1.27.4/1688462130578-9568fde3-d0fc-40cb-b751-46835a9f2a98.png)

可以提前拉取这些需要的镜像（每个节点都拉取！！！）

![img](assets/kubeadm部署k8s1.27.4/1688462192358-256374ba-643c-4a81-be5b-b48ab54fa611.png)

这里添加两行内容：

![img](assets/kubeadm部署k8s1.27.4/1688462454427-1a41a897-e6c6-4d19-af38-1c0253eeb524.png)

```plain
[root@k8s-master ~]# kubectl apply -f kube-flannel.yml
```

## 2.配置网络组件calico

calico官网：https://docs.tigera.io/calico/latest/getting-started/kubernetes/quickstart#install-calico

注释：一定要学会看官方文档

![img](assets/kubeadm部署k8s1.27.4/1690367127188-6577f2ca-328a-4fb5-a211-58feee5ddba1.png)

```plain
[root@k8s-master ~]# kubectl apply -f tigera-operator.yaml
```

![img](assets/kubeadm部署k8s1.27.4/1690422287912-d5cc2e6a-30ba-4552-a806-03a65efc756f.png)

![img](assets/kubeadm部署k8s1.27.4/1690422339310-5d2da378-e85e-467c-981f-3fb656d5c528.png)

那这里，自己要查一下kubectl apply 和 kubectl create的区别了

```plain
[root@k8s-master ~]# kubectl delete -f tigera-operator.yaml
[root@k8s-master ~]# kubectl create -f tigera-operator.yaml
```

![img](assets/kubeadm部署k8s1.27.4/1690422469737-855f57bc-a069-46cf-b006-b167f612ea59.png)

```plain
[root@k8s-master ~]# vim custom-resources.yaml
```

![img](assets/kubeadm部署k8s1.27.4/1690427288163-beffbfaf-aea0-4f10-bdb1-9b6ac4903333.png)

```plain
[root@k8s-master ~]# kubectl apply -f custom-resources.yaml
```

![img](assets/kubeadm部署k8s1.27.4/1690422500689-06e8d0b9-3056-4bb1-ab8e-e80eade534e6.png)

![img](assets/kubeadm部署k8s1.27.4/1690422918940-7b4cfdaf-8b60-41b7-8286-fa1dcd227ee1.png)

```plain
[root@k8s-master ~]# kubectl get pod -n calico-system
```

![img](assets/kubeadm部署k8s1.27.4/1690423666897-3b8c0181-dd12-4d48-b513-874564fdffc2.png)

如果上面pod出现异常，请同步ntp时间，再次查看，时间不同步，pod会初始化失败，亲身经历

参考错误链接：https://blog.51cto.com/u_15567202/6149595

![img](assets/kubeadm部署k8s1.27.4/1690423765693-6a74a9c1-d225-43cc-a1dd-a686e6c58a1a.png)

## 11.验证集群可用性

```plain
[root@k8s-master ~]# kubectl get node
NAME         STATUS   ROLES           AGE   VERSION
k8s-master   Ready    control-plane   18h   v1.27.4
k8s-node1    Ready    <none>          16h   v1.27.4
k8s-node2    Ready    <none>          16h   v1.27.4
```

## 12.查看集群健康情况,理想状态

```plain
[root@k8s-master ~]# kubectl get cs
Warning: v1 ComponentStatus is deprecated in v1.19+
NAME                 STATUS    MESSAGE   ERROR
controller-manager   Healthy   ok        
scheduler            Healthy   ok        
etcd-0               Healthy 
```

## 13.查看kubernetes集群pod运⾏情况

```plain
[root@k8s-master ~]# kubectl get pod -n kube-system
NAME                                 READY   STATUS    RESTARTS        AGE
coredns-7bdc4cb885-s67tc             1/1     Running   0               18h
coredns-7bdc4cb885-x9q9z             1/1     Running   0               18h
etcd-k8s-master                      1/1     Running   0               18h
kube-apiserver-k8s-master            1/1     Running   2 (4m30s ago)   18h
kube-controller-manager-k8s-master   1/1     Running   7 (3m31s ago)   18h
kube-proxy-nnjxs                     1/1     Running   0               16h
kube-proxy-smnlb                     1/1     Running   0               18h
kube-proxy-vgjx7                     1/1     Running   0               16h
kube-scheduler-k8s-master            1/1     Running   5 (3m30s ago)   18h
```

## 14.所用安装包百度网盘链接

```plain
链接：https://pan.baidu.com/s/1N3Gn-rpmaxhPuDp8dMsDHQ?pwd=s3jl
提取码：s3jl
```

## 15.部署nginx页面测试

```plain
[root@k8s-master ~]# kubectl create deployment nginx --image=nginx
[root@k8s-master ~]# kubectl get pod
NAME                     READY   STATUS    RESTARTS   AGE
nginx-77b4fdf86c-rfh4w   1/1     Running   0          3s

# 开放80端口
[root@k8s-master ~]# kubectl expose deployment nginx --port=80 --type=NodePort

# 下载choujiang项目
[root@k8s-master ~]# wget https://files.cnblogs.com/files/simuhunluo/choujiang.zip

# 解压拷贝到Nginx站点目录
[root@k8s-master ~]# yum -y install unzip
[root@k8s-master ~]# mkdir choujiang
[root@k8s-master ~]# unzip choujiang.zip choujiang
[root@k8s-master ~]# kubectl cp choujiang nginx-77b4fdf86c-rfh4w:/usr/share/nginx/html

# 进入容器执行操作
[root@k8s-master ~]# kubectl exec -it nginx-77b4fdf86c-rfh4w /bin/bash
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
root@nginx-77b4fdf86c-rfh4w:/# cd /usr/share/nginx/html/
root@nginx-77b4fdf86c-rfh4w:/usr/share/nginx/html# ls
choujiang
root@nginx-77b4fdf86c-rfh4w:/usr/share/nginx/html# mv choujiang/* .

#查看service状态
[root@k8s-master ~]# kubectl get service
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP        19h
nginx        NodePort    10.98.227.163   <none>        80:30254/TCP   24m
```

\# 访问网页：http://nodeIP:30254

![img](assets/kubeadm部署k8s1.27.4/1690426936796-f9d97403-0db9-4da2-ad46-8f517f45bdfd.png)

## 16.Dashboard界面

Dashboard 是基于网页的 Kubernetes 用户界面。 你可以使用 Dashboard 将容器应用部署到 Kubernetes 集群中，也可以对容器应用排错，还能管理集群资源。 你可以使用 Dashboard 获取运行在集群中的应用的概览信息，也可以创建或者修改 Kubernetes 资源 （如 Deployment，Job，DaemonSet 等等）。 例如，你可以对 Deployment 实现弹性伸缩、发起滚动升级、重启 Pod 或者使用向导创建新的应用。Dashboard 同时展示了 Kubernetes 集群中的资源状态信息和所有报错信息



![img](assets/kubeadm部署k8s1.27.4/1690427418494-f8e6d948-0aab-4dde-9bd0-6cb69b403a59.png)

```plain
[root@k8s-master ~]# kubectl apply -f recommended.yaml
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created
[root@k8s-master ~]# kubectl get ns
NAME                   STATUS   AGE
calico-apiserver       Active   60m
calico-system          Active   82m
default                Active   19h
kube-node-lease        Active   19h
kube-public            Active   19h
kube-system            Active   19h
kubernetes-dashboard   Active   55s   #这里
tigera-operator        Active   86m
[root@k8s-master ~]# kubectl get pod -n kubernetes-dashboard
NAME                                         READY   STATUS              RESTARTS   AGE
dashboard-metrics-scraper-5cb4f4bb9c-2ghd7   0/1     ContainerCreating   0          2m3s
kubernetes-dashboard-6967859bff-xtxdk        0/1     ContainerCreating   0          2m3s

#一个是dashboard页面应用，一个是dashboard管理插件。两个都需要拉取镜像，需要等待，。，。，
等待5分钟左右
[root@k8s-master ~]# kubectl get pod -n kubernetes-dashboard
NAME                                         READY   STATUS    RESTARTS   AGE
dashboard-metrics-scraper-5cb4f4bb9c-2ghd7   1/1     Running   0          4m50s
kubernetes-dashboard-6967859bff-xtxdk        1/1     Running   0          4m50s
[root@k8s-master ~]# kubectl get svc -n kubernetes-dashboard
NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
dashboard-metrics-scraper   ClusterIP   10.107.118.24   <none>        8000/TCP        5m28s
kubernetes-dashboard        NodePort    10.99.76.240    <none>        443:31879/TCP   5m28s

#dashboard暴露了一个端口号30764，使用https://nodeIP:31879访问
```

![img](assets/kubeadm部署k8s1.27.4/1690427819841-57127631-0dbd-436d-9181-6c1a6e23c6c0.png)

## 17.创建服务账户和ClusterRoleBinding

```plain
[root@k8s-master dashboard]# cat dashboard-adminuser.yaml 
#写入如下,创建了一个admin-user的服务账户并与集群绑定
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
[root@k8s-master dashboard]# kubectl apply -f dashboard-adminuser.yaml
serviceaccount/admin-user created
clusterrolebinding.rbac.authorization.k8s.io/admin-user created
[root@k8s-master dashboard]# kubectl -n kubernetes-dashboard create token admin-user
eyJhbGciOiJSUzI1NiIsImtpZCI6IlVjbkFuVzZTaWNIeXg4SWphRk8teXF3Nk5vZzV2QlllenN0VlRhQ0lPaUUifQ.eyJhdWQiOlsiaHR0cHM6Ly9rdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNsdXN0ZXIubG9jYWwiXSwiZXhwIjoxNjkwNDMxNjcxLCJpYXQiOjE2OTA0MjgwNzEsImlzcyI6Imh0dHBzOi8va3ViZXJuZXRlcy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsIiwia3ViZXJuZXRlcy5pbyI6eyJuYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsInNlcnZpY2VhY2NvdW50Ijp7Im5hbWUiOiJhZG1pbi11c2VyIiwidWlkIjoiMWFhODg5ZGYtZGI0Mi00YTBhLTkwYjMtZTcxODZmOGRjMjI4In19LCJuYmYiOjE2OTA0MjgwNzEsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlcm5ldGVzLWRhc2hib2FyZDphZG1pbi11c2VyIn0.A_XQqYUIWWrR10nx8jt7z79-14QLhx7BT7tz9Uj12z3gkmCDYgFbvflZlVve-Ek6vLVBZk2lsPcIpPknl0_Siv6XhxcYjS5CVf27Ht-E-c00tG84wFkRAxl3yjtVglkEvhaf2zJEka-IYL3XUgpcx65ESlrBvLmkWX0J0Vd0OUahmy_xSRQrWFS3YvzD_dI8qBFJZYe2kYLIK7INkBbWKkx5-Ap5vuSBc9Sse7D_HE_5fX-wCMGhOYFfxXS8VasWP2wN_0hac-IhrHg14jHUR1ctU9xYueAri_3Jn3HJ9VYXVHKWBntR_8AIbGjJwkmeQgRtnOzViLtcrJi83Z8HKA
```

![img](assets/kubeadm部署k8s1.27.4/1690428097521-37c27640-f934-4dbd-b553-8cf46112ebae.png)

## 18.清理、删除管理员ServiceAccount和ClusterRoleBinding

```plain
kubectl -n kubernetes-dashboard delete serviceaccount admin-user
kubectl -n kubernetes-dashboard delete clusterrolebinding admin-user
```

说明：在以前的版本中，可以通过类似于**kubectl get secret -n kubernetes-dashboard | grep token | grep admin**的命令先查询secret，再使用**kubectl describe secret <上条命令查询的结果> -n kubernetes-dashboard**即可显示token信息，但现如今版本（我也不知道从哪个版本开始的）这种方式查看无效，并且**kubectl get secret -n kubernetes-dashboard**也不会显示，所以这里建议生成了token就创建一个文件复制到里面保存下来，如：

```plain
[root@k8s-master ~]# cd dashboard/
[root@k8s-master dashboard]# vi admin-user.token 
#将token复制到admin-user.token文件里保存
```