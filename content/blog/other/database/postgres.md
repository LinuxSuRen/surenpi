---
title: Postgres
description: Postgres
toc: true
---

## 安装

`helm install --namespace default stable/postgresql --name postgre`

## 获取密码

`export POSTGRES_PASSWORD=$(kubectl get secret postgre-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)`

## 连接数据库

`kubectl run postgre-postgresql-client --rm --tty -i --restart='Never' --image docker.io/bitnami/postgresql:11.3.0 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host postgre-postgresql -U postgres`

## 查询

创建数据库 `create database`

查询数据库 `\l`