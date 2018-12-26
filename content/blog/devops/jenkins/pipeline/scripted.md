---
title: 脚本式流水线
description: 本文介绍 Jenkins 的脚本式流水线
keywords: jenkins pipeline
toc: true
tags:
- Jenkins
- Pipeline
---

# 触发器

{{< exfile "static/codes/jenkinsfile/script-trigger.groovy" "groovy" >}}

# 参数化执行

```groovy
properties([[$class: 'JobRestrictionProperty'],
	parameters([run(description: '',
		filter: 'ALL', 
		name: 'Name', 
		projectName: 'Project')]),
	pipelineTriggers([])]
)
```

# 异常

你可以利用处理异常的方式来实现类似于[申明式流水线中]({{< ref "declarative.md" >}})的 `post`

{{< exfile "static/codes/jenkinsfile/try-catch.groovy" "groovy" >}}

# 循环

{{< exfile "static/codes/jenkinsfile/loop.groovy" "groovy" >}}
