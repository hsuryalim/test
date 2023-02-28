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
            	retry(3){
			echo "Before error"
			error "error in retry"
		}
		echo "never come out"
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
