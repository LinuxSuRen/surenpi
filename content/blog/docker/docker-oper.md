---
title: Docker Operation
description: Docker Operation
---

# Filter Images

当前支持的过滤如下：

* dangling（布尔值，true 或 false）
* label（label=<key> 或 label=<key>=<value>

删除包含制定字符串的镜像 `docker rmi $(docker images | grep jenkinsx | awk '{print $3}')`