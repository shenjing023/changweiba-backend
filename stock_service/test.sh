#!/bin/bash

CURRDIR=`dirname "$0"`
BASEDIR=`cd "$CURRDIR"; pwd`
# 配置文件挂载
HCONF=$BASEDIR/conf
CCONF=/app/conf
# delete running services
docker rm -f stock-service_v1
# delete local docker image
docker rmi stock-service:v1
# build docker image
docker build -t stock-service:v1 .
# run docker image

docker run -d -it -p 8021:8021 --name stock-service_v1  --network=host -v $HCONF/config.yaml:$CCONF/config.yaml stock-service:v1