---
title: 参数化流水线
description: 参数化流水线
keywords: pipeline 
draft: false
toc: true
---

参数化可以避免创建很多相似的流水线，例如：同一个代码仓库，每次希望构建的分支不一样，就可以把分支的名称作为参数。

# 类型

流水线支持很多类型的参数，包括：字符串、布尔值、密码、文件，甚至可以获取到其他任务（job）的数据。

[点击这里](https://jenkins.io/doc/developer/extensions/jenkins-core/#parameterdefinition)查看更多的参数类型

# 位置

参数可以直接写在 Jenkinsfile 中，也可以在流水线配置界面上添加。下面给出一个在 Jenkinsfile 中添加参数的的例子：

{{< exfile "static/codes/jenkinsfile/parameter.groovy" "groovy" >}}

# 执行历史

每次执行时的参数值，都可以在执行历史中看到。