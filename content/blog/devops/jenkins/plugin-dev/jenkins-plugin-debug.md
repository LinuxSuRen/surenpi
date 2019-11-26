---
title: Jenkins 插件调试
description: Jenkins 插件调试
date: 2019-11-26
keywords:
- jenkins
- plugin
toc: true
---

## 运行

如果是在 Jenkins core 的源码目录下的话，需要进入到 `war` 目录；如果是插件的话，通常在根目录下执行：

`mvn hpi:run`

程序正常启动后，会在 `target` 下生成一个运行时的目录：`target/jetty/webapp/`

## 跳过测试

`mvn hpi:run -DskipTests` 该命令会跳过测试的执行，但还是会编译测试代码的。如果确实想要加快构建速度而跳过测试的话，建议使用这个命令，而不是下面的：

`mvn hpi:run -Dmaven.test.skip=true` 该命令既不执行测试代码，也不会编译。

## 指定端口

`mvn hpi:run -Djetty.port=1234'

`mvn hpi:run -Djenkins.install.runSetupWizard=true`

### 指定 Jenkins 版本

`mvn hpi:run -Djenkins.version=2.121.1`

## 参考

[maven-hpi-plugin](https://jenkinsci.github.io/maven-hpi-plugin/run-mojo.html)
