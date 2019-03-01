---
title: "Git GC"
description: "清理无用的文件并优化本地仓库"
date: 2019-01-29T15:28:28+08:00
draft: false
toc: true
---

在当前库中运行一些清理任务，例如：压缩文件版本（减少磁盘占用，提高性能），移除垃圾对象。

鼓励用户在每个库中定期运行该任务，这样能保持更好的磁盘利用率，以及运行性能。

一些 `git` 命令会自动执行 `git gc`。查看参数 `--auto` 了解更多细节。如果你希望永久禁用该功能并知道后果的话，可以执行如下命令：

`git config --global gc.auto 0`

# 本地分支

查看已经合并过的分支

`git branch --merged`

查看未合并过的分支

`git branch --no-merged`

# 远程分支

清理远程过期的分支引用

`git remote prune origin`

我们也可以在拉取远程分支信息的时候执行清理任务

`git fetch -p`

查看已经合并过的远程分支

`git branch -r --merged`

[更多的分支操作]({{< ref "git-branch.md" >}})