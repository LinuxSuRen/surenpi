---
title: "Token"
description: "Token"
date: 2019-05-07T11:54:58+08:00
draft: false
toc: true
---

```
curl http://localhost:8080/user/admin/descriptorByName/jenkins.security.ApiTokenProperty/generateNewToken --header "Authorization: Basic `echo -n admin:1182fc5ef9b28a630986cb50c0319eb9d0 | base64`" -X POST -d 'newTokenName=name'
```

返回的数据如下所示：

```
{"status":"ok","data":{"tokenName":"test","tokenUuid":"3a07e7d7-c864-4474-bc3c-d217938e460e","tokenValue":"1123acbf5d3e93e628a38dc0a671cf0194"}}
```