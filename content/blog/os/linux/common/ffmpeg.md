---
title: ffmpeg
description: ffmpeg
toc: true
keywords: ffmpeg watermark
---

## 静态水印

`ffmpeg -y -i movieFile -vf "movie=watermark.png [watermark]; [in][watermark]  overlay=x:y:opaque [out]" movieWatermark.mp4`

`ffmpeg -i jenkins.mov -i watermark.png -filter_complex "overlay=10:10" jenkins.mp4`

## 动态水印

`ffmpeg -y -i movieFile -ignore_loop 0 -i gif.gif -ss start -t duration  -filter_complex overlay=0:0:1 movieGifFile`

http://ksloan.net/watermarking-videos-from-the-command-line-using-ffmpeg-filters/

## 合并视频

把视频文件按照顺序写入 `files.txt`，然后执行命令：

`ffmpeg -f concat -i files.txt -c copy jenkins-basic-install.mp4`

## 格式转换

通过下面的命令，可以实现视频格式的转换：

`ffmpeg -i demo.mkv -codec copy demo.mp4`

## 修改分辨率

通过下面的命令可以实现视频文件的分辨率修改：

`ffmpeg -i demo.mp4 -vf scale=1920:1080 demo.high.mp4`

## 图片缩放

假设图片 `input.jpg` 的尺寸为：320⨉207

`ffmpeg -i input.jpg -vf scale=320:240 output_320x240.png` 强制设置图片的尺寸

`ffmpeg -i input.jpg -vf scale=320:-1 output_320.png` 保留缩放比（输入和输出文件不能相同），输出的图片尺寸为：320⨉207

`ffmpeg -i input.jpg -vf scale=320:-2 output_320.png` 这样可以让图片的宽、高都被 `n` 整除，输出的图片尺寸为：320⨉206

`ffmpeg -i input.jpg -vf scale=iw*2:ih input_double_width.png` 使用变量设置宽、高，`iw` 代表输入的宽度，`ih` 代表输入的高度

`ffmpeg -i input.jpg -vf "scale=iw/2:ih/2" input_half_size.png` 使得图片的宽、高为原来的一半

`ffmpeg -i input.jpg -vf "scale='min(320,iw)':'min(240,ih)'" input_not_upscaled.png`

`ffmpeg -i input.jpg -vf scale=w=320:h=240:force_original_aspect_ratio=decrease output_320.png`

`ffmpeg -i input.jpg -vf "scale=320:240:force_original_aspect_ratio=decrease,pad=320:240:(ow-iw)/2:(oh-ih)/2" output_320_padding.png`
