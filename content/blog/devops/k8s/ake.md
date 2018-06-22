---
title: Ake
description: 一键安装 K8S
---

# 更新依赖

执行命令 `yum upgrade libsepol libsemanage libselinux libsemanage libselinux audit-libs`

`sudo apt-get update && apt-get install python -y`

# 安装 AKE

通过命令 `wget http://get.alauda.cn/deploy/ake/ake && chmod u+x ake && ./ake up` 会安装 AKE 以及 Kubernetes。

如果发生错误的，需要执行该命令来清理 `curl http://get.alauda.cn/deploy/ake/cleanup | bash`


# 修改配置文件

使用命令 `vim ~/.kube/config`，添加如下内容：

```
apiVersion: v1
clusters:
- cluster:
    insecure-skip-tls-verify: true
    server: http://localhost:8080
  name: local
contexts:
- context:
    cluster: local
    namespace: jx
    user: ""
  name: local
current-context: local
kind: Config
preferences: {}
users:
- name: local
  user:
    as-user-extra: {}
```
