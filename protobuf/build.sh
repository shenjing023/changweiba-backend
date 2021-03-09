#!/bin/bash

# protoc --go_out=../account_service/pb --go-grpc_out=../account_service/pb account_service.proto enums.proto
protoc --go_out=$1 --go-grpc_out=$1 $2 $3