---
title: "让我们一起了解开源"
description: "致广大开源用户"
type: slide
date: 2022-08-27
draft: false
toc: true
---

class: center, middle
# 让我们一起了解开源

.black[[@LinuxSuRen](https://github.com/linuxsuren)]

---
name: catalog
class: center, middle

# 大纲
#### .black[[个人职业发展](#career)]
--
count: false
#### .black[[企业与开源](#company-open-source)]
--
count: false
#### .black[[开源是什么](#what-is-open-source)]
--
count: false
#### .black[[开源在哪里](#where-is-open-source)]
--
count: false
#### .black[[如何迈出第一步](#first-step-open-source)]
--
count: false
#### .black[[深入了解](#godeep)]

---
name: career
class: left, middle
.left-column[
## 职业发展
### 困境
]

.right-column.panding[
* 得不到充足的锻炼、成长
  * 复杂的、高难度的任务**始终**交由经验丰富的人做
* 技术止于商业诉求
  * 业务线生命周期短
  * **可用**优先，质量、性能按需分配
  * 技术选型**保守**，迭代**缓慢**
* 基于职位的决策机制，缺少创新
  * 谁拍板，谁负责，谁受益
* 人脉圈小，社交范围局限
  * 以同学、（前）同事为主要**社交圈**
]

.footnote.left.black[[目录](#catalog)]
---
class: left, middle、
.left-column[
## 职业发展
### .gray[困境]
### 突破
]

.right-column.panding[
* 跳出**舒适区**
* 有机会**精益求精**
* 持久**深耕**特定领域
* 结识不同领域的有志之士
]

---
name: company-open-source
.center[
# 企业与开源
]

.left-column-eq.panding[
.right[产品免费，损失营收]
.right[源代码无法作为核心资产]
.right[内部团队之间难以**复用**]
.right[大量修改开源代码，没有回馈上游，需要自行维护]
.right.red[潜在的合规风险]
]

.right-column-eq.panding[
.left[节省大量的产品宣传资源]
.left[通用、基础性功能免费，**企业级**需求付费]
.left[业务紧密相关的代码才是核心资产]
.left[企业修改的开源代码交由全球开发者维护]
.left[提升领域、行业内的技术品牌影响力]
]

.footnote.left.black[[目录](#catalog)]

---
.left-column-eq[
# 企业合规
#### 潜在风险
]

.right-column-eq.vertical-center[
.center[引用了 **GPL** 等商业不友好的开源项目]
.center[没有按规定指明依赖的开源项目]
]

.footnote.left.black[[目录](#catalog)]

---
.left-column-eq[
# 企业合规
#### .gray[潜在风险]
#### 应对策略
]

.right-column-eq.vertical-center[
.center[普及合规常识]
.center[建立开源项目清单]
.center[聘请（全职、兼职）开源法律顾问]
]

.footnote.left.black[[目录](#catalog)]

---
.left-column-eq[
# 企业合规
#### .gray[潜在风险]
#### .gray[应对策略]
#### 案例
]

* 2021年9月29日，**广州**知识产权法院就**罗盒网络科技**有限公司诉**玩友网络科技**有限公司等侵害开源软件著作权纠纷一案作出一审判决，判决侵权行为成立，被告须立即停止通过互联网平台提供含有被侵权开源代码的相关软件，并赔偿原告经济损失及维权合理开支共计**50万元**。
https://www.zhichanli.com/p/158105840

--

* TODO

--

* TODO

--

* TODO

.footnote.left.black[[目录](#catalog)]

---
# 企业参与开源

.left-column-eq[
### 默许
]

.right-column-eq.panding[
.left[对开源的认知停留在**免费**、公开的层面]
.left[能解决自身业务问题即可]
.left[项目类、定制类工程居多]
.left[技术氛围淡薄]
.left[初、中级程序员比重较大]
]

---
# 企业参与开源

.left-column-eq[
### .gray[默许]
### 鼓励
]

.footnote.left.black[通用的贡献指南 https://linuxsuren.github.io/open-source-best-practice/how-to-contribute.md/]

.left[核心产品**深度依赖**一些开源项目]
--
.center[==========]
.left[骨干：带头了解开源**文化**、参与**上游**，起到表率、示范效果]
.left[**新人**招聘：优先选择有开源经验的候选人]
--
.center[==========]
.left[不止步于解决问题本身，探索**通用**方案]
.left[提倡改进工作**效率**，预留更多时间去解决业务底层问题]
.left[提高代码**质量**要求：单元测试、静态检查、代码风格]
.left[**核心代码**严格执行代码 review 流程]
.left[提倡使用**英文**编写注释]
--
.center[==========]
.left[大型开源项目背后往往会有个社区]
.left[要熟悉、接纳社区的制度、流程，入“乡”随俗]
.left[通过社区会议、活动、issue 等与社区成员**互动**、交流]
.left[熟悉、遵循开源**协作**流程]

---
# 企业参与开源

.left-column-eq[
### .gray[默许]
### .gray[鼓励]
### 上游优先
]

对所依赖的开源项目有深度依赖，有能力修复缺陷、功能改进


--
对项目架构、设计提出了更高要求，可以相对容易地在 fork 和 upstream 版本之间切换


--
抽象到 upstream 的核心价值在于降本增效

---
# 企业参与开源

.left-column-eq[
### .gray[默许]
### .gray[鼓励]
### .gray[上游优先]
### 误区
]

**无差别**地激励、鼓励员工参与开源，不必要地浪费资源
--


公司层面缺少开源**战略指导**，政策摇摆不定、难以形成规模效应
--


**简单化**地理解、认识开源，以为只要把任务下发到研发即可
--


只要把代码放到 GitHub 上，就会有很多开发者来参与
--


开源协议只是君子协议，没有法律效益

---
# 企业发起开源

.left-column-eq[
### 常见形态
]

.center[大厂成熟解决方案开源，打造技术品牌效应]
.center[大厂为避免其发起的技术轻易被淘汰而以开源的方式创建生态]
.center[大厂竞争]
.center[大厂内部出于职位晋升而开源的 KPI 项目]
.center[特定技术领域，缺乏开源项目，抢占市场]
.center[开源创业公司，以开源作为主要的商业模式]
.center[小规模零星开源（试水、偶然）]

---
# 企业发起开源

.left-column-eq[
### .gray[常见形态]
### 商业模式
]

Open Core：社区版免费、企业版收费

SaaS：项目免费，提供云服务

双许可证（例如：MySQL）

建立生态，设立插件、应用市场准入

---
# 企业发起开源

.left-column-eq[
### .gray[常见形态]
### .gray[商业模式]
### 案例
]

TODO

---
name: what-is-open-source
# *开源是什么*

.center[
### Just for fun
### 开放源代码软件
### 新兴的产品获客形式
### 极具潜力的商业模式
### 开放式协作
]

.footnote.left.black[[目录](#catalog)]

---
name: osd
# The Open Source Definition

.left-column-eq.panding[
1. Free Redistribution
2. Source Code
3. Derived Works
4. Integrity of The Author's Source Code
5. No Discrimination Against Persons or Groups
6. No Discrimination Against Fields of Endeavor
7. Distribution of License
8. License Must Not Be Specific to a Product
9. License Must Not Restrict Other Software
10. License Must Be Technology-Neutral
]

.right-column-eq[
[![](https://opensource.org/sites/default/files/public/osi_keyhole_300X300_90ppi_0.png)](https://opensource.org/osd)
]

---
class: center

# 我认为的开“源”

--
### 以平等、公开、透明为基础的异步的协作方式

--
### 过程 > 结果


--
### 源，是源头活水的源


--
### 倡导互相成就，而不是控制

---
name: where-is-open-source
background-image: url(FishDiagram.png)
# 开源在哪里

.footnote.left.black[[目录](#catalog)]

---
name: first-step-open-source
# 如何迈出开源第一步

* 找一个**理由**，让自己至少能坚持**一年**
* 设定一个具有一定挑战性的**目标**
* 体验、融入是重点，**抛弃功利**的考虑

.footnote.left.black[[目录](#catalog)]

---
class: center, middle

## 具体有哪些方向呢？
--

## 从做好一名用户开始
--

## 本地化文档
--

## good-first-issue
--

## 参与社区活动
--

## 了解开源文化

.footnote.left.black[[目录](#catalog)]

---
.left-column[
## 深入了解
### 文章
]

.center[
#### [夜天之书](https://mp.weixin.qq.com/s/CD9hiunqOnwaKqsuaPDFdg)
#### [卫sir说](https://mp.weixin.qq.com/s/-h4qCCPnxfkkPGeMUNMOpw)
]

---
.left-column[
## 深入了解
### .gray[文章]
### 视频
]

.center[
#### [Koala聊开源](https://space.bilibili.com/489667127)
#### [OpenTEKr](https://space.bilibili.com/1177021647)
#### [开源之夏](https://space.bilibili.com/578074510)
]

.footnote.left.black[[目录](#catalog)]
---
name: godeep
.left-column[
## 深入了解
### .gray[文章]
### .gray[视频]
### 播客
]

.center[
#### [ALC Beijing Podcast](https://www.ximalaya.com/keji/37853515/)
#### [OpenTEKr 大话开源](https://www.ximalaya.com/album/55050616)
#### [开源社·陪你看开源](https://www.xiaoyuzhoufm.com/podcast/6286388483ddae2d9e72a79d)
#### [开源雨林—大咖访谈](https://www.xiaoyuzhoufm.com/podcast/62d5329cc606feab61b21815)
#### [开源之道Talking](https://www.ximalaya.com/album/32915040)
#### [开源面对面](https://www.ximalaya.com/album/53320813)
]

.footnote.left.black[[目录](#catalog)]
---
class: center, middle

# 欢迎走进开源的世界
