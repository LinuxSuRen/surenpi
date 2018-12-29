---
title: Jenkins Agent
description: Jenkins Agent
layout: suren
keywords: jenkins agent
toc: true
---

Jenkins 的运行模式，分为 master 和 agent。配置、管理、调度等会由 master 节点来完成，而具体的构建任务则由计算节点（agent）完成。

默认安装的 Jenkins 会包含一个 master 节点。通常，不会在 master 节点上执行构建任务，只用于调度。

Jenkins 的节点，通常可以配置如下属性：

* 名称 *唯一标示*
* 描述
* 并发构建数 *同一时刻可执行的任务数*
* 远程工作目录
* 标签 *空格分割的字符串，通常包含环境信息*
* 启动方式 *agent 节点如何启动，通常与 agent 类型相关*
* 工具位置 *构建工具的安装位置配置*
* 环境变量

# 类型

Jenkins 支持多种与节点通信的方式。

{{< youtube Y9Aa1U3wJrI >}}

## JNLP

默认支持。

## SSH

需要[安装插件]({{< relref "/blog/devops/jenkins/jenkins-plugin/ssh-slaves.md" >}})。

## Docker

这里指的是在 Docker 容器里运行 Jenkins 计算节点。我们可以分为两类：静态、动态。

## 静态 Docker

静态的 Docker 容器，是通过命令来启动一个 `jnlp` 容器。这种方式与 JNLP 其实是相同的机制，都是需要提前在 Jenkins 的节点管理界面添加一个静态节点，然后再启动。

通常的命令 `docker run jenkins/jnlp-slave -url http://jenkins-server:port <secret> <agent name>`

设置 Tunnel 端口 `docker run jenkins/jnlp-slave -tunnel :<tunnel port> -url http://jenkins-server:port <secret> <agent name>`

{{% note %}}
*注意，上面的端口是 Jenkins 的管理界面的端口，secret 则要从 JNLP 节点的状态页面获取，最后需要的是节点的名称。*
{{% /note %}}

想了解源码的朋友，可以查看[这里](https://github.com/jenkinsci/docker-jnlp-slave)。

## 动态 Docker

这种方式需要 Jenkins 所在的服务器上安装 Docker，它会在流水线执行时动态地创建、销毁容器。优点：无需提前启动计算节点，不依赖 Java 环境，不依赖 JNLP；缺点：架构上与 Docker 紧紧地绑定了，无法做到扩缩容。下面是一种直接指定 Docker 镜像的方式，还有另外一种方式支持动态构建出一个镜像：

{{< exfile "static/codes/jenkinsfile/docker-agent.groovy" "groovy" >}}

## Kubernetes

需要[安装插件]({{< relref "/blog/devops/jenkins/jenkins-plugin/kubernetes.md" >}})。

这种动态调度节点的方式，在和使用静态节点的没有区别，都是通过 Jenkins 强大的[标签]({{< ref "jenkins-label.md" >}})机制来调度。而节点的扩缩容则交给了 K8S 集群。Jenkins 的 master 节点可以运行在任意环境里，这一点是没有任何限制。

目前，也可能是较长时间内的一个缺点就是：复杂。不过，对于配置复杂的问题，可以借助[配置及代码]({{< relref "/blog/devops/jenkins/jenkins-plugin/configuration-as-code.md" >}})这个插件来解决。异常诊断的问题，也在积极的解决。
