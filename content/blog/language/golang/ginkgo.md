---
title: "Ginkgo"
description: "Ginkgo"
date: 2019-01-04T16:17:49+08:00
draft: false
toc: true
---

{{< ghref "Ginkgo" "onsi/ginkgo" >}} 是 Golang 的 BDD 测试框架。
从它的{{< exref "官方文档" "http://onsi.github.io/ginkgo/#getting-started-writing-your-first-test" >}}
中可以看到大致的用法。可惜的是，写文本时还没有中文的文档。

{{% note %}}
️⚠️值得注意的是：测试用例的文件必须是以 `test` 结尾的，例如：`foo_test.go`
{{% /note %}}

{{< ghref "Gomega" "onsi/gomega" >}} 是 Ginkgo 推荐的逻辑判断库，查看{{< exref "官方文档" "http://onsi.github.io/gomega/" >}}，了解更多内容。

# 执行测试

`ginkgo test pkg/registry/generic`

# IDE 插件

VS-Code 的插件 `vscode-ginkgo`
