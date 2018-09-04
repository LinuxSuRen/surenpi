---
title: Jenkins JUnit
description: Jenkins JUnit
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