pipeline {
    agent {
        dockerfile {
            filename 'Dockerfile'
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

