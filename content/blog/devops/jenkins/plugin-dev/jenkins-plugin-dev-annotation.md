---
title: Annotation
description: Jenkins 中的注解类介绍
keywords: jenkins plugin
---

# CheckForNull

# Extension

# NoExternalUse

# Initializer

Usage: `@Initializer(after=InitMilestone.PLUGINS_PREPARED, fatal=false)`

通过该注解，可以实现在 Jenkins 的特定生命周期阶段来初始化特定方法。

# Terminator

Usage: `@Terminator(after=TermMilestone.STARTED)`

上面的示例中，表明了使得某个方法在 Jenkins 开始关闭时调用。

# Restricted

Usage: `@Restricted(NoExternalUse.class)`

上面的示例中，表明了某个方法不允许被直接调用，通常在接口回调中使用。