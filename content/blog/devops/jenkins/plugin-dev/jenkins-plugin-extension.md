---
title: Jenkins 插件扩展点
description: Jenkins 插件扩展点
keywords: jenkins-plugin
---

# 接口

接口 `hudson.ExtensionPoint` 用于自动生成文档

# 注解

注解 `@hudson.Extension` 加在实现类上，就可以获取扩展点的实现类列表。

# Demo

```
Jenkins.getInstance().getExtensionList(ManagementLink.class).forEach(System.out::println);
```