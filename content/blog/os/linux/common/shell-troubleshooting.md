---
title: "Shell Troubleshooting"
description: "Shell Troubleshooting"
date: 2019-04-30T21:50:19+08:00
draft: false
toc: true
keywords:
- troubleshooting
---

## line 33: syntax error: unexpected end of file

通常情况下面，遇到这个问题，可能是由于换换行符导致的，我们可以通过执行下面的命令来解决：

`dos2unix yourfile.sh`

但是，如发现报错的行数根本不存在（没有那么多）的话，就可能是由于一些语法错误导致的。

例如：`if` 对应的结束符是 `fi`，如果没有用对，就会报错。
