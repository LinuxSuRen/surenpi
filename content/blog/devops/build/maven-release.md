---
title: 发布Maven工程到中央仓库
description: 本文介绍如何把自己的maven项目发布到OSS中央仓库中
toc: true
---


注册帐号
https://issues.sonatype.org/secure/Signup!default.jspa

新建项目
https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134

sonatype对发布的jar有一些要求，必须包括：源码（既然是开源，那么源码肯定是少不了的）、javaDoc、gpg签名。下面介绍如何安装gpg工具，然后生成签名文件。

# 签名

参考这里安装gpg。完成后，按照下面的步骤生成密钥。

`gpg --gen-key` 生成密钥，成功后，会有类似如下的输出。

gpg: key 69EC4448C6CE32EB marked as ultimately trusted

gpg: directory '/Users/mac/.gnupg/openpgp-revocs.d' created

69EC4448C6CE32EB就是我们需要的keyname，会在maven的settings.xml中配置。

修改settings.xml
增加账户信息：

```
<server>
  <id>hello</id>
  <username>name</username>
  <password>passwd</password>
</server>
```
设置gpg的profile：
```
<profile>
  <id>hello</id>
  <activation>
    <activeByDefault>true</activeByDefault>
  </activation>
  <properties>
    <gpg.executable>gpg2</gpg.executable>
    <gpg.passphrase>pass</gpg.passphrase>
  </properties>
</profile>
```

需要注意的是，gpg有两个版本gpg和gpg2。maven的gpg插件，默认使用gpg命令来作签名，如果你需要使用gpg2的话，可以在参考上面的配置来修改（gpg.executable）。另外，上面的配置中，直接写了密钥，这样可能是不安全的。你可以在配置文件中配置keyname，如下所示：

<gpg.keyname>69EC4448C6CE32EB</gpg.keyname>

修改pom文件
```
<distributionManagement>
    <repository>
        <id>oss-sonatype-staging</id>
        <name>Sonatype Staging Repository</name>
        <url>https://oss.sonatype.org/service/local/staging/deploy/maven2</url>
    </repository>
    <snapshotRepository>
        <id>oss-sonatype-staging</id>
        <name>OSS Sonatype Nexus Snapshots</name>
        <url>https://oss.sonatype.org/content/repositories/snapshots</url>
    </snapshotRepository>
</distributionManagement>
```

上面的配置，描述了版本（release、snapshot）的发布地址。另外，这里的id需要和settings.xml中server的id保持一致。然后，添加生成javaDoc、源码等的Maven插件：

```
<profiles>
<profile>
  <id>release</id>
  <build>
    <plugins>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-gpg-plugin</artifactId>
        <executions>
          <execution>
            <id>sign-artifacts</id>
            <phase>verify</phase>
            <goals>
              <goal>sign</goal>
            </goals>
          </execution>
        </executions>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-enforcer-plugin</artifactId>
        <version>1.3.1</version>
        <executions>
          <execution>
            <id>enforce-no-snapshots</id>
            <goals>
              <goal>enforce</goal>
            </goals>
            <configuration>
              <rules>
                <requireReleaseDeps>
                  <message>No Snapshots Allowed!</message>
                </requireReleaseDeps>
              </rules>
              <fail>true</fail>
            </configuration>
          </execution>
        </executions>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-javadoc-plugin</artifactId>
        <version>2.10.3</version>
        <configuration>
          <additionalparam>${javadoc.opts}</additionalparam>
        </configuration>
        <executions>
          <execution>
            <id>attach-javadocs</id>
            <goals>
              <goal>jar</goal>
            </goals>
          </execution>
        </executions>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-source-plugin</artifactId>
        <version>2.4</version>
        <executions>
          <execution>
            <id>attach-sources</id>
            <goals>
              <goal>jar</goal>
            </goals>
          </execution>
        </executions>
      </plugin>
    </plugins>
  </build>
</profile>
<profile>
  <id>doclint-java8-disable</id>
  <activation>
    <jdk>[1.8,)</jdk>
  </activation>
  <properties>
    <javadoc.opts>-Xdoclint:none</javadoc.opts>
  </properties>
</profile>
</profiles>
```
发布
经过上面繁琐的配置后，现在终于可以发布项目到Maven公共仓库了。

在项目根目录下（根据需要自行选择），执行命令mvn deploy

如果没有任何报错的话，会先把jar等文件上传（upload）。等上传完成后，你需要登录[https://oss.sonatype.org](https://oss.sonatype.org) ，点击左侧菜单Staging Repositories。从列表中找到你刚刚上传的工程，选中后点击Close。如果你不想每次手动完成该步骤，可以在pom.xml中添加下面的插件配置来实现自动化：
```
<plugin>
  <groupId>org.sonatype.plugins</groupId>
  <artifactId>nexus-staging-maven-plugin</artifactId>
  <version>1.6.7</version>
  <extensions>true</extensions>
  <configuration>
      <serverId>nexus-snapshots</serverId>
      <nexusUrl>https://oss.sonatype.org/</nexusUrl>
      <autoReleaseAfterClose>true</autoReleaseAfterClose>
  </configuration>
</plugin>
```
多模块
有些Maven工程是多模块工程，这时候，如果我们希望把特定的模块在deploy的时候排除掉，可以通过下面的配置来实现：
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
你可以把上面的配置添加到需要排除的模块中。