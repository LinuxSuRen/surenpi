---
title: "Remote IDE"
description: "Remote IDE"
date: 2024-07-22
draft: false
toc: true
---

IDE（Intelligent Development Environment） 对于软件开发者来说，是一个非常重要的工具。

## GitPod

## VSCode Server

[VSCode](https://code.visualstudio.com/Download) 除了有桌面版，还提供了 CLI 版本（支持 Windows、Linux、MacOS），可以作为服务启动，用户通过浏览器来访问并使用。

打开官网下载地址，选择 CLI 版本，解压后能看到 `code` 命令。执行下面的命令会启动一个 Web 服务：

```shell
code serve-web
```

你可以从终端输出中看到访问地址，例如：`http://127.0.0.1:8000`。对于在家庭网络、办公网络情况下，这种模式非常适合。但如果，你希望把开发环境暴露到公网，他还提供了隧道的方式，执行下面的命令，可以得到一个公网访问地址：

```shell
code tunnel
```

根据提示打开服务地址后，需要使用 GitHub 或 Microsoft 账号登录，然后就可以访问了。

这种模式，适合个人使用，难以满足团队使用的需求，例如：

* 无法为团队不同人员自动化地提供 IDE 实例
* 没有用户认证，无法保证使用安全
* 缺少实例资源（CUP、内存、硬盘等）控制

## OpenVSCode Server

[OpenVSCode Server](https://github.com/gitpod-io/openvscode-server) 是一个基于 [VSCode](https://code.visualstudio.com/) 的开源 IDE，由 [GitPod](https://www.gitpod.io/) 团队维护。


他默认提供了一个 Docker 镜像，可以直接使用 Docker 运行。

```shell
docker run -it --init -p 3000:3000 -v "$(pwd):/home/workspace:cached" gitpod/openvscode-server
```

你也可以通过自定义 Dockerfile 来构建更加符合自己需求的镜像。

```dockerfile
FROM gitpod/openvscode-server:latest

ENV OPENVSCODE_SERVER_ROOT="/home/.openvscode-server"
ENV OPENVSCODE="${OPENVSCODE_SERVER_ROOT}/bin/openvscode-server"

SHELL ["/bin/bash", "-c"]
RUN \
    # Direct download links to external .vsix not available on https://open-vsx.org/
    # The two links here are just used as example, they are actually available on https://open-vsx.org/
    urls=(\
        https://github.com/rust-lang/rust-analyzer/releases/download/2022-12-26/rust-analyzer-linux-x64.vsix \
        https://github.com/VSCodeVim/Vim/releases/download/v1.24.3/vim-1.24.3.vsix \
    )\
    # Create a tmp dir for downloading
    && tdir=/tmp/exts && mkdir -p "${tdir}" && cd "${tdir}" \
    # Download via wget from $urls array.
    && wget "${urls[@]}" && \
    # List the extensions in this array
    exts=(\
        # From https://open-vsx.org/ registry directly
        gitpod.gitpod-theme \
        # From filesystem, .vsix that we downloaded (using bash wildcard '*')
        "${tdir}"/* \
    )\
    # Install the $exts
    && for ext in "${exts[@]}"; do ${OPENVSCODE} --install-extension "${ext}"; done
```

可以看到，你可以通过自定义镜像的方式，预先安装插件、工具、语音 SDK 等等。

我在 https://github.com/LinuxSuRen/openvscode-server-images 这个项目中维护了：Golang、Java、NodeJS 以及包含常见前后端的全量镜像，欢迎各位使用。

但需要注意的是，OpenVSCode Server 会从 https://open-vsx.org/ 这里获取插件列表，和 VSCode 的插件中心地址不同。当然，格式上是完全兼容的。
