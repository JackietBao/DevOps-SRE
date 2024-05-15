清理缓存

```
sudo sync
sudo sh -c "echo 3 > /proc/sys/vm/drop_caches"
```

移除旧版本go

```
sudo rm -rf /usr/local/go
```

下载新版本go

```
wget https://go.dev/dl/go1.21.10.linux-amd64.tar.gz
```

解压并安装到`/usr/local`目录：

```
sudo tar -C /usr/local -xzf go1.21.10.linux-amd64.tar.gz
```

设置环境变量

```
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
source ~/.bash_profile
```

重新加载环境变量

```
source ~/.bash_profile
```

检查go版本

```
go version
```

编译安装caddy+naive：

```
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
~/go/bin/xcaddy build --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
```

Caddyfile配置：

```
:443, naive.buliang0.tk #你的域名
tls example@example.com #你的邮箱
route {
 forward_proxy {
   basic_auth user pass #用户名和密码
   hide_ip
   hide_via
   probe_resistance
  }
 #支持多用户
 forward_proxy {
   basic_auth user2 pass2 #用户名和密码
   hide_ip
   hide_via
   probe_resistance
  }
 reverse_proxy  https://demo.cloudreve.org  { #伪装网址
   header_up  Host  {upstream_hostport}
   header_up  X-Forwarded-Host  {host}
  }
}
```

开放80,443端口

```
iptables -I INPUT -p tcp --dport 443 -j ACCEPT
iptables -I INPUT -p tcp --dport 80 -j ACCEPT
```

caddy常用指令：

```
前台运行caddy：./caddy run
后台运行caddy：./caddy start
停止caddy：./caddy stop
重载配置：./caddy reload
```



















