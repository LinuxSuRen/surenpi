---
title: "Blue Ocean"
description: "Blue Ocean"
date: 2019-07-15
draft: false
toc: true
---

## API 列表
获取 JUnit 测试包括的 API 包括如下：

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?state=REGRESSION&start=0&limit=101`

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=FAILED&state!=REGRESSION&start=0&limit=101`

`ttps://129.28.182.197/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=SKIPPED&start=0&limit=101`

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?state=FIXED&start=0&limit=101`

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=PASSED&start=0&limit=101`

响应体的结构如下：

```
{
	Age int
	// Duration the time of test
	Duration float32
	// ErrorDetails error details of test
	ErrorDetails string
	// ErrorStackTrace if the status is erro then error stack trace of test
	ErrorStackTrace string
	// HasStdLog indicate whether has standard log outpupt
	HasStdLog bool
	// ID id for the test report item
	ID string
	// Name is the name of test report item
	Name  string
	State string
	// Status indicates the status of report item
	Status string
}
```

## 在 Jenkins 的测试报告中，可以拿到五种状态：
* REGRESSION
* FAILED
* SKIPPED
* FIXED
* PASSED

其中 `REGRESSION` 和 `FIXED` 使用的字段为 `state`

`FAILED` 状态则会和 `REGRESSION` 配合使用

`PASSED` 和 `SKIPPED` 则使用的字段为 `status`

上面的 API 返回的状态码可能会包括：400、404、200。当给出的状态值为非法时，例如：`status=abc` 就会返回 `400`。当该流水线中没有测试包括时会返回 `404`。错误时的响应体结构如下：

```
{
  "message" : "bad state PASSED",
  "code" : 400,
  "errors" : [ ]
}
```

更多来自 Blue Ocean 的 Restful API 请参考 https://github.com/jenkinsci/blueocean-plugin/tree/master/blueocean-rest
