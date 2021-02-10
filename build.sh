#!/usr/bin/env bash

echo ${DOCKER_PASS} | docker login --username ${DOCKER_USER} --password-stdin

HASH=$(git rev-parse --short HEAD)

docker build -t s1moe2/poc-jobs:$HASH -t s1moe2/poc-jobs:latest .
docker push s1moe2/poc-jobs:$HASH
docker push s1moe2/poc-jobs:latest
