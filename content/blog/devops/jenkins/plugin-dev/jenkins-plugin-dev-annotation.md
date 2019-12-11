---
title: Annotation
description: Jenkins 中的注解类介绍
keywords: jenkins plugin
toc: true
---

## CheckForNull

## Extension

## NoExternalUse

## Initializer

Usage: `@Initializer(after=InitMilestone.PLUGINS_PREPARED, fatal=false)`

通过该注解，可以实现在 Jenkins 的特定生命周期阶段来初始化特定方法。

## Terminator

Usage: `@Terminator(after=TermMilestone.STARTED)`

上面的示例中，表明了使得某个方法在 Jenkins 开始关闭时调用。

## Restricted

Usage: `@Restricted(NoExternalUse.class)`

上面的示例中，表明了某个方法不允许被直接调用，通常在接口回调中使用。

## Whitelisted

`org.jenkinsci.plugins.scriptsecurity.sandbox.whitelists.Whitelisted`

Jenkins 流水线中的 DSL 语句，会收到沙盒机制的保护。你可能会见到类似如下的错误：

```
Scripts not permitted to use method groovy.lang.GroovyObjecgt invokeMethod
Administrators can decide whether to approve or reject this signature.
```

如果对应的方法是在自定义插件中实现的话，通过给对应的方法上增加 `@Whitelisted` 解决该问题。
 
