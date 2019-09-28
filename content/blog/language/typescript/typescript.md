---
title: TypeScript
description: TypeScript
---

`Error: Cannot find module '../lib/utils/unsupported.js'`

当遇到上面的错误时，可以试试下面的办法：
```
sudo rm -rf /usr/local/lib/node_modules/npm
brew reinstall node
```