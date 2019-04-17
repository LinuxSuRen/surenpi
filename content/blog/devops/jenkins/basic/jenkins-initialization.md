---
title: "Jenkins Initialization"
description: "Jenkins Initialization"
date: 2019-03-26T14:35:09+08:00
draft: false
keywords:
- groovy
toc: true
---

你可以创建一个 Groovy 脚本文件，例如 `$JENKINS_HOME/init.groovy` 或者目录 `$JENKINS_HOME/init.groovy.d/` 下的任意脚本 `.grooy`. 它就会在 Jenkins 启动后执行。