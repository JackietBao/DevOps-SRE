# Jenkins构建nodejs项目

前端vue项目。需要nodeJS基础环境；

```shell
[root@es-3-head-kib ~]# wget https://nodejs.org/dist/v14.17.6/node-v14.17.6-linux-x64.tar.xz
[root@es-3-head-kib ~]# tar xf node-v14.17.6-linux-x64.tar.xz -C /usr/local/
[root@es-3-head-kib nodejs]# vim /etc/profile
# 添加 如下配置
NODE_HOME=/usr/local/node-v14.17.6-linux-x64
JAVA_HOME=/usr/local/java
PATH=$NODE_HOME/bin:$JAVA_HOME/bin:$PATH
export JAVA_HOME PATH
```

设置淘宝源

```shell
npm config set registry https://registry.npm.taobao.org
#还原默认源：npm config set registry https://registry.npmjs.org/
```

![img](assets/Jenkins构建nodejs前端项目/1678711079979-300e3fb2-4dec-4d95-b9b6-30d60d2cec68.png)

```shell
[root@docker-server vue-admin-wygl]# nohup npm run dev &
```

![img](assets/Jenkins构建nodejs前端项目/1678711124516-f6c553e2-3f0e-4135-aba3-97e804be9b36.png)

![img](assets/Jenkins构建nodejs前端项目/1678781302659-cc901445-53f4-42f2-bfb0-65406033d3df.png)



![img](assets/Jenkins构建nodejs前端项目/1678757246430-81ee8cb5-9362-4a19-8181-a924d989b2c5.png)

![img](assets/Jenkins构建nodejs前端项目/1678757229170-68118222-1b3e-425a-ad45-809595f65174.png)