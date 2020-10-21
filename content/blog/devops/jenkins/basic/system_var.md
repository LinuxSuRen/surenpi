---
title: Jenkins 系统属性配置
description: Jenkins 系统属性配置
lastmod: 2020-03-06
toc: true
---

Jenkins 有很多不包含在 UI 中的配置，可以应对一些特殊情况：

| Key | Values |
|---|---|
| `hudson.model.DownloadService.noSignatureCheck` | `true` |
| `hudson.model.DirectoryBrowserSupport.CSP` | |
| `hudson.security.csrf.DefaultCrumbIssuer.EXCLUDE_SESSION_ID` | `true` |
| `kubernetes.websocket.ping.interval` | `10000` |
| `org.jenkinsci.plugins.gitclient.Git.timeOut` | `60` |
| `-Dhudson.security.csrf.GlobalCrumbIssuerConfiguration.DISABLE_CSRF_PROTECTION` | `true` |
