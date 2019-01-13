---
title: Maven
description: Maven 是非常流行的一个 Java 持续构建工具
toc: true
keywords: maven
---

{{< youtube "cynDKA03bWA" >}}

# 构建

跳过JavaDoc `mvn package -Dmaven.javadoc.skip=true`

跳过测试 `mvn package -DskipTests=true`

`mvn package -U` 强制更新快照版本的依赖

# 离线

通常情况下，Maven 需要从网络上下载依赖文件，这样的话，执行效率会比较低。如果你确保本地仓库已经有了所有的依赖文件的话，可以添加参数使 Maven 不再从网络上下载，直接使用本地的依赖。例如：

`mvn clean package --offline`

# 部署

`mvn deploy`

# 多模块

## 部署

如果有一个多模块的 Maven 工程，但是，在发布的时候有些工程不希望发布，例如：example 工程。这时候，可以在不需要发布的模块工程 `pom.xml` 中添加如下插件配置：

```
<plugin>
	<groupId>org.apache.maven.plugins</groupId>
	<artifactId>maven-deploy-plugin</artifactId>
	<version>2.8.2</version>
	<configuration>
		<skip>true</skip>
	</configuration>
</plugin>
```


```sequence
Alice->Bob: Hello Bob, how are you?
Note right of Bob: Bob thinks
Bob-->Alice: I am good thanks!
```

# 查看依赖

`mvn dependency:tree` 查看依赖树

`mvn dependency:resolve`

`mvn dependency:analyze` 依赖分析

`mvn help:effective-pom`

# 调试

`mvn package -X` 查看调试信息