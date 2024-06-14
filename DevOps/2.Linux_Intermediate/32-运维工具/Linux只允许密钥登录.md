# Linux只允许密钥登录

# Linux服务器远程连接只允许密钥文件



1.点击Xshell菜单栏的工具，选择新建用户密钥生成向导，进行密钥对生成操作。



![img](assets/Linux只允许密钥登录/1683016186099-df2db822-4827-4609-bba5-8ed0aad412a3.png)



![img](assets/Linux只允许密钥登录/1683016186214-6071bc26-d320-4e57-bd37-c571dc2e6c90.png)



![img](assets/Linux只允许密钥登录/1683016186266-f4b8d79a-88cd-4613-b185-c0ca46e97db9.png)



![img](assets/Linux只允许密钥登录/1683016186313-95f0669e-de85-4886-b3d1-745f097ec4e8.png)



![img](assets/Linux只允许密钥登录/1683016186388-7977abaf-f93e-4071-817f-f832efc40359.png)



![img](assets/Linux只允许密钥登录/1683016187255-22e19ef0-aa4c-4210-9769-53a5ff43b28f.png)



![img](assets/Linux只允许密钥登录/1683016187396-e24b6007-9d85-4655-b819-baedc676c57b.png)



![img](assets/Linux只允许密钥登录/1683016187617-1fcb3f39-c186-4cff-8386-b903c78cc2e8.png)



2.这个时候，你已经有了一个密钥，需要开始设定服务器的配置，启用密钥认证登录，同时为了系统安全着想，关闭密码认证的方式！



```shell
# vim /etc/ssh/sshd_config
```



修改下面几处：



```shell
PubkeyAuthentication  yes  #启用PublicKey认证。
AuthorizedKeysFile       .ssh/authorized_keys  #PublicKey文件路径。
PasswordAuthentication  no  #不适用密码认证登录。
```



![img](assets/Linux只允许密钥登录/1683016187736-f6dfa0d3-6ca8-4650-894f-14cb0bac66f1.png)



```shell
[root@192 .ssh]# vim /root/.ssh/authorized_keys
将公钥内容复制到此文本中
```



![img](assets/Linux只允许密钥登录/1683016187666-40818dc5-94ff-4b40-9b6c-9a66f95f82b0.png)



3.接着，修改该文件的权限，否则可能会遇到问题！



```shell
[root@192 .ssh]# chmod 600 authorized_keys
```



4.重启ssh服务



```shell
[root@192 .ssh]# systemctl restart sshd
```



5.至此，登录测试吧！你会发现输入完用户，密码一栏是灰色的！

![img](assets/Linux只允许密钥登录/1683016188388-498ce309-49fa-4615-90c3-c2c1c0da205a.png)



登录成功，配置完成！



禁止root用户远程登录

```shell
# vim /etc/ssh/sshd_config
```

![img](assets/Linux只允许密钥登录/1664501580073-dcd83cc1-4db5-4de8-97c4-2e279955a0e7.png)

重启sshd服务

# Centos7下Rinetd安装与应用



两台虚拟机



```shell
192.168.62.131

192.168.62.231
放火墙均处于开启状态，selinux可以关闭
```



### 在192.168.62.131操作：



1. rpm安装



```shell
# wget  http://li.nux.ro/download/nux/misc/el7/x86_64//rinetd-0.62-9.el7.nux.x86_64.rpm
# yum -y install rinetd-0.62-9.el7.nux.x86_64.rpm
```



2.使用配置文件的用法

rpm安装的配置文件默认路径是/etc/rinetd.conf

```shell
# vim /etc/rinetd.conf
0.0.0.0     10022       192.168.62.231      1022
0.0.0.0     20022       192.168.62.232      1022
```



3.启动rinetd

```shell
# systemctl start rinetd
```



4.开放端口10022

```shell
# firewall-cmd --zone=public --add-port=10022/tcp --permanent
```



### 在192.168.62.231操作：



1.配置防火墙规则



指定源地址开放端口



```shell
# firewall-cmd --permanent --add-rich-rule="rule family="ipv4" source address="192.168.62.131" port protocol="tcp" port="1022" accept"
```



在Xshelll上测试



![img](assets/Linux只允许密钥登录/1683016188687-a06b709f-2fc9-4411-a7bf-db368bbe5e29.png)



![img](assets/Linux只允许密钥登录/1683016188717-8034beb4-4ec3-43e3-8dab-f96e616be780.png)



登录成功！！