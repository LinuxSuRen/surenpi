---
title: 配置
description: 配置
toc: true
---

## 代理
给 Docker 设置代理，该代理是在拉取镜像时的网络代理设置

```
mkdir -p /etc/systemd/system/docker.service.d
cat > /etc/systemd/system/docker.service.d/http-proxy.conf << EOF
[Service]
Environment="HTTP_PROXY=http://your.proxy:8080"
Environment="HTTPS_PROXY=http://your.proxy:8080"
Environment="NO_PROXY=127.0.0.1,localhost"
EOF
systemctl daemon-reload
systemctl restart docker
```

## 运行时代理

```
{
        "HttpHeaders": {
                "User-Agent": "Docker-Client/19.03.4 (linux)"
        },
        "proxies": {
                "default": {
                        "httpProxy": "http://172.17.212.62:3128",
                        "httpsProxy": "http://172.17.212.62:3128",
                        "noProxy": "127.0.0.1,localhost,172.17.212.63"
                }
        }
}
```

## Insecure

`vim /etc/systemd/system/docker.service.d/docker.conf`

```
[Service]
ExecStart=
ExecStart=/usr/bin/dockerd -H fd:// --insecure-registry 10.101.101.45:5000
```

`vim /etc/docker/daemon.json`

`"insecure-registries" : ["192.168.0.180"]`

Ubuntu 下修改配置文件 `vim /etc/default/docker` 添加如下内容：

`DOCKER_OPTS="$DOCKER_OPTS --insecure-registry registry_server_name/ip:port"`

## User

让普通用户可以执行 docker

`sudo groupadd docker`

`sudo gpasswd -a ${USER} docker`

`sudo service docker restart`

## 并发下载（推送）镜像

下面的示例中给出了如何设置（`/etc/docker/daemon.json`）并发下载或者推送镜像层的配置：
```
{
	"max-concurrent-uploads": 20,
	"max-concurrent-downloads": 20,
	"registry-mirrors": [
		"https://qtzsrp4m.mirror.aliyuncs.com",
		"https://hub-mirror.c.163.com",
		"https://mirror.baidubce.com"
	]
}
```

## 重启服务

执行命令 `systemctl daemon-reload && systemctl stop docker && systemctl start docker`
