---
title: "钉钉通知"
description: "钉钉通知"
date: 2019-01-08T11:08:06+08:00
draft: false
toc: true
---

支持的特性包括：

* 发送任务状态
* 发送文本类型的消息

# 安装插件

在 *系统管理》插件管理》可选插件* 中搜索 `dingding-notifications` 并安装。

# 机器人

在钉钉中新建一个群，然后在群里添加一个机器人，选择自定义类型。

# 自由风格

在 Jenkins 中新建自由风格任务

在 *构建后操作* 中 *添加钉钉通知器配置* 。在表单中需要把 Jenkins 的 URL 设置正确——必须让钉钉可以访问到。Token 是刚才在钉钉机器人里的 webhook 里。

# 流水线

你也可以在流水线（Jenkinsfile）中使用下面是 DSL 来实现发送消息：

`dingTalk accessToken: 'token', imageUrl: 'imageUrl', jenkinsUrl: 'http://localhost/jenkins', message: 'messages', notifyPeople: '18211192837'
`
