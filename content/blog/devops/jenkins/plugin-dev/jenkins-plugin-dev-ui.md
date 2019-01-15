---
title: Jenkins UI
description: Jenkins UI 介绍
keywords: jenkins plugin jelly
toc: true
---

# 按钮

```
<j:jelly  xmlns:f="/lib/form"></j:jelly>
<f:validateButton title="Test Connection" method="verifyConnect" with="server,credentialsId" progress="Testing..."></f:validateButton>
<f:textbox/>
<f:checkbox/>
```

# 下拉框

在 `jelly` 文件添加如下标签：

```
<f:entry field="method" title="${%Method}">
    <f:select/>
</f:entry>
```

在对应的后端类的 `DescriptorImpl` 中添加如下方法：

```
public ListBoxModel doFillMethodItems() {
    ListBoxModel listBoxModel = new ListBoxModel();
    listBoxModel.add(HttpPost.METHOD_NAME);
    listBoxModel.add(HttpGet.METHOD_NAME);
    return listBoxModel;
}
```

细心的童鞋肯定能看出来，`doFillMethodItems` 中间的 `Method` 就对应页面上的字段名称。

# 校验

如果你的表单里有一个字段 `crontab` 希望能添加数据格式校验的话，可以在对应类的 `DescriptorImpl` 中添加方法来实现：

```
<f:entry title="${%Crontab}" field="crontab">
<f:textbox/>
</f:entry>
```

```
public FormValidation doCheckCrontab(@QueryParameter String value, @AncestorInPath Item item) {
    try {
        CronTabList ctl = CronTabList.create(fixNull(value), item != null ? Hash.from(item.getFullName()) : null);
        Collection<FormValidation> validations = new ArrayList<>();
        updateValidationsForSanity(validations, ctl);
        updateValidationsForNextRun(validations, ctl);
        return FormValidation.aggregate(validations);
    } catch (ANTLRException e) {
        if (value.trim().indexOf('\n')==-1 && value.contains("**"))
            return FormValidation.error(Messages.TimerTrigger_MissingWhitespace());
        return FormValidation.error(e.getMessage());
    }
}
```

# 凭据

```
<j:jelly xmlns:c="/lib/credentials">
<c:select/>
</j:jelly>
```

# 高级

当有些字段具有默认值，通常情况下也无需修改时，你可以把这些字段隐藏起来（只有点击高级选项的时候才会展示）。

```
<f:advanced>
    <f:entry field="format" title="Format">
        <f:textbox/>
    </f:entry>
</f:advanced>
```

# 静态资源文件

当我们需要在 Jenkins 插件中添加静态文件时，可以把静态访问放到 `webapp` 下。然后，通过类似下面的地址访问：

`http://localhost:8080/jenkins/plugin/sample-plugin-id/test.css`

# 帮助信息

如果想给你的页面添加帮助信息的话，可以[参考这里]({{< ref "jenkins-help.md" >}})。
