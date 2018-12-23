pipeline {
    agent {
        docker {
            image 'alpine:3.7'
        }
    }

    stages{
        stage("one"){
            steps{
                echo "hello"
                sh 'uname -a'
            }
        }
    }
}