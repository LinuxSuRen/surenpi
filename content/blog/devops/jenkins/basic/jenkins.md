---
title: Jenkins
description: Jenkins 是一个开源的 CI 服务
keywords: jenkins
author: linuxsuren
toc: true
categories:
- CICD
- Jenkins
tags:
- Jenkins
---

本文介绍 [Jenkins](https://jenkins.io/zh) 相关的内容，包括：如何启动、配置等等。

# 部署

Jenkins 支持几种安装（启动）方式：

{{< youtube EtlhC3vw8pA >}}

## War

由于 Jenkins 是用 Java 语言来开发的，自然就会有 War 的打包方式。你需要首先在[下载页面](https://jenkins.io/zh/download/)，选择一个版本。

然后，执行命令 `java -jar jenkins.war` 启动后，就可以通过浏览器访问 `http://localhost:8080/jenkins`

这种方式，方便简单，但是需要有 Java 环境。

## Docker

用容器的方式来运行 Jenkins 也很方便。只要执行命令 `docker run -p 8080:8080 jenkins/jenkins:lts` 即可。

## Kubernetes

## Helm

执行命令 `helm install stable/jenkins --name my-jenkins`

如果希望看到具体内容[请查看这里](https://github.com/helm/charts/tree/master/stable/jenkins)。

[KubeApps](https://hub.kubeapps.com/charts?q=Jenkins) 上也可以找到 Jenkins

# 版本选择

从 Jenkins 官网的下载页面，看到有很多版本可以下载，但是，具体该怎么选择呢。我下面给出一些建议：

* 生产环境：最近一年内的 LTS（长期支持版）
* 相对正式的环境：最新的 LTS
* 体验：最新的 Weekly
* [插件研发](../../plugin-dev/)：尽可能选择较老的 LTS

# 参数

Jenkins 在启动过程中，大多数的参数是具有默认值的。但是，对于一些特定的场景下，我们需要设置对应的参数值。

## 端口

`java -jar jenkins.war --httpPort=1234` 设置 Jenkins 的监听端口。这样 Jenkins 会监听 0.0.0.0 上的端口，也就是外部网络也可以访问。

## JENKINS_HOME

我们可以在启动 Jenkins 时指定家目录，默认值为用户目录下的 `.jenkins`。参数（`-D` 参数必须要在 `-jar` 之前）如下：

`java -DJENKINS_HOME=. -jar jenkins.war` 或者 `java -Duser.home=. -jar jenkins.war`

## 时区

如果以 Docker 容器的方式来启动 Jenkins 的话，它默认的是 UTC 时区。 我们可能就需要设置为我们所在的当地时区。

`docker run -v /etc/localtime:/etc/localtime -p 8080:8080 jenkins/jenkins:lts`

## 向导

Jenkins 在首次运行时，会有一个配置向导，可以通过下面的参数跳过这个向导：

`java -Djenkins.install.runSetupWizard=false -jar jenkins.war`

## 用户

`--argumentsRealm.passwd.admin=admin --argumentsRealm.roles.admin=admin`

# 重启

如果，我们需要重新启动 Jenkins，没有必要把进程杀掉后再次启动。可以直接在 Jenkins 界面上就实现重启。

访问地址 `http://localhost:8080/restart` 后点击确认。

# 公钥

访问下面的地址可以得到 Jenkins 的公钥：

`http://localhost:8080/instance-identity/`
