#!/bin/bash

gen_proto() {
  rm -rf /tmp/protos && mkdir -p /tmp/protos

  for f in src/*/*.proto ; do
    echo "gen proto File -> $f";
    protoc -Isrc -I/usr/local/include -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:/tmp/protos \
    --govalidators_out=/tmp/protos \
    --grpc-gateway_out=logtostderr=true:/tmp/protos \
    $f
  done ;

  cp -rf /tmp/protos/git.fogcdn.top/axe/protos/goout .
}

gen_swagger() {
  mkdir -p swagger
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
    "host": "axe-develop.fogcdn.top:8969,
    "swagger": "2.0",
    "info": {
      "description": "AXE运维平台RESTful API",
      "title": "AXE运维API",
      "version": "1.0.0"
    }
  }' > swagger/swagger.json 

  for f in src/*/*.proto ; do
    echo "gen swagger File -> $f";
    protoc -Isrc -I/usr/local/include -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:swagger \
    $f
  done ;
}

gen_doc() {
  mkdir -p doc

  echo "gen doc File -> doc.markdown";
  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=markdown,doc.md src/*/*.proto

  echo "gen doc File -> doc.html";
  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=html,doc.html src/*/*.proto

  echo "gen doc File -> doc.json";
  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=json,doc.json src/*/*.proto

  echo "gen doc File -> doc.docbook";
  protoc -Isrc -I/usr/local/include -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --doc_out=./doc --doc_opt=docbook,doc.docbook src/*/*.proto
}

gen_mock() {
  for f in goout/*/*.pb.go ; do
    services=$(cat $f |grep interface|grep -o "[A-Z].*Client")
    dirname=$(basename $(dirname $f));
    savedirname="mock_${dirname}";
    savefilename="mock_${dirname}.go";
    if  [ $services ]; then
      echo "gen mock File -> $f";
      mockgen -destination=gooutmock/$savedirname/$savefilename  git.fogcdn.top/axe/protos/goout/$dirname $services;
    fi;
  done ;
}

if [ "all" == "$2" ]; then
  gen_all $1
else
  gen_$1
fi;