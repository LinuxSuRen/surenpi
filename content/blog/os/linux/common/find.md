---
title: Find
description: Find
---

# 查找大文件

`find / -type f -size +500M` 这里的 `f` 指的是文件类型，并且，大于500M的。而且，字母 `M` 是大写。

# 查找文件夹

`find ~ -type d -name jx`

# 查找并执行命令

`find . -name *.java | xargs echo` 查找完后执行

# 查找特定文件

`cp $(find core/src/main/resources -type f -name "*_zh_CN.properties") /tmp`