---
title: Jenkins Kubernetes Plugin
description: Jenkins Kubernetes Plugin
keywords: jenkins plugin kubernetes
draft: false
tags:
- k8s
- plugin
- docker
- agent
toc: true
---

{{< exref "Kubernetes" "https://kubernetes.io/" >}} 是一个容器编排的系统，而 {{< exref "Kubernetes 插件" "https://plugins.jenkins.io/kubernetes" >}}则是 Jenkins 用来调度 K8S 以实现动态计算节点的方式。对{{< ghref "插件源码" "jenkinsci/kubernetes-plugin" >}}感兴趣的童鞋可以了解下。

# 预备知识

要能够顺利地使用 Kubernetes 的 Jenkins 插件，必须对相关的一些基本知识所有了解。

Docker 中的基本概念，例如：镜像、容器、卷（Volume）。熟悉如何拉取镜像，启动、停止、查看容器。

Kubernetes 中的基本概念，例如：命名空间（Namespace）、Pod、PodTemplate、Secret 等。熟悉命令行工具 kubectl 的一些基本操作，查看 Pod 列表、事件、日志等。

JNLP 的基本原理。

# 基本原理

当[标签规则]({{< relref "/blog/devops/jenkins/basic/jenkins-label.md" >}})匹配到某个 Pod 模板时，插件会尝试根据该模板的配置启动一个 Pod。而该 Pod 在启动过程中，会通过 JNLP 协议来连接 master 节点。

为了能够让 Pod 连接到 Jenkins 的 master 节点，我们首先需要在“全局安全配置”中开启 JNLP 代理协议，并记住端口。

然后，在“系统配置”页面中的 Kubernetes 区域配置正确的 Jenkins 以及 JNLP 的地址以及端口。

# 认证

有多种方式可以配置 Jenkins 与 Kubernetes 集群集成认证。最简单的就是添加凭据时选择 `Secret File` 类型，并把文件 `~/.kube/config` 上传即可。这里有[如何导出 k8s 集群认证文件]({{< relref "blog/devops/k8s/kubectl.md" >}})的介绍。

导出 Kubernetes 认证文件的命令：

`curl https://raw.githubusercontent.com/LinuxSuRen/surenpi/master/static/codes/shell/k8s/export-k8s-crt.sh|bash`

# Pod 模板

一个 Pod 对应 Jenkins 中的一个执行节点，

# K8S 代理节点

该插件在流水线中以支持多种方式使用 K8S 代理节点。通常情况下，我们需要在 Pod 模板中添加一个容器模板，也就是所需要的运行环境，例如：Java、Golang 等语言相关的。在运行期，Kubernetes 插件会自动添加一个 JNLP 的容器。

对于希望使用自定义 JNLP 镜像的用户来说，可以多添加一个 JNLP 镜像就可以覆盖了。

# 多容器

Pod 里如果包含多个容器的话，在使用的过程中，要注意根据实际情况进行切换。对应的 DSL 语句如下：

```
container("container_name"){
    echo "hello"
}
```

# 排错