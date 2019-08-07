---
title: Git Remote
description: Git 远程库操作
toc: true
---

`git remote show origin` 可以查看仓库的分支信息，虽然另外一条命令 `git branch -a` 能看到所有的远程分支列表，但缺点是无法看到哪些分支是被删除了的。

这里的场景是：当你建立的feature或者fix分支被合并到了master后，通常会把该分支（远程）删除。那么，本地的分支实际上已经没有了远程所对应的分支了。如果，我们不把本地的分支删除的话，就会导致本地分支的数量越来越多，难以管理。通过上面的命令，我们可以很清晰地看到有哪些远程分支已经被删除。然后，通过下面的命令来清理本地的陈旧分支。

`git remote prune origin` 该命令会立即把所有需要修剪（prune）的分支删除。但，如果我们还希望能在删除之前查看有哪些会被删除的话可以执行下面的命令：

`git remote prune origin --dry-run` 当然，origin 这里就是远程库的名称，你可以换成你实际的远程库名称。

最后，你还需要再把本地的分支删除，命令是 `git branch --delete dev`。

## 多个远程库

我们可以在本地 git 仓库中，添加多个远程仓库，请参考下面的命令：

```
git remote add origin https://github.com/jenkins-zh/jenkins-cli
git remote add gitee https://gitee.com/jenkins-zh/jenkins-cli.git
```

当我们执行命令 `git push` 时，默认会把当前分支推送到所关联的远程仓库中，这里通常会是 `origin`。

如果，我们希望能同时把代码推送到多个远程仓库中的话，可以先执行如下的命令

```
git remote set-url --add --push origin https://github.com/jenkins-zh/jenkins-cli
git remote set-url --add --push origin https://gitee.com/jenkins-zh/jenkins-cli.git
```

## 删除远程仓库

我们可以通过下面的命令，删除一个远程仓库：

`git remote remove gitee`
