---
title: Syntax
description: Syntax
---

# 集合

判断 Map 中是否存在指定 Key

```
package main

import "fmt"

func main() {
  mapTest := map[string]string{
    "hao": "hao value",
  }

  if _, ok := mapTest["hao"]; ok {
    fmt.Println("ok")
  }
}
```