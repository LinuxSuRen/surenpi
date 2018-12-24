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