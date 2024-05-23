v2rayN 下载(6.23)：[https://github.com/2dust/v2rayN/releases/latest](https://bulianglin.com/g/aHR0cHM6Ly9naXRodWIuY29tLzJkdXN0L3YycmF5Ti9yZWxlYXNlcy9sYXRlc3Q)
Hysteria 2下载(2.0.1)：[https://github.com/apernet/hysteria/releases](https://bulianglin.com/g/aHR0cHM6Ly9naXRodWIuY29tL2FwZXJuZXQvaHlzdGVyaWEvcmVsZWFzZXM)
Hysteria 2文档：[https://v2.hysteria.network/zh/](https://bulianglin.com/g/aHR0cHM6Ly92Mi5oeXN0ZXJpYS5uZXR3b3JrL3poLw)

sing-box文档：[https://sing-box.sagernet.org/zh/](https://bulianglin.com/g/aHR0cHM6Ly9zaW5nLWJveC5zYWdlcm5ldC5vcmcvemgv)
Android客户端（SFA）：[https://install.appcenter.ms/users/nekohasekai/apps/sfa/distribution_groups/publictest](https://bulianglin.com/g/aHR0cHM6Ly9pbnN0YWxsLmFwcGNlbnRlci5tcy91c2Vycy9uZWtvaGFzZWthaS9hcHBzL3NmYS9kaXN0cmlidXRpb25fZ3JvdXBzL3B1YmxpY3Rlc3Q)
IOS客户端（TestFlight）：[https://testflight.apple.com/join/AcqO44FH](https://bulianglin.com/g/aHR0cHM6Ly90ZXN0ZmxpZ2h0LmFwcGxlLmNvbS9qb2luL0FjcU80NEZI) （1.5.0 beta版支持Hysteria 2）
IOS客户端（AppStore）：[https://apps.apple.com/us/app/sing-box/id6451272673](https://bulianglin.com/g/aHR0cHM6Ly9hcHBzLmFwcGxlLmNvbS91cy9hcHAvc2luZy1ib3gvaWQ2NDUxMjcyNjcz) （暂不支持Hysteria 2）

## 服务器相关指令

```bash
#一键安装Hysteria2
bash <(curl -fsSL https://get.hy2.sh/)

#生成自签证书
openssl req -x509 -nodes -newkey ec:<(openssl ecparam -name prime256v1) -keyout /etc/hysteria/server.key -out /etc/hysteria/server.crt -subj "/CN=bing.com" -days 36500 && sudo chown hysteria /etc/hysteria/server.key && sudo chown hysteria /etc/hysteria/server.crt

#启动Hysteria2
systemctl start hysteria-server.service
#重启Hysteria2
systemctl restart hysteria-server.service
#查看Hysteria2状态
systemctl status hysteria-server.service
#停止Hysteria2
systemctl stop hysteria-server.service
#设置开机自启
systemctl enable hysteria-server.service
#查看日志
journalctl -u hysteria-server.service
```

## 服务器配置文件

```shell
cat << EOF > /etc/hysteria/config.yaml
listen: :443 #监听端口

#使用CA证书
#acme:
#  domains:
#    - a.com #你的域名，需要先解析到服务器ip
#  email: test@sharklasers.com

#使用自签证书
#tls:
  cert: /etc/hysteria/server.crt
  key: /etc/hysteria/server.key

auth:
  type: password
  password: 123456 #设置认证密码
  
masquerade:
  type: proxy
  proxy:
    url: https://bing.com #伪装网址
    rewriteHost: true
EOF
```

## 客户端配置文件

```yaml
server: 192.253.238.52:443
auth: 123456

bandwidth:
  up: 20 mbps
  down: 100 mbps
  
tls:
  sni: bing.com
  insecure: true #使用自签时需要改成true

socks5:
  listen: 127.0.0.1:1080
http:
  listen: 127.0.0.1:8080
```

## sing-box配置文件(Android/IOS)

```json
{
  "dns": {
    "servers": [
      {
        "tag": "cf",
        "address": "https://1.1.1.1/dns-query"
      },
      {
        "tag": "local",
        "address": "223.5.5.5",
        "detour": "direct"
      },
      {
        "tag": "block",
        "address": "rcode://success"
      }
    ],
    "rules": [
      {
        "geosite": "category-ads-all",
        "server": "block",
        "disable_cache": true
      },
      {
        "outbound": "any",
        "server": "local"
      },
      {
        "geosite": "cn",
        "server": "local"
      }
    ],
    "strategy": "ipv4_only"
  },
  "inbounds": [
    {
      "type": "tun",
      "inet4_address": "172.19.0.1/30",
      "auto_route": true,
      "strict_route": false,
      "sniff": true
    }
  ],
  "outbounds": [
    {
      "type": "hysteria2",
      "tag": "proxy",
      "server": "192.253.238.52",
      "server_port": 443,
      "up_mbps": 20,
      "down_mbps": 90,
      "password": "123456",
      "tls": {
        "enabled": true,
        "server_name": "bing.com",
        "insecure": true
      }
    },
    {
      "type": "direct",
      "tag": "direct"
    },
    {
      "type": "block",
      "tag": "block"
    },
    {
      "type": "dns",
      "tag": "dns-out"
    }
  ],
  "route": {
    "rules": [
      {
        "protocol": "dns",
        "outbound": "dns-out"
      },
      {
        "geosite": "cn",
        "geoip": [
          "private",
          "cn"
        ],
        "outbound": "direct"
      },
      {
        "geosite": "category-ads-all",
        "outbound": "block"
      }
    ],
    "auto_detect_interface": true
  }
}
```

## 网速优化

安装并配置NTP同步时间，以确保系统时间的准确性，有助于优化网络性能：

```
sudo yum install -y ntp
sudo systemctl enable ntpd
sudo systemctl start ntpd
```

编辑`/etc/resolv.conf`文件，使用快速且可靠的DNS服务器，例如Google DNS或Cloudflare DNS：

```
nameserver 8.8.8.8
nameserver 1.1.1.1
```

Google BBR（Bottleneck Bandwidth and Round-trip propagation time）是一种拥塞控制算法，旨在提高网络传输效率和速度。它通过测量网络瓶颈带宽和往返时间（RTT）来确定发送数据的速率，不依赖于传统的丢包信号。

要在CentOS 7上启用BBR算法，请按以下步骤操作：

检查当前内核版本

BBR需要Linux内核版本4.9或更高版本。可以通过以下命令检查当前内核版本：

```
bash
复制代码
uname -r
```

如果内核版本低于4.9，需要升级内核。可以参考官方文档或使用ELRepo仓库进行内核升级。

### 升级内核

可以通过ELRepo仓库来升级CentOS 7的内核。以下是具体步骤：

#### 1. 安装ELRepo仓库

```
sudo rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
sudo yum install -y https://www.elrepo.org/elrepo-release-7.0-4.el7.elrepo.noarch.rpm
```

#### 2. 安装最新的主线稳定内核

```
sudo yum --enablerepo=elrepo-kernel install -y kernel-ml
```

#### 3. 更新GRUB配置

```
sudo grub2-set-default 0
sudo grub2-mkconfig -o /boot/grub2/grub.cfg
```

#### 4. 重启服务器

```
sudo reboot
```

#### 5. 验证内核版本

重启后，登录到服务器并验证新的内核版本：

```
uname -r
```

如果内核版本已经更新到4.9或更高版本，接下来可以启用BBR。

### 启用BBR

按以下步骤操作：

1. **编辑sysctl配置文件** 打开`/etc/sysctl.conf`文件，添加以下行：

   ```
   net.core.default_qdisc=fq
   net.ipv4.tcp_congestion_control=bbr
   ```

2. **应用更改** 运行以下命令以应用更改：

   ```
   sudo sysctl -p
   ```
   
3. **验证BBR是否启用** 使用以下命令验证BBR是否已启用：

   ```
   sysctl net.ipv4.tcp_congestion_control
   ```
   
   预期输出应该是：

   ```
   net.ipv4.tcp_congestion_control = bbr
   ```
   
   然后使用以下命令检查BBR是否正在运行：
   
   ```
   lsmod | grep bbr
   ```

   如果BBR已启用并运行，你将看到类似以下的输出：
   
   ```
   tcp_bbr                20480  12
   ```

### Hysteria配置文件优化

```
tls:
  cert: /etc/hysteria/server.crt
  key: /etc/hysteria/server.key

auth:
  type: password
  password: 123456
  
masquerade:
  type: proxy
  proxy:
    url: https://bing.com
    rewriteHost: true

# 优化Hysteria配置
server:
  listen: :443
  protocol: udp
  obfs:
    type: basic
    password: "your_obfs_password"  # 替换为你自己的混淆密码
  udp:
    timeout: 30s
  congestion:
    enable: true
    minRTT: 10ms
    lossThreshold: 0.01
    gain: 1.25
    maxGain: 2.5
  conn:
    maxConnections: 10

# 调整TCP/IP参数
sysctl:
  net.core.rmem_max: 16777216
  net.core.wmem_max: 16777216
  net.ipv4.tcp_rmem: 4096 87380 16777216
  net.ipv4.tcp_wmem: 4096 65536 16777216
  net.core.somaxconn: 4096
  net.ipv4.tcp_tw_reuse: 1
  net.core.default_qdisc: fq
  net.ipv4.tcp_congestion_control: bbr
  net.core.netdev_max_backlog: 5000
```

### 应用和验证

1. **确保系统参数生效** 确保你之前设置的TCP/IP参数已经生效：

   ```
   sudo sysctl -p
   ```
   
2. **重启Hysteria服务** 重新启动Hysteria服务以应用新的配置：

   ```
   sudo systemctl restart hysteria-server
   ```

## udp优化

vi /etc/sysctl.conf

```
net.core.rmem_max=26214400
net.core.wmem_max=26214400
net.ipv4.udp_mem=26214400 26214400 26214400
net.ipv4.udp_rmem_min=26214400
net.ipv4.udp_wmem_min=26214400
```

