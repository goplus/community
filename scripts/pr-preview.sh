#!/usr/bin/env bash
###########################################
# This script is used to start the application in the PR environment
###########################################
set -ex

echo "WORKSPACE: ${PWD}"

# compile the application
cd cmd/gopcomm
gop mod tidy
gop build -o ./community

# Copy the environment file
cp /config/.env .env

# build the docker image
cat > Dockerfile << EOF
FROM aslan-spock-register.qiniu.io/library/ubuntu:22.04
WORKDIR /app
COPY ./community .
COPY ./.env .
COPY yap ./
EXPOSE 8080
CMD ["./community"]
EOF

export CONTAINER_IMAGE=aslan-spock-register.qiniu.io/goplus/goplus-community-pr:${PULL_NUMBER}-${PULL_PULL_SHA:0:8}
docker build -t ${CONTAINER_IMAGE} -f ./Dockerfile . --builder="kube" --push


# generate kubernetes yaml with unique flag for PR
cat > community.yaml << EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goplus-community-pr-${PULL_NUMBER}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: goplus-community-pr-${PULL_NUMBER}
  template:
    metadata:
      labels:
        app: goplus-community-pr-${PULL_NUMBER}
    spec:
      containers:
        - name: my-container
          image: ${CONTAINER_IMAGE}
          ports:
            - containerPort: 8080
---            
apiVersion: v1
kind: Service
metadata:
  name: goplus-community-pr-${PULL_NUMBER}
spec:
  selector:
    app: goplus-community-pr-${PULL_NUMBER}
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: ClusterIP
EOF

kubectl -n goplus-pr-review apply -f community.yaml 

kubectl -n goplus-pr-review get pods

# comment on the PR
Preview_URL=http://goplus-community-pr-${PULL_NUMBER}.goplus-pr-review.svc.jfcs-qa1.local
message=$'The PR environment is ready, please check the [PR environment]('${Preview_URL}')

[Attention]: This environment will be automatically cleaned up after 6 hours, please make sure to test it in time. If you have any questions, please contact the author of the PR or the community team. 
'
gh_comment -org=${REPO_OWNER} -repo=${REPO_NAME} -num=${PULL_NUMBER} -p=${Preview_URL} -b "${message}"


