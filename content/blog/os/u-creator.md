---
title: U Creator
description: U Creator
---

```
$ diskutil list #查看当前分区情况，记住你的U盘的名称
 
$ diskutil unmountDisk /dev/disk1 #卸载U盘。这里的disk1是刚刚显示的名称
 
$ dd if={iso_dir} of=/dev/disk1 bs=1m #进行拷贝
```

```
diskutil list #1.找到U盘的代号 比如disk1
diskutil unmountDisk /dev/disk1 #2.
diskutil eraseDisk JHFS+ iDisk disk1 #3.
```

```
brew install e2fsprogs
diskutil unmountdisk /dev/disk2s1
sudo $(brew --prefix e2fsprogs)/sbin/mkfs.ext3 /dev/disk2s1
```