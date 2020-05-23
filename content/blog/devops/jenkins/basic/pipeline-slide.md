---
title: "流水线众生相之 Jenkins 对话 Tekton"
description: "分别介绍 Jenkins 和 Tekton 的流水线特性"
type: slide
date: 2020-05-21
draft: false
toc: true
---

class: center, middle, inverse
# 流水线众生相之
Jenkins 对话 Tekton

---
class: inverse
## About Me

.left-column-eq[
* [CDF Ambassador](https://cd.foundation/ambassador-program-overview-application/community-ambassador-cohort20/)
* [GSoC mentor](https://www.jenkins.io/projects/gsoc/)
* [Jenkins Press Contact in China](https://www.jenkins.io/press/)
* [Jenkins Chinese Localization](https://www.jenkins.io/sigs/chinese-localization/)
]

.right-column-eq[
* [Jenkins CLI](https://github.com/jenkins-zh/jenkins-cli)
* [Jenkins Chinese Simplified Plugin](https://github.com/jenkinsci/localization-zh-cn-plugin)
* [Jenkins Distribution Formulas](https://github.com/jenkins-zh/jenkins-formulas)
* [Wechat backend](https://github.com/jenkins-zh/wechat-backend)
]

---
class: inverse
## 流水线 vs 手工

.left-column-eq[
* 易执行，可复用
* 代码
* 可视化
* 统计信息
* review
]

.right-column-eq[
* 繁琐，易出错
* 文档 + 脚本
* 依赖个人
]

---
class: inverse
## 相关产品

.left-column-eq[
.img-max-width-200[
![github](/images/logo/github.jpeg)
]
.img-max-width-200[
![jenkins](/images/logo/jenkins.png)
]
.img-max-width-200[
![travisci](/images/logo/travisci.png)
]
]

.right-column-eq[
.img-max-width-200[
![tekton](/images/logo/tekton.png)
]
.img-max-width-200[
![gitlab](/images/logo/gitlab.png)
]
.img-max-width-200[
![circleci](/images/logo/circleci.png)
]
]

---
class: inverse
## Jenkins vs Tekton

.left-column-eq[
* 传统部署、容器部署
* 非常成熟、丰富的插件
* UI、CLI 功能完整
* Jenkinsfile、groovy
* 动态构建节点
* 单点故障
* 本地文件系统存储
]

.right-column-eq[
* Kubernetes
* v1beta1
* CRD、yaml
* UI Dashboard
* Task、Pipeline 使用灵活
]

---
class: center, middle
## Jenkins

Jenkins automation server

---
class: inverse
## Jenkinsfile

```Jenkinsfile
pipeline {
    agent none
    stages {
        stage('Example Build') {
            steps {
                echo 'Hello World'
            }
        }
        stage('Example Deploy') {
            when {
                beforeOptions true
                branch 'testing'
            }
            options {
                lock label: 'testing-deploy-envs', quantity: 1, variable: 'deployEnv'
            }
            steps {
                echo "Deploying to ${deployEnv}"
            }
        }
    }
}
```

---
class: inverse
## 专业的组件做专业的事情

* 凭据 -> k8s
    * 备份恢复
* 系统配置 -> git
    * 配置修改的审核
* 数据文件 -> 共享存储
* 用户认证 -> oidc(ladp)

---
class: inverse
## Jenkins 节点标签

* 与节点连接类型结偶
* 每个节点可以有多个标签
* 名称命名要有业务意义
    * linux、java、mem-high
* 逻辑关系
    * linux & java
* 有哪些标签可以用？

---
class: inverse
## Jenkins 节点管理

* 当好 master
* 推荐使用动态节点
* 一定要有 limit
* k8s 节点的限制
* k8s 节点的伸缩

---
class: inverse
## 构建优化

* 避免串行构建
* 内存、CPU 资源合理分配
* Docker registry mirror
* Maven mirror
* HTTP proxy
* 构建状态监控

---
class: center, middle
## Tekton

A K8s-native Pipeline resource.

---
class: inverse
## Tekton

.left-column-eq[
### Task
```
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: echo-hello-world
spec:
  steps:
    - name: echo
      image: ubuntu
      command:
        - echo
      args:
        - "Hello World"
```
]

.right-column-eq[
### Task Run
```
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  name: echo-hello-world-task-run
spec:
  taskRef:
    name: echo-hello-world
```
]

---
class: inverse
## Tekton Pipeline

.left-column-eq[
```
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: tutorial-pipeline
spec:
  resources:
    - name: source-repo
      type: git
    - name: web-image
      type: image
  tasks:
    - name: build-skaffold-web
      taskRef:
        name: build-docker-image-from-git-source
      params:
        - name: pathToDockerFile
          value: Dockerfile
        - name: pathToContext
          value: /workspace/docker-source/examples/microservices/leeroy-web #configure: may change according to your source
      resources:
        inputs:
          - name: docker-source
            resource: source-repo
        outputs:
          - name: builtImage
            resource: web-image
```
]

.right-column-eq[
```
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
metadata:
  name: tutorial-pipeline-run-1
spec:
  serviceAccountName: tutorial-service
  pipelineRef:
    name: tutorial-pipeline
  resources:
    - name: source-repo
      resourceRef:
        name: skaffold-git
    - name: web-image
      resourceRef:
        name: skaffold-image-leeroy-web
```
]

---
class: inverse
## Tekton Resources

```
apiVersion: tekton.dev/v1alpha1
kind: PipelineResource
metadata:
  name: skaffold-image-leeroy-web
spec:
  type: image
  params:
    - name: url
      value: gcr.io/<use your project>/leeroy-web
```

---
class: inverse
## Installation

* Pipeline
* Dashboard
* Trigger
* [CLI](https://github.com/tektoncd/cli)

```
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
kubectl apply --filename https://github.com/tektoncd/dashboard/releases/download/v0.6.1.4/tekton-dashboard-release.yaml
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

brew tap tektoncd/tools
brew install tektoncd/tools/tektoncd-cli
```

---
class: center, middle

## Thanks

.img-max-height-200[
![jenkins-wechat](/images/qrcode/jenkins-wechat-banner.png)
![jenkins-wechat](/images/qrcode/jenkins-assistant.jpeg)
]