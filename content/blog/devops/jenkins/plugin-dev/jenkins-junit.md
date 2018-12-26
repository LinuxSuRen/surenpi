---
title: Jenkins JUnit
description: Jenkins JUnit
keywords: jenkins junit
draft: false
toc: true
---

# 断言

```
@Test
public void basic() {
    assertThat("abc", StringContains.containsString("a"));
    assertThat("abc", Matchers.not(StringContains.containsString("d")));
    assertTrue(true);
    assertFalse(false);
    assertNull(null);
    assertNotNull("");
    assertEquals("a", "a");
}
```

# 数据源

注解 `@LocalData` 会按照一定的规则去加载 `src/test/resources` 中的配置文件。

例如：该注解加载类 `org.test.TestAgent` 的方法 `test` 上的话，就会去加载目录 `src/test/resources/org/test/TestAgent/test` 中的配置文件。另外，它还有一个 `value` 参数可以替代使用方法名称。

# 框架

Jenkins 的测试框架为 [acceptance-test-harness](https://github.com/jenkinsci/acceptance-test-harness)。

在 `pom.xml` 中可以修改框架的版本，如下所示：

```
<properties><jenkins-test-harness.version>2.28</jenkins-test-harness.version></properties>
```

通过下面的方式修改 Jenkins 运行测试的版本：

```
<properties><jenkins.version>2.150.1</jenkins.version></properties>
```

# 环境变量

如果需要修改运行测试时的 Jenkins 版本，可以在启动 JVM 时指定参数 `-Djth.jenkins-war.path=/tmp/jenkins.war`