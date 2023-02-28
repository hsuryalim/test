@Library('shared-library') _
pipeline {
    agent any
    stages {
	stage('Hello world') {
		tools {
			maven 'maven 3.6.3'
		}
		steps {
			sh 'mvn -version'
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
