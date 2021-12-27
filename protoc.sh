#!/bin/bash
echo "pwd:"`pwd`
echo "============================="
cd ..
path=`pwd`
echo "work path:" $path
find . -name "Proto"
echo "============================="

# 属于对外的客户端proto
str2="client"

for k in $(find . -name "Proto"); do
   for file in $( find $k -name "*.proto"); do

    if [[ $file == *"third_party"* ]]; then
      echo "third_party"
      continue
    fi

    echo "start protoc file:$file"
    protoc --proto_path=./Proto \
           --proto_path=./Proto/third_party \
           --go_out=paths=source_relative:./Proto $file \
           --go-grpc_out=paths=source_relative:./Proto $file \

    if [ $file = ./Proto/login/login.proto ]; then
      protoc --proto_path=./Proto \
             --proto_path=./Proto/third_party \
             --go-http_out=paths=source_relative:./Proto $file \

    fi

  done
done