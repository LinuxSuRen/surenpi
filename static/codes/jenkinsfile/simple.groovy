pipeline {
    agent any

    stages{
        stage("one") {
            steps{
                echo "one"
            }
        }
        stage("one") {
            steps{
                echo "one"
            }
        }
    }
}