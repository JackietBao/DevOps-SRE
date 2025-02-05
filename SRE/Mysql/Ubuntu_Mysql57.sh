#!/bin/bash
# 此脚本将在 Ubuntu 24.04 LTS (GNU/Linux 6.8.0-31-generic x86_64) 上从源码编译安装 MySQL 5.7
# 并将 MySQL root 用户的登录密码设置为 "root"。

set -e  # 出现错误则退出

# ===============================
# 1. 定义变量
# ===============================
MYSQL_VERSION="5.7.37"
SRC_TARBALL="mysql-${MYSQL_VERSION}.tar.gz"
DOWNLOAD_URL="https://dev.mysql.com/get/Downloads/MySQL-5.7/${SRC_TARBALL}"
INSTALL_PREFIX="/usr/local/mysql5.7"

# ===============================
# 2. 安装编译所需依赖
# ===============================
sudo apt update
sudo apt install -y build-essential cmake libncurses-dev bison \
    libssl-dev libaio-dev wget pkg-config \
    zlib1g-dev libzstd-dev liblz4-dev libevent-dev \
    libprotobuf-dev protobuf-compiler libtirpc-dev

# 下载并编译 OpenSSL 1.1.1
OPENSSL_VERSION="1.1.1w"
OPENSSL_FILE="openssl-${OPENSSL_VERSION}.tar.gz"
OPENSSL_DIR="${PWD}/openssl-${OPENSSL_VERSION}"
OPENSSL_INSTALL_DIR="${PWD}/openssl"

if [ ! -f "${OPENSSL_FILE}" ]; then
    echo "下载 OpenSSL ${OPENSSL_VERSION}..."
    wget https://www.openssl.org/source/${OPENSSL_FILE}
fi

if [ ! -d "${OPENSSL_INSTALL_DIR}" ]; then
    echo "编译安装 OpenSSL..."
    tar xzf ${OPENSSL_FILE}
    cd ${OPENSSL_DIR}
    ./config --prefix=${OPENSSL_INSTALL_DIR} shared
    make -j$(nproc)
    make install
    cd ..
fi

# 获取 OpenSSL 安装路径
OPENSSL_ROOT_DIR="/usr"
OPENSSL_LIBRARIES="/usr/lib/x86_64-linux-gnu"

# 下载并解压正确版本的 Boost
BOOST_VERSION="1.59.0"
BOOST_FILE="boost_1_59_0.tar.gz"
BOOST_DIR="boost_1_59_0"

if [ ! -f "${BOOST_FILE}" ]; then
    echo "下载 Boost ${BOOST_VERSION}..."
    wget https://sourceforge.net/projects/boost/files/boost/${BOOST_VERSION}/${BOOST_FILE}
fi

# 解压 Boost
if [ ! -d "${BOOST_DIR}" ]; then
    echo "解压 Boost..."
    tar xzf ${BOOST_FILE}
fi

# ===============================
# 3. 下载 MySQL 源码包（若不存在则下载）
# ===============================
if [ ! -f "${SRC_TARBALL}" ]; then
    echo "下载 MySQL ${MYSQL_VERSION} 源码包..."
    wget ${DOWNLOAD_URL}
else
    echo "源码包 ${SRC_TARBALL} 已存在，跳过下载。"
fi

# ===============================
# 4. 解压源码包并进入源码目录
# ===============================
tar -zxvf ${SRC_TARBALL}
cd mysql-${MYSQL_VERSION}

# ===============================
# 5. 创建构建目录并配置编译参数
# ===============================
# 如果 build 目录已存在，先删除它
rm -rf build
mkdir build && cd build
echo "运行 cmake 配置..."

# 设置 OpenSSL 相关环境变量
export OPENSSL_ROOT_DIR=${OPENSSL_INSTALL_DIR}
export OPENSSL_LIBRARIES=${OPENSSL_INSTALL_DIR}/lib

cmake .. \
  -DCMAKE_INSTALL_PREFIX=${INSTALL_PREFIX} \
  -DDEFAULT_CHARSET=utf8 \
  -DDEFAULT_COLLATION=utf8_general_ci \
  -DWITH_SSL=${OPENSSL_INSTALL_DIR} \
  -DWITH_ZLIB=system \
  -DDOWNLOAD_BOOST=1 \
  -DWITH_BOOST=../boost_1_59_0 \
  -DCMAKE_C_FLAGS="-I/usr/include/tirpc" \
  -DCMAKE_CXX_FLAGS="-I/usr/include/tirpc"

# ===============================
# 6. 编译并安装
# ===============================
echo "开始编译（使用所有CPU核心）..."
make -j$(nproc)
echo "安装 MySQL 到 ${INSTALL_PREFIX} ..."
sudo make install

# ===============================
# 7. 添加 MySQL 用户及相关目录
# ===============================
# 若 mysql 用户组不存在则创建
if ! getent group mysql >/dev/null; then
  sudo groupadd mysql
fi
# 若 mysql 用户不存在则创建（禁止登录）
if ! id -u mysql >/dev/null 2>&1; then
  sudo useradd -r -g mysql -s /bin/false mysql
fi

# 创建数据目录
sudo mkdir -p ${INSTALL_PREFIX}/data
# 修改安装目录所有权
sudo chown -R mysql:mysql ${INSTALL_PREFIX}

# ===============================
# 8. 初始化数据库
# ===============================
echo "初始化数据库..."
sudo ${INSTALL_PREFIX}/bin/mysqld --initialize-insecure --user=mysql \
    --basedir=${INSTALL_PREFIX} --datadir=${INSTALL_PREFIX}/data

# ===============================
# 9. 安装 MySQL 服务脚本（可选）
# ===============================
echo "安装服务脚本..."
sudo cp ${INSTALL_PREFIX}/support-files/mysql.server /etc/init.d/mysql5.7
sudo chmod +x /etc/init.d/mysql5.7

# ===============================
# 10. 启动 MySQL 服务
# ===============================
echo "启动 MySQL 服务..."
# 使用 mysqld_safe 启动，后台运行
sudo ${INSTALL_PREFIX}/bin/mysqld_safe --user=mysql &
# 等待 10 秒让服务启动完成（可根据实际情况调整等待时间）
sleep 10

# ===============================
# 11. 设置 root 用户密码为 "root"
# ===============================
echo "设置 root 用户密码为 'root'..."
sudo ${INSTALL_PREFIX}/bin/mysqladmin -u root password "root"

echo "MySQL ${MYSQL_VERSION} 编译安装完成！"
echo "安装目录：${INSTALL_PREFIX}"
echo "数据库数据目录：${INSTALL_PREFIX}/data"
echo "请使用用户名 root 和密码 'root' 登录。"
