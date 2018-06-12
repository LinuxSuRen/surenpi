---
title: 如何操作 Git 分支
description: Git 有很强大的分支支持
---

# checkout
该命令可以切换分支、新建分支等。

# 切换分支

# 切换标签

# 新建分支

# 放弃修改

`git checkout somefile` 通过该命令，可以放弃对某个文件的修改。

* 场景1 当我使用命令 `git pull` 更新代码时，发现由于和本地对某个文件的修改发生冲突，而且确定可以放弃本地的修改时。

# 标签

`git tag v0.1.2-light`

`git tag -a v0.2 -m 'here is comment'`

切换到标签 `git checkout v0.2`

删除标签 `git tag -d v0.2`

基于提交记录打标签及 `git tag -a v0.2 beab786 -m 'from branch hello beab786'`

发布指定的标签 `git push origin v0.2`

把所有的标签一次性发布 `git push origin --tags`

git可以在同一个提交上打多个名称不一样的标签