---
title: Git Reset
description: Git Reset
categories:
- scm
keywords:
- reset
toc: true
---

# 撤销对某个文件的修改

当发现某次 Git 提交中，包含了错误的修改记录，可能误删文件、多提交了文件等。这时候，可以按照如下的流程把提交记录进行修改：

首先，获取某个文件的提交记录，命令为：`git log <filename>`

找到需要恢复到的那次提交记录，执行撤销命令： `git reset <commit-id> <filename>` 然后 `git checkout <filename>`

最后，确认修改：`git commit --amend && git push`

如果在上述的操作之前，`push` 过的话，记得再次 `push` 时需要添加 `-f` 参数。