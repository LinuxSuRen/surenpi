---
title: "CrumIssuer"
description: "跨站请求伪造保护"
date: 2019-01-09T17:35:18+08:00
draft: false
toc: true
---

如果配置了跨站请求伪造保护的话，可以通过下面的命令来获取相应的 Token

`curl -u user:passwd  http:/localhost:8080/jenkins/crumbIssuer/api/json`

返回值类似于，而如果没有启用该配置的话会返回 `404`

```
{"_class":"hudson.security.csrf.DefaultCrumbIssuer","crumb":"6e78c5725ea0c522bdcb787d548465aa","crumbRequestField":"Jenkins-Crumb"}
```

使用该 Crumb 的方式如下：

`curl -u user:passwd http:/localhost:8080/jenkins/job/jobName/build -X POST --header "Jenkins-Crumb: 6e78c5725ea0c522bdcb787d548465aa"`

使用 Python 获取 Crumb 值：

```
export issuer=$(curl -u admin:$JENKINS_TOKEN $JENKINS_URL/crumbIssuer/api/json -s)
issuer=$(python -c "import json;import os;issuer=os.getenv('issuer');issuer=json.loads(issuer);print issuer['crumb']")
```