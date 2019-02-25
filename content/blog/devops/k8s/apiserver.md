---
title: ApiServer
description: ApiServer
toc: true
keywords:
- curl
- ServiceAccount
- base64
---

# 类型查看

查看所有的资源组 `curl --insecure https://localhost:6443/apis/ | grep groupVersion | sort`

执行命令 `curl --insecure https://localhost:6443/apis/{api.group}/{version}` 可以看到所支持的资源类型

# 认证

首先，需要从 `Secret` 中获取 `Token`，参考命令如下：

`kubectl get secret default-token-m99lz -o jsonpath={.data.token} | base64 --d > token`

上面的命令，我们从当前的命名空间（namespace）里默认的 `Secret` 中获取 `Token`，进行 `base64` 解密后保存为文件。

然后，通过如下类似的命令即可访问到集群的 API：

```
curl -k  https://localhost:6443/api/v1/namespaces --header "Authorization: Bearer `cat token`"
```

{{% note %}}
需要确保的是，所使用 `Token` 对应的 `ServiceAccount` 具有所访问的资源的权限。
{{% /note %}}