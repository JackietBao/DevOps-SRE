环境准备

```
客户端：192.168.77.129

服务端：192.168.77.133
```

关闭防火墙

```
[root@192 ~]# systemctl stop firewalld
[root@192 ~]# vi /etc/sysconfig/selinux 
[root@192 ~]# setenforce 0
```

### 服务端

安装nfs-server服务

```
[root@192 ~]# yum -y install nfs-utils
```

配置文件

```
[root@192 ~]# vim /etc/exports
/data 192.168.77.0/24(rw,sync,all_squash)
```

nfs服务器建立用于nfs文件共享目录，并设置权限

```
[root@192 ~]# mkdir /data
[root@192 ~]# chown -R nfsnobody.nfsnobody /data
```

开机自启

```
[root@192 ~]# systemctl enable rpcbind nfs-server
```

启动服务

```
[root@192 ~]# systemctl restart rpcbind nfs-server
```

### 客户端

安装客户端工具

```
[root@192 ~]# yum -y install nfs-utils
```

使用showmount -e查看远程服务器rpc提供的可挂载nfs信息

```
[root@192 ~]# showmount -e 192.168.77.133
Export list for 192.168.77.133:
```

挂载

```
[root@192 ~]# mkdir /nfsdir
[root@192 ~]# mount -t nfs 192.168.77.133:/data /nfsdir
```

查看挂载信息

```
[root@192 ~]# df -Th
Filesystem              Type      Size  Used Avail Use% Mounted on
devtmpfs                devtmpfs  898M     0  898M   0% /dev
tmpfs                   tmpfs     910M     0  910M   0% /dev/shm
tmpfs                   tmpfs     910M  9.6M  901M   2% /run
tmpfs                   tmpfs     910M     0  910M   0% /sys/fs/cgroup
/dev/mapper/centos-root xfs        17G  1.8G   16G  11% /
/dev/sda1               xfs      1014M  183M  832M  19% /boot
tmpfs                   tmpfs     182M     0  182M   0% /run/user/0
192.168.77.133:/data    nfs4       17G  1.3G   16G   8% /nfsdir
```

验证

```
#客户端往nfs写入
[root@192 ~]# echo "hello jackietbao" >> /nfsdir/test.txt

#服务端查看
[root@192 ~]# ls /data/
test.txt
```

nfs文件共享服务一直有效，需要写入到fstab文件当中

```
[root@192 ~]# vim /etc/fstab

#
# /etc/fstab
# Created by anaconda on Sun Mar 10 15:59:49 2024
#
# Accessible filesystems, by reference, are maintained under '/dev/disk'
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
#
/dev/mapper/centos-root /                       xfs     defaults        0 0
UUID=ea771271-8b04-40f4-a37b-e573298005a0 /boot                   xfs     defaults        0 0
/dev/mapper/centos-swap swap                    swap    defaults        0 0
192.168.77.133:/data /nfsdir nfs defaults 0 0
```

卸载nfs共享存储

```
unmount /nfsdir
```

通过mount -o指定挂载参数禁止使用suid，exec增加安全性能

```
mount -t nfs -o nosuid,noexec,nodev 192.168.77.133:/data /nfsdir
```

性能考虑

```
mount -t nfs -o noatime,nodiratime 192.168.77.133:/data /nfsdir
```