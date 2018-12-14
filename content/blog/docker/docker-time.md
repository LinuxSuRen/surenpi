---
title: Docker 时间设置
description: Docker 时间设置
---

通过如下配置，就可以实现对容器内的时区设置：

```
ENV TZ=Hongkong
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
```

# Java

如何运行的是 Java 程序的话，还可以通过命令行参数在运行时指定：

`docker run -p 8888:8080  -e JAVA_OPTS="-Duser.timezone=HongKong" jenkins/jenkins:lts`