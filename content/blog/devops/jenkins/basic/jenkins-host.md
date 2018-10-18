---
title: Jenkins 插件托管
description: 本文介绍如何把一个 Jenkins 插件托管到 Jenkins 的 Github 上。
---

# 创建库

你需要在 [GitHub](https://github.com/) 上申请一个账号，然后，把你的工程提交上去。

# 提交请求

[详细阅读托管要求](https://wiki.jenkins.io/display/JENKINS/Hosting+Plugins#HostingPlugins-Contents)

在 [Jira](https://issues.jenkins-ci.org) 上申请一个账号，然后[创建](https://issues.jenkins-ci.org/projects/HOSTING)一个 issues，等待管理员的审核。

## 托管

## 上传

首先，在 Jenkins 的[成品库](https://repo.jenkins-ci.org/)上，用上面 Jira 的账号登陆一次，保证该账号是可用的。

你需要 Fork [上传权限申请](https://github.com/jenkins-infra/repository-permissions-updater)，然后，根据要求编写对应的 YAML 文件。发起的 PR 通过以后，你就有了上传插件的权限了。

# 发布

等一切就绪，你的插件要发布时，只要执行 `mvn release:prepare release:perform`

跳过测试 `mvn -Darguments="-DskipTests" release:prepare`

执行发布准备而且无需输入参数 `mvn -Darguments="-DskipTests" release:prepare -B`