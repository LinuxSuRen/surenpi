---
title: "Sort"
description: "Sort"
date: 2019-05-05T09:30:25+08:00
draft: false
toc: true
---

## 文件列表排序

使用降序对当前目录中的文件（夹）进行排序 `ls | sort -r`

## 版本号

我们可以借助 `sort` 命令对版本号进行排序，只需要使用参数 `-V` 即可，下面的参数 `-r` 是降序的意思

```
echo "2.0
10.0
3.0" | sort -Vr
```

从下面的例子中可以看到，`sort` 命令对带字符前缀的情况也同样是支持的：

```
echo "v2.0
v10.0
v3.0" | sort -V
```
