---
title: "Tiller"
description: "Tiller"
date: 2019-02-14T11:15:54+08:00
draft: true
toc: true
---

第一次使用 helm 时需要进行初始化（如果服务端没有对应的容器启动的话），初始化命令为 `helm init`

当执行 `helm list` 时遇到了下面的报错信息

`Error: configmaps is forbidden: User "system:serviceaccount:kube-system:default" cannot list configmaps in the namespace "kube-system"`

你可以通过添加集群管理员权限来修复这个问题：

`kubectl create clusterrolebinding add-on-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:default`