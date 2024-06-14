# K8S集群证书过期解决



 \#检测证书是否过期



```shell

$ openssl x509 -in /etc/kubernetes/pki/apiserver.crt -noout -text |grep Not

$ kubeadm certs check-expiration 

```







\#先备份/etc/kubernetes下的所有文件



```shell

$ cp -r /etc/kubernetes /etc/kubernetes.bak

```





\#自动更新所有证书并检查证书状态(这个命令适用1.2*版本)



```shell

$ kubeadm certs renew all 

$ kubeadm certs check-expiration （证书的有效期都是invaild）

```







\#如果使用kubectl报错：【You must be logged in to the server (Unauthorized)】

\#进行集群授权（对于root用户的操作）



```shell

$ echo “export KUBECONFIG=/etc/kubernetes/admin.conf” >> ~/.bash_profile

$ source ~/.bash_profile

```







\#对于非root用户



```shell

mkdir -p $HOME/.kube       #输入y确认

sudo cp -i /etc/kubernetes/admin.conf   $HOME/.kube/config       #输入y，强制覆盖

sudo chown  (id−u):(id -g)  $HOME/.kube/config

```





\#上述问题解决后，执行kubectl apply、kubectl create命令可以正常执行，但无法实际操作资源

换句话说：就是执行了，但没生效

\#重启kubelet



```shell

$ systemctl restart kubelet

```





\#重启kube-apiserver、kube-controller-manage、kube-scheduler



```shell

$ docker ps |grep kube-apiserver|grep -v pause|awk '{print $1}'|xargs -i docker restart {}

$ docker ps |grep kube-controller-manage|grep -v pause|awk '{print $1}'|xargs -i docker restart {}

$ docker ps |grep kube-scheduler|grep -v pause|awk '{print $1}'|xargs -i docker restart {}

```







\#或者docker ps找出一下三列的容器ID ,再restart  