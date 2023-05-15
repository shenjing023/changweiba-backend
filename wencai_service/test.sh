#!/bin/bash

# delete running services
docker rm -f wencai-service_v1
# delete local docker image
docker rmi wencai-service:v1
# build docker image
docker build -t wencai-service:v1 .
# run docker image

docker run -d -it -p 8022:8022 --name wencai-service_v1 --network=host wencai-service:v1