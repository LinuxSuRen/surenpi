---
title: Jenkins JUnit
description: Jenkins JUnit
keywords: jenkins junit
draft: true
---

# Assert

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