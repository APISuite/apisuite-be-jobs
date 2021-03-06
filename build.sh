#!/usr/bin/env bash

echo ${DOCKER_PASS} | docker login --username ${DOCKER_USER} --password-stdin

HASH=$(git rev-parse --short HEAD)

VERSION=$(cat package.json | grep version | head -1 | awk -F ": " '{ print $2 }' | sed 's/[",]//g')

docker build \
  -t cloudokihub/apisuite-be-jobs:$HASH \
  -t cloudokihub/apisuite-be-jobs:latest \
  -t cloudokihub/apisuite-be-jobs:$VERSION .

docker push cloudokihub/apisuite-be-jobs:$HASH
docker push cloudokihub/apisuite-be-jobs:$VERSION
docker push cloudokihub/apisuite-be-jobs:latest
