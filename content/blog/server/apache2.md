查看工作方式 `apachectl -V | grep -i mpm`

配置文件如下：

```
<IfModule mpm_prefork_module>
    StartServers         10         # 启动时进程数
    MinSpareServers      5          # 最小空闲进程数
    MaxSpareServers      10         # 最大空闲进程数
    MaxRequestWorkers    100        # 最大并发进程数
    MaxConnectionsPerChild   10000  # 最大连接数限制
</IfModule>
```
