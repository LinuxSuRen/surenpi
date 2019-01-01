---
title: Ake
description: 一键安装 K8S
draft: true
---

# 更新依赖

执行命令 `yum upgrade libsepol libsemanage libselinux libsemanage libselinux audit-libs`

`sudo apt-get update && sudo apt-get install python -y`

# 安装

Support Docker versions `['1.12.6', '17.03.0', '17.03.1', '17.03.2', '17.12.0']`

通过命令 `curl https://raw.githubusercontent.com/LinuxSuRen/getGo/master/jx-k8s.sh | bash -s ip` 安装 Kubernetes


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
