---
title: Sync
description: Sync
---

```
package main

import (
  "sync"
)

type Boot struct {
  count int
  once  sync.Once
}

func (b *Boot) Init() {
  b.once.Do(func() {
    b.count += 1
  })
}

func main() {
  b := Boot{
    count: 0,
  }

  b.Init()
}
```