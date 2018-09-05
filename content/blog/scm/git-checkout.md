---
title: 如何操作 Git 分支
description: Git 有很强大的分支支持
---

# 检出分支

`git checkout --track origin/fea/test` 检出 git 远程分支

# 切换标签

`git checkout master` 切换到 master 分支

# 标签

`git tag v0.1.2-light`

`git tag -a v0.2 -m 'here is comment'`

切换到标签 `git checkout v0.2`

删除标签 `git tag -d v0.2`

基于提交记录打标签及 `git tag -a v0.2 beab786 -m 'from branch hello beab786'`

发布指定的标签 `git push origin v0.2`

把所有的标签一次性发布 `git push origin --tags`

git可以在同一个提交上打多个名称不一样的标签