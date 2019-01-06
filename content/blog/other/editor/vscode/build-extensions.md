---
title: VS Code 扩展
description: VS Code 扩展
toc: true
draft: true
---

你需要有 NodeJs 和 Git，然后通过下面的命令安装 Yeoman 和 VS-Code 生成器：

`npm install -g yo generator-code`

执行命令 `yo code` 选择一个模板工程，可以快速地生成一个插件工程。更多的特性可以查看 {{< exref "API 文档" "https://code.visualstudio.com/api" >}}

# 打包

打包 `vsce package`

# 发布

首先，需要{{< exref "创建一个账户" "https://code.visualstudio.com/api/working-with-extensions/publishing-extension" >}}。然后，获取“个人访问 Token”。https://dev.azure.com/vscode

安装工具 `npm install -g vsce`

`vsce create-publisher (publisher name)`

`vsce login (publisher name)`

`vsce publish -p <token>`

# 本地安装

你可以通过下面的命令或者在 VS-Code 的插件安装界面中手动安装：

`code --install-extension my-extension-0.0.1.vsix`