---
title: Linux Service
description: Linux Service
---

```
#!/bin/sh
### BEGIN INIT INFO
# Provides: suren
# Required-Start: $network $remote_fs $syslog
# Required-Stop: $network $remote_fs $syslog
# Should-Start: jenkins
# Should-Stop: jenkins
# Default-Start: 2 3 4 5
# Default-Stop: 0 1 6
# Short-Description: http://surenpi.com
# Description: http://surenpi.com
### END INIT INFO
 
case $1 in
    start)
        echo "start"
        nohup java -jar /var/www/jenkins/jenkins.war --httpPort=9099 > /var/www/jenkins/jenkins.log 2> /var/www/    jenkins/jenkins-error.log &
        ;;  
    stop)
        echo "stop"
        ;;  
    *)  
        echo "Usage: $0 (start|stop)"
        ;;  
esac
```