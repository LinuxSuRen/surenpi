---
title: Git 代理
description: 配置 Git 使用代理
keywords:
- Git代理
toc: true
---

## 全局代理

`git config --global http.proxy http://proxyUsername:proxyPassword@proxy.server.com:port`

## 部分代理
我们也可以对指定的域名配置代理，格式为：`git config --global http.https://github.com.proxy http://proxyUsername:proxyPassword@proxy.server.com:port`

```
git config --global http.https://github.com.proxy 127.0.0.1:1087
git config --global http.https://github.com.sslVerify false
```

## 查看配置

```
git config --global --get-regexp http.*
git config --get-regexp http.*
```

## 取消代理

```
git config --global --unset http.proxy
git config --global --unset http.https://domain.com.proxy

git config --global --unset http.sslVerify
git config --global --unset http.https://domain.com.sslVerify
```
