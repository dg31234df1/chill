def isTimeTriggeredBuild = currentBuild.getBuildCauses('hudson.triggers.TimerTrigger$TimerTriggerCause').size() != 0
def regressionTimeout = isTimeTriggeredBuild ? "180" : "60"
timeout(time: "${regressionTimeout}", unit: 'MINUTES') {
    container('deploy-env') {
        dir ('milvus-helm-chart') {
            sh " helm version && \
                 helm repo add stable https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts && \
                 helm repo add bitnami https://charts.bitnami.com/bitnami && \
                 helm repo add minio https://helm.min.io/ && \
                 helm repo update"

            def milvusHelmURL = "https://github.com/zilliztech/milvus-helm-charts.git"
            checkout([$class: 'GitSCM', branches: [[name: "${env.HELM_BRANCH}"]], userRemoteConfigs: [[url: "${milvusHelmURL}"]]])

            dir ('charts/milvus-ha') {
                sh script: "kubectl create namespace ${env.HELM_RELEASE_NAMESPACE}", returnStatus: true

                def helmCMD = ""
                if ("${REGRESSION_SERVICE_TYPE}" == "distributed") {
                    helmCMD = "helm install --wait --timeout 300s \
                                   --set standalone.enabled=false \
                                   --set image.all.repository=${env.TARGET_REPO}/milvus-distributed \
                                   --set image.all.tag=${env.TARGET_TAG} \
                                   --set image.all.pullPolicy=Always \
                                   --namespace ${env.HELM_RELEASE_NAMESPACE} ${env.HELM_RELEASE_NAME} ."
                } else {
                    helmCMD = "helm install --wait --timeout 300s \
                                   --set image.all.repository=${env.TARGET_REPO}/milvus-distributed \
                                   --set image.all.tag=${env.TARGET_TAG} \
                                   --set image.all.pullPolicy=Always \
                                   --namespace ${env.HELM_RELEASE_NAMESPACE} ${env.HELM_RELEASE_NAME} ."
                }

                try {
                    sh "${helmCMD}"
                } catch (exc) {
                    def helmStatusCMD = "helm get manifest -n ${env.HELM_RELEASE_NAMESPACE} ${env.HELM_RELEASE_NAME} | kubectl describe -n ${env.HELM_RELEASE_NAMESPACE} -f - && \
                                         helm status -n ${env.HELM_RELEASE_NAMESPACE} ${env.HELM_RELEASE_NAME}"
                    sh script: helmStatusCMD, returnStatus: true
                    throw exc
                }
            }
        }
    }

    container('test-env') {
        try {
            dir ('tests/python_test') {
                sh "python3 -m pip install --no-cache-dir -r requirements.txt"
                if (isTimeTriggeredBuild) {
                    echo "This is Cron Job!"
                    sh "pytest --tags=0331 -n 2 --ip ${env.HELM_RELEASE_NAME}-milvus-ha.${env.HELM_RELEASE_NAMESPACE}.svc.cluster.local"
                } else {
                    sh "pytest --tags=smoke -n 2 --ip ${env.HELM_RELEASE_NAME}-milvus-ha.${env.HELM_RELEASE_NAMESPACE}.svc.cluster.local"
                }
            }
        } catch (exc) {
            echo 'PyTest Regression Failed !'
            throw exc
        } finally {
            container('deploy-env') {
                def milvusLabels = "app.kubernetes.io/instance=${env.HELM_RELEASE_NAME}"
                def componentLabels = "release=${env.HELM_RELEASE_NAME}"
                def namespace = "${env.HELM_RELEASE_NAMESPACE}"
                def artifactsPath = "${env.DEV_TEST_ARTIFACTS_PATH}"


                sh "mkdir -p $artifactsPath"
                sh "for pod in \$(kubectl get pod -n $namespace -l ${milvusLabels} -o jsonpath='{range.items[*]}{.metadata.name} '); do kubectl logs --all-containers -n $namespace \$pod > $artifactsPath/\$pod.log; done"
                sh "for pod in \$(kubectl get pod -n $namespace -l ${componentLabels} -o jsonpath='{range.items[*]}{.metadata.name} '); do kubectl logs --all-containers -n $namespace \$pod > $artifactsPath/\$pod.log; done"
                archiveArtifacts artifacts: "$artifactsPath/**", allowEmptyArchive: true
            }
        }
    }
}
