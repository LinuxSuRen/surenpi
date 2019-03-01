---
title: "Cause"
description: "Cause"
date: 2019-01-17T10:41:03+08:00
draft: true
toc: true
---

`hudson.model.Cause` 用于申明构建触发的原因。下面介绍几个实现类：

`BranchIndexingCause` 由扫描分支而触发

`CommitHookCause` 由 WebHook 触发

`TimerTriggerCause` 由定时计划触发

`SCMTriggerCause` 由代码变更触发