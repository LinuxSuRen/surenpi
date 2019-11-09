---
title: user
description: user
---

# 添加用户

`useradd user-name` 添加用户 `user-name` ，并没有添加用户目录、密码、shell等

`useradd user-name --home /home/user-name --create-home`

`su user-name`

`userdel user-name`

## sudo

`xx is not in the sudoers file.`

这时候可以使用 root 账号修改配置文件来解决：`vim /etc/sudoers`

或者把用户加到用户组 `sudo` 中。
`usermod -a -G sudo jenkins`

## 设置用户

`usermod --shell /bin/bash user-name`
