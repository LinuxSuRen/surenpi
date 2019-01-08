pipeline {
    agent {
        label 'master'
    }

    stages {
        stage('one') {
            when {
                branch 'master'
            }
            steps {
                echo 'Hello World'
            }
        }
    }
    post { 
        always { 
            echo 'I will always say Hello again!'
        }
    }
}