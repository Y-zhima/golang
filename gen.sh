#!/bin/bash
mkdir swagger
echo '{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "AXE运维平台RESTful API",
    "title": "AXE运维API",
    "version": "1.0.0"
  }
}' > swagger/swagger.json 

mkdir -p /tmp/protos/javaout

for f in src/*/*.proto ; do
    echo "gen proto File -> $f";
    protoc -Isrc -I/usr/local/include \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src \
    --go_out=plugins=grpc:/tmp/protos \
    --govalidators_out=/tmp/protos \
    --java_out=/tmp/protos/javaout \
    --grpc-gateway_out=logtostderr=true:/tmp/protos \
    --swagger_out=logtostderr=true:swagger \
    $f
done ;

cp -rf /tmp/protos/javaout .
cp -rf /tmp/protos/git.fogcdn.top/axe/protos/goout .