@Library('shared-library') _
pipeline {
    agent {
        docker {
            image 'golang'
        }
    }

    stages {
        stage('Hello world') {
            steps {
                helloWorld()
            }
        }
        stage ('Testing') {
            steps {
                echo "Testing the application"
                sh 'go test ./...'
            }
        }

        stage ('Building') {
            steps {
                sh 'go build -o app'
                sh './app'
            }
        }
    }

}