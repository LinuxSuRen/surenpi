---
title: Docker 镜像管理
description: Docker 镜像管理
keywords:
- docker
toc: true
---

使用 `Docker` 一段时间后，就会发现 `Docker` 镜像会占用很多的存储空间，我们可以通过 `prune` 命令来清理掉不再使用的镜像 `docker image prune`

# 修改

Docker 镜像可以和 Git 一样，进行修改，然后重新推送回去。

首先，运行一个镜像 `docker run -it demo:v1.0 bash`

然后，通过名称等方式获取对应容器的 ID `docker ps | grep demo:v1.0`

在之前的容器中，做一些修改，例如：添加一个文件 `touch demo`

最后，把修改提交 `docker commit -m 'fix something' f6591bcdce2c demo:v1.0`。如果有需要的话，可以再把修改后的镜像推送回去 `docker push demo:v1.0`
