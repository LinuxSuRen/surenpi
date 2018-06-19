---
title: Typora
---

Typora 支持用 Markdown 的方式来“画”流程图。

# 流程图

​```flow
st=>start: Start
op=>operation: Your Operation
cond=>condition: Yes or No?
e=>end

st->op->cond
cond(yes)->e
cond(no)->op
​```
