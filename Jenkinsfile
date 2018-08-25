pipeline {
    agent {
        label "hugo"
    }

    stages{
        stage("Clone") {
            steps {
                checkout scm
            }
        }

        stage("Build") {
            steps{
                hugo baseUrl: 'https://surenpi.gitee.io/blog', destination: 'gitee', hugoHome: ''
            }
        }
        
        stage("Lunr Index") {
            steps {
                sh '''
                #!/bin/sh
                pwd
                node themes/gohugoTheme/src/hugo-lunr.js
                '''
            }
        }

        stage("Gitee Publish") {
            steps {
                hugoGitPublsh authorEmail: '', authorName: '', commitLog: 'Auto commit by hugo-plugin.', committerEmail: 'zxjlwt@126.com', committerName: 'suren', credentialsId: 'a832798a-0513-4d45-a47a-95d152dc915a', publishBranch: 'master', publishDir: 'gitee', targetUrl: 'https://gitee.com/surenpi/blog'
            }
        }

        stage("Github Publish") {
            steps {
                hugoGitSubmodulePublsh authorEmail: '', authorName: '', commitLog: 'Auto commit by hugo-plugin.', committerEmail: 'zxjlwt@126.com', committerName: 'suren', credentialsId: '5e1609b1-c9a8-4bb2-bc9c-1a2e0eda4856', publishBranch: 'master', publishDir: 'public'
            }
        }
    }
}