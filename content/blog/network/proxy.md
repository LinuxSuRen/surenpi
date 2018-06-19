# 代理

# 分类

# ssh

# shadowsocks

Ubuntu `sudo apt-get install shadowsocks`

CentOS `yum install shadowsocks`

`pip install shadowsocks`

`sudo pip install shadowsocks --upgrade`

# 使用

`sudo sslocal -c shadowsocks.json -d start`

# polipo

## 安装

`sudo apt-get install polipo`

`brew install polip`

安装依赖 `yum install  texinfo  -y` 然后下载源码 `git clone https://github.com/jech/polipo.git` 最后编译 `cd polipo && make all && make install`

## 使用

polipo socksParentProxy=127.0.0.1:1080

## 配置终端的代理

```
HTTP_PROXY=http://127.0.0.1:8123
HTTPS_PROXY=http://127.0.0.1:8123
SOCKS_PROXY=socks5://127.0.0.1:1080
no_proxy=*.baidu.com,localhost,127.0.0.1,192.18.*.*,

export HTTP_PROXY HTTPS_PROXY no_proxy
```
