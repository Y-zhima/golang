#!/bin/bash
rm -rf /tmp/protos/javaout
rm -rf /tmp/protos/git.fogcdn.top/axe/protos/goout

mkdir swagger
mkdir -p /tmp/protos/javaout

for f in src/*/*.proto ; do
    echo "gen proto File -> $f";
    protoc -Isrc -I/usr/local/include \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:/tmp/protos \
    --java_out=/tmp/protos/javaout \
    --grpc-gateway_out=logtostderr=true:/tmp/protos \
    --swagger_out=logtostderr=true:swagger \
    $f
done ;

cp -rf /tmp/protos/javaout .
cp -rf /tmp/protos/git.fogcdn.top/axe/protos/goout .