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
                helloWorld(name: "Hanjaya", day: "Friday")
	    	timeout(time: 1, unit: 'SECONDS') {
			echo "sleeping in timeout"
			sleep 2
		}
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
