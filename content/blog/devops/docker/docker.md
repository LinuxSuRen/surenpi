---
title: "Docker 命令"
description: "Docker 命令"
toc: true
---

# 拉取镜像

`docker pull jenkins/jnlp-slave:alpine`

`docker pull hub.alauda.cn/jenkins-jnlp-slave`

# 启动

## 自启动

当 `Docker` 服务启动后，默认情况下，所有的容器都不会启动。如果在每次启动时 `Docker` 服务时，自动启动容器，可以通过添加下面的参数实现：

`docker run jenkins --restart=always`

最简单的启动命令 `docker run surenpi/hugo-docker`

容器停止后，自动删除 `docker run --rm surenpi/hugo-docker`

## 设置名称

当你启动 `Docker` 容器后，容器的名称时随机生成的。为了方便后续的查找，我们可以通过参数给容器设置一个名称：

`docker run jenkins --name ci-jenkins`

# 查看容器

我们可以使用命令 `docker ps` ，查看所有运行状态的容器。

如果要查看所有（包括已经停止的）容器话，可以使用参数：`docker ps -a`。

# 停止容器

我们可以使用命令 `docker stop 94a5fa599b05` 来停止容器。这里的参数，是容器的ID。当然，我们也可以把容器的名称作为参数。

# 日志

`docker logs -f 94a5fa599b05`

# 删除容器

我们可以通过命令 `docker rm 94a5fa599b05` 删除一个容器。当然，前提条件是该容器处于已停止的状态。

通过过滤镜像名称批量删除容器 `docker stop $(docker ps | grep hello | awk '{print $1}') && docker rm $(docker ps | grep hello | awk '{print $1}')`

删除所有已经停止的容器 `docker rm $(docker ps -a | grep Exited | awk '{print $1}')`

# 标签

`docker tag source-image target-image`

# 端口映射

`docker -p 127.0.0.1:1234:8080 run jenkins`

# 卷

查看容器所挂载的卷信息 `docker inspect 421b094bf94a | grep Mounts -A 20`

讲本地的目录绑定到容器中 `docker run -v /your/home:/var/jenkins_home jenkins:jenkins:lts`

# 构建

设置构建过程中的代理 `docker build . --build-arg http_proxy=http://192.168.2.9:8123 --build-arg https_proxy=http://192.168.2.9:8123`
