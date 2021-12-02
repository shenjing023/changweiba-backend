#!/bin/bash

CURRDIR=`dirname "$0"`
BASEDIR=`cd "$CURRDIR"; pwd`
# 配置文件挂载
HCONF=$BASEDIR/conf
CCONF=/app/conf
# delete running services
docker rm -f post-service_v1
# delete local docker image
docker rmi post-service:v1
# build docker image
docker build -t post-service:v1 .
# run docker image

docker run -d -it -p 8019:8019 --name post-service_v1  -v $HCONF/config.yaml:$CCONF/config.yaml post-service:v1