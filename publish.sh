#!/usr/bin/env bash
set -e

if [[ -z $(git status -s) ]]
then
  echo "tree is clean"
else
  echo "tree is dirty, please commit changes before running this"
  echo $(git diff)
  exit 1
fi

# configurations
REPO=link.sakaba.repo.api
ECR_REPO=823135059493.dkr.ecr.ap-northeast-1.amazonaws.com/${REPO}
CUR_DATE=`date +%Y%m%d`
SHORT_HASH=$(git rev-parse --short HEAD)
TAG=${CUR_DATE}-${SHORT_HASH}

# Build Docker image
go build *.go
docker build . -t ${REPO}:${TAG}

# Login to ECR
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 823135059493.dkr.ecr.ap-northeast-1.amazonaws.com

# Push the built image to ECR
docker tag  ${REPO}:${TAG} ${ECR_REPO}:${TAG}
docker push ${ECR_REPO}:${TAG}

# Clean up the buit image
docker rmi ${REPO}:${TAG}

# Output artifact
echo ${ECR_REPO}:${TAG} > ./build/deploy.txt

# Set env for the following actions
echo "IMAGE_TAG=${TAG}" >> $GITHUB_ENV
