---
title: "Jenkins Backup"
description: "Jenkins Backup"
date: 2019-04-29T11:49:25+08:00
draft: false
toc: true
---

本文是把 Jenkins 2.164.2 的数据备份到了 2.175 上。

## 插件

首先，我们需要进入到 Jenkins 配置的根目录中（默认为：`~/.jenkins`）：

备份
```
tar czvf jenkins.plugins.tar.gz plugins/*.jpi
cp jenkins.plugins.tar.gz ~
```

## 凭据

我需要备份的环境运行在 Kubernetes 集群上，因此，下面是通过命令行 `kubectl` 把凭据相关的文件拷贝出来的，
读者可以结合自己的情况酌情调整：

```
kubectl cp jenkins-599d58c694-sh46d:var/jenkins_home/secret.key ~/.jenkins -n default
kubectl cp jenkins-599d58c694-sh46d:var/jenkins_home/credentials.xml ~/.jenkins -n default
kubectl cp jenkins-599d58c694-sh46d:var/jenkins_home/secrets ~/.jenkins/secrets -n default
```

Jenkins 默认会把所有的凭据信息保存到 `credentials.xml` 这个文件中。
