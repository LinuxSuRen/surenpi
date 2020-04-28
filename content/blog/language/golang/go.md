---
title: golang
description: golang
toc: true
---

## Environment

环境变量 `GOROOT` 是 Golang 的安装根目录

环境变量 `GOPATH` 是 Golang 的包依赖目录，包括的目录有：

* src 源码
* bin 执行 go get 等命令的二进制文件下载目录
* pkg 生成编译文件的目录

### GOGC
`GOGC` 是用来控制内存垃圾回收（GC）被触发时的比例，默认值为：`GOGC=100`。`GOGC=off`可以完全关闭垃圾回收。

## Proxy

在使用命令 `go get xxx` 时，如果需要翻墙的话，可以考虑下面的方案

设置全局代理 `git config --global http.proxy 'http://192.168.2.9:8123'`

设置忽略的代理 `git config --global no.proxy '*.github.com'`

## Package import
If you want to use a custom domain in your Go import path, please read more about [govanit](https://github.com/GoogleCloudPlatform/govanityurls). Here is [a blog](https://tonybai.com/2017/06/28/set-custom-go-get-import-path-for-go-package/) related with it.

A online service of [Go Vanity Server](https://github.com/thumbai/thumbai).

## Troubleshooting

`--oom-kill-disable=false`
