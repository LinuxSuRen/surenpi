---
title: Postgres
description: Postgres
toc: true
---

## 安装

`helm install --namespace default stable/postgresql --name postgre`

`docker run -p 5432:5432 -u root --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres`

## 获取密码

`export POSTGRES_PASSWORD=$(kubectl get secret postgre-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)`

## 连接数据库

`kubectl run postgre-postgresql-client --rm --tty -i --restart='Never' --image docker.io/bitnami/postgresql:11.3.0 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host postgre-postgresql -U postgres`

## 查询

创建数据库 `CREATE DATABASE testdb;`

查询数据库 `\l`

查询当前数据库中的表 `dt`

选择数据库 `\c postgres;`

## 备份和恢复

`pg_dump -U [postgres-user] [database-name] > database.sql`

`cat database.sql | psql -U [postgres-user] -d [database-name]`
