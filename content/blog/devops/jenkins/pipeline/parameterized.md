---
title: 参数化流水线
description: 参数化流水线
keywords: pipeline 
draft: false
toc: true
---

参数化可以避免创建很多相似的流水线，例如：同一个代码仓库，每次希望构建的分支不一样，就可以把分支的名称作为参数。

参数可以直接写在 Jenkinsfile 中，也可以在流水线配置界面上添加。


{{< exfile "static/codes/jenkinsfile/parameter.groovy" "groovy" >}}