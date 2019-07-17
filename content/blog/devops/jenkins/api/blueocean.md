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

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=FAILED&start=0&limit=101`

`ttps://129.28.182.197/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=SKIPPED&start=0&limit=101`

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?state=FIXED&start=0&limit=101`

`http://localhost:8080/jenkins/blue/rest/organizations/jenkins/pipelines/demo/runs/7/tests/?status=PASSED&start=0&limit=101`

## 在 Jenkins 的测试报告中，可以拿到五种状态：
* REGRESSION
* FAILED
* SKIPPED
* FIXED
* PASSED

更多来自 Blue Ocean 的 Restful API 请参考 https://github.com/jenkinsci/blueocean-plugin/tree/master/blueocean-rest
