# Proxy

如果执行 Docker 构建时，需要设置网络代理的化，可以在 `Dockerfile` 中添加下面的指令：

`ENV HTTP_PROXY http://192.168.1.1:1234`

上面的方法，在容器运行期，也会走代理。如果只是希望在构建期间设置网络代理的化，可以把 `ENV` 设置为空：

`ENV HTTP_PROXY ""`

但是，需要在构建时指定参数：`docker build --build-arg HTTP_PROXY=http://192.168.1.1:1234`
