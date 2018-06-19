---
title: 安装 Kubernetes
---

启用 Docker 服务 `systemctl enable docker.service`

安装 go 依赖 `go get github.com/kubernetes-incubator/cri-tools/cmd/crictl`

安装客户端

```
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF
setenforce 0
yum install -y kubelet kubeadm kubectl
systemctl enable kubelet && systemctl start kubelet
```

设置 kubelet `sed -i "s/cgroup-driver=systemd/cgroup-driver=cgroupfs/g" /etc/systemd/system/kubelet.service.d/10-kubeadm.conf`

重启 kubelet `systemctl daemon-reload && systemctl restart kubelet`

禁用Swap `swapoff -a`

安装 master `kubeadm init`

`export LC_ALL=en_US.utf-8`
`export LANG=en_US.utf-8`

```
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/etcd-amd64:3.0.17
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/pause-amd64:3.0
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/kube-proxy-amd64:v1.7.2
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/kube-scheduler-amd64:v1.7.2
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/kube-controller-manager-amd64:v1.7.2
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/kube-apiserver-amd64:v1.7.2
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/kubernetes-dashboard-amd64:v1.6.1
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-sidecar-amd64:1.14.4
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-kube-dns-amd64:1.14.4
docker pull registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-dnsmasq-nanny-amd64:1.14.4

docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/etcd-amd64:3.0.17    gcr.io/google_containers/etcd-amd64:3.0.17
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/pause-amd64:3.0     gcr.io/google_containers/pause-amd64:3.0
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/kube-proxy-amd64:v1.7.2   gcr.io/google_containers/kube-proxy-amd64:v1.7.2
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/kube-scheduler-amd64:v1.7.2  gcr.io/google_containers/kube-scheduler-amd64:v1.7.2
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/kube-controller-manager-amd64:v1.7.2 gcr.io/google_containers/kube-controller-manager-amd64:v1.7.2
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/kube-apiserver-amd64:v1.7.2 gcr.io/google_containers/kube-apiserver-amd64:v1.7.2
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/kubernetes-dashboard-amd64:v1.6.1 gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.1
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-sidecar-amd64:1.14.4 gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.4
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-kube-dns-amd64:1.14.4 gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.4
docker tag registry.cn-hangzhou.aliyuncs.com/surenpi/k8s-dns-dnsmasq-nanny-amd64:1.14.4 gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.4
```

https://www.cnblogs.com/ericnie/p/7749588.html
