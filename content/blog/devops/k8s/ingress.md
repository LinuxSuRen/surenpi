---
title: Ingress
description: Ingress
---

# 413 Request Entity Too Large

添加如下 Annotations

```
ingress.kubernetes.io/proxy-body-size: "0"
nginx.ingress.kubernetes.io/proxy-body-size: "0"
nginx/client_max_body_size: "0"
```