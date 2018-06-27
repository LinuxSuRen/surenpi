---
title: SSH
description: SSH
---

配置文件 `vim /etc/ssh/sshd_config`

#  

如果你希望允许 `root` 用户可以通过 `ssh` 远程登陆的话，可以找到配置文件，并将 `PermitRootLogin` 设置为 `yes`。

# 服务重启

`service ssh restart`