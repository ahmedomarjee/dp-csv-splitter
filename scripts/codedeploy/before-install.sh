#!/bin/bash

CONTAINER_ID=$(docker ps -a | grep dp-csv-splitter | awk '{print $1}')
IMAGE_ID=$(docker images -a | grep dp-csv-splitter | awk '{print $3}')

if [[ -n $CONTAINER_ID ]]; then
  docker rm $CONTAINER_ID
fi

if [[ -n $IMAGE_ID ]]; then
  docker rmi $IMAGE_ID
fi
