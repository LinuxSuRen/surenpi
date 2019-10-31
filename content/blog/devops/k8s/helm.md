---
title: "helm"
description: "helm"
date: 2019-02-14T11:15:54+08:00
toc: true
---

`helm repo list` 列出所有已经配置的 chart 仓库

`helm status jenkins` 插件 chart 安装后的状态，和首次安装后打印出来的内容一致

`helm get values jenkins` 显示 chart 在安装时指定的自定义值

`helm get values jenkins --revision 1` 显示 chart 的历史安装自定义值
