---
title: Nginx
description: Nginx
---

# 重定向

```
server {
        listen 80;
        server_name surenpi.com;

        location / {
                rewrite ^ https://linuxsuren.github.io/blog/;
        }
}
```