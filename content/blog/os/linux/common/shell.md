---
title: Shell
description: Shell
---

# 下载并执行

`curl http://test.com/test.sh | bash -s arg1 arg2`

# 文件路径

从文件路径中提取文件名和目录

```
path=/var/www/1.log
echo $(basename $path)
echo $(dirname $path)
```