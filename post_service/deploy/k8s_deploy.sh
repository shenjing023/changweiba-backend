#!/bin/bash

# build docker image
docker build -t post-service:vxxx .

# load image to kind cluster node
kind load docker-image post-service:vxxx --name changweiba

# create k8s configMap
kubectl create cm post-service-config --from-file=conf/config.yaml

# deploy
kubectl apply -f post-service.yaml