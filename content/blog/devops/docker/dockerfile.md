# 基础镜像

如果没有特殊要求的话，建议尽可能选择较小的基础镜像。例如：[alpine](https://hub.docker.com/_/alpine/)。您可以参考下面的 Dockerfile 写法：

`FROM alpine:3.8`

# Proxy

如果执行 Docker 构建时，需要设置网络代理的化，可以在 `Dockerfile` 中添加下面的指令：

`ENV HTTP_PROXY http://192.168.1.1:1234`

上面的方法，在容器运行期，也会走代理。如果只是希望在构建期间设置网络代理的化，可以把 `ENV` 设置为空：

`ENV HTTP_PROXY ""`

但是，需要在构建时指定参数：`docker build --build-arg HTTP_PROXY=http://192.168.1.1:1234`

# EXPOSE

指定容器要暴露的端口，例如：`EXPOSE 9898`。在执行容器时，可以通过参数来自动暴露给一个本地的随机端口，例如：`docker run -P sample`

# RUN

RUN 为 Docker 镜像在构建时执行的命令

# CMD

CMD 为 Docker 镜像在运行时的执行入口，例如：`CMD ["echo", "hello"]`