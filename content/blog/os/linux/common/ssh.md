---
title: SSH
description: SSH
toc: true
---

配置文件 `vim /etc/ssh/sshd_config`

## 允许 Root 登陆

如果你希望允许 `root` 用户可以通过 `ssh` 远程登陆的话，可以找到配置文件，并将 `PermitRootLogin` 设置为 `yes`。

## 无密码登录

1. 在客户端执行命令ssh-keygen生成id_rsa（私钥）、id_rsa_pub（公钥）

2. 在服务器端~/.ssh/authorized_keys文件中添加id_rsa.pub的内容

或者，在需要无密码登录的机器上执行：`ssh-copy-id user@ip`

## 证书登陆

`ssh -i test.pem root@localhost`

## 服务重启

`service ssh restart`

## 密码传递

我们可以通过 `sshpass` 给 `ssh` 传递密码，下面给出一个例子：

`sshpass -p password ssh root@192.168.1.1`

通过 brew 安装的话，可以执行下面的命令：

`brew install https://raw.githubusercontent.com/kadwanev/bigboybrew/master/Library/Formula/sshpass.rb`
