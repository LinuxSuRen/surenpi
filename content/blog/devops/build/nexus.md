---
title: Nexus
description: Nexus
---

## 初始密码

Nexus3 的 `admin` 的初始密码默认会自动生成，并保存在文件 `/nexus-data/admin.password` 中。

为了自动化 Nexus 的安装和初始化配置工作，我们需要能够将管理员的初始密码固定或者可以设置。
这时候，你可以通过下面的两个途径来做到：

* 修改 `$data-dir/etc/nexus.properties` 文件中的值 `nexus.security.randompassword` 为 `false`
* 设置环境变量 `NEXUS_SECURITY_RANDOMPASSWORD` 的值为 `false`

> See Also  
[Sonatype community discuss](https://community.sonatype.com/t/is-that-possible-to-install-a-nexus-chart-with-a-given-initial-password/2962/1)  
[Source code](https://github.com/sonatype/nexus-public/blob/26aa3a2b54ffe524a5e7df164d847e268b624d21/components/nexus-security/src/main/java/org/sonatype/nexus/security/config/StaticSecurityConfigurationSource.groovy#L54)
