@Library('shared-library') _
pipeline {
    agent any
    stages {
	stage('Hello world') {
		tools {
			maven 'maven 3.6.3'
		}
		steps {
			sh 'mvn --version'
		}
	}
    }

}
