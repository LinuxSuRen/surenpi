---
title: curl
description: curl
keywords: curl
---

# GET

使用 GET 方法请求一个URL地址 `curl http://baidu.com`

# POST

`curl -d "param" http://baidu.com`

`curl -H "Content-Type:applicatioin/json" -X POST -d "param" http://baidu.com`

curl --header "Authorization: Bearer $token" --insecure https://localhost:6443/apis/devops.alauda.io/v1alpha1namespaces/zxj/jenkinsbindings/hao/proxy -X POST -H "Content-Type:application/json" --data '{"url":"alauda/cronTabCheck?cronText=sdf"}'

# 管道

我们可以通过 curl 下载并执行脚本，例如： `curl http://test.com/test.sh | bash -s arg1`

curl http://admin:admin@chartmuseum.jx.k1.surenpi.com/api/charts -F "chart=@jenkins-x-platform-0.0.1607.tgz"
