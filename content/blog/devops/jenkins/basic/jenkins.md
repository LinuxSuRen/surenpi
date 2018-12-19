---
title: Jenkins
description: Jenkins 是一个开源的 CI 服务
keywords: jenkins
---

本文介绍 [Jenkins](https://jenkins.io/zh) 相关的内容，包括：如何启动、配置等等。

# 启动

`java -jar jenkins.war` 这样启动后，就可以通过浏览器访问 `http://localhost:8080/jenkins`

`java -jar jenkins.war --httpPort=1234` 设置 Jenkins 的监听端口。这样 Jenkins 会监听 0.0.0.0 上的端口，也就是外部网络也可以访问。
