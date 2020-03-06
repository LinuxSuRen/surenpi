---
title: Jenkins 事件监听
description: Jenkins 事件监听
date: 2020-03-06
keywords:
- 事件监听
toc: true
---

| 类型 | 接口 | 描述 |
|---|---|---|
| 监听队列事件 | `hudson.model.queue.QueueListener` | |
| 任务监听器 | `hudson.model.listeners.ItemListener` | 接收 `Item` 的 CRUD 操作事件 |
| 任务执行状态 | `hudson.model.listeners.RunListener` | 接收任务构建的事件 |
