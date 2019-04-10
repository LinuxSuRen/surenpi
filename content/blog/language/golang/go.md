---
title: golang
description: golang
toc: true
---

# Environment

环境变量 `GOROOT` 是 Golang 的安装根目录

环境变量 `GOPATH` 是 Golang 的包依赖目录，包括的目录有：

* src 源码
* bin 执行 go get 等命令的二进制文件下载目录
* pkg 生成编译文件的目录

# Proxy

在使用命令 `go get xxx` 时，如果需要翻墙的话，可以考虑下面的方案

设置全局代理 `git config --global http.proxy 'http://192.168.2.9:8123'`

设置忽略的代理 `git config --global no.proxy '*.github.com'`

# Troubleshooting

`--oom-kill-disable=false`
