---
title: Ingress
description: Ingress
keywords: 413, ingress
toc: true
---

## Controller

如何要利用 `Ingress` 来监听 `80` 端口的话，需要按照 `Controller`.

## Rewrite

Ingress 支持地址重写，具体请[查看参考](https://github.com/kubernetes/ingress-nginx/blob/master/docs/examples/rewrite/README.md)。

## 413

在通过 Ingress 访问的系统中上传较大文件时，可能会遇到错误 `Request Entity Too Large`。

这时，找到对应的 Ingress 添加如下 Annotations 即可解决：

```
ingress.kubernetes.io/proxy-body-size: "0"
nginx.ingress.kubernetes.io/proxy-body-size: "0"
nginx/client_max_body_size: "0"
```

## HTTPS

如果不希望自动跳转到 HTTPS 的话，可以添加如下的注解：

`
nginx.ingress.kubernetes.io/ssl-redirect: "false"
`