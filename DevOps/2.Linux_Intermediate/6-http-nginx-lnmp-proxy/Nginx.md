### yum安装nginx

安装依赖

```
yum -y install gcc gcc-c++ autoconf pcre-devel make automake httpd-tools
```

配置官方yum源

vim /etc/yum.repos.d/nginx.repo

```
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
```

安装nginx服务，启动并加入开机自启

```
yum -y install nginx
systemctl enable nginx
systemctl start nginx
```

检查nginx软件版本以编译参数

```
nginx -v
nginx -V
```

rpm -ql查看整体目录结构和功能

```
[root@localhost ~]# rpm -ql nginx
/etc/logrotate.d/nginx
/etc/nginx
/etc/nginx/conf.d
/etc/nginx/conf.d/default.conf
/etc/nginx/fastcgi_params
/etc/nginx/mime.types
/etc/nginx/modules
/etc/nginx/nginx.conf
/etc/nginx/scgi_params
/etc/nginx/uwsgi_params
/usr/lib/systemd/system/nginx-debug.service
/usr/lib/systemd/system/nginx.service
/usr/lib64/nginx
/usr/lib64/nginx/modules
/usr/libexec/initscripts/legacy-actions/nginx
/usr/libexec/initscripts/legacy-actions/nginx/check-reload
/usr/libexec/initscripts/legacy-actions/nginx/upgrade
/usr/sbin/nginx
/usr/sbin/nginx-debug
/usr/share/doc/nginx-1.26.0
/usr/share/doc/nginx-1.26.0/COPYRIGHT
/usr/share/man/man8/nginx.8.gz
/usr/share/nginx
/usr/share/nginx/html
/usr/share/nginx/html/50x.html
/usr/share/nginx/html/index.html
/var/cache/nginx
/var/log/nginx
```

### nginx基本配置

/etc/nginx/nginx.conf

> coremodule(核心模块)
>
> eventmodule(事件驱动模块)
>
> httpcoremodule(http内核模块)
