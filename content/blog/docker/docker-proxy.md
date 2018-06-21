---
title: 配置
description: 配置
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

```
[Service]
ExecStart=/usr/bin/docker daemon -H fd:// --insecure-registry 10.101.101.45:5000
```

`vim /etc/docker/daemon.json`

`"insecure-registries" : ["192.168.0.180"]`

# User

让普通用户可以执行 docker

`sudo groupadd docker`

`sudo gpasswd -a ${USER} docker`

`sudo service docker restart`

# 重启服务

执行命令 `systemctl daemon-reload && systemctl stop docker && systemctl start docker`
