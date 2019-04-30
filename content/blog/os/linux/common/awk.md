---
title: AWK
description: AWK
toc: true
---

# 列出第一列

`docker ps | grep hello | awk '{print $1}'`

# 特殊字符

有些情况下，我们要打印出一些特殊字符，下面给出一些示例来。

## 冒号

`docker images | grep jenkins | awk '{print $1":"$2}'`
