---
title: Docker Mirror
description: Docker Mirror  
---

`vim /etc/docker/daemon.json`

```
{
  "registry-mirrors": ["https://mirror.gcr.io"]
}
```

或者，获取在 daemon 启动时传递参数 `dockerd --registry-mirror=https://mirror.gcr.io`

或者，在文件 `/etc/default/docker` 中添加如下配置：

`DOCKER_OPTS="${DOCKER_OPTS} --registry-mirror=https://mirror.gcr.io"`