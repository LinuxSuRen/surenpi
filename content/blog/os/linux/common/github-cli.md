---
title: GitHub CLI
description: GitHub CLI
toc: true
---

[GitHub CLI](https://github.com/cli/cli) is a very powerful too. It provides almost all the ability that you might need with GitHub. 
And it also allows you send a GitHub API request easily.

## Get started
Read [the official document](https://cli.github.com/), then install it. Or, you could install it with [the following command](https://github.com/LinuxSuRen/http-downloader):

```shell
hd i cli/cli
```

## Get Status list of a PR
First, you need to get the sha of a PR:

```shell
gh api repos/kubesphere/ks-devops/pulls/818 | jq .head.sha
```

then, you could get the status list of a PR:
```shell
gh api repos/kubesphere/ks-devops/statuses/e5dff7a8b45320ffa0a7869d8c8458151687074e
```
