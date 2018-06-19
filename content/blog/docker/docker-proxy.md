---
title: Proxy
---

给 Docker 设置代理

创建目录 `mkdir /etc/systemd/system/docker.service.d`


编辑配置文件 `vim /etc/systemd/system/docker.service.d/http-proxy.conf`

```
[Service]

Environment="HTTP_PROXY=http://ip:port"
```

重启服务 `systemctl daemon-reload && systemctl stop docker && systemctl start docker`
