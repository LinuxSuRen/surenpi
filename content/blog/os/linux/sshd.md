---
title: SSDH
description: SSHD
---

# 禁止root用户远程登录

修改sshd配置文件 `/etc/ssh/sshd_config`

`PermitRootLogin no`

然后重启服务 `service ssh restart`