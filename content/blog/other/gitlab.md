---
title: Gitlab
description: Gitlab 使用介绍
---

# Docker

我们可以在 Docker 中运行 Gitlab。首先，要拉取镜像 `docker pull gitlab/gitlab-ce`。

启动容器的命令 `docker run -d -p 1234:80 --rm gitlab/gitlab-ce`

Gitlab 在 DockerHub 上的地址 `https://hub.docker.com/r/gitlab/gitlab-ce/`