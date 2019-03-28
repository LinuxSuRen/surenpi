---
title: ffmpeg
description: ffmpeg
keywords: ffmpeg watermark
---

# 静态水印

`ffmpeg -y -i movieFile -vf "movie=watermark.png [watermark]; [in][watermark]  overlay=x:y:opaque [out]" movieWatermark.mp4`

`ffmpeg -i jenkins.mov -i watermark.png -filter_complex "overlay=10:10" jenkins.mp4`

# 动态水印

`ffmpeg -y -i movieFile -ignore_loop 0 -i gif.gif -ss start -t duration  -filter_complex overlay=0:0:1 movieGifFile`

http://ksloan.net/watermarking-videos-from-the-command-line-using-ffmpeg-filters/

# 合并视频

把视频文件按照顺序写入 `files.txt`，然后执行命令：

`ffmpeg -f concat -i files.txt -c copy jenkins-basic-install.mp4`

# 格式转换

通过下面的命令，可以实现视频格式的转换：

`ffmpeg -i demo.mkv -codec copy demo.mp4`