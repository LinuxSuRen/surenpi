---
title: Git Branch
description: Git Branch
---

`git branch` 列出本地分支

# 删除分支

`git branch --delete abc` 删除分支 abc

`git push origin --delete abc` 删除远程分支

`git branch -v | grep gone | awk '//{print $1}' | xargs git branch -d` 删除远程分支不存在的`

# 拉取分支

`git checkout -b abc` 基于当前分支创建新分支

`git checkout --track origin/abc` 拉取远程分支
