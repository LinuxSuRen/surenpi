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

Note: If you still have files to push or commit, please make sure that these actions are performed or the files are stashed before you perform the next commands.

## 重置 GitAttributes
```
git rm --cached -r
git reset --hard
```

The above commands will now update the files for the repo using the newly defined line ending as specified in the .gitattributes.

Any changes or new changes will automatically use the line endings specified for that file type.

Next step is to inform any team mate or collaborator, to run the GitAttributes Reset commands.

Now prettier won't complain about CR and all developers can now live in peace! ☮️

Code File With No Linting Errors - .gitattributes

Code File With No Prettier Linting Errors
