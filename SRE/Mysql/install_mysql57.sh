#!/bin/bash

# 安装MySQL 5.7脚本
# 适用于Ubuntu 24.04 LTS

# 确保脚本以root权限运行
if [ "$(id -u)" != "0" ]; then
   echo "此脚本必须以root权限运行" 1>&2
   exit 1
fi

echo "====================================================="
echo "开始安装MySQL 5.7..."
echo "====================================================="

# 更新系统
echo "正在更新系统包..."
apt-get update
apt-get upgrade -y

# 安装依赖
echo "正在安装依赖项..."
apt-get install -y wget gnupg lsb-release apt-transport-https ca-certificates

# 添加MySQL APT仓库
echo "正在添加MySQL APT仓库..."
wget https://dev.mysql.com/get/mysql-apt-config_0.8.22-1_all.deb
# 使用debconf预设MySQL 5.7选择
echo "mysql-apt-config mysql-apt-config/select-server select mysql-5.7" | debconf-set-selections
DEBIAN_FRONTEND=noninteractive dpkg -i mysql-apt-config_0.8.22-1_all.deb

# 更新仓库
echo "正在更新仓库..."
apt-get update

# 预设root密码
echo "正在预设MySQL root密码..."
debconf-set-selections <<< "mysql-community-server mysql-community-server/root-pass password root"
debconf-set-selections <<< "mysql-community-server mysql-community-server/re-root-pass password root"

# 安装MySQL服务器
echo "正在安装MySQL服务器5.7..."
apt-get install -y mysql-server

# 启动MySQL服务
echo "正在启动MySQL服务..."
systemctl start mysql
systemctl enable mysql

# 配置MySQL允许远程连接
echo "正在配置MySQL允许远程连接..."
# 备份原始配置文件
cp /etc/mysql/mysql.conf.d/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf.bak

# 修改bind-address为0.0.0.0以允许远程连接
sed -i 's/bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf

# 创建SQL脚本以允许root用户远程访问
cat > /tmp/mysql_remote_access.sql << EOF
CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF

# 执行SQL脚本
echo "正在创建远程访问用户并授权..."
mysql -u root -proot < /tmp/mysql_remote_access.sql

# 重启MySQL服务以应用更改
echo "正在重启MySQL服务..."
systemctl restart mysql

# 清理临时文件
rm -f mysql-apt-config_0.8.22-1_all.deb
rm -f /tmp/mysql_remote_access.sql

# 显示MySQL状态
echo "====================================================="
echo "MySQL 5.7安装完成！"
echo "MySQL服务状态："
systemctl status mysql --no-pager
echo ""
echo "MySQL版本信息："
mysql --version
echo ""
echo "MySQL监听地址信息："
netstat -tulnp | grep mysql
echo "====================================================="
echo "登录信息："
echo "用户名：root"
echo "密码：root"
echo "已启用远程访问"
echo "=====================================================" 