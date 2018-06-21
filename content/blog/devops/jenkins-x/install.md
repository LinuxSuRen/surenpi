---
title: Install
---

# 预备

Jenkins X 是基于容器化来实现的 CI/CD 平台，因此，我们首先需要[搭建 Kubernetes 环境](/blog/devops/k8s/install/)。

# 安装

安装 Git `yum install git -y`

拉取 Docker 镜像 

```
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/tiller:v2.9.1
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/tiller:v2.9.1 gcr.io/kubernetes-helm/tiller:v2.9.1

docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/monocular-api:v0.6.1
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/monocular-api:v0.6.1 bitnami/monocular-api:v0.6.1

docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/addon-resizer:1.7
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/addon-resizer:1.7 k8s.gcr.io/addon-resizer:1.7

docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/heapster:v1.3.0
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/heapster:v1.3.0 k8s.gcr.io/heapster:v1.3.0

docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/jenkinsx:0.0.25
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/jenkinsx:0.0.25 jenkinsxio/jenkinsx:0.0.25

docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/mongodb:3.4.9-r1
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/mongodb:3.4.9-r1 docker.io/bitnami/mongodb:3.4.9-r1

docker pull registry.cn-beijing.aliyuncs.com/surenpi/nexus:0.0.14
docker tag registry.cn-beijing.aliyuncs.com/surenpi/nexus:0.0.14 docker.io/jenkinsxio/nexus

docker pull registry.cn-beijing.aliyuncs.com/suren-jenkinsci/jnlp-slave:3.14-1
docker tag registry.cn-beijing.aliyuncs.com/suren-jenkinsci/jnlp-slave:3.14-1 jenkinsci/jnlp-slave:3.14-1
```

获取 jx `curl -L https://github.com/jenkins-x/jx/releases/download/v1.3.8/jx-linux-amd64.tar.gz | tar xzv`

移动到系统目录中 `mv jx /usr/local/bin`

安装 jx `jx install --username suren --provider=kubernetes  --on-premise`
