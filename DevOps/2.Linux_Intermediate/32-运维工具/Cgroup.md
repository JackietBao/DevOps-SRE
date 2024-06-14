# Cgroup

# 使用 cgroups 进行资源控制



## 什么是Cgroup



自 Linux 2.6.24 以来已被合并以来，用来限制、控制与分离一个进程组的资源（如CPU、内存、磁盘、输入输出等）的内核功能。



主要子系统有



-  Cpu：CPU使用 
-  Cpuacct：自动生成报告统计控制组所使用的CPU资源 
-  Cpuset：通过调度程序来管理cgroup中任务对CPU访问 
-  Memory：用于任务内存使用限制，并自动生成控制组内内存资源使用报告 
-  Blkio：块输入输出子系统 
-  Devices：允许或者拒绝cgroup的任务对指定设备的访问 
-  Net_cls：对网络包进行标签化 
-  Net_prio：动态设置应用程序网络流量优先级 
-  Freezer：用于挂起和恢复cgroup中的进程 



## 安装Cgroup



### 安装Cgroup



```shell
yum install libcgroup libcgroup-tools
```



### Cgroup启动设定



```shell
systemctl enable cgconfig
systemctl enable cgred
```



## 通过命令设置Cgroup



### 首先，创建一个组



创建一个名为 [CPU75_Mem_Group] 的控制组，目标为 CPU 和内存。



```shell
cgcreate -g cpu,memory:/CPU75_Mem_Group
```



### 设置限制值



为创建的控制组设置资源限制。



限制的使用率 75% x 核心数（在 4 核心的情况下）



当你计算[cpu.cfs_quota_us]



![img](https://g.yuque.com/gr/latex?%5Bcpu.cfs_quota_us%5D%3D%200.75*4core*100000(%E5%B0%86%20s%20%E8%BD%AC%E6%8D%A2%E4%B8%BA%20%C2%B5s)%20%3D%20300000%0A)



以下为设定方式



```shell
cgset -r cpu.cfs_quota_us=300000 CPU75_Mem_Group
cgset -r cpu.cfs_period_us=1000000 CPU75_Mem_Group
```



如果要限制内存分配



需要设置memory.limit_in_bytes（用户内存上限）和memory.memsw.limit_in_bytes（用户内存上限+swap）的值。



例如，如果将用户内存（物理内存）限制为 100MB，swap为 500MB



```shell
cgset -r memory.limit_in_bytes=100M CPU75_Mem_Group
cgset -r memory.memsw.limit_in_bytes=500M CPU75_Mem_Group
```



### 其他 cgroups 相关命令



命令功能



- cgcreate 创建一个控制组
- cgdelete 删除控制组
- cgset子系统的参数设置
- cgget 控制组参数显示
- 在 cgclassify 控制组中移动任务
- 在cgexec 控制组中启动一个进程
- cgsnapshot 从现有子系统生成配置文件
- cgconfigparser 解析配置文件和挂载层次结构
- cgclear 清除整个 cgroup 文件系统
- lscgroup 列出所有cgroups
- lssubsys 子系统挂载点显示



## 通过配置文件设置Cgroup



### 配置文件的配置方法



这是一种编辑和配置cgconfig服务引用的（/etc/cgconfig.conf）的方法。



如果要在与命令相同的条件下进行限制，则需要编辑其文件，如下所示。



/etc/cgconfig.conf



```shell
group CPU75_Mem_Group {
	cpu {
		cpu.cfs_quota_us = 300000;
		cpu.cfs_period_us = 1000000;
	}
	memory {
		memory.limit_in_bytes = 100M
		memory.memsw.limit_in_bytes = 500M
	}
}
```



### 如何仅限制特定用户



在 cgred 服务引用的 （/etc/cgrules.conf） 中进行以下设置。



将 （/etc/cgconfig.conf]）中设置的控制组与特定用户关联。



/etc/cgrules.conf



```shell
hoge_user cpu CPU75_Mem_Group
```



需要重启服务才能使设置生效。



### 如何仅限制特定进程



通过cgroup启动对应software



```shell
cgexec -g cpu:CPU75_Mem_Group <进程启动命令>
```



### 如何仅限制特定用户的特定程序



在/etc/cgrules.conf中进行以下设置并在/etc/cgconfig.conf 中设置
将控制组与特定用户或特定程序相关联。



/etc/cgrules.conf



```shell
hoge_user:<进程启动命令> cpu CPU75_Mem_Group
```



### 如何限制磁盘 I/O 带宽



可以通过更改以下参数来限制磁盘 I/O



#### 以字节/秒为单位指定特定设备的访问速度上限



指定 0 为删除限制



指定方法为"主要和次要设备号  设定值"（例如："8:0 1048576"）



- blkio.throttle.read_bps_device
- blkio.throttle.write_bps_device



#### 以 IO/Sec 为单位指定特定设备的访问速度上限



指定 0 为删除限制。



指定方法为"主要和次要设备号  设定值"（例如："8:0 2048"）



- blkio.throttle.read_iops_device
- blkio.throttle.write_iops_device



#### 指定对/dev/sda设备进行限制



首先，检查目标块设备的主设备号



```shell
[root@localhost ~]# lsblk -p
NAME        MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
/dev/sda      8:0    0   20G  0 disk
├─/dev/sda1   8:1    0  300M  0 part /boot
├─/dev/sda2   8:2    0    2G  0 part [SWAP]
└─/dev/sda3   8:3    0 17.7G  0 part /
/dev/sr0     11:0    1 1024M  0 rom
```



在/etc/cgconfig.conf设置，限制写入为1MB/s



[/etc/cgconfig.conf]



```shell
group DiskIO_Group  {
    blkio {
        blkio.throttle.read_bps_device = "8:0 1048576";
        blkio.throttle.write_bps_device = "8:0 1048576";
    }
}
```



同样如果限制IOPS上限为999，应该如下设定



[/etc/cgconfig.conf]



```shell
group DiskIO_Group  {
    blkio {
        blkio.throttle.read_iops_device = "8:0 999";
        blkio.throttle.write_iops_device = "8:0 999";
    }
}
```



## 参考



https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/chap-introduction_to_control_groups