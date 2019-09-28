---
title: Sonarqube
description: Sonarqube
---

## Docker

```
docker run --restart=always -d --name sonarqube -p 9000:9000 -p 9092:9092 sonarqube
```

## Helm

我们可以通过 helm chart 来快速地安装一个 Sonarqube。下面的例子中，我们使用 `postgresql` 作为数据库：

```
helm install --name postgresql --set postgresqlUsername=sonarUser  \
  --set postgresqlPassword=sonarPass,postgresqlDatabase=sonarDB \
  --set persistence.enabled=false \
  --set service.type=NodePort \
    stable/postgresql
```

安装 Sonarqube：

```
helm install --name sonaqube \
  --set postgresql.postgresServer=10.111.207.76 \
  --set postgresql.service.port=30180 \
  --set service.type=NodePort \
   stable/sonarqube
```

默认的用户名和密码都是：`admin`

## 安装插件

`http://129.28.186.10:30870/admin/marketplace`
