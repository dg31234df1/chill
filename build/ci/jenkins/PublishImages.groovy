#!/usr/bin/env groovy

pipeline {
    agent {
        kubernetes {
            label "milvus-e2e-test-kind"
            defaultContainer 'main'
            yamlFile "build/ci/jenkins/pod/krte.yaml"
            customWorkspace '/home/jenkins/agent/workspace'
            // We allow this pod to remain active for a while, later jobs can
            // reuse cache in previous created nodes.
            idleMinutes 120
        }
    }

    options {
        timestamps()
        timeout(time: 30, unit: 'MINUTES')
        // parallelsAlwaysFailFast()
    }

    environment {
        DOCKER_CREDENTIALS_ID = "ba070c98-c8cc-4f7c-b657-897715f359fc"
        DOKCER_REGISTRY_URL = "registry.zilliz.com"
        TARGET_REPO = "${DOKCER_REGISTRY_URL}/milvus"
    }

    stages {
        stage ('Publish Milvus Images') {
            steps {
                container('main') {
                    script {
                        sh "build/builder.sh /bin/bash -c \"make install\""

                        def date = sh(returnStdout: true, script: 'date +%Y%m%d').trim()
                        def gitShortCommit = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()

                        withCredentials([usernamePassword(credentialsId: "${env.DOCKER_CREDENTIALS_ID}", usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                            sh """
                                docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD} ${DOKCER_REGISTRY_URL}
                                export MILVUS_IMAGE_REPO="${env.TARGET_REPO}/milvus"
                                export MILVUS_IMAGE_TAG="${env.BRANCH_NAME}-${date}-${gitShortCommit}"
                                build/build_image.sh
                                docker push \${MILVUS_IMAGE_REPO}:\${MILVUS_IMAGE_TAG}
                                docker tag \${MILVUS_IMAGE_REPO}:\${MILVUS_IMAGE_TAG} \${MILVUS_IMAGE_REPO}:${env.BRANCH_NAME}-latest
                                docker push \${MILVUS_IMAGE_REPO}:${env.BRANCH_NAME}-latest
                                docker logout ${DOKCER_REGISTRY_URL}
                            """
                        }
                    }
                }
            }
        }
    }
}
