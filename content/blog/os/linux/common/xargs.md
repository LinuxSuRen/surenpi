---
title: xargs
description: xargs
---

`docker images | grep nginx | awk '{print $3}' | xargs docker rmi` 删除包含指定字符的 Docker 镜像