---
title: Base64
description: Base64
toc: true
keywords: base64
---

# 加密

`echo -n 1234 | base64` 这里的参数 `n` 很重要，如果没有的话，会把换行符也包含进去。

# 解密

`echo MTIzNA== | base64 --d`