---
title: Mac 快捷键
---

# 快捷键

删除文件 `command + delete`

## 浏览器

关闭选项卡 `command + w`

前后切换选项卡 `command + <-`

指定顺序切换选项卡 `command + 1`

关闭浏览器 `command + q`

## 截图

* 截取全屏并保存文件 `shift + command + 3`

* 截取部分并保存文件 `shift + command + 4`

* 截取全屏并保存粘贴板 `shift + control + command + 3`

* 截取部分并保存粘贴板 `shift + control + command + 4`

# 丢失管理身份

当你的误操作导致丢失了 mac 用户的管理员身份时，可以按照下面的步骤找回管理。

1. 开机后按下 `command+s`

2. 进入终端后输入命令 `/sbin/mount -uaw``

3. 执行命令 `rm /var/db/.applesetupdone`

操作系统重启后，会提示用户重新创建一个管理员身份。注意，这不会删除之前的用户信息。登陆系统后，你可以用户管理功能再次把之前的用户身份找回来。
