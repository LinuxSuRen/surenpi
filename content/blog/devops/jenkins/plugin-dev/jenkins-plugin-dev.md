---
title: Jenkins 插件开发
description: Jenkins 插件开发教程
keywords: jenkins plugin
toc: true
---

## 环境准备

你需要安装好Java和Maven。

## 创建工程

执行命令 `mvn --update-snapshots archetype:generate -Dfilter=io.jenkins.archetypes:`，然后，根据提示选择骨架类型，以及当前 Maven 工程的信息。

{{< ghref "插件工程的父 POM" "jenkinsci/plugin-pom" >}} 里定义了很多信息，包括：Jenkins 的版本、Java 版本、单元测试框架版本等。

## 运行

在工程根目录下执行 `mvn hpi:run`

## 访问

`http://localhost:8080/jenkins/`

## Jenkins 版本选择

理论上，你可以任意选择一个喜欢的 Jenkins 版本来开发。但是，处于严谨的做法，同时也是 Jenkins 官方推荐的。你可以到 Jenkins 的[下载页面](https://jenkins.io/download/)找一个长期支持版本（LTS）。

## 单元测试

Jenkins 官方提供了[单元测试](https://wiki.jenkins.io/display/JENKINS/Unit+Test)相关的内容。

## 发布

`mvn -Darguments="-DskipTests" release:prepare`

`mvn -Darguments="-DskipTests" release:perform`

### 实验性版本发布
为了简化发布 beta 版本的插件到特定感兴趣的人群，Jenkins 有一个实验性的更新中心。它包括 alpha 和 beta 版本的插件，这些通常不会包含在常规的更新中心里。

如果插件的版本中包含 `alpha` 或 `beta` 的插件，发布后就会出现在实验性的更新中心里，而不会出现在常规的更新中心。但是，实验性的更新中心也包含常规的版本发布。在所有的更新中心，较新的发布会隐藏掉较老的版本。例如：1.4版本会使得1.3-beta-2在实验中心里看不到。

用户要使用实验性更新中心的话，只需要在插件管理-》高级的选项卡页面中，设置更新中心地址为 `https://updates.jenkins.io/experimental/update-center.json`。

常规的更新中心地址为 `https://updates.jenkins.io/update-center.json`。

## 参考

[Extend Jenkins](https://wiki.jenkins.io/display/JENKINS/Extend+Jenkins)
