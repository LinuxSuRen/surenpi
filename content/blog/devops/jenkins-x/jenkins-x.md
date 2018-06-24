---
title: Jenkins X
description: Jenkins X 特性介绍
---

添加 Gitee 为 Git 服务器

`jx create git server gitea http://localhost:10080`

# DevPod

目前只支持 overlay 类型的 Docker 驱动，通过命令 `docker info | grep "Storage Driver"` 查看 Docker 当前支持的类型。

因此，我们需要的 Docker 版本为 `v17.06`。
