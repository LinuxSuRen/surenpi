---
title: "Git Config"
description: "Git Config"
date: 2019-02-26T18:04:00+08:00
draft: false
toc: true
---

# 修改作者

git commit --amend --author="Author Name <name@email.com>"

# 修改提交记录

`git commit --amend` 执行该命令后直接修改上次的提交记录即可

# 本地/全局

通过参数可以指定是对全局还是当前项目的修改

`git config --local user.email <name@email.com>`

`git config --global user.email <name@email.com>`