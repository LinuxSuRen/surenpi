---
title: Jenkins 插件开发
description: Jenkins 插件开发教程
keywords: jenkins-plugin
---

# 环境准备

你需要安装好Java和Maven。

# 创建工程

执行命令 `mvn --update-snapshots archetype:generate -Dfilter=io.jenkins.archetypes:`，然后，根据提示选择骨架类型，以及当前 Maven 工程的信息。

# 运行

在工程根目录下执行 `mvn hpi:run`

# 访问

`http://localhost:8080/jenkins/`

# Jenkins 版本选择

理论上，你可以任意选择一个喜欢的 Jenkins 版本来开发。但是，处于严谨的做法，同时也是 Jenkins 官方推荐的。你可以到 Jenkins 的[下载页面](https://jenkins.io/download/)找一个长期支持版本（LTS）。

# 单元测试

Jenkins 官方提供了[单元测试](https://wiki.jenkins.io/display/JENKINS/Unit+Test)相关的内容。

# 发布

`mvn -Darguments="-DskipTests" release:prepare`

`mvn -Darguments="-DskipTests" release:perform`

# 参考

[Extend Jenkins](https://wiki.jenkins.io/display/JENKINS/Extend+Jenkins)
