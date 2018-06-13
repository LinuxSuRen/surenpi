---
title: Java8
date: 2018-06-10T19:12:20+08:00
draft: true
---

# Lambda

## Collection

判断是否存在特定元素：

```
Set<String> set = new HashSet<>(); set.add("a"); set.add("b");
boolean exists = set.stream().anyMatch((a) -> a.startsWith("a"));
System.out.println(exists);
```