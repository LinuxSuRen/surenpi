---
title: Ingress
description: Ingress
keywords: 413, ingress
---

# 413

在通过 Ingress 访问的系统中上传较大文件时，可能会遇到错误 `Request Entity Too Large`。

这时，找到对应的 Ingress 添加如下 Annotations 即可解决：

```
ingress.kubernetes.io/proxy-body-size: "0"
nginx.ingress.kubernetes.io/proxy-body-size: "0"
nginx/client_max_body_size: "0"
```