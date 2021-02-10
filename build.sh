#!/usr/bin/env bash

echo ${DOCKER_PASS} | docker login --username ${DOCKER_USER} --password-stdin

HASH=$(git rev-parse --short HEAD)

docker build -t cloudokihub/apisuite-be-jobs:$HASH -t s1moe2/apisuite-be-jobs:latest .
docker push cloudokihub/apisuite-be-jobs:$HASH
docker push cloudokihub/apisuite-be-jobs:latest
