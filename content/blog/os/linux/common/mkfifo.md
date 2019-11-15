---
title: mkfifo
description: make fifos
---

nc -lkv localhost 9090 -c 'tee /dev/stderr | nc -v baidu.com 80 | tee /dev/stderr'