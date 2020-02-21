---
title: Brew
description: Brew
---

[brew](https://brew.sh/) 是 MacOS（或者 Linux 操作系统）下的缺省包管理器。

在有 ruby 环境的操作系统中，可以使用下面的命令快速安装 brew：

`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

其他的安装方式，请参考[这里](https://docs.brew.sh/Homebrew-on-Linux)。

## 安装软件

根据来源的不同，brew 包括两类软件，分别是：core、thridparty。

### core

如果你不知道有哪些软件是可以直接安装的话，可以通过搜索命令来搜索：

`brew search curl`

如果可以找到的话，就能用下面的命令来安装：

`brew install curl`

### thridparty

如何你想要把自己开发的命令行工具可以使用 brew 来安装，也是非常的简单，只要在 GitHub 中提供一个 ruby 文件即可，具体可以参考 [homebrew-jcli](https://github.com/jenkins-zh/homebrew-jcli)。那么，对于这种第三方的命令行如何安装呢，也是非常简单，请参考下面 [jcli](https://github.com/jenkins-zh/jenkins-cli) 的例子：

```
brew tap jenkins-zh/jcli
brew install jcli
```

## 其他操作

| 命令 | 描述 |
|---|---|
| `brew remove jcli` | 删除软件 |
| `brew upgrade jcli` | 升级软件 |
| `brew list` | 列出来已经安装的 |
| `brew info jcli` | 查看安装信息 |
| `brew deps jcli` | 查看依赖 |
| `brew outdated` | 查看可更新的软件 |
| `brew --repo` | 查看本地目录 |

## 代理

替换上游：

```
git -C "$(brew --repo)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git

git -C "$(brew --repo homebrew/core)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git

git -C "$(brew --repo homebrew/cask)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-cask.git

brew update
```

复原：

```
git -C "$(brew --repo)" remote set-url origin https://github.com/Homebrew/brew.git

git -C "$(brew --repo homebrew/core)" remote set-url origin https://github.com/Homebrew/homebrew-core.git

git -C "$(brew --repo homebrew/cask)" remote set-url origin https://github.com/Homebrew/homebrew-cask.git

brew update
```

参考资料：https://mirror.tuna.tsinghua.edu.cn/help/homebrew/
