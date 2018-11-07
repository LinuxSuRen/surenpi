---
title: Docker api
description: Docker Api
---

获取 Token：

`curl -X POST   https://hub.docker.com/v2/users/login   -d '{"username":"user","password":"password"}' -H 'Content-Type: application/json'`