---
title: "Pipeline Input Step"
description: "Pipeline Input Step"
date: 2019-04-04T11:14:39+08:00
draft: false
toc: true
---

{{< exref "该插件" "https://plugins.jenkins.io/pipeline-input-step" >}}实现了一个在构建过程中
允许人工干预的功能，{{% ghref "源代码" "jenkinsci/pipeline-input-step-plugin" %}}托管在 Jenkins 的官方仓库中。

它提供的 DSL 语句为 `input`，执行到该语句的时候，需要人工决定“通过”或者“拒绝”。它还支持设置允许点击”通过“的人或者组，以及
确认该操作时输入一些参数。

# 示例

在下面的例子中，通过的话会打印用户输入的参数 `name` 的值；如果拒绝的话，当前流水线会直接失败

```
script {
    def approver = input id: 'approveId', message: 'need your approve', parameters: [string(defaultValue: '', description: 'Your name', name: 'name', trim: false), string(defaultValue: '', description: 'Your age', name: 'age', trim: false)]

    if (approver != null) {
        print approver.name
    } else {
        print "no data from approver"
    }
}
```

# API

对于“通过”或者“拒绝“这样的操作，你可以[使用 API 调用]({{< ref "blog/devops/jenkins/api/approve.md" >}})的方式来实现。
