---
title: Install
---

安装 Git `yum install git -y`

拉取 Docker 镜像 `docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/tiller:v2.9.1 && docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/tiller:v2.9.1 gcr.io/kubernetes-helm/tiller:v2.9.1`

获取 jx `curl -L https://github.com/jenkins-x/jx/releases/download/v1.3.8/jx-linux-amd64.tar.gz | tar xzv`

`mv jx /usr/local/bin`

安装 jx `jx install --username suren --provider kubernetes`
