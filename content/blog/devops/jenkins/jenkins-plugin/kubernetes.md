---
title: Jenkins Kubernetes Plugin
description: Jenkins Kubernetes Plugin
keywords: jenkins plugin kubernetes
draft: false
toc: true
---

[Kubernetes](https://kubernetes.io/) 是一个容器编排的系统，而 [Kubernetes 插件](https://plugins.jenkins.io/kubernetes)则是 Jenkins 用来调度 K8S 以实现动态计算节点的方式。

# 多容器

Pod 里可以配置多个容器，在使用的过程中，要注意根据实际情况进行切换。对应的 DSL 语句如下：


```
container("container_name"){
    echo "hello"
}
```