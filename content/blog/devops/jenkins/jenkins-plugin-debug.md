---
title: Jenkins 插件调试
description: Jenkins 插件调试
---

# 运行

在 `war` 目录下运行：

`mvn hpi:run`

`mvn jenkins-dev:run`

`mvn jenkins-dev:run -Djenkins.install.runSetupWizard=true -DJENKINS_HOME=.` 需要在目录 `war` 下运行，并保证没有 `work` 目录

# 跳过测试

`mvn hpi:run -DskipTests` 该命令会跳过测试的执行，但还是会编译测试代码的。如果确实想要加快构建速度而跳过测试的话，建议使用这个命令，而不是下面的：

`mvn hpi:run -Dmaven.test.skip=true` 该命令既不执行测试代码，也不会编译。

# 指定端口

`mvn hpi:run -Djetty.port=1234'

`mvn hpi:run -Djenkins.install.runSetupWizard=true`

# 参考

[maven-hpi-plugin](https://jenkinsci.github.io/maven-hpi-plugin/run-mojo.html)
