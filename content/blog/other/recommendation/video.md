---
title: "Video"
description: "Video"
date: 2019-04-15T12:18:24+08:00
draft: false
toc: true
---

## 格式转换

{{< exref "这个网站" "https://www.onlinevideoconverter.com" >}}可以把 {{< exref "Youtube" "https://www.youtube.com" >}} 上的视频转换为 `mp4` 格式并提供下载链接。

https://youtubemp4.to/

### 下载

我们可以通过一个命令行工具来从 [Youtube](https://youtube.com) 上下载视频，首先，使用 [brew](/blog/os/mac/brew) 来安装命令行：

`brew install youtube-dl`

然后，就可以开始下载了：

`youtube-dl "https://www.youtube.com/watch?v=EtlhC3vw8pA" --proxy 127.0.0.1:1087`
