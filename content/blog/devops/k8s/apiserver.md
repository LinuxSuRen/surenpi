---
title: ApiServer
descrption: ApiServer
---

# ApiServer

# 类型查看

查看所有的资源组 `curl --insecure https://localhost:6443/apis/ | grep groupVersion | sort`

执行命令 `curl --insecure https://localhost:6443/apis/{api.group}/{version}` 可以看到所支持的资源类型
