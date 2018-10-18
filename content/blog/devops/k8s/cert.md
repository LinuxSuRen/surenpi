---
title: K8S 证书
description: K8S 证书
---

根证书公钥(ca.pem)与私钥(ca-key.pem)

API Server 公钥和私钥

集群管理员公钥和私钥

从节点公钥和私钥

# 根证书

```
openssl genrsa -out ca-key.pem 2048
openssl req -x509 -new -nodes -key ca-key.pem -days 10000 -out ca.pem -subj "/CN=kube-ca"
```

# API Server 证书

创建 `openssl.cnf`

```
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = kubernetes
DNS.2 = kubernetes.default
DNS.3 = kubernetes.default.svc
DNS.4 = kubernetes.default.svc.cluster.local
IP.1 = ${K8S_SERVICE_IP}
IP.2 = ${MASTER_IPV4}
```

```
openssl genrsa -out apiserver-key.pem 2048
openssl req -new -key apiserver-key.pem -out apiserver.csr -subj "/CN=kube-apiserver" -config openssl.cnf
openssl x509 -req -in apiserver.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out apiserver.pem -days 365 -extensions v3_req -extfile openssl.cnf
```

一般生成的根证书(ca-key.pem, ca.pem)与 apiserver 证书(apiserver-key.pem,apiserver.pem)放置在 Master 节点的 `/etc/kubernetes/ssl/` 路径下

# Ref

https://blog.csdn.net/shenshouer/article/details/53035948