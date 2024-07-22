---
title: Git Stage
description: Git Stage
---

我们提交文件，要先执行 `git add` 然后再执行 `git commit`。但是，如果在执行 `commit` 时，发现由于误操作导致一些不希望提交的文件也被 `add` 了。这时候，我们可以执行命令 `git reset HEAD readme.md` 来取消该操作。

`git checkout readme.md` 可以把 `unstage` 状态的文件修改丢弃