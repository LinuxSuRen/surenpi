// dependecy plugin is https://plugins.jenkins.io/phoenix-autotest
pipeline {
    agent {
        label "master"
    }
    
    stages {
      stage('test') {
          steps{
              script {
                  http url: "http://baidu.com"
                  http url: "http://baidu.com", responseBody: "body.txt"
                  
                  archiveArtifacts "body.txt"
                  
                  def quy = 'https://www.baidu.com/s?wd=' + URLEncoder.encode('a b', 'UTF-8')
                  echo quy
                  http url: quy
              }
          }
      }
    }
}
