#!/bin/bash

# MySQL 5.7 编译安装脚本
# 适用于 Ubuntu 24.04 LTS
# 设置登录密码为root并允许远程访问

set -e  # 遇到错误立即退出
set -x  # 显示执行的命令

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}开始安装MySQL 5.7...${NC}"

# 创建工作目录
WORK_DIR="/usr/local/src/mysql_build"
mkdir -p $WORK_DIR
cd $WORK_DIR

# 安装依赖
echo -e "${YELLOW}安装编译依赖...${NC}"
apt-get update
apt-get install -y build-essential cmake libncurses5-dev bison pkg-config libssl-dev libreadline-dev \
                   libsystemd-dev libcurl4-openssl-dev libldap2-dev libnuma-dev libjemalloc-dev \
                   libaio-dev wget git libtirpc-dev

# 下载MySQL 5.7源代码
echo -e "${YELLOW}下载MySQL 5.7源代码...${NC}"
wget https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-boost-5.7.44.tar.gz
tar xzf mysql-boost-5.7.44.tar.gz
cd mysql-5.7.44

# 创建MySQL用户和组
echo -e "${YELLOW}创建MySQL用户和组...${NC}"
groupadd -r mysql || echo "mysql组已存在"
useradd -r -g mysql -s /bin/false mysql || echo "mysql用户已存在"

# 创建MySQL数据目录
mkdir -p /usr/local/mysql/data
chown -R mysql:mysql /usr/local/mysql

# 编译安装MySQL
echo -e "${YELLOW}开始编译MySQL...${NC}"
mkdir -p build
cd build

cmake .. \
    -DCMAKE_INSTALL_PREFIX=/usr/local/mysql \
    -DMYSQL_DATADIR=/usr/local/mysql/data \
    -DSYSCONFDIR=/etc \
    -DWITH_BOOST=../boost/boost_1_59_0 \
    -DWITH_SYSTEMD=1 \
    -DENABLED_LOCAL_INFILE=1 \
    -DWITH_INNOBASE_STORAGE_ENGINE=1 \
    -DWITH_ARCHIVE_STORAGE_ENGINE=1 \
    -DWITH_BLACKHOLE_STORAGE_ENGINE=1 \
    -DWITH_PERFSCHEMA_STORAGE_ENGINE=1 \
    -DMYSQL_TCP_PORT=3306 \
    -DMYSQL_UNIX_ADDR=/tmp/mysql.sock \
    -DDEFAULT_CHARSET=utf8mb4 \
    -DDEFAULT_COLLATION=utf8mb4_general_ci

# 编译和安装
echo -e "${YELLOW}编译MySQL，这将花费一段时间...${NC}"
make -j$(nproc)

echo -e "${YELLOW}安装MySQL...${NC}"
make install

# 设置环境变量
echo -e "${YELLOW}设置环境变量...${NC}"
echo 'export PATH=$PATH:/usr/local/mysql/bin' > /etc/profile.d/mysql.sh
source /etc/profile.d/mysql.sh

# 创建配置文件
echo -e "${YELLOW}创建MySQL配置文件...${NC}"
cat > /etc/my.cnf << EOF
[mysqld]
basedir = /usr/local/mysql
datadir = /usr/local/mysql/data
socket = /tmp/mysql.sock
bind-address = 0.0.0.0
port = 3306
user = mysql
character-set-server = utf8mb4
collation-server = utf8mb4_general_ci
default-storage-engine = InnoDB
max_allowed_packet = 16M
max_connections = 1000
explicit_defaults_for_timestamp = 1

[client]
socket = /tmp/mysql.sock
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4
EOF

# 设置权限
chown -R mysql:mysql /usr/local/mysql
chown -R mysql:mysql /etc/my.cnf

# 初始化数据库
echo -e "${YELLOW}初始化MySQL数据库...${NC}"
cd /usr/local/mysql
bin/mysqld --initialize-insecure --user=mysql

# 创建systemd服务文件
echo -e "${YELLOW}创建systemd服务文件...${NC}"
cat > /etc/systemd/system/mysql.service << EOF
[Unit]
Description=MySQL Server
After=network.target

[Service]
Type=forking
User=mysql
Group=mysql
PIDFile=/usr/local/mysql/data/mysql.pid
ExecStart=/usr/local/mysql/bin/mysqld --daemonize --pid-file=/usr/local/mysql/data/mysql.pid
ExecReload=/bin/kill -s HUP \$MAINPID
ExecStop=/bin/kill -s QUIT \$MAINPID
Restart=on-failure
RestartPreventExitStatus=1
PrivateTmp=false

[Install]
WantedBy=multi-user.target
EOF

# 重载systemd配置
systemctl daemon-reload

# 启动MySQL服务
echo -e "${YELLOW}启动MySQL服务...${NC}"
systemctl start mysql
systemctl enable mysql

# 等待MySQL服务启动
echo -e "${YELLOW}等待MySQL服务启动...${NC}"
sleep 10

# 设置root密码为"root"并允许远程访问
echo -e "${YELLOW}设置root密码并允许远程访问...${NC}"
/usr/local/mysql/bin/mysql -u root --connect-expired-password << EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY 'root';
CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF

# 配置防火墙(如果启用)
echo -e "${YELLOW}配置防火墙...${NC}"
if command -v ufw > /dev/null; then
    ufw allow 3306/tcp
    echo -e "${GREEN}防火墙已配置，允许3306端口访问${NC}"
fi

echo -e "${GREEN}MySQL 5.7 安装完成！${NC}"
echo -e "${GREEN}数据库用户: root${NC}"
echo -e "${GREEN}数据库密码: root${NC}"
echo -e "${GREEN}MySQL服务已配置允许远程访问${NC}"
echo -e "${GREEN}您可以使用以下命令连接到MySQL:${NC}"
echo -e "${YELLOW}mysql -u root -p${NC}"

# 显示MySQL版本信息
/usr/local/mysql/bin/mysql --version

echo -e "${GREEN}安装过程完成！${NC}" 