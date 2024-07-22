---
title: "Github"
description: "Github"
date: 2019-01-08T21:48:43+08:00
draft: false
toc: true
---

# webhook

该插件提供的默认 `webhook` 地址为：`http://localhost/github-webhook/`。

同时，也提供修改该地址的配置。

我们需要在 GitHub 的 webhook 设置页面添加，`Content type` 要选择 `application/json`

如果要支持多分支流水线的话，主要需要确保勾选事件：`Pull Request`