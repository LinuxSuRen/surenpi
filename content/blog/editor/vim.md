---
title: Vim
description: Vim
---

# 故障排除

## 粘贴缩进错乱

让 `vim` 进入 `paste` 模式，命令如下：

`:set paste`

解除 `paste` 模式的命令如下：

`:set nopaste`

# 样例配置

```
set number
set incsearch
set hlsearch
set autoindent
"set smartindent
set autochdir
set ignorecase
"set textwidth=80
set colorcolumn=+1
set tabstop=4
set shiftwidth=4
set expandtab
set backspace=indent,start,eol
set showmatch
set matchtime=1
set ruler
set foldmethod=syntax
"set foldmethod=indent
set foldlevel=0
set encoding=utf-8
set termencoding=utf-8
set langmenu=zh_CN.utf-8

language message zh_CN.utf-8

"set titlestring=中文

set nolist
set nocursorline
set nocursorcolumn

syntax on
filetype on
```
