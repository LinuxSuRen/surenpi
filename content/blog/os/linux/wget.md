---
title: wget
description: wget
---

# 代理

通过命令参数设置代理 `wget https://google.com -e "https_proxy=http://localhost:1234"`

为了方便使用，我们可以为 `wget` 添加一个别名，如下所示：

`alias wall-wget='wget -e "https_proxy=http://localhost:8123" '`
