---
title: curl
description: curl
toc: true
keywords:
- curl
---

## GET

使用 GET 方法请求一个URL地址 `curl http://baidu.com`

## POST

`curl -d "param" http://baidu.com`

`curl -H "Content-Type:application/json" -X POST -d "param" http://baidu.com`

```
curl --header "Authorization: Bearer $token" --insecure https://localhost:6443/apis/devops.alauda.io/v1alpha1namespaces/zxj/jenkinsbindings/hao/proxy -X POST -H "Content-Type:application/json" --data '{"url":"alauda/cronTabCheck?cronText=sdf"}'
```

## 管道

我们可以通过 curl 下载并执行脚本，例如： `curl http://test.com/test.sh | bash -s arg1`

`curl http://admin:admin@chartmuseum.jx.k1.surenpi.com/api/charts -F "chart=@jenkins-x-platform-0.0.1607.tgz"`

## 状态码

`curl -I -m 10 -o /dev/null -s -w %{http_code} www.baidu.com` 这样获取 HTTP 状态码

## 输出格式

支持的变量包括：

* content_type
* filename_effective
* http_code
* http_connect
* http_version

## 代理
`curl` 默认会从环境变量中获取代理配置信息，但如果给定了参数 `--proxy` 的话，就会使用参数中给定的。

Linux 下，添加环境环境变量为：`export http_proxy=127.0.0.1:8080` 以及 `export https_proxy=127.0.0.1:8080`。这里需要注意的是：`HTTP_PROXY` 是不支持的。

语法：`curl --proxy <[protocol://][user:password@]proxyhost[:port]> url`

示例：`curl --proxy http://user:password@proxyhost:port http://yourserver.com`

我们可以使用 Nginx 作为一个代理服务器，参考配置如下：
```
server {
  resolver 8.8.8.8;
  listen 80;

  location ~ .*/releases/.* {
    proxy_pass https://nexus.xxx.xxx/repository/github-proxy$request_uri;
  }

  location / {
    proxy_pass http://$http_host$request_uri;
  }
}
```
