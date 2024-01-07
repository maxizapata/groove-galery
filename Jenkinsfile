pipeline {
    agent {
        label 'java-docker-slave'
    }
    stages {
        stage('Build') {
            steps { 
                echo 'Building'
                echo 'Building2'
                which go
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