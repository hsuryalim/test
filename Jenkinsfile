@Library('shared-library') _
pipeline {
    agent any
    stages {
	stage('Hello world') {
		tools {
			maven 'maven_3_6_3'
		}
		steps {
			sh 'mvn -version'
		}
	}
    }

}
