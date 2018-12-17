---
title: HTTP
description: HTTP
---

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