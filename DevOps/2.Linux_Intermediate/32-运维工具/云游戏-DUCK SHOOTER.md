# 云游戏-DUCK SHOOTER

## 下载nginx服务

nginx能将页面出来，渲染显示

```shell
[root@youngfit ~]# yum -y install nginx
[root@youngfit ~]# rm -rf /usr/share/nginx/html/*  #删除默认的页面，以便于放后续的项目页面
```

## 上传项目包

![img](assets/云游戏-DUCK SHOOTER/1663496952286-08238b70-a3da-45ab-88f7-343b0dfa1a49.png)

```shell
[root@youngfit ~]# tar -xvzf http-not.tar.gz
```

![img](assets/云游戏-DUCK SHOOTER/1663497067090-ed7d8cd5-8cbe-428a-985e-245c8f017a6d.png)

```shell
[root@youngfit ~]# mv http-not/* /usr/share/nginx/html/
```

## 启动nginx服务

```shell
[root@youngfit ~]# systemctl start nginx
```

## 访问测试

![img](assets/云游戏-DUCK SHOOTER/1663497113932-8f58b6b7-2a54-4b43-8b8e-6aa6077f1de9.png)



## 自定义配置访问页面

Nginx的主配置文件

/etc/nginx/nginx.conf



学会配置默认发布目录，默认访问的页面

```shell
[root@jspgou ~]# vim /etc/nginx/nginx.conf
```

![img](assets/云游戏-DUCK SHOOTER/1663559245323-ae629b9a-1f36-4b8c-b29a-d9acda596d2c.png)

注释：

root   指定的是默认发布目录

index 指定的是默认访问的文件



创建对应的默认目录和默认访问的文件

```shell
[root@jspgou ~]# mkdir -p /var/www/nginx
[root@jspgou ~]# vim /var/www/nginx/index.html
HeNan Gong Xue Yuan
保存并退出
```

重新启动nginx服务

```shell
[root@jspgou ~]# systemctl restart nginx
```

再次访问，多刷新几次

![img](assets/云游戏-DUCK SHOOTER/1663559376397-f8264471-ee40-40f5-a518-a960554e1dff.png)



## 自定义显示图片

准备自己的照片，上传到服务器上，显示出来；

上传图片到/root目录

![img](assets/云游戏-DUCK SHOOTER/1663559810547-5416a21c-cdf5-4007-a007-d2e65a37d598.png)

定义nginx的默认发布目录，和默认显示文件

![img](assets/云游戏-DUCK SHOOTER/1663559864668-ac461617-bc6f-4a8f-998f-952f437814b5.png)

将图片移动到默认发布目录下

```shell
[root@jspgou ~]# mv /root/gailun.jpg /var/www/nginx/
```

重新启动nginx

```shell
[root@jspgou ~]# systemctl restart nginx
```

再次访问测试，多刷新几次

![img](assets/云游戏-DUCK SHOOTER/1663559934919-ee63b925-7c61-4993-a496-3cf4b79cd606.png)