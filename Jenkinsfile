pipeline {
    agent {
        label "hugo"
    }

    stages{
        stage("Build") {
            steps{
                hugo baseUrl: '', destination: ''
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
                    credentialsId: '84862d34-5d17-4521-a37e-4200b8c99afe',
                    publishBranch: 'master', publishDir: 'github',
                    targetUrl: 'https://github.com/linuxsuren/linuxsuren.github.io'
            }
        }
    }
}
