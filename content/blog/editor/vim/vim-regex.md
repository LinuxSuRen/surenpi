---
title: VIM Replace
description: VIM Replace
---

简单文本替换
`:[range]s[ubtitute]/{pattern}/{string}/[flags] [count]`

全局文本替换
`:[range]g[lobal]/{pattern}/[cmd]`

`:[range]g[lobal]!/{pattern}/[cmd]`

示例：

将找到的第一个字符串abc替换为def
`s/abc/def`

将找到的所有字符串abc替换为def
`s/abc/def/g`

将包含有字符串abc的行删除
`g/abc/ del`

将行首增加序列号
`let i = 1 | g/./s/^/\=i/ | let i += 1`

将所有包含image331的行中，替换331为从331开始递增的数字
`let i = 331 | g/image331/s/331/\=i/ | let i += 1`

从1358行开始进行替换
`let i = 792 | 1358,$g/resources\/image/s/\d\d\d/\=i/ | let i += 2`

将所有已<para>开头的行在结尾添加</para>
`g/^\s*<para>/ s/$/<\/para>`

把以<para>开头，若干个空白</para>结尾的行删除
`g/^\s*<para>\s*<\/para>$/ del`

把以空白字符开头，并且第一个字符不是<的行首添加<para>
`g!/^\s*</ s/^/<para>`

把<para></para>中间还有尖括号(<>)的行删除
`g/<para>.*[<>].*<\/para>/ del`

找到没有以</para>结尾的行
`s/para>.*[^(<\/para>)]$`