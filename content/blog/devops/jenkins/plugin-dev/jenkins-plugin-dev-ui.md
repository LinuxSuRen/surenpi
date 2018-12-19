---
title: Jenkins UI
description: Jenkins UI 介绍
keywords: jenkins-plugin
---

# 按钮

`<j:jelly  xmlns:f="/lib/form"></j:jelly>`

`<f:validateButton title="Test Connection" method="verifyConnect" with="server,credentialsId" progress="Testing..."></f:validateButton>` 验证按钮

`<f:textbox/>`

`<f:checkbox/>`

# 凭据

```
<j:jelly xmlns:c="/lib/credentials">
<c:select/>
</j:jelly>
```

https://ci.adoptopenjdk.net/

# 静态资源文件

当我们需要在 Jenkins 插件中添加静态文件时，可以把静态访问放到 `webapp` 下。然后，通过类似下面的地址访问：

`http://localhost:8080/jenkins/plugin/sample-plugin-id/test.css`
