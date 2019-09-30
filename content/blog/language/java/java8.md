---
title: Java8
date: 2018-06-10T19:12:20+08:00
draft: false
---

## Lambda

### Collection

判断是否存在特定元素：

```
Set<String> set = new HashSet<>(); set.add("a"); set.add("b");
boolean exists = set.stream().anyMatch((a) -> a.startsWith("a"));
System.out.println(exists);
```

## brew

我们可以通过 [brew](/blog/os/mac/brew) 来安装 java

`brew cask install adoptopenjdk8`
