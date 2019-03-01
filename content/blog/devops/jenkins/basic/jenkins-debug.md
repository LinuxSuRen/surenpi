---
title: Jenkins 调试
description: Jenkins 调试
lastmod: 2019-03-01
toc: true
keywords:
- runSetupWizard
- httPort
---

`mvn jenkins-dev:run`

`mvn jenkins-dev:run -Djenkins.install.runSetupWizard=true -DJENKINS_HOME=.` 需要在目录 `war` 下运行，并保证没有 `work` 目录

`java -server -Xms200m -Xmx512m -jar jenkins.war --webroot=/var/cache/jenkins --httpPort=8090 -DJENKINS_HOME=./jenkins-home/` 设置 Jenkins 启动参数

`mvn clean install -DskipTests -Dfindbugs.skip=true -Dpmd.skip=true -Dmaven.javadoc.skip=true`

# Docker

如果需要调试运行在 `Docker` 中的 `Jenkins` 实例的话，可以通过增加 `JVM` 参数来实现。对 Jenkins 启动过程感兴趣的同学，可以自行翻阅[jenkins.sh](https://github.com/jenkinsci/docker/blob/master/jenkins.sh)来了解更多细节。

添加如下环境变量

`JAVA_OPTS=-Xdebug -Xrunjdwp:server=y,transport=dt_socket,address=5005,suspend=y`