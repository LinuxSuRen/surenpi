---
title: "Unzip"
description: "Unzip"
date: 2019-04-30T15:26:05+08:00
draft: false
toc: true
---

## 解压

`unzip git.zip`

## 列出压缩包中的文件

`unzip -l git.zip`

## 输入压缩包中的文件内容

我们可以把 zip 压缩包中的文件内容解压后读取，但也可以做到直接读取而无需解压 `unzip -c git.zip META-INF/MANIFEST.MF`
