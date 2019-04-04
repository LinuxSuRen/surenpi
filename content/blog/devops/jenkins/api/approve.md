---
title: "Approve"
description: "Approve"
date: 2019-04-04T10:49:47+08:00
draft: false
toc: true
---

这里介绍的是[核准插件]({{< ref "blog/devops/jenkins/jenkins-plugin/pipeline-input-step.md" >}})中的 API 使用。

{{% note %}}
该 DSL 是基于某一次构建的，因此，下面的 API 都需要对应的构建 ID。另外，还需要对应的 `inputid`。
{{% /note %}}

# 通过

`input` 这个 DSL 支持接收参数。

## 无参数

对于没有参数的情况，可以使用下面的两个 API 来核准。

```
curl http://localhost:8080/job/pipeline/1/input/inputid/proceed \
--header "Authorization: Basic `echo -n admin:token | base64`" -X POST
```

下面的 API 是直接无参数的调用

```
curl http://localhost:8080/job/pipeline/1/input/inputid/proceedEmpty \
--header "Authorization: Basic `echo -n admin:token | base64`" -X POST
```

## 有参数

参数需要以 JSON 数组的形式传递。

```
curl http://localhost:8080/job/pipeline/1/input/inputid/proceed \
--header "Authorization: Basic `echo -n admin:token | base64`" \
-X POST -d 'json={"parameter": [{"name": "name", "value": "bob"}, {"name": "age", "value": "12"}]}'
```

如果缺少参数的话，会返回 `400`

# 拒绝

拒绝执行，可以调用下面的 API

```
curl http://localhost:8080/job/pipeline/1/input/inputid/abort \
--header "Authorization: Basic `echo -n admin:token | base64`" -X POST
```
