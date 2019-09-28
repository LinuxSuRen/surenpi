---
title: Action
description: Jenkins Plugin Action
keywords: jenkins plugin
---

本文介绍如果开发一个 Jenkins 的后端，也就是 Action。

# 接口

你首先需要实现接口 `hudson.model.Action` 的三个方法:

* getIconFileName
* getDisplayName
* getUrlName

hudson.model.queue.FoldableAction

# 约定

## QueueAction
hudson.model.Queue.QueueAction

方法 `shouldSchedule` 会决定某一次构建是否会被调度，如果返回 `false` 的话，这次构建
动作就会被合并。通常情况下，对于用户触发两次构建，这两次之间没有任何变化的话，我们就
只需要构建一次即可，这样可以减少不必要的构建资源浪费。

## FoldableAction

hudson.model.queue.FoldableAction