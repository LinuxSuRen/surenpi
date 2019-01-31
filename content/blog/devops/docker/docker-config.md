---
title: 配置
description: 配置
toc: true
---

# 代理
给 Docker 设置代理

创建目录 `mkdir /etc/systemd/system/docker.service.d`

编辑配置文件 `vim /etc/systemd/system/docker.service.d/http-proxy.conf`

```
[Service]
Environment="HTTP_PROXY=http://ip:port"
```

# Insecure

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

# User

让普通用户可以执行 docker

`sudo groupadd docker`

`sudo gpasswd -a ${USER} docker`

`sudo service docker restart`

# 重启服务

执行命令 `systemctl daemon-reload && systemctl stop docker && systemctl start docker`
