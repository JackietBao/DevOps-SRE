v2rayN 下载：[https://github.com/2dust/v2rayN/releases/latest](https://bulianglin.com/g/aHR0cHM6Ly9naXRodWIuY29tLzJkdXN0L3YycmF5Ti9yZWxlYXNlcy9sYXRlc3Q)
Hysteria 2下载：[https://github.com/apernet/hysteria/releases](https://bulianglin.com/g/aHR0cHM6Ly9naXRodWIuY29tL2FwZXJuZXQvaHlzdGVyaWEvcmVsZWFzZXM)
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