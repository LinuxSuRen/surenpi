---
title: Docker
description: Docker
---

# 内存限制

镜像 `8u151-alpine` 在 JVM 的内存使用上有一些限制，它会限制 JVM 使用的最大内存为机器的 1/4。我们可以使用 `8u171` 以后的版本来解决这个问题。

# 参考

[Docker OpenJDK](https://hub.docker.com/r/library/openjdk/tags/)