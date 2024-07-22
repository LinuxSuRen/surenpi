---
title: Nginx
description: Nginx
toc: true
keywords:
- rewrite
---

Nginx 配置文件的一大特点是：必须要以分号结尾。

## 变量

| Name | Description |
|---|---|
| `$scheme` | The scheme of HTTP request, could be `http`, `https` |
| `$host` |  |
| `$request_uri` | ] |

## 逻辑判断

## ngx_http_rewrite_module

### rewrite

```
server {
        listen 80;
        server_name surenpi.com;

        location / {
                rewrite ^ https://linuxsuren.github.io/blog/;
        }
}
```

### return

```
Syntax: return code [text];
        return code URL;
        return URL;
Default: -
Context: server, location, if
```

```
if ($host = "github.com") {
        return 301 https://nexus-b.alauda.cn/repository/github-proxy$request_uri;
}
```

## HTTPS