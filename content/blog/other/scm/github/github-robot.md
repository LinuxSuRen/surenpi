---
title: "GitHub Robot"
description: "GitHub Robot"
date: 2019-04-06T20:57:38+08:00
draft: true
toc: true
---

这里列出一些 GitHub 中的“机器人”。

# Stale

[Stale](https://probot.github.io/apps/stale/)可以监控已有的 PR 列表，对于长期没有更新的 PR 会自动
添加“过期”的标签。当这类 PR 再次有更新或者评论时，自动取消“过期”标签；如果还是没有任何变化的话，会讲对应
的 PR 关闭。

这样就能做到，提供贡献者及时查看自己相关的 PR 动态。
