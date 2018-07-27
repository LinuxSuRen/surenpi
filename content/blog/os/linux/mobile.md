---
title: Connect to Mobile
description: How to connect mobile device
---

# iPad

安装依赖的工具：

```
add-apt-repository ppa:pmcenery/ppa
apt-get update
apt-get install libimobiledevice-utils
apt-get install ideviceinstaller
```

安装应用 `ideviceinstaller -i xxx.ipa`

卸载应用 `ideviceinstaller -U [bundleID]`

查看系统日志 `idevicesyslog`

查看当前已连接的设备的UUID `idevice_id --list`

截图 `idevicescreenshot`

查看设备信息 `videviceinfo`

获取设备时间 `idevicedate`

参考 https://github.com/LinuxSuRen/libimobiledevice