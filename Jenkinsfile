pipeline {
    agent {
        label "hugo"
    }

    stages{
        stage("clone") {
            steps{
                checkout scm
            }
        }
        
        stage("Build") {
            steps{
                hugo baseUrl: '', destination: 'github'
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
        
        stage("branch"){
            steps{
                echo env.BRANCH_NAME
                echo env.CHANGE_ID
                echo env.CHANGE_TARGET
                echo env.CHANGE_TITLE
                echo 'hello'
                echo 'good'
                echo 'pipeline'
            }
        }

        stage("Github Publish") {
            steps {
                hugoGitPublish authorEmail: 'linuxsuren@gmail.com', authorName: 'suren',
                    commitLog: 'Auto commit by hugo-plugin.',
                    committerEmail: 'linuxsuren@gmail.com', committerName: 'suren',
                    credentialsId: '84862d34-5d17-4521-a37e-4200b8c99afe',
                    publishBranch: 'hello', publishDir: 'github',
                    targetUrl: 'https://github.com/linuxsuren/linuxsuren.github.io'
            }
        }
    }
}
