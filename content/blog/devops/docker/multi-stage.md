---
title: Docker multi-stage technology
description: Docker multi-stage technology
toc: true
---

We can build our application in Dockerfile. But we don't want the whole thing in our final docker image. Docker multi-stage technology could help us the reduce the size.
```Dockerfile
FROM golang:1.12 AS builder

WORKDIR /work
RUN git clone https://github.com/devops-ws/learn-pipeline-go && \
  cd learn-pipeline-go && go build -o go-server
RUN ls -hal

FROM alpine:3.10
COPY --from=builder /work/learn-pipeline-go/go-server .
CMD ['go-server']
```

After built this: `docker build . -t go-server`

then we can test it: `docker run go-server`
