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
            	script {
			def name = "jeff"

			if(name=="jeff")
				println("hi ${name}")
			else 
				println("hi human!")

			sleep 2
			echo "end of script"
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
