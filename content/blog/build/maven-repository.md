---
title: Maven Repository
description: Maven 仓库
---

Maven 仓库（Repository）大致可以分为私有、公有两类。我们暂且认为企业内部的就是私有的，但是，公有的仓库也会有很多（除 Maven 官方的中央仓库以外）。

当我们的 Maven 项目发布到特定仓库以后，可以在需要依赖的 `pom.xml` 文件中添加明确的仓库地址，这样就可以避免在构建构成中因为找不到而报错了。下面给出一个参考配置：

```
<project>
  <repositories>
    <repository>
      <id>repo.jenkins-ci.org</id>
      <url>http://repo.jenkins-ci.org/public/</url>
    </repository>
    <repository>
      <id>nexus-oss-snapshots</id>
      <url>https://oss.sonatype.org/content/repositories/snapshots/</url>
      <snapshots>
        <enabled>true</enabled>
      </snapshots>
      <releases>
        <enabled>false</enabled>
      </releases>
    </repository>
    <repository>
      <id>release</id>
      <url>https://oss.sonatype.org/content/repositories/releases</url>
    </repository>
  </repositories>
</project>
```

*注意* `central` 这个 `id` 是 Maven 官方中央仓库的默认值，如果你设置的仓库使用了该值的话就会覆盖中央仓库。
