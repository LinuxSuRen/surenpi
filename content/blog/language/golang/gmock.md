---
title: "GoMock"
description: "GoMock"
date: 2019-01-04T16:17:15+08:00
draft: false
toc: true
---

{{< ghref "GoMock" "golang/mock" >}} 是 Go 语言的 Mock 框架。它可以很很好地和 Go 内置的 `testing` 包集成。

# 安装

你只要安装了 Go 就可以通过下面的命令来安装 `gomock` 和 `mockgen` 工具。

```
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

# 编写接口（测试用）

{{< exfile "static/codes/golang/gmock/foo.go" "golang" >}}

# 生成 Mock 接口

执行命令 `mockgen --source=foo.go --destination foo-mock.go`

# 测试用例

{{< exfile "static/codes/golang/gmock/foo_test.go" "golang" >}}
