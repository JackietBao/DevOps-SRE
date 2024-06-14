# ES集群

# 概念



## 投票机制



-  每个 Elasticsearch 集群都有一个投票配置，用于决定诸如选举新主节点或提交新集群状态等决策，当投票配置中超过一半的节点做出响应时，才能做出决定。 
-  为确保集群保持可用，不得同时停止投票配置中的一半或更多节点。例如 3 个或 4 个符合 master 条件的节点，则集群可以容忍一个不可用的节点。 
-  添加或删除集群节点时，Elasticsearch 通过自动对投票配置进行相应更改来做出反应，以确保集群尽可能具有弹性。 
-  当前的投票配置不一定与集群中所有可用的主合格节点的集合相同。 
-  集群中通常应该有奇数个符合主节点资格的节点。如果有偶数，Elasticsearch 会将其中一个排除在投票配置之外，以确保其大小为奇数。 
-  第一次启动集群，需要设置初始投票配置进行引导，以此选举它的第一个主节点 



## 引导集群



首次启动 Elasticsearch 集群需要在集群中的一个或多个符合主节点的节点上明确定义初始的主节点集。这称为集群引导。这仅在集群第一次启动时需要：已加入集群的节点将此信息存储在其数据文件夹中，以用于完整的集群重新启动，而新启动的节点在加入运行中的集群时，从集群选出的主节点获得这一信息。



-  集群形成前，需要在`cluster.initial_master_nodes` 中设置定义初始合格节点列表。集群形成后，从每个节点的配置删除此配置。 
-  建议使用至少三个符合条件的主节点进行引导，每个节点都有一个包含所有三个节点的 `cluster.initial_master_nodes` 设置。 
-  默认每个节点在第一次启动时，会自动引导到单节点集群，若需要将新节点添加到现有集群中，需要进行以下设置。 

-  - `discovery.seed_providers`
   - `discovery.seed_hosts`
   - `cluster.initial_master_nodes`



## 发布集群状态



-  主节点的职责是维护全局集群状态并在节点加入或离开集群时重新分配分片。每次更改集群状态时，都会将新状态发布到集群中的所有节点。 
-  Elasticsearch 是一种基于点对点的系统，其中节点直接相互通信。高吞吐量 API（索引、删除、搜索）通常不与主节点交互。 
-  主节点允许在有限的时间内将每个集群状态更新完全发布到所有节点。由 `cluster.publish.timeout` 设置定义，默认为 30 秒 
-  如果主节点没有收到更新状态的确认（即一些节点还没有确认它们已经应用了当前的更新），这些节点被称为滞后。主节点等待滞后的节点再追上一段时间`cluster.follower_lag.timeout`，默认为90s。 



# 集群设计



## 自身配置



Elasticsearch 只需很少的配置即可开始使用，但在生产环境中使用集群之前必须考虑一些事项：



- 路径设置
- 集群名称设置
- 节点名称设置
- 网络主机设置
- 发现设置
- 堆大小设置
- JVM堆转储路径设置
- GC 日志记录设置
- 临时目录设置
- JVM致命错误日志设置
- 集群备份



## 系统配置



理想情况下，Elasticsearch 应该在服务器上单独运行并使用它可用的所有资源。为此，您需要配置您的操作系统以允许运行 Elasticsearch 的用户访问比默认允许的资源更多的资源。



在投入生产之前必须考虑以下设置：



- 配置系统限制ulimit
- 关闭swap交换内存
- 增加文件描述符
- 确保足够的虚拟内存
- 确保足够的线程
- TCP重传超时，修改 `net.ipv4.tcp_retries2` 参数



# ES集群部署



本集群配置 SSL、TLS 以及 HTTPS确保Elasticsearch的安全，暂不考虑集群角色划分的问题。

| 主机名                  | IP地址          | 配置  |
| ----------------------- | --------------- | ----- |
| elasticsearch-43、nginx | 192.168.197.143 | 2C 4G |
| elasticsearch-44、nginx | 192.168.197.144 | 2C 4G |
| elasticsearch-45、nginx | 192.168.197.145 | 2C 4G |



## 安装ES软件



```shell
gosh cmd -i /tmp/test.ip "mkdir -p /opt/{src,release,apps}"
gosh push -i /tmp/test.ip /opt/src/elasticsearch-8.1.3-linux-x86_64.tar.gz /opt/src/
gosh cmd -i /tmp/test.ip "tar -xf /opt/src/elasticsearch-8.1.3-linux-x86_64.tar.gz -C /opt/release"
gosh cmd -i /tmp/test.ip "ln -s /opt/release/elasticsearch-8.1.3 /opt/apps/elasticsearch"
gosh cmd -i /tmp/test.ip "mkdir /data/elasticsearch/ /opt/logs/elasticsearch -p"                       #创建数据目录和日志目录
gosh cmd -i /tmp/test.ip "groupadd -g 9200 es;useradd -g 9200 -u 9200 -M -s /sbin/nologin es"          #创建用户
gosh cmd -i /tmp/test.ip "chown -R es.es /opt/release/elasticsearch-8.1.3"
gosh cmd -i /tmp/test.ip "chown -R es.es  /opt/logs/elasticsearch"
gosh cmd -i /tmp/test.ip "chown -R es.es /data/elasticsearch"
```



## 配置system



```bash
gosh push -i /tmp/test.ip  /opt/config/elasticsearch/elasticsearch.service  /usr/lib/systemd/system/
gosh cmd -i /tmp/test.ip "systemctl daemon-reload"
```



```yaml
[Unit]
Description=Elasticsearch
Documentation=https://www.elastic.co
After=network.target

[Service]
User=es
Group=es
LimitNOFILE=65535
LimitNPROC=4096
LimitAS=infinity
LimitFSIZE=infinity
TimeoutStopSec=0
KillMode=control-group
Restart=on-failure
RestartSec=60
ExecStart=/opt/apps/elasticsearch/bin/elasticsearch
StandardOutput=null
StandardError=null

[Install]
WantedBy=multi-user.target
```



## 配置ssl证书



该证书用于集群节点之间认证的，避免随意加入新的 es 节点。一般教程都是使用 es 自带的命令行工具来生成，其实只需要搞懂证书的原理，就可以用openssl来生成。证书原理参考以下博客。需要注意的是，es不支持多域名证书！



https://www.yuque.com/danhuangpai/nginx/smgh7e



- 生成 CA 和服务器证书



```bash
[root@devops-7-3 ~]# mkdir -p ssl/elk
[root@devops-7-3 ~]# cd ssl/elk
[root@devops-7-3 elk]# openssl genrsa -out ca-key.pem 2048
[root@devops-7-3 elk]# openssl req -new -out ca-req.csr -key ca-key.pem
[root@devops-7-3 elk]# openssl x509 -req -in ca-req.csr -out ca-cert.pem -signkey ca-key.pem -days 3650
```



```bash
# 图省事可以仅仅生成一个空域名证书, 有CA证书就可以随意签署N多个es的证书
[root@devops-7-3 elk]# openssl genrsa -out es.key 2048
[root@devops-7-3 elk]# openssl req -new -out es-req.csr -key es.key 
[root@devops-7-3 elk]# openssl x509 -req -in es-req.csr -out es.pem -signkey es.key -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650
```



```bash
[root@devops-7-3 elk]# ll | grep -E "key|pem"
-rw-r--r-- 1 root root 1103 2020-10-04 20:49:27 ca-cert.pem
-rw-r--r-- 1 root root 1679 2020-10-04 20:48:56 ca-key.pem
-rw-r--r-- 1 root root 1679 2020-10-04 21:01:31 es.key
-rw-r--r-- 1 root root 1249 2020-10-04 21:03:02 es.pem
[root@devops-7-3 elk]# tar -zcf es-cert.tar ca-cert.pem es.key es.pem
```



- 解压证书



```bash
gosh cmd -i /tmp/test.ip "mkdir /opt/apps/elasticsearch/config/certs"
gosh push -i /tmp/test.ip es-cert.tar /opt/apps/elasticsearch/config/certs/
gosh cmd -i /tmp/test.ip "cd /opt/apps/elasticsearch/config/certs;tar -xf es-cert.tar"
gosh cmd -i /tmp/test.ip "rm -f /opt/apps/elasticsearch/config/certs/es-cert.tar"
gosh cmd -i /tmp/test.ip "cd /opt/apps/elasticsearch/config/certs;chown -R es.es * ; chmod 400 es.key"
```



## 配置elasticsearch



```plain
gosh push -i /tmp/test.ip /opt/config/elasticsearch/elasticsearch.yml /opt/apps/elasticsearch/config/
```



文件内容如下



```bash
[root@elasticsearch-46 certs]# vim /opt/apps/elasticsearch/config/elasticsearch.yml
cluster.name: log-cluster
node.name: node-2 # 不同节点需要修改
path.data: /data/elasticsearch/data 
path.logs: /opt/logs/elasticsearch 
network.host: 10.4.7.46 # 不同节点需要修改，生产中如果有DNS服务器，推荐使用域名，制作证书也用得上
http.port: 9200
# https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-discovery-settings.html
discovery.seed_hosts: ["10.4.7.45","10.4.7.46","10.4.7.47"] # 静态指定具备master资格的节点地址
cluster.initial_master_nodes: ["node-1", "node-2", "node-3"] # 初次形成集群时，显示指定具备master节点资格的列表

# 开启用户认证，仅对节点加入集群做验证
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
xpack.security.transport.ssl.verification_mode: certificate
xpack.security.transport.ssl.key: certs/es.key
xpack.security.transport.ssl.certificate: certs/es.pem
xpack.security.transport.ssl.certificate_authorities: certs/ca-cert.pem
```



## 配置系统设置



```bash
# 不要修改根 jvm.options 文件。请在 config/jvm.options.d/ 目录下添加文件。
gosh push -i /tmp/test.ip /opt/config/elasticsearch/log-center.optins /opt/apps/elasticsearch/config/jvm.options.d/
```



文件如下



```bash
[root@elasticsearch-46 certs]# vim /opt/apps/elasticsearch/config/jvm.options.d/log-center.optins
# 官方建议为物理内存的一半
-Xms2g
-Xmx2g
# 为JVM致命错误日志指定替代路径
-XX:ErrorFile=/opt/logs/elasticsearch/hs_err_pid%p.log
# 9-表示为JDK版本从9到最新，设置垃圾收集（GC）日志记录
9-:-Xlog:gc*,gc+age=trace,safepoint:file=/opt/logs/elasticsearch/gc.log:utctime,pid,tags:filecount=32,filesize=64m
```



**系统参数调整**



参考官网的参数设置（https://www.elastic.co/guide/en/elasticsearch/reference/current/setting-system-settings.html#limits.conf）



注： `/etc/security/limits.conf` 中调整对 systemctl 方式启动的进程无效，因此此处不做配置，采用非 systemctl 启动的需要参考官方文档配置。



```bash
# 切记要禁用swap分区，否则会影响到ES节点的性能，也会影响节点的稳定性。
# swapping会导致Java GC的周期延迟从毫秒级恶化到分钟，更严重的是会引起节点响应延迟甚至脱离集群。
gosh cmd -i /tmp/test.ip "swapoff -a"
gosh cmd -i /tmp/test.ip "cat >> /etc/sysctl.conf<<EOF
vm.max_map_count=262144          
vm.swappiness=0
net.ipv4.tcp_retries2=5
EOF"
gosh cmd -i /tmp/test.ip "sysctl -p  /etc/sysctl.conf"
```



**启动服务**



```plain
gosh cmd -i /tmp/test.ip "systemctl start elasticsearch.service;systemctl enable elasticsearch.service"
```



## 初始化密码



登陆到其中一台机器上进行操作



```plain
[root@alert-webhook sbin]#  cd /opt/apps/elasticsearch/
[root@alert-webhook elasticsearch]# bin/elasticsearch-setup-passwords auto
Please confirm that you would like to continue [y/N]y
Changed password for user apm_system
PASSWORD apm_system = joltoaU96pin4EmoRlFp

Changed password for user kibana_system
PASSWORD kibana_system = ivttfJVzR6Lq23ydh29F

Changed password for user kibana
PASSWORD kibana = ivttfJVzR6Lq23ydh29F

Changed password for user logstash_system
PASSWORD logstash_system = lWo0j04UPjQ8jDcgtYmc

Changed password for user beats_system
PASSWORD beats_system = tOOkdvgTxKnPqgOdn3B0

Changed password for user remote_monitoring_user
PASSWORD remote_monitoring_user = TfGUvRgrXjhOZ8Qah8q4

Changed password for user elastic
PASSWORD elastic = BIIHn2TlOj3Tg8KM7V6g
```







## 查看集群状态



```bash
[root@alert-webhook elasticsearch]#  curl -u elastic:BIIHn2TlOj3Tg8KM7V6g 192.168.197.145:9200/_cluster/health?pretty
{
  "cluster_name" : "log-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 3,
  "number_of_data_nodes" : 3,
  "active_primary_shards" : 2,
  "active_shards" : 4,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```



## 配置ES的HTTP LB



### nginx 安装



```bash
# 下载软件包到部署机器，并推送到nginx节点
[root@danhuangpai monitor]# wget http://tengine.taobao.org/download/tengine-2.3.3.tar.gz -O /opt/src/tengine-2.3.3.tar.gz
[root@danhuangpai monitor]# gosh push -i /tmp/test.ip /opt/src/tengine-2.3.3.tar.gz  /opt/src/
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "tar -xf /opt/src/tengine-2.3.3.tar.gz -C /opt/src"

# 编译tengine
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "yum install -y gcc gcc-c++ pcre pcre-devel zlib zlib-devel openssl openssl-devel"
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "groupadd -g 8080 nginx; useradd -g 8080 -u 8080 -s /sbin/nologin -M nginx"
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "cd /opt/src/tengine-2.3.3 && ./configure --user=nginx --group=nginx --prefix=/opt/release/tengine-2.3.3 --with-threads --with-http_ssl_module --with-http_v2_module --with-http_realip_module --with-stream --add-module=modules/ngx_http_upstream_check_module >/dev/null"
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "cd /opt/src/tengine-2.3.3 && make >/dev/null && make install >/dev/null"
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "ln -sf /opt/release/tengine-2.3.3 /opt/apps/nginx"
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "mkdir /opt/logs/nginx;chown nginx.nginx /opt/logs/nginx/"
```



```bash
# 配置service，并启动nginx
[root@danhuangpai monitor]# gosh push -i /tmp/test.ip /opt/config/nginx/nginx.service /usr/lib/systemd/system/
[root@danhuangpai monitor]# gosh push -i /tmp/test.ip /opt/config/nginx/nginx.conf /opt/apps/nginx/conf/
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "systemctl daemon-reload;systemctl start nginx ;systemctl enable nginx"
```



**system配置文件**



```yaml
[Unit]
Description=The nginx HTTP and reverse proxy server
After=network-online.target remote-fs.target nss-lookup.target chronyd.service
Wants=network-online.target

[Service]
Type=forking
PIDFile=/opt/apps/nginx/logs/nginx.pid
ExecStartPre=/usr/bin/rm -f /opt/apps/nginx/logs/nginx.pid
ExecStartPre=/opt/apps/nginx/sbin/nginx -t
ExecStart=/opt/apps/nginx/sbin/nginx
ExecReload=/opt/apps/nginx/sbin/nginx -s reload
KillSignal=SIGQUIT
TimeoutStopSec=5
KillMode=process
PrivateTmp=true
LimitNOFILE=655350
LimitNPROC=65535

[Install]
WantedBy=multi-user.target
```



**nginx主配置文件**



```yaml
worker_processes  1;
events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;
    error_log /opt/logs/nginx/error.log error ;
    log_format access '$time_local|$http_x_real_ip|$remote_addr|$http_x_forwarded_for|$upstream_addr|'
                      '$request_method|$server_protocol|$host|$request_uri|$http_referer|$http_user_agent|'
                      '$proxy_host|$status' ;
    access_log  /opt/logs/nginx/access.log  access ;

    sendfile        on;
    keepalive_timeout  65;
    include     conf.d/*.conf;
}
```



### 配置es的HTTP LB



- 生成es 的域名证书



需要注意的是，logstash需要连接该域名，而 elk 组件不支持解析多域名证书，因此需要单独为 elasticsearch.huanle.com 签发证书。



```bash
[root@devops-7-3 elk]# openssl genrsa -out elasticsearch-ddn.key 2048
[root@devops-7-3 elk]# openssl req -new -out elasticsearch-ddn-req.csr -key elasticsearch-ddn.key
[root@devops-7-3 elk]# openssl x509 -req -in elasticsearch-ddn-req.csr -out elasticsearch-ddn.pem -signkey elasticsearch-ddn.key -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650
Signature ok
subject=/C=CN/ST=JS/L=NJ/O=Default Company Ltd/CN=elasticsearch.ddn.com
Getting Private key
Getting CA Private Key
[root@devops-7-3 elk]# chmod 400 elasticsearch-ddn*
[root@devops-7-3 elk]# gosh push -i /tmp/test.ip /opt/config/ssl/elk/es-http* /opt/apps/nginx/conf/ssl_key/
```



- 配置es的LB



```bash
[root@danhuangpai monitor]# gosh push -i /tmp/test.ip /opt/config/nginx/conf.d/elasticsearch.conf  /opt/apps/nginx/conf/conf.d/
[root@danhuangpai monitor]# gosh cmd -i /tmp/test.ip "systemctl restart nginx"
[root@danhuangpai monitor]# cat /opt/apps/nginx/conf/conf.d/elasticsearch.conf 
server {
    listen 443 ssl;
    server_name elasticsearch.huanle.com;
    keepalive_timeout 100s ;    ## 增加长连接的时间
    keepalive_requests 200 ;

    ssl_certificate ssl_key/es-http.pem ;
    ssl_certificate_key ssl_key/es-http.key ;
    ssl_ciphers HIGH:!aNULL:!MD5 ;
    ssl_session_cache shared:SSL:30m ;  ## 设置SSL session缓存
    ssl_session_timeout 10m ;
    location / {
        proxy_pass http://elasticsearch;
    }
}

upstream elasticsearch {
    server 192.168.197.143:9200;
    server 192.168.197.144:9200;
    server 192.168.197.145:9200;
}
```



# Kibana部署



## 安装kibana



需要注意的时，此项目中，kibana直接连接ES的LB，因此在走 https 协议情况，es 的域名证书不能是多域名证书。



```bash
cd /opt/src/
curl -O https://artifacts.elastic.co/downloads/kibana/kibana-8.1.3-linux-x86_64.tar.gz
gosh push -i /tmp/test.ip  /opt/src/kibana-8.1.3-linux-x86_64.tar.gz /opt/src
gosh cmd -i /tmp/test.ip "tar -xf /opt/src/kibana-8.1.3-linux-x86_64.tar.gz -C /opt/release/"
gosh cmd -i /tmp/test.ip "ln -s /opt/release/kibana-8.1.3 /opt/apps/kibana"
gosh push  -i /tmp/test.ip /opt/config/ssl/elk/ca-cert.pem  /opt/apps/kibana/config/certs/
gosh push  -i /tmp/test.ip  /opt/config/kibana/kibana.yml /opt/apps/kibana/config/
gosh push  -i /tmp/test.ip  /opt/config/kibana/kibana.service  /usr/lib/systemd/system/
gosh cmd  -i /tmp/test.ip "systemctl daemon-reload"
gosh cmd  -i /tmp/test.ip "systemctl start kibana.service ; systemctl enable kibana.service"
```



kibana配置文件



```yaml
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.hosts: ["https://elasticsearch.huanle.com:443"]
#elasticsearch.ssl.certificateAuthorities: ["/opt/apps/kibana/config/certs/ca-cert.pem"]
elasticsearch.username: "kibana_system"
elasticsearch.password: "RZ2uPE1BkHOKJJP6A4Xu"
i18n.locale: zh-CN              # 设置中文
```



system配置文件



```yaml
[root@kibana-48 src]# vim /usr/lib/systemd/system/kibana.service
[Unit]
Description=kibana service
After=network.target

[Service]
User=nobody
Group=nobody
KillMode=control-group
Restart=on-failure
RestartSec=60
ExecStart=/opt/apps/kibana/bin/kibana

[Install]
WantedBy=multi-user.target
[root@kibana-48 src]# systemctl daemon-reload
[root@kibana-48 src]# systemctl start kibana.service ; systemctl enable kibana.service
```



**需要注意**：在启动 Kibana 的时候日志会有一些报错(warning和error)，在 /var/log/message 中可以查看，读者可以自行解决(解决报错时，建议直接命令行执行 `/opt/apps/kibana/bin/kibana --allow-root` )，即使不处理，也能使用 kibana 的基本功能，常用操作不受影响。



## 配置kibana的代理



```bash
[root@devops-7-3 elk]# openssl genrsa -out kibana-ddn.key 2048
[root@devops-7-3 elk]# openssl req -new -out kibana-ddn-req.csr -key kibana-ddn.key
[root@devops-7-3 elk]# openssl x509 -req -in kibana-ddn-req.csr -out kibana-ddn.pem -signkey kibana-ddn.key -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650
Signature ok
subject=/C=CN/ST=JS/L=NJ/O=Default Company Ltd/CN=kibana.ddn.com
Getting Private key
Getting CA Private Key
[root@devops-7-3 elk]# chmod 400 kibana-ddn*
[root@devops-7-3 elk]# ll kibana-ddn*
-r-------- 1 root root 1675 2020-10-11 21:54:18 kibana-ddn.key
-r-------- 1 root root 1139 2020-10-11 21:55:36 kibana-ddn.pem
-r-------- 1 root root  993 2020-10-11 21:55:06 kibana-ddn-req.csr
[root@devops-7-3 elk]# scp kibana-ddn* 10.4.7.41:/opt/apps/nginx/conf/ssl_key/
```



```plain
gosh push -i /tmp/test.ip  /opt/config/nginx/conf.d/kibana.conf  /opt/apps/nginx/conf/conf.d/
gosh cmd -i /tmp/test.ip "systemctl restart nginx"
```



```bash
[root@filebeat-41 ~]# cat /opt/apps/nginx/conf/conf.d/kibanna.conf 
server {
    listen 443 ssl;
    server_name kibana.ddn.com;
    keepalive_timeout 100s ;    ## 增加长连接的时间
    keepalive_requests 200 ;

    ssl_certificate ssl_key/kibana-ddn.pem ;
    ssl_certificate_key ssl_key/kibana-ddn.key ;
    ssl_ciphers HIGH:!aNULL:!MD5 ;
    ssl_session_cache shared:SSL:30m ;  ## 设置SSL session缓存
    ssl_session_timeout 10m ;
    location / {
        proxy_pass http://10.4.7.48:5601;
    }
}
```



# 插件



新版未同步，暂不安装



```bash
gosh cmd -i /tmp/test.ip "mkdir  /opt/apps/elasticsearch/plugins/{ik,hanlp}"
cd /opt/src
wget https://github.com/kepmov/elasticsearch-analysis-hanlp/releases/download/v7.x/analysis-hanlp.zip
gosh push -i /tmp/test.ip /opt/src/analysis-hanlp.zip /opt/src/
gosh cmd -i /tmp/test.ip "unzip /opt/src/analysis-hanlp.zip -d /opt/apps/elasticsearch/plugins/hanlp/"
# 修改配置
sed -i 's#^\(root=\).*$#\1/opt/release/analysis-hanlp/data#' hanlp.properties
sed -i 's#^\(elasticsearch.version=\).*$#\18.1.3#' plugin-descriptor.properties
gosh push -i /tmp/test.ip /opt/release/analysis-hanlp/hanlp.properties /opt/apps/elasticsearch/plugins/hanlp/analysis-hanlp
gosh push -i /tmp/test.ip /opt/release/analysis-hanlp/plugin-descriptor.properties /opt/apps/elasticsearch/plugins/hanlp/analysis-hanlp
# 修改es的jvm.options文件
gosh cmd -i /tmp/test.ip "cp /opt/apps/elasticsearch/config/jvm.options /opt/apps/elasticsearch/config/jvm.options.bak"
gosh cmd -i /tmp/test.ip "echo "-Djava.security.policy=/opt/apps/elasticsearch/plugins/analysis-hanlp/plugin-security.policy" >> /opt/apps/elasticsearch/config/jvm.options"
```