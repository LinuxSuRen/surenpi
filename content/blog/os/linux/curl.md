---
title: curl
description: curl
---

# GET

使用 GET 方法请求一个URL地址 `curl http://baidu.com`

# POST

`curl -d "param" http://baidu.com`

`curl -H "Content-Type:applicatioin/json" -X POST -d "param" http://baidu.com`

# 管道

我们可以通过 curl 下载并执行脚本，例如： `curl http://test.com/test.sh | bash -s arg1`