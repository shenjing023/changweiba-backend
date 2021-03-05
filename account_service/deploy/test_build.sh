#!/bin/bash

# delete local docker image
docker rmi account-service:v1
# build docker image
docker build -t account-service:v1 ..
# delete docker image in k8s node
docker exec -it changweiba-control-plane crictl rmi docker.io/library/account-service:v1
# load image to k8s node
kind load docker-image account-service:v1 --name changweiba
# delete pod to restart pod
kubectl delete po $(kubectl -n default get pod -l app=account-service -o jsonpath='{.items[0].metadata.name}')
