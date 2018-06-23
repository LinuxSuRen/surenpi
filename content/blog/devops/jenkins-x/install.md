---
title: Install
description: Jenkins X 部署教程
---

# 预备

Jenkins X 是基于容器化来实现的 CI/CD 平台，因此，我们首先需要[搭建 Kubernetes 环境](/blog/devops/k8s/install/)。

# 安装依赖

安装 Git `yum install git -y`
安装 Maven `yum install maven -y`

`sudo apt-get install maven make -y`

为了方便，我们把当前用户加入 Docker 用户组 `sudo gpasswd -a ${USER} docker`

重启 Docker 服务 `sudo service docker restart`

添加权限 `sudo chmod a+rw /var/run/docker.sock`

# 拉取镜像 

执行命令 `curl https://raw.githubusercontent.com/LinuxSuRen/getGo/master/jx-images.sh | bash` 拉取 jx 所需要的镜像

## 安装 jx

获取 jx `curl -LO https://github.com/jenkins-x/jx/releases/download/$(curl -s https://api.github.com/repos/jenkins-x/jx/releases/latest | grep tag_name | cut -d '"' -f 4)/jx-linux-amd64.tar.gz | tar xzv`


移动到系统目录中 `mv jx /usr/local/bin`

安装 jx `jx install --username suren --provider=kubernetes  --on-premise`

# 连接远程

如果我们想要从本地连接远程的 Kubernetes 集群的话，需要[导出用户证书](/blog/devops/k8s/kuberctl/)。
