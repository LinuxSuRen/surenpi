# 流水线

本文介绍 Jenkins 的脚本式流水线。

## 触发器

```groovy
node('local') {
	echo 'hello'
}

properties([
	buildDiscarder(
		logRotator(
			artifactDaysToKeepStr: '',
			artifactNumToKeepStr: '',
			daysToKeepStr: '5',
			numToKeepStr: '10'
		)
	),
	pipelineTriggers([
		cron('H 3,12,17 * * *'),
		scm('H 3,12,17 * * *')
	])
])
```

## 参数化执行

```groovy
properties([[$class: 'JobRestrictionProperty'],
	parameters([run(description: '',
		filter: 'ALL', 
		name: 'Name', 
		projectName: 'Project')]),
	pipelineTriggers([])]
)
```
