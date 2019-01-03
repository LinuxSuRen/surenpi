---
title: "Core"
description: "Core"
date: 2019-01-03T16:24:48+08:00
draft: false
toc: true
---

You can access Jenkins by url, like below:

`curl http://localhost:8080/job/jobName/api/json -u user:passwd`

# Header Way

You can also put your user and token into http header:

`curl http://localhost:8080/job/jobName/api/json --header "Authorization: Basic `echo -n 'user:passwd' | base64`"`