# 切割文件

指定每个文件的大小和名称前缀 `split -b 10m  minikube-v0.26.0.iso minikube-v0.26.0.iso-part-`

合并文件，可以使用命令 `cat minikube-v0.26.0.iso-part-* > minikube.iso`
