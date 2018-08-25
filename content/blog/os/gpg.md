---
title: GPG
description: GPG
---

# 安装

Windows版本下载地址 https://www.gpg4win.org/download.html

Ubuntu下安装 `sudo apt-get install gpg -y`

Mac下安装 `brew install gpg` 或者从 https://gpgtools.org/ 下载

从 https://sourceforge.net/p/gpgosx/docu/Download/ 下载 gpg2

# 命令

`gpg --keyserver hkp://pool.sks-keyservers.net --send-k` 上传公钥

`gpg --export-secret-key` 标识名 > 导出文件名

`gpg -o keyfilename --export-secret-keys KeyID`

`gpg -o keyfilename --export KeyID`

`gpg --import` 密钥文件

`gpg --delete-secret-key hello`

`gpg --delete-key hello`

`gpg –keyserver pool.sks-keyservers.net –recv-keys surenkey`

# 异常

```
gpg: cannot open /dev/tty': No such device or address
can't connect to /root/.gnupg/S.gpg-agent’: No such file or directory
gpg: can’t connect to the agent: IPC connect call failed
gpg: signing failed: No secret key
```