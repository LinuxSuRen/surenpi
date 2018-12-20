---
title: GPG
description: GPG
---

# 安装

Windows版本下载地址 https://www.gpg4win.org/download.html

Ubuntu下安装 `sudo apt-get install gpg -y`

Mac下安装 `brew install gpg` 或者从 https://gpgtools.org/ 下载

从 https://sourceforge.net/p/gpgosx/docu/Download/ 下载 gpg2

# 命令

`gpg --export-secret-key` 标识名 > 导出文件名

`gpg -o keyfilename --export-secret-keys KeyID`

`gpg -o keyfilename --export KeyID`

`gpg --import` 密钥文件

`gpg --list-keys` 列出密钥

`gpg --delete-secret-key hello`

`gpg --delete-key hello`

# 上传公钥

`gpg --keyserver hkp://pool.sks-keyservers.net --send-key 1234567`

# 接收公钥

`gpg --keyserver pool.sks-keyservers.net --recv-key 1234567`

`gpg --keyserver hkp://keyserver.ubuntu.com --search-keys 1234567` 搜索公钥

# 公钥服务器

`hkp://pool.sks-keyservers.net`

`hkp://subkeys.gpg.net`

`hkp://keyserver.ubuntu.com`

# 加解密

`gpg --recipient 1234567 --output a-encrypt.log --encrypt a.log` 加密文件

`gpg a-encrypt.log` 解密文件

`gpg --decrypt a-encrypt.log` 解密文件

# 签名

`gpg --sign a.log` 生成二进制签名文件 a.log.gpg

`gpg --clearsign a.log` 生成 ASCII 签名文件 a.log.asc

`gpg --detach-sign a.log` 生成单独的二进制签名文件 a.log.sig

`gpg --armor --detach-sign a.log` 生成单独的 ASCII 签名文件 a.log.asc

`gpg --verify a.log.asc` 验证签名

`gpg --local-user 1234567 --recipient 1234567 --armor --sign --encrypt a.log` 签名并加密

# 异常

```
gpg: cannot open /dev/tty': No such device or address
can't connect to /root/.gnupg/S.gpg-agent’: No such file or directory
gpg: can’t connect to the agent: IPC connect call failed
gpg: signing failed: No secret key
```

如果是用 Maven 插件的话，可以参考下面的配置：

```
<plugin>
<groupId>org.apache.maven.plugins</groupId>
<artifactId>maven-gpg-plugin</artifactId>
<version>1.6</version>
<executions>
    <execution>
    <id>sign-artifacts</id>
    <phase>verify</phase>
    <goals>
        <goal>sign</goal>
    </goals>
    <configuration>
        <gpgArguments>
        <arg>--pinentry-mode</arg>
        <arg>loopback</arg>
        <arg>--no-tty</arg>
        </gpgArguments>
        <interactive>false</interactive>
        <passphrase>1233</passphrase>
    </configuration>
    </execution>
</executions>
</plugin>
```