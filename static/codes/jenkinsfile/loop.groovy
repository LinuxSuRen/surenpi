node('suren') {
    def dev_path = '/opt/suren/bin'
    def services = [
        [
            'name': 'admin',
            'project': 'admin',
            'port': '7002',
            'jarName': 'admin'
        ]
    ];
    
    stage('Copy Artifact') {
        for(service in services){
            step([$class: 'CopyArtifact', fingerprintArtifacts: true, flatten: true,
                projectName: service.project,
                selector: [$class: 'StatusBuildSelector', stable: false],
                target: dev_path + '/' + service.name
            ])
        }
    }
    
    stage('Stop Service') {
        for(service in services){
           sh 'fuser -n tcp -k ' + service.port + ' > redirection &'
        }
    }
    
    stage('Start Service') {
        for(service in services){
            sh 'cd ' + pass_bin + '/' + service.name + ' && nohup nice java -server -Xms128m -Xmx384m \
                -jar ' + service.jarName + '.jar \
                --server.port=' + service.port + ' $> initServer.log 2>&1 &'
        }
    }
}