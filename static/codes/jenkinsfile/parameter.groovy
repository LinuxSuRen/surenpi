pipeline {
    agent {
        label "master"
    }

    parameters {
        text defaultValue: 'Rick', description: '', name: 'name'
    }

    stages{
        stage("one"){
            steps{
                echo name
                sh 'echo ${name}'
            }
        }
    }
}