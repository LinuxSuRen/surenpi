---
title: tcpdump
description: dump traffic on a network
---

`tcpdump -i lo0 -s0 -v port 8080`

`tcpdump dst jenkins-zh.gitee.io -s0 -v`

`tcpdump src localhost`

`tcpdump -o lo0 -s0 -v port 8080`

`sudo tcpflow -p -i lo0 -c port 8080`