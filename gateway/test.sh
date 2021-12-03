#!/bin/bash

CURRDIR=`dirname "$0"`
BASEDIR=`cd "$CURRDIR"; pwd`
# 配置文件挂载
HCONF=$BASEDIR/conf
CCONF=/app/conf
# delete running services
docker rm -f gateway-service_v1
# delete local docker image
docker rmi gateway-service:v1
# build docker image
docker build -t gateway-service:v1 .
# run docker image

docker run -d -it -p 8020:8020 --name gateway-service_v1  -v $HCONF/config.yaml:$CCONF/config.yaml gateway-service:v1