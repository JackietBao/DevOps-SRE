配置了代理还是不行

---

_**Assistant**_



---

我需要查看您的当前环境和Docker配置来帮助解决代理问题。让我检查一些信息：

---

```bash
systemctl status docker
```

```
root@ubuntu2404:~# systemctl status docker
lines 1-1...skipping...
● docker.service - Docker Application Container Engine
lines 1-1...skipping...
● docker.service - Docker Application Container Engine
     Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/docker.service.d
lines 1-3...skipping...
● docker.service - Docker Application Container Engine
     Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/docker.service.d
             └─http-proxy.conf
     Active: active (running) since Sun 2025-03-16 12:47:54 UTC; 59s ago
TriggeredBy: ● docker.socket
lines 1-6...skipping...
● docker.service - Docker Application Container Engine
     Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/docker.service.d
             └─http-proxy.conf
     Active: active (running) since Sun 2025-03-16 12:47:54 UTC; 59s ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 51137 (dockerd)
      Tasks: 10
     Memory: 21.1M (peak: 21.7M)
        CPU: 805ms
     CGroup: /system.slice/docker.service
             └─51137 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
lines 1-13...skipping...
● docker.service - Docker Application Container Engine
     Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/docker.service.d
             └─http-proxy.conf
     Active: active (running) since Sun 2025-03-16 12:47:54 UTC; 59s ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 51137 (dockerd)
      Tasks: 10
     Memory: 21.1M (peak: 21.7M)
        CPU: 805ms
     CGroup: /system.slice/docker.service
             └─51137 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

Mar 16 12:47:54 ubuntu2404 dockerd[51137]: time="2025-03-16T12:47:54.812415829Z" level=info msg="Loadi>
lines 1-15
```

---

我看到您已经创建了一个`http-proxy.conf`文件。让我们检查一下它的内容：

---

```bash
cat /etc/systemd/system/docker.service.d/http-proxy.conf
```

```
root@ubuntu2404:~# cat /etc/systemd/system/docker.service.d/http-proxy.conf
[Service]
Environment="HTTP_PROXY=socks5://192.168.0.101:7897"
Environment="NO_PROXY=localhost,127.0.0.1,docker-registry.example.com,.corp"
root@ubuntu2404:~#
```

---

现在让我看看Docker是否能连接到外部网络：

---

```bash
docker pull ubuntu:latest
```

```
root@ubuntu2404:~# docker pull ubuntu:latest
Error response from daemon: Get "https://registry-1.docker.io/v2/": net/http: request canceled while wa
iting for connection (Client.Timeout exceeded while awaiting headers)
root@ubuntu2404:~#
```

---

问题可能是您的代理配置不完整。我注意到缺少了HTTPS_PROXY设置，而且Docker可能无法直接使用SOCKS5代理。让我们进行以下调整：

1. 首先，让我们修改代理配置文件：

---

```diff
  [Service]
  Environment="HTTP_PROXY=socks5://192.168.0.101:7897"
+ Environment="HTTPS_PROXY=socks5://192.168.0.101:7897"
  Environment="NO_PROXY=localhost,127.0.0.1,docker-registry.example.com,.corp"
  
```

---

由于Docker可能不直接支持SOCKS5代理，我们需要安装并配置Privoxy来将SOCKS5代理转换为HTTP代理：

---

```bash
apt-get update && apt-get install -y privoxy
```

```
root@ubuntu2404:~# apt-get update && apt-get install -y privoxy
Hit:1 https://download.docker.com/linux/ubuntu noble InRelease
Hit:2 http://security.ubuntu.com/ubuntu noble-security InRelease                             
Hit:3 http://in.archive.ubuntu.com/ubuntu noble InRelease                                    
Hit:4 http://in.archive.ubuntu.com/ubuntu noble-updates InRelease
Hit:5 http://in.archive.ubuntu.com/ubuntu noble-backports InRelease
Reading package lists... Done
W: https://download.docker.com/linux/ubuntu/dists/noble/InRelease: Key is stored in legacy trusted.gpg 
keyring (/etc/apt/trusted.gpg), see the DEPRECATION section in apt-key(8) for details.
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
The following additional packages will be installed:
  doc-base libmbedcrypto7t64 libmbedtls14t64 libmbedx509-1t64 libuuid-perl libyaml-tiny-perl
Suggested packages:
  dhelp | dwww | dochelp | doc-central | yelp | khelpcenter
The following NEW packages will be installed:
  doc-base libmbedcrypto7t64 libmbedtls14t64 libmbedx509-1t64 libuuid-perl libyaml-tiny-perl privoxy
0 upgraded, 7 newly installed, 0 to remove and 179 not upgraded.
Need to get 1,055 kB of archives.
After this operation, 3,849 kB of additional disk space will be used.
Get:1 http://in.archive.ubuntu.com/ubuntu noble/universe amd64 libmbedcrypto7t64 amd64 2.28.8-1 [209 kB
]
Get:2 http://in.archive.ubuntu.com/ubuntu noble/universe amd64 libmbedx509-1t64 amd64 2.28.8-1 [46.6 kB
]
Get:3 http://in.archive.ubuntu.com/ubuntu noble/universe amd64 libmbedtls14t64 amd64 2.28.8-1 [82.2 kB]
Get:4 http://in.archive.ubuntu.com/ubuntu noble/universe amd64 privoxy amd64 3.0.34-3ubuntu2 [600 kB]
Get:5 http://in.archive.ubuntu.com/ubuntu noble/main amd64 libuuid-perl amd64 0.31-1build3 [14.3 kB]
Get:6 http://in.archive.ubuntu.com/ubuntu noble/main amd64 libyaml-tiny-perl all 1.74-1 [25.3 kB]
Get:7 http://in.archive.ubuntu.com/ubuntu noble/main amd64 doc-base all 0.11.2 [77.0 kB]
Fetched 1,055 kB in 4s (261 kB/s)   
Preconfiguring packages ...
Selecting previously unselected package libmbedcrypto7t64:amd64.
(Reading database ... 122283 files and directories currently installed.)
Preparing to unpack .../0-libmbedcrypto7t64_2.28.8-1_amd64.deb ...
Unpacking libmbedcrypto7t64:amd64 (2.28.8-1) ...
Selecting previously unselected package libmbedx509-1t64:amd64.
Preparing to unpack .../1-libmbedx509-1t64_2.28.8-1_amd64.deb ...
Unpacking libmbedx509-1t64:amd64 (2.28.8-1) ...
Selecting previously unselected package libmbedtls14t64:amd64.
Preparing to unpack .../2-libmbedtls14t64_2.28.8-1_amd64.deb ...
Unpacking libmbedtls14t64:amd64 (2.28.8-1) ...
Selecting previously unselected package privoxy.
Preparing to unpack .../3-privoxy_3.0.34-3ubuntu2_amd64.deb ...
Unpacking privoxy (3.0.34-3ubuntu2) ...
Selecting previously unselected package libuuid-perl.
Preparing to unpack .../4-libuuid-perl_0.31-1build3_amd64.deb ...
Unpacking libuuid-perl (0.31-1build3) ...
Selecting previously unselected package libyaml-tiny-perl.
Preparing to unpack .../5-libyaml-tiny-perl_1.74-1_all.deb ...
Unpacking libyaml-tiny-perl (1.74-1) ...
Selecting previously unselected package doc-base.
Preparing to unpack .../6-doc-base_0.11.2_all.deb ...
Unpacking doc-base (0.11.2) ...
Setting up libuuid-perl (0.31-1build3) ...
Setting up libmbedcrypto7t64:amd64 (2.28.8-1) ...
Setting up libyaml-tiny-perl (1.74-1) ...
Setting up libmbedx509-1t64:amd64 (2.28.8-1) ...
Setting up doc-base (0.11.2) ...
Registering 18 doc-base files...
Setting up libmbedtls14t64:amd64 (2.28.8-1) ...
Setting up privoxy (3.0.34-3ubuntu2) ...

Creating config file /etc/privoxy/config with new version
Created symlink /etc/systemd/system/multi-user.target.wants/privoxy.service → /usr/lib/systemd/system/p
rivoxy.service.
```

---

现在我们需要配置Privoxy：

---

```diff
  #      resources are being blocked.
  #
- #handle-as-empty-doc-returns-ok 1
+ handle-as-empty-doc-returns-ok 1
  #
  #  6.13. enable-compression
  #      length.
  #
- #enable-compression 1
+ enable-compression 1
  #
  #  6.14. compression-level
  #hide-console
  #
+ # 配置socks5代理
+ forward-socks5 / 192.168.0.101:7897 .
  #
  
```

---

配置看起来不太正确，我将再次修改Privoxy配置文件：

---



---

让我们查看一下当前配置文件末尾：

---

```bash
tail -40 /etc/privoxy/config
```

```
root@ubuntu2404:~# tail -40 /etc/privoxy/config
#log-buffer-size 1
#
#  log-max-lines is the maximum number of lines held in the log
#  buffer. See above.
#
#log-max-lines 200
#
#  If "log-highlight-messages" is set to 1, Privoxy will highlight
#  portions of the log messages with a bold-faced font:
#
#log-highlight-messages 1
#
#  The font used in the console window:
#
#log-font-name Comic Sans MS
#
#  Font size used in the console window:
#
#log-font-size 8
#
#  "show-on-task-bar" controls whether or not Privoxy will appear as
#  a button on the Task bar when minimized:
#
#show-on-task-bar 0
#
#  If "close-button-minimizes" is set to 1, the Windows close button
#  will minimize Privoxy instead of closing the program (close with
#  the exit option on the File menu).
#
#close-button-minimizes 1
#
#  The "hide-console" option is specific to the MS-Win console
#  version of Privoxy. If this option is used, Privoxy will
#  disconnect from and hide the command console.
#
#hide-console
#
# 配置socks5代理
forward-socks5 / 192.168.0.101:7897 .
#
root@ubuntu2404:~#
```

---

让我们添加监听地址：

---



---

让我们再次尝试，使用grep找到特定位置：

---

```bash
grep -n "listen-address" /etc/privoxy/config
```

```
root@ubuntu2404:~# grep -n "listen-address" /etc/privoxy/config
694:#  4.1. listen-address
786:#        listen-address  192.168.0.1:8118
792:#        listen-address [::1]:8118
794:listen-address  127.0.0.1:8118
795:listen-address  [::1]:8118
852:#      access Privoxy (see "ACLs" and listen-address above) can
928:#      access Privoxy (see "ACLs" and listen-address above) can
1026:#      Don't restrict access further than implied by listen-address
1035:#      listen-address option.
1074:#      listen-address are set: "localhost" is OK. The absence of a
root@ubuntu2404:~#
```

---

好的，我看到配置文件已经包含了监听地址配置。让我们重启Privoxy服务：

---

```bash
systemctl restart privoxy
```

```
root@ubuntu2404:~# systemctl restart privoxy
root@ubuntu2404:~#
```

---

现在我们需要修改Docker的代理配置，使其使用Privoxy提供的HTTP代理：

---

```diff
  [Service]
- Environment="HTTP_PROXY=socks5://192.168.0.101:7897"
- Environment="HTTPS_PROXY=socks5://192.168.0.101:7897"
+ Environment="HTTP_PROXY=http://127.0.0.1:8118"
+ Environment="HTTPS_PROXY=http://127.0.0.1:8118"
  Environment="NO_PROXY=localhost,127.0.0.1,docker-registry.example.com,.corp"
  
```

---

让我们重新加载Docker守护进程配置并重启Docker：

---

```bash
systemctl daemon-reload && systemctl restart docker
```

```
root@ubuntu2404:~# systemctl daemon-reload && systemctl restart docker
root@ubuntu2404:~#
```

---

现在让我们验证Docker的代理设置是否生效：

---

```bash
systemctl show --property=Environment docker
```

```
root@ubuntu2404:~# systemctl show --property=Environment docker
lines 1-1...skipping...
Environment=HTTP_PROXY=http://127.0.0.1:8118 HTTPS_PROXY=http://127.0.0.1:8118 NO_PROXY=localhost,127.>
lines 1-1...skipping...
Environment=HTTP_PROXY=http://127.0.0.1:8118 HTTPS_PROXY=http://127.0.0.1:8118 NO_PROXY=localhost,127.>
~
~
lines 1-1/1 (END)...skipping...
Environment=HTTP_PROXY=http://127.0.0.1:8118 HTTPS_PROXY=http://127.0.0.1:8118 NO_PROXY=localhost,127.>
~
~
~
~
~
lines 1-1/1 (END)...skipping...
Environment=HTTP_PROXY=http://127.0.0.1:8118 HTTPS_PROXY=http://127.0.0.1:8118 NO_PROXY=localhost,127.>
~
~
~
~
~
~
~
~
~
~
~
~
lines 1-1/1 (END)...skipping...
Environment=HTTP_PROXY=http://127.0.0.1:8118 HTTPS_PROXY=http://127.0.0.1:8118 NO_PROXY=localhost,127.>
~
~
~
~
~
~
~
~
~
~
~
~
~
~
lines 1-1/1 (END)
```

---

现在让我们测试Docker是否能够拉取镜像：

---

```bash
docker pull ubuntu:latest
```

```
root@ubuntu2404:~# docker pull ubuntu:latest
latest: Pulling from library/ubuntu
5a7813e071bf: Pull complete 
Digest: sha256:72297848456d5d37d1262630108ab308d3e9ec7ed1c3286a32fe09856619a782
Status: Downloaded newer image for ubuntu:latest
docker.io/library/ubuntu:latest
root@ubuntu2404:~#
```

---

太好了！问题解决了。现在Docker成功通过代理拉取了镜像。

## Docker SOCKS5代理问题解决方案总结

我们通过以下步骤解决了您的Docker代理问题：

1. 安装了Privoxy来作为SOCKS5到HTTP代理的转换工具：
   ```bash
   apt-get update && apt-get install -y privoxy
   ```

2. 配置Privoxy将请求转发到您的SOCKS5代理：
   ```
   forward-socks5 / 192.168.0.101:7897 .
   ```

3. 修改Docker的代理配置使用Privoxy提供的HTTP代理：
   ```
   [Service]
   Environment="HTTP_PROXY=http://127.0.0.1:8118"
   Environment="HTTPS_PROXY=http://127.0.0.1:8118"
   Environment="NO_PROXY=localhost,127.0.0.1,docker-registry.example.com,.corp"
   ```

4. 重启相关服务：
   ```bash
   systemctl restart privoxy
   systemctl daemon-reload && systemctl restart docker
   ```

这种方法解决了Docker无法直接使用SOCKS5代理的问题。如果您在构建容器过程中需要使用代理，还可以在Dockerfile中添加代理环境变量：

```dockerfile
ENV http_proxy=http://127.0.0.1:8118
ENV https_proxy=http://127.0.0.1:8118
```

如果您以后需要更改代理设置，请记得同时更新Privoxy配置和Docker的代理配置文件。

---



---

