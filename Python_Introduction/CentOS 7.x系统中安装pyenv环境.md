### 一、pyenv安装

1、安装git

```bash
yum install git -y
1
```

2、[安装Python](https://so.csdn.net/so/search?q=安装Python&spm=1001.2101.3001.7020)编译依赖

```bash
yum -y install gcc make patch gdbm-devel openssl-devel sqlite-devel readline-devel zlib-devel bzip2-devel
1
```

3、git下载pyenv

pyenv在github上的地址：https://github.com/pyenv/pyenv

默认安装在当前用户$HOME目录下

```bash
git clone https://github.com/pyenv/pyenv.git ~/.pyenv
1
```

4、添加环境变量

```bash
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bash_profile
echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bash_profile
echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.bash_profile
123
```

5、使环境变量生效

```bash
source ~/.bash_profile
1
```

### 二、pyenv使用

1、[查看python版本](https://so.csdn.net/so/search?q=查看python版本&spm=1001.2101.3001.7020)

```bash
python --version
python -V
12
```

2、查看可安装python版本

```bash
pyenv install --list
1
```

3、不同版本python安装

指定版本直接安装

```bash
pyenv install 3.6.4
pyenv versions
12
```

推荐自己下载python安装包，这样比较快http://npm.taobao.org/mirrors/python/，下载需要安装的版本.tar.gz结尾的压缩包

```bash
cd .pyenv/
mkdir cache
cd cache/
123
```

放入下载好的安装包，执行`pyenv install 3.6.4 -v`

4、pyenv的python版本控制

- global 全局设置

  pyenv global 3.6.4 可以看到所有受pyenv控制的窗口中都是3.6.4的python版本了。 这里用global是作用于非root用户python用户上，如果是root用户安装，请不要使用global，否则影响太大。

- shell 会话设置

  影响只作用于当前会话 pyenv shell 3.6.4

- local 本地设置

  使用pyenv local设置从当前工作目录开始向下递归都继承这个设置。 pyenv local 3.6.4

### 三、Virtualenv 虚拟环境设置

pyenv-virtualenv在github上的地址pyenv-virtualenv

1、安装

```bash
git clone https://github.com/pyenv/pyenv-virtualenv.git $(pyenv root)/plugins/pyenv-virtualenv
echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bash_profile
source ~/.bash_profile
123
```

2、使用

```bash
pyenv virtualenv 3.6.4 test363
1
```

使用python3.6.4版本创建出一个独立的虚拟空间

3、其它命令

```bash
pyenv virtualenvs 　　　　　　　 　　　　　　　 # 列出所有虚拟环境
pyenv virtualenv-delete                      # 删除虚拟环境
pyenv activate your_env_name 　　 　　　　　　# 使用某虚拟环境
pyenv deactivate 　　　　　　　　　　　　　　　# 退出虚拟环境，回到系统环境
1234
```

### 四、修改pip 通用配置

```bash
mkdir ~/.pip
vim ~/.pip/pip.conf
[global]
index-url=https://mirrors.aliyun.com/pypi/simple/
trusted-host=mirrors.aliyun.com
```