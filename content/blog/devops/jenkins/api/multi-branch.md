---
title: "多分支流水线"
description: "多分支流水线"
date: 2019-01-09T15:44:33+08:00
draft: false
toc: true
---

# 扫描多分支

`curl -X POST -u user:passwd http://localhost:8080/jenkins/job/jobName/build --header "Jenkins-Crumb: 6e78c5725ea0c522bdcb787d548465aa"`

[Jenkins-Crumb]({{< ref "crumIssuer.md" >}}) 需要单独获取。