---
title: HTTP
description: HTTP
toc: true
keywords: http
---

{{< exfile "static/codes/golang/http/main.go" "golang" >}}

# 创建 Response 结构

```
import (
    "net/http"
    "io/ioutil"
)
http.Response{
    StatusCode: http.StatusOK,
    Body:       ioutil.NopCloser(bytes.NewReader([]byte("test"))),
}
```