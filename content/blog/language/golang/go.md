---
menu:
  blog:
    parent: "golang"
description: golang
---

# Proxy

在使用命令 `go get xxx` 时，如果需要翻墙的话，可以考虑下面的方案

设置全局代理 `git config --global http.proxy 'http://192.168.2.9:8123'`

设置忽略的代理 `git config --global no.proxy '*.github.com'`
