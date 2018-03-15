/**
 * Copyright 2018, Boundless, https://boundlessgeo.com - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

node {
  withCredentials([string(credentialsId: 'boundlessgeoadmin-token', variable: 'GITHUB_TOKEN'),
                   string(credentialsId: 'sonar-jenkins-pipeline-token', variable: 'SONAR_TOKEN')
                   string(credentialsId: 'AWS_ACCESS_KEY_ID', variable:'AWS_ACCESS_KEY_ID')
                   string(credentialsId: 'AWS_SECRET_ACCESS_KEY', variable:'AWS_SECRET_ACCESS_KEY')]) {

    currentBuild.result = "SUCCESS"

    try {
      stage('Checkout'){
        checkout scm
          echo "Running ${env.BUILD_ID} on ${env.JENKINS_URL}"
      }

      stage('Deploy'){
        // make build
        sh """
          docker run -v \$(pwd -P):/go/src/github.com/boundlessgeo/wfs3 \
                     -w /code golang:1.9.2-alpine3.7 sh \
                     -c 'apk add --no-cache git build-base bash zip && \
                         bash -c "go get -u github.com/golang/dep/cmd/dep" && \
                         bash -c "cd /go/src/github.com/boundlessgeo/wfs3; \
												 dep ensure; \
												 go build -ldflags -v -o /code/target/wfs3" && \
												 bash -c "zip deployment.zip /code/target/wfs3" && \
                         bash -c "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}"
												 bash -c "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}"
                         bash -c "aws lambda update-function-code \
												              --region us-east-1
																			--function-name wfs3
																			--zip-file fileb://./deployment.zip"'
           """
      }

    }
    catch (err) {

      currentBuild.result = "FAILURE"
        throw err
    } finally {
      // Success or failure, always send notifications
      echo currentBuild.result
      notifyBuild(currentBuild.result)
    }

  }
}

// Slack Integration

def notifyBuild(String buildStatus = currentBuild.result) {

  // generate a custom url to use the blue ocean endpoint
  def jobName =  "${env.JOB_NAME}".split('/')
  def repo = jobName[0]
  def pipelineUrl = "${env.JENKINS_URL}blue/organizations/jenkins/${repo}/detail/${env.BRANCH_NAME}/${env.BUILD_NUMBER}/pipeline"
  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}\nJob: ${env.JOB_NAME}\nBuild: ${env.BUILD_NUMBER}\nJenkins: ${pipelineUrl}\n"
  def summary = (env.CHANGE_ID != null) ? "${subject}\nAuthor: ${env.CHANGE_AUTHOR}\n${env.CHANGE_URL}\n" : "${subject}"

  // Override default values based on build status
  if (buildStatus == 'SUCCESS') {
    colorName = 'GREEN'
    colorCode = '#228B22'
  }

  // Send notifications
  slackSend (color: colorCode, message: summary, channel: '#cto-svc-bots')
}

