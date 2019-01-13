#!/bin/bash

IMAGE_NAME=example
BRANCH_NAME=$TRAVIS_BRANCH

if [ $BRANCH_NAME == "master" ]; then
	BRANCH_NAME="latest"
fi

if [ -z $DOCKER_USERNAME ]; then
	echo "Missing DOCKER_USERNAME env var"
	exit 1
fi

DOCKER_IMAGE=$IMAGE_NAME:$BRANCH_NAME
REPO_IMAGE=$DOCKER_USERNAME/$DOCKER_IMAGE

docker build -t $REPO_IMAGE .
