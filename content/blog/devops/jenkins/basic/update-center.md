---
title: Update Center
description: Update Center
keywords: plugin update
draft: true
---

构建项目 `mvn package appassembler:assemble -DskipTests`

生成网站目录并缓存插件文件 `sh target/appassembler/bin/app -id default -www www -key surenpi.key -certificate surenpi.crt -root-certificate surenpi.crt -root-certificate rootCA/jenkins-update-center-root-ca.crt -whitelist list.properties -cache /Users/mac/IdeaProjects/Github/jenkins-infra/update-center2/tmp -skip-release-history -cacheServer http://blog.surenpi.com/www/download/`