#!/bin/bash

# python3 -m grpc_tools.protoc --python_out=../wencai_service/pb 
# --pyi_out=../wencai_service/pb --grpc_python_out=../wencai_service/pb wencai_service.proto
python3 -m grpc_tools.protoc -I. --python_out=$1 --pyi_out=$1 --grpc_python_out=$1 $2