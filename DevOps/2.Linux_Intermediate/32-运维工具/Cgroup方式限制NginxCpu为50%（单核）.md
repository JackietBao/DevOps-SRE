# Cgroup方式限制NginxCpu为50%（单核）

### Cgroup方式限制NginxCpu为50%（单核）



**首先，创建一个组**



```shell
cgcreate -g cpu，memory:/CPU50_Nginx_Gtoup
```



**计算公式**



![img](https://g.yuque.com/gr/latex?%5Bcpu%2Ccfs_quota_us%5D%3D%200.50*1coce*100000%3D50000(%E5%8D%95%E6%A0%B8)%0A)



**设定方式为：**



```shell
cgset -r cpu.cfs_quota_us=50000 CPU50_Nginx_Group
cgset -r cpu.cfs_period_us=100000 CPU50_Nginx_Group
```



**写到  .conf**



```shell
group CPU50_Nginx_Group {
     cpu {
            cpu.cfs_quota_us = 50000;
            cpu.cfs_period_us = 100000;
      }
}
```



**限制特定的进程**



```shell
cgexec -g cpu:CPU50_Nginx_Grouop systemctl start nginx
```