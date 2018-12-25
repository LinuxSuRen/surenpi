pipeline {
    agent {
        label "master"
    }

    stages {
        stage('one') {
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