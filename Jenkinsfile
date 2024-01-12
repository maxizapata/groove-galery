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
                    def result = sh(script: 'ls /home/jenkins/', returnStatus: true)
                    if (result == 0) {
                        def home_jenkins = sh(script: 'ls /home/jenkins/', returnStdout: true).trim()
                        echo "${home_jenkins}"
                    } else {
                        error "'java' command not found"
                    }
                }
            }
        }
        stage('Test') {
            steps { 
                echo 'Testing'
                sh 'go version'
            }
        }
        stage('Deploy') {
            steps { 
                echo 'Deploying'
            }
        }
    }
}