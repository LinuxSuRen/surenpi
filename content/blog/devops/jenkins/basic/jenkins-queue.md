---
title: "Jenkins 队列"
description: "队列在 Jenkins 中是非常重要的环节"
date: 2019-06-05T12:52:11+08:00
draft: false
toc: true
---

我们知道，在 Jenkins 中执行任务的话，执行过程是：
1. 进入队列
2. 查找是否有可以执行该任务的节点
3. 启动节点执行任务

这里队列的默认行为是：先进先出。这对于很多常规场景下是正常合理的，毕竟首先触发的任务应该先执行。

但是，在实际的企业（或组织）应用中，往往会有更加复杂的场景需求。例如：部分任务的优先级相对比较高、持久化队列等等。有一些需求，已经有相对应的插件可以实现，本文会对默写场景进行描述。

https://github.com/jenkinsci/purge-build-queue-plugin
https://github.com/jenkinsci/priority-sorter-plugin
https://github.com/jenkinsci/persistent-build-queue-plugin
https://github.com/jenkinsci/merge-queue-plugin