---
title: Please Add .gitattributes To Your Git Repository
origin: https://dev.to/deadlybyte/please-add-gitattributes-to-your-git-repository-1jld
translator: linuxsuren
draft: true
---

##　什么是 .gitattributes ？
当执行 git 动作时，.gitattributes 文件允许你指定由 git 使用的文件和路径的属性，例如：`git commit` 等。

换句话说，每当有文件保存或者创建时，git 会根据指定的属性来自动地保存。

其中的一个属性是 *eol*(end of line)，用于配置文件的结尾。本文将会深入了解如何配置文件的结尾行，这样的话，即使在不同的机器、操作系统
上都可以使得每个开发者都可以使用相同的值。

## 为什么是 .gitattributes（开发者之间的争议 ⚔️ ）?
不是所有的开发者都是一样的，不能因为你是在 Windows 上使用 Visual Studio Code 开发的，就期望别的 Pull Request 也是基于相同的开发
环境完成的（在 MacOS 上可能用的是 Sublime Text2）。

正如上面提到的，开发者使用不同的操作系统，默认的文件结尾行就会不同。在 Windows 上默认的是回车换行（Carriage Return Line 
Feed, CRLF），然而，在 Linux/MacOS 上则是换行（Line Feed, LF）。

从表面看起来，内容都是一样的，那我们为什么还会困扰呢？??

好的，如果你启用了该功能，并设置属性 `endOfLine` 为 `lf` 的话。

```
{
  "endOfLine": "lf"
}
```

在 Windows 上，开发者会看到如下的检查警告。

!(linting-errors.png)[linting-errors.png]

这时候 `.gitattributes` 就可以派上用场了 🦸！

## 新仓库
首先，在仓库的根目录下创建名为 `.gitattributes` 的文件。

下面，是一份 `.gitattributes` 文件的样例内容。
```
*.js    eol=lf
*.jsx   eol=lf
*.json  eol=lf
```

把该文件提交并推送到服务器上。

```
git add .
git commit -m "Added .gitattributes to repo"
git push
```

现在，任何人从仓库获取代码后，创建以及修改文件时，git 都会自动地设置好正确的文件结尾。

## 增加到已有的 Git 仓库
正如上面提到的，在仓库的根目录下创建名为 `.gitattributes` 的文件。一旦文件推送到 git 服务器后，请确保你的本地
仓库是干净的、无需提交的。使用命令 `git status` 可以检查是否你的仓库是干净的。

**注意**：如果你还有未提及或推送的文件，请确保这些动作已经被执行过了，或者在执行下面的命令前 `stash` 过。

## 重置 GitAttributes
```
git rm --cached -r
git reset --hard
```

上面的命令就会根据文件 `.gitattributes` 中的定义，更新文件的结尾行。

任何变更都会自动使用指定文件的文件结尾行格式。

下一步，可以通知团队成员或者协作者去执行 Git 属性重置的命令。

现在，prettier 就不会在提示有关 CR 的问题了，所有的开发者都可以安心写代码了! ☮️

