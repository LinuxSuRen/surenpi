pipeline {
    agent any

    stages{
        stage("one") {
            steps{
                echo "one"
            }
        }
        stage("two") {
            steps{
                echo "two"
            }
        }
    }
}