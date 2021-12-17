#!/bin/bash
echo "pwd:"`pwd`
echo "============================="
cd ..
path=`pwd`
echo "work path:" $path
find . -name "Protocol"
echo "============================="

# 属于对外的客户端proto
str2="client"

for k in $(find . -name "Protocol"); do
   for file in $( find $k -name "*.proto"); do

    if [ $file = ./Protocol/third_party/validate/validate.proto ]; then
      echo "continue protoc file:$file"
      continue
    fi

    echo "start protoc file:$file"
    protoc --proto_path=. \
           --proto_path=./Protocol/third_party \
           --go_out=paths=source_relative:. $file \
           --go-grpc_out=paths=source_relative:. $file \

    if [ $file = ./Protocol/login/login.proto ]; then
      protoc --proto_path=. \
           --proto_path=./Protocol/third_party \
           --go-http_out=paths=source_relative:. $file \

    fi

  done
done