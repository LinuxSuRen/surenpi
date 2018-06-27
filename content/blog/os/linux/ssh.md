---
title: SSH
description: SSH
---

配置文件 `vim /etc/ssh/sshd_config`

# 允许 Root 登陆

如果你希望允许 `root` 用户可以通过 `ssh` 远程登陆的话，可以找到配置文件，并将 `PermitRootLogin` 设置为 `yes`。

# 无密码登录

1. 在客户端执行命令ssh-keygen生成id_rsa（私钥）、id_rsa_pub（公钥）

2. 在服务器端~/.ssh/authorized_keys文件中添加id_rsa.pub的内容

# 服务重启

`service ssh restart`
