---
title: JVM
description: JVM
toc: true
---

## 查看最大内存

执行命令 `java -XX:+PrintFlagsFinal -version | grep HeapSize` 或 `java -XshowSettings:all`

## 切换 JAVA_HOME
如果需要在多个版本的 Java 之间切换的话，可以试试下面的命令（在 MacOS 下做的测试）：
```
export JAVA_HOME=`/usr/libexec/java_home -v 1.8`
```
