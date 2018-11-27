---
title: Git 标签（tag）
description: Git 标签（tag）的相关操作
---

# 标签

`git tag v0.1.2-light`

`git tag -a v0.2 -m 'here is comment'`

切换到标签 `git checkout v0.2`

删除标签 `git tag -d v0.2`

基于提交记录打标签及 `git tag -a v0.2 beab786 -m 'from branch hello beab786'`

发布指定的标签 `git push origin v0.2`

把所有的标签一次性发布 `git push origin --tags`

git可以在同一个提交上打多个名称不一样的标签

# 获取远程标签

`git fetch --all --tags` 这样会获取所有远程库的所有标签（tag）