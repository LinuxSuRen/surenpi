---
title: xargs
description: xargs
toc: true
---

`docker images | grep nginx | awk '{print $3}' | xargs docker rmi` 删除包含指定字符的 Docker 镜像

## 参数

`ls | xargs -I {} mv {} {}.bak` 多次使用参数

## 分批使用参数

```shell
gh api /orgs/coredns/repos --paginate | jq '.[].full_name' | xargs -n1 gh dev clone
```
