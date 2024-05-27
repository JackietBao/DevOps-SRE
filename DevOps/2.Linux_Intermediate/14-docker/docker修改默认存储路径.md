1、停掉 Docker 进程:

```
systemctl stop docker。
```


2、将 Docker 存储目录拷贝到要迁移的目录中去

（例如，此处为/home/docker/lib/）：

```
mkdir -p /home/docker/lib/
rsync -r -avz /var/lib/docker /home/docker/lib/
```


3、链接迁移目录到原目录：

```
mv /var/lib/docker /var/lib/docker-old

ln -s /home/docker/lib/docker /var/lib/docker
```


##编辑 /usr/lib/systemd/system/docker.service

##加入新的工作目录。

4、修改 Docker 配置文件，将工作目录指向新创建的目录

```
vi /etc/docker/daemon.json
在该文件中添加以下内容：
{ "data-root": "/home/docker/lib/docker/" }
```

5、重新载入系统服务: 

```
systemctl daemon-reload
```


6、启动 Docker: 

```
systemctl start docker
```

7、查看 Docker 的 info 信息，确认工作目录已经修改过来

![image-20231008091035556](assets/docker修改默认存储路径/image-20231008091035556.png)

8、（可选）在确定镜像和容器正常运行后，可以删除 /var/lib/docker-old 目录



![image-20231008090426945](assets/docker修改默认存储路径/image-20231008090426945.png)