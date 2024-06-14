# Watchtower自动更新容器镜像

对开发运维人员来说，保持 Docker 容器为最新版本是重要工作之一。手动更新 Docker 容器是一项耗时的工作。这篇文章解释了 **Watchtower** 是什么，如何安装它，以及在 Linux 中如何 **使用 Watchtower 自动更新正在运行的 Docker 容器** 。

## Watchtower 是什么？

**Watchtower** 是一款自由开源的应用，用来监控运行中的 Docker 容器，并且当它发现基础镜像被更改后，可以自动的更新容器。

若 Watchtower 发现一个运行中的容器需要更新，它会以发送 SIGTERM 信号的方式，优雅的结束运行中容器的运行。

它会下载新镜像，然后以最初部署时使用的方式，重启容器。所有文件会在后台自动下载，因此不需要用户的介入。

在这份指南中，我们将会明白如何在类 Unix 系统中使用 Watchtower 自动更新正在运行的 Docker 容器。



官网：https://containrrr.dev/watchtower/

| **传统方式** | **Watchtower** |
| ------------ | -------------- |
| 耗时费力。   |                |



【1】stop，停止容器；rm删除容器（docker compose启容器也只会缩减一步）
【2】rmi删除旧镜像
【3】pull拉取新镜像
【4】重新启动容器 | 能自动拉取最新的docker镜像并将其自动运行，能在很大程度上减少运维的工作量 |



watchtower部署在主服务器，通过次服务器进行新镜像的build以及push到私仓



## 1、Watchtower部署(本次采用手动更新)



【1】手动方式，次服务器通过ssh远端控制主服务器更新容器镜像，存在不安全隐患
【2】采用设置更新频率的方式，对于特定需求下，存在更新不及时的情况
【3】可采用HTTP API模式，通过HTTP请求及时触发容器镜像的更新



```shell
vim compose.yml
version: "3.9"

services:
  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    deploy:
      restart_policy:
        condition: on-failure
        delay: 5s
        window: 60s
    environment:
      - WATCHTOWER_TIMEOUT=300s
      #- WATCHTOWER_HTTP_API_TOKEN=mytoken
    labels:
      - "com.centurylinklabs.watchtower.enable=false"
    volumes:
      - /root/.docker/config.json:/config.json
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 18080:8080
    privileged: true
    cap_add:
      - ALL
    command: --cleanup --label-enable --run-once #--http-api-periodic-polls --interval 60


--run-once   #采用手动（容器启动一次，自动退出）
--interval 60  #自动更新检查频率方式，每60秒自动检测一次更新

*非手动模式，容器重启策略建议always

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
若采用HTTP API模式，仅允许由 HTTP 请求触发映像更新   --http-api-update

参考：https://containrrr.dev/watchtower/http-api-mode/


若开始上述环境变量- WATCHTOWER_HTTP_API_TOKEN=mytoken

次服务器（或任意可访问主服务器的服务器）通过
curl -H "Authorization: Bearer mytoken" 主服务器IP:18080/v1/update


发送HTTP请求，触发主服务器的watchtower进行容器镜像的更新
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
```



## 2、测试容器准备



```shell
vim compose.yml
x-config:
  &default-config
  deploy:
      restart_policy:
        condition: always
        delay: 5s
        window: 60s
  labels:
    - "com.centurylinklabs.watchtower.enable=true"   #true表示更新
  environment:
    - BIMINI_ENV=PROD
  volumes:
    - /etc/localtime:/etc/localtime:ro
    - /etc/timezone:/etc/timezone:ro
  logging:
    driver: "json-file"
    options:
      max-size: 4G
  network_mode: host
  working_dir: /


services:
  test-ubuntu:
    << : *default-config
    image: 192.168.10.81:8082/test/test-ubuntu:v1
    container_name: test-ubuntu
    tty: true



也可
version: "3"

services:
  test-ubuntu:
    restart: always
    image: 192.168.10.81:8082/test/test-ubuntu:v1
    container_name: test-ubuntu
    tty: true
    labels:
      - "com.centurylinklabs.watchtower.enable=true"   #true表示更新
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - /etc/timezone:/etc/timezone:ro
    logging:
      driver: "json-file"
      options:
        max-size: 4G
```



## 3、远程工具准备（另一台）



**sshpass**
安装



```shell
$ sudo apt update
$ sudo apt install sshpass
```



基本使用



```shell
$ sshpass -p '123456' ssh cwy@192.168.10.81


********若没有任何反应，可以通过ssh cwy@192.168.10.81先建立连接
```



sudo: no tty present and no askpass program specified



```shell
visudo

%sudo   ALL=(ALL:ALL) ALL

变为
%sudo   ALL=(ALL:ALL) NOPASSWD:ALL
```



## 4、编写脚本（另一台）



```shell
vim test.sh
#/bin/bash
img=192.168.10.81:8082/test/test-ubuntu
ver=v1

docker push $img:$ver

sleep 2

sshpass -p '123456' ssh cwy@192.168.10.81 'sudo docker compose -f /opt/services/watchtower/compose.yml up -d'


chmod +x test.sh
```



## 5、测试（另一台）



```shell
#通过以下命令得到镜像
docker build -t 192.168.10.81:8082/test/test-ubuntu:v1 .


#执行脚本
./test.sh


观察镜像发生改变
```



![img](https://cdn.nlark.com/yuque/0/2023/png/29105411/1675914003068-8520d788-6fb4-4e50-86bd-9ec1a45b8aa6.png#averageHue=%23170f0d&clientId=uef3ec712-bb21-4&from=paste&height=116&id=uacd0a7fe&name=image.png&originHeight=145&originWidth=1798&originalType=binary&ratio=1&rotation=0&showTitle=false&size=41629&status=done&style=none&taskId=u61b67403-000e-4bfc-af80-63855e4da26&title=&width=1438.4)