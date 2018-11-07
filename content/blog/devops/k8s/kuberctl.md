---
title: Kubectl
description: Kubectl 的使用
---

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