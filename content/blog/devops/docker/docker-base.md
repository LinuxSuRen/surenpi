---
title: Docker 基础镜像
description: Docker 基础镜像
toc: true
---

通常情况下，我们自己的镜像都会依赖于某个基础镜像。下面，是我总结的一些基础镜像。

## Alpine

`alpine` 是一个非常轻量级的 Docker 镜像，大概只有 `5.55MB` 的大小。它包含一些比较常用的 Linux 命令。

`bash` 是很多人使用的 `shell`，但是，默认（`sh`）情况下，`alpine` 并不包含 `bash`。
如果我们直接在 `alpine` 中执行 `bash` 命令的话，可能会报一些貌似没有道理的错误。其实，
你只要在 `Dockerfile` 中安装一下 `bash` 即可，下面给出一个示例：

```
FROM alpine:latest
RUN apk add bash
```