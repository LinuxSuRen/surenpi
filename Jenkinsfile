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
                hugoGitSubmodulePublsh authorEmail: '', authorName: '', commitLog: 'Auto commit by hugo-plugin.', committerEmail: 'zxjlwt@126.com', committerName: 'suren', credentialsId: '5e1609b1-c9a8-4bb2-bc9c-1a2e0eda4856', publishBranch: 'master', publishDir: 'public'
            }
        }
    }
}