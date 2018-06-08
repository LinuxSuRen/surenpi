# 构建

跳过JavaDoc `-Dmaven.javadoc.skip=true`

跳过测试 `-DskipTests=true`

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
