---
title: Jenkins 调试
description: Jenkins 调试
---

`mvn jenkins-dev:run`

`mvn jenkins-dev:run -Djenkins.install.runSetupWizard=true -DJENKINS_HOME=.` 需要在目录 `war` 下运行，并保证没有 `work` 目录

`java -server -Xms200m -Xmx512m -jar jenkins.war --webroot=/var/cache/jenkins --httpPort=8090 -DJENKINS_HOME=./jenkins-home/` 设置 Jenkins 启动参数

`mvn clean install -DskipTests -Dfindbugs.skip=true -Dpmd.skip=true -Dmaven.javadoc.skip=true`
