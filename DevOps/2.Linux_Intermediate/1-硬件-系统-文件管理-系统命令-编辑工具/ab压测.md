## ab 压力测试工具

### 1、介绍

- 吞吐率（Requests per second）
  概念：服务器并发处理能力的量化描述，单位是reqs/s，指的是某个并发用户数下单位时间内处理的请求数。某个并发用户数下单位时间内能处理的最大请求数，称之为最大吞吐率。
  计算公式：总请求数(PV量) / 处理完成这些请求数所花费的时间，即 Request per second = Complete requests
  / Time taken for tests
- 并发连接数（The number of concurrent connections）
  概念：某个时刻服务器所接受的请求数目，简单的讲，就是一个会话。
- 并发用户数（The number of concurrent users，Concurrency Level）
  概念：要注意区分这个概念和并发连接数之间的区别，一个用户可能同时会产生多个会话，也即连接数。
- 用户平均请求等待时间（Time per request） 计算公式：处理完成所有请求数所花费的时间/ （总请求数 / 并发用户数），即
  Time per request = Time taken for tests /（ Complete requests /Concurrency Level）
- 服务器平均请求等待时间（Time per request: across all concurrent requests）
  计算公式：处理完成所有请求数所花费的时间 / 总请求数，即 Time taken for / testsComplete requests
  可以看到，它是吞吐率的倒数。 同时，它也=用户平均请求等待时间/并发用户数，即 Time per request /Concurrency Level
- Apache Benchmark(简称ab) 是Apache安装包中自带的压力测试工具 ，简单易用
- 在此提供 ab 在 centOS7 下的安装和使用方法

### 2、安装

```shell
[root@qfedu.com ~]# yum -y install httpd-tools 
```

### 4、命令使用

```shell
[root@qfedu.com ~]# ab -n 1000 -c 10 http://www.baidu.com # -n访问1000次, -c并发10个
```

- 更多使用方法详见 [ab 官方文档](https://links.jianshu.com/go?to=http%3A%2F%2Fhttpd.apache.org%2Fdocs%2F2.0%2Fprograms%2Fab.html)

### 5、执行结果

```bash
Server Software:        Apache          #服务器软件
Server Hostname:        www.taoquan.ink #域名
Server Port:            80              #请求端口号
Document Path:          /               #文件路径
Document Length:        40888 bytes     #页面字节数
Concurrency Level:      10              #请求的并发数
Time taken for tests:   27.300 seconds  #总访问时间
Complete requests:      1000            #请求成功数量
Failed requests:        0               #请求失败数量
Write errors:           0
Total transferred:      41054242 bytes  #请求总数据大小（包括header头信息）
HTML transferred:       40888000 bytes  #html页面实际总字节数
Requests per second:    36.63 [#/sec] (mean)  #每秒多少请求，这个是非常重要的参数数值，服务器的吞吐量
Time per request:       272.998 [ms] (mean)     #用户平均请求等待时间 
Time per request:       27.300 [ms] (mean, across all concurrent requests) # 服务器平均处理时间，也就是服务器吞吐量的倒数                  
Transfer rate:          1468.58 [Kbytes/sec] received  #每秒获取的数据长度
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       43   47   2.4     47      53
Processing:   189  224  40.7    215     895
Waiting:      102  128  38.6    118     794
Total:        233  270  41.3    263     945

Percentage of the requests served within a certain time (ms)
  50%    263    #50%用户请求在263ms内返回
  66%    271    #66%用户请求在271ms内返回
  75%    279    #75%用户请求在279ms内返回
  80%    285    #80%用户请求在285ms内返回
  90%    303    #90%用户请求在303ms内返回
  95%    320    #95%用户请求在320ms内返回
  98%    341    #98%用户请求在341ms内返回
  99%    373    #99%用户请求在373ms内返回
 100%    945 (longest request)
```

# ab常用参数

-n ：总共的请求执行数，缺省是1；

-c： 并发数，缺省是1；

-t：测试所进行的总时间，秒为单位，缺省50000s

-p：POST时的数据文件

-w: 以HTML表的格式输出结果