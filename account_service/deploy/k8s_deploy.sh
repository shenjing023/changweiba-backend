#!/bin/bash

# build docker image
docker build -t account-service:vxxx .

# load image to kind cluster node
kind load docker-image account-service:vxxx --name changweiba

# create k8s configMap
kubectl create cm account-service-config --from-file=conf/config.yaml

# deploy
kubectl apply -f account-service.yaml