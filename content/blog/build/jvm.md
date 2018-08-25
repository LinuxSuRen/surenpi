---
title: JVM Arguments
description: JVM Arguments
---

有些情况下，我们需要修改启动 Maven 时的 JVM 参数。

`MAVEN_OPTS="-Xmx1024m" mvn clean` 或 `mvn -DjvmArgs="-Xmx1024m" clean`