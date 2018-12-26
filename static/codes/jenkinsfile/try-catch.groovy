node {
    stage("one"){
        try {
            echo "hello"
        }catch(error){
            echo error.getMessage()
        }finally{
            echo "finally"
        }
    }
}