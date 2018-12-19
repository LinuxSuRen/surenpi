---
title: Action
description: Jenkins Plugin Action
keywords: jenkins-plugin
---

本文介绍如果开发一个 Jenkins 的后端，也就是 Action。

# 接口

你首先需要实现接口 `hudson.model.Action` 的三个方法:

* getIconFileName
* getDisplayName
* getUrlName

# 约定