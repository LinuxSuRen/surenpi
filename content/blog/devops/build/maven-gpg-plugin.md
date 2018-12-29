---
title: maven-gpg-plugin
description: maven-gpg-plugin
---

[INFO] --- maven-gpg-plugin:1.5:sign (sign-artifacts) @ kubernetes-model-generator ---
gpg: directory '/home/jenkins/.gnupg' created
gpg: keybox '/home/jenkins/.gnupg/pubring.kbx' created
gpg: no default secret key: No secret key
gpg: signing failed: No secret key

`gpg: signing failed: Inappropriate ioctl for device`

https://github.com/pstadler/keybase-gpg-github/issues/11
https://github.com/jenkinsci/remoting/pull/280
https://gist.github.com/michaelajr/ff4693bce9fc20e5200b34683aa4ba51