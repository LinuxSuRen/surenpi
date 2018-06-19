---
title: Ake
---

`yum upgrade libsepol libsemanage libselinux libsemanage libselinux audit-libs`


`wget http://get.alauda.cn/deploy/ake/ake && chmod u+x ake && ./ake up`

如何发生错误的，需要执行该命令来清理 `curl http://get.alauda.cn/deploy/ake/cleanup | bash`


修改配置文件 `~.kube/config`
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
