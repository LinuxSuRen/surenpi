---
title: "Interface"
description: "Interface"
date: 2019-01-03T10:15:44+08:00
draft: true
toc: true
---

# 定义

```
type InterfaceTypeName interface {
    Print()
}
```

# 实现方法的区别

```
type Demo struct{}

func (i InterfaceTypeName) Print(){}

type Demo2 struct {}

func (i *InterfaceTypeName) Print(){}
```