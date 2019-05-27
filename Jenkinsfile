pipeline {
    agent {
        label "hugo"
    }

    options {
        disableConcurrentBuilds()
        // rateLimitBuilds([count: 6, durationName: 'hour', userBoost: true])
        quietPeriod 10
    }

    stages{
        stage("Build") {
            steps{
                hugo()
            }
        }
        
        // stage("Lunr Index") {
        //     steps {
        //         sh '''
        //         #!/bin/sh
        //         pwd
        //         node themes/gohugoTheme/src/hugo-lunr.js
        //         '''
        //     }
        // }

        stage("Github Publish") {
            steps {
                hugoGitPublish authorEmail: 'linuxsuren@gmail.com', authorName: 'suren',
                    commitLog: 'Auto commit by hugo-plugin.',
                    committerEmail: 'linuxsuren@gmail.com', committerName: 'suren',
                    credentialsId: 'a832798a-0513-4d45-a47a-95d152dc915a',
                    publishBranch: env.BRANCH_NAME,
                    targetUrl: 'https://github.com/linuxsuren/linuxsuren.github.io'
            }
        }
    }
    post { 
        always { 
            mail from: 'admin',
                to: 'rick',
                subject: 'SuRenPi 上线了结果：' + currentBuild.currentResult + '！', 
                body: '访问地址 http://surenpi.com/'
        }
    }
}
