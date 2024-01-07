pipeline {
    agent {
        label 'java-docker-slave'
    }
    stages {
        stage('Build') {
            steps { 
                echo 'Building'
                echo 'Building2'
                script {
                    def result = sh(script: 'which java', returnStatus: true)
                    if (result == 0) {
                        def location = sh(script: 'which java', returnStdout: true).trim()
                        echo "Location of 'java' command: ${location}"
                    } else {
                        error "'java' command not found"
                    }
                }
            }
        }
        stage('Test') {
            steps { 
                echo 'Testing'
            }
        }
        stage('Deploy') {
            steps { 
                echo 'Deploying'
            }
        }
    }
}