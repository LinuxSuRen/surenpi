---
title: Java8
date: 2018-06-10T19:12:20+08:00
draft: false
---

## Lambda

### Collection

判断是否存在特定元素：

```
Set<String> set = new HashSet<>(); set.add("a"); set.add("b");
boolean exists = set.stream().anyMatch((a) -> a.startsWith("a"));
System.out.println(exists);
```

## brew

我们可以通过 [brew](/blog/os/mac/brew) 来安装 java

`brew cask install adoptopenjdk8`

## Exceptions

如果遇到下面在请求 HTTPS 时遇到证书的问题，可能是因为由于自制私有证书导致的。我们可以把私有证书导入 JRE 来解决：

```
Caused by: sun.security.validator.ValidatorException: PKIX path building failed: sun.security.provider.certpath.SunCertPathBuilderException: unable to find valid certification path to requested target
```

首先，要先找到 JRE 的位置：`which java`。目录通常是在：`{{Installation_directory}}/{{JDK_version}}/jre/lib/security`

然后，导入证书：

`keytool -importcert -file selfsigned.crt -keystore {{Installation_directory}}/{{JDK_version}}/jre/lib/security/cacerts`

在导入过程中，可能是需要有 root 权限，并要求输入证书密码。如果你在安装 Java 后没有修改过的话，默认的密码为：`changeit`。
