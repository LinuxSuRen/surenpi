---
title: Template
description: Go Template
toc: true
---

# Action

1.条件语句
2.迭代
3.封装
4.引用

## 条件语句

```
{{ if arg }}
  code
{{ end }}

{{ if arg }}
  code-1
{{ else }}
  code-2
{{ end }}
```

## 迭代

```
{{ range . }}
{{ end }}
```

```
{{ range $key, $value := . }}
  {{ $key }} {{ $value }}
{{ end }}
```

## 数组

```
{{ index "abc" 0 }}
```
