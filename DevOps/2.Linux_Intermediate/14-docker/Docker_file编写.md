vim Dockerfile

 

FROM  #指定基础镜像

 

WORKDIR     #将工作目录切换指定目录

 

RUN    #在镜像中执行命令或脚本

 

ADD #复制指定目录里的文件到容器

 

COPY  #将本地文件或目录复制到镜像中

 

EXPOSE  #允许外界访问容器的端口

 

ENV     #设置环境变量

 

CMD  #定义容器启动后执行的命令或脚本每个Dockerfile中只能有一个CMD如果有多个那么只执行最后一个。

 

ENTRYPOINT  #定义容器启动后执行的命令或脚本，与 CMD 类似，但可以在运行时添加参数只能存在一个，若存在多个那么只执行最后一个。

 

USER  #指定容器运行时的用户

 

VOLUME #定义容器挂载的数据卷。

 

 

 

docker run # -d 镜像名 启动容器

docker pull # 下载镜像

docker images# 展示已下载的镜像名

docker load < nginx.tar #将镜像导入 

docker save # -o nginx.tar daocloud.io/library/nginx:latest 将镜像打包为tar包

docker start #让容器运行在后台

docker stop #中止容器

docker kill # 强制终止容器

docker restart# 重启容器

docker commit #把容器做成镜像 

docker inspect #展示容器详细信息

docker search #搜索镜像

docker rm #删除容器 先stop再删除

docker rmi #-f 删除指定镜像 加空格隔离删除多个镜像

docker ps # 展示当前运行容器 -a展示历史运行容器

docker attach #连接容器 但是容器创建的时候必须指定了交互shell

docker exec #-it 容器id /bin/bash 创建交互型任务 #容器id touch /testfile 后台型任务 不进入容器里面执行任务

docker build #构建容器

docker push #上传镜像

docker import #导入镜像归档文件倒其他宿主机

docker stats #展示容器资源使用统计信息的实时流（包括cpu，内存，网络I/O）

docker rename# 修改容器名称

docker pause #暂停容器中的所有进程

docker unpause #回复容器内暂停的进程

docker tag #给镜像打标签

 

\#杀死所有runing状态的容器

docker kill $(docker ps -q) 

docker stop `docker ps -q`

\#删除所有容器

docker rm $(docker ps -qf status=exited)

参数

 

--name 给容器定义名称

 

-i 标准输入输出

 

-t 分配终端

 

-d 后台运行

 

-p 映射端口

 

-e 修改环境配置

 

exec 连接容器

 

 

 

registry.cn-hangzhou.aliyuncs.com/obr/centos:7.9   #centos7.9+jdk13.02环境

 

 