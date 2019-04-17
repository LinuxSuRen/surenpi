---
title: "OpenAPI"
description: "OpenAPI"
date: 2019-04-16T11:28:49+08:00
draft: false
toc: true
---

# 类型约束
```
F0416 03:23:25.690611       1 openapi.go:40] Failed to register open api spec for root: non-body Restful parameter type should be a simple type, but got :
```
遇到这个错误，可以把复杂对象的数组变成为一个指针对象，然后，在对象里添加数组。