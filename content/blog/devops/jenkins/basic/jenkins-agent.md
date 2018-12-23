---
title: Jenkins Agent
description: Jenkins Agent
keywords: jenkins agent
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

需要插件。

## Docker

## Kubernetes
