---
title: Jenkins UI
description: Jenkins UI 介绍
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