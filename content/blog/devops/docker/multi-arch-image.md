---
title: How to build a multi-arch docker image
description: How to build a multi-arch docker image
toc: true
---

## Install multi-arch binary files in one Dockerfile
```Dockerfile
FROM ubuntu

RUN ARCH= && ARCH_ALIAS= && dpkgArch="$(dpkg --print-architecture)" \
  && case "${dpkgArch##*-}" in \
    amd64) ARCH='amd64';; \
    ppc64el) ARCH='ppc64le';; \
    s390x) ARCH='s390x';; \
    arm64) ARCH='arm64';; \
    armhf) ARCH='armv7l';; \
    i386) ARCH='x86';; \
    *) echo "unsupported architecture"; exit 1 ;; \
  esac && \
  apt-get update && apt-get install wget -y && \
  wget https://cdn.jsdelivr.net/gh/jenkins-zh/jcli-repo/jcli-linux-$ARCH.tar.gz && \
  tar xzvf jcli-linux-$ARCH.tar.gz && rm -rf jcli-linux-$ARCH.tar.gz
```

Save those content into `Dockerfile`, then execute command `docker buildx build --platform linux/amd64,linux/arm64 . -f Dockerfile -t jcli:test`
