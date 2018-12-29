---
title: VIM Command
description: VIM Command
---

# 命令模式

## 删除文本

`x` 删除当前字符 `X` 删除前一个字符

`dd` 删除当前行 `D` 删除当前字符后的文本

`dw` 删除当前单词

## 字符替换

`R` 开启字符串替换模式，知道按下 `ESC`

`r` 替换当前字符

## 转换为大写

通过命令 `v` 选中文本，然后用 `U` 变成大写，用 `u` 变为小写。

通过命令 `~` 可以切换所选或光标所在的字符大小写

`g~~` 切换当前行的大小写

`guu` 切换当前行为小写

`gUU` 切换当前行为大写

## 撤销

`u` 撤销一次更改 `ctrl+r` 恢复一次撤销

## 剪贴板

`:reg` 或 `:register` 或 `:dis` 或 `:display` 显示剪贴板

`yy` 拷贝当前行

`p` 粘贴

## 格式化

`:left` 居左对齐 `:right` 居右对齐

## 数字

`ctrl+a` 数字+1

`ctrl+x` 数字-1

## 其他

`ctrl+g` 显示当前文件路径


# 编辑模式

# Reference

[Vim documentation: help](http://vimdoc.sourceforge.net/htmldoc/help.html)
[Vim documentation: change](http://vimdoc.sourceforge.net/htmldoc/change.html)