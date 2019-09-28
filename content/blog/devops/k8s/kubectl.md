---
title: Kubectl
description: Kubectl 的使用
keywords:
- kubectl
- patch
categories:
- kubernetes
- k8s
toc: true
---

# 查看集群信息

`kubectl cluster-info`

# 下载

`wget https://storage.googleapis.com/kubernetes-release/release/v1.10.0/bin/linux/amd64/kubectl`

# 切换上下文

`kubectl config use-context dev`

# 查看上下文

`kubectl config current-context`

# 导出用户证书

```
curl https://raw.githubusercontent.com/LinuxSuRen/getGo/master/k8s-create-user.sh | bash -s username
```

# Troubleshooting

journalctl -u kubelet -f

# Delete pod

`kubectl get pods -a | grep Error | cut -d' ' -f 1 | xargs kubectl delete pod`

# 节点

查看节点列表 `kubectl get node`

查看节点详情 `kubectl get node -o wide`

显示标签 `kubectl get node --show-labels=true`

添加标签 `kubectl label node nodeName jenkinsAgent=maven`

# Patch

在调试的时候，我们可能需要反复地修改 `k8s` 资源中的某几个字段。如果，每次都使用 `kubectl edit` 命令操作的话，会比较麻烦。通过 `kubectl` 的子命令 `patch` 可以实现快速地实现这个过程。

该命令，可以通过 `json` 或者 `yaml` 格式来对资源*打补丁*。下面，给出一个使用 `json` 的例子：

`kubectl patch deploy template -n namespace --patch '{"status":{"phase":"Pending"}}'`

## 更新镜像

`kubectl set image deploy/your-deployment -n namespace imageName=surenpi/image`
