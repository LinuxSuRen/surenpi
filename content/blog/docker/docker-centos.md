---
title: 在 CentOS 安装 Docker
---

```
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine
```

安装依赖：
```
yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
```

添加库
```
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```

查看版本 `yum list docker-ce --showduplicates | sort -r`

安装 Docker `yum install docker-ce` 安装执行版本 `yum install docker-ce17.12.0.ce-1.el7.centos`