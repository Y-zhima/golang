FROM golang:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache protobuf git make
ENV GO111MODULE 'on'
ENV GOPROXY https://athens.azurefd.net
# 安装依赖
RUN go get google.golang.org/grpc
RUN go get github.com/micro/protoc-gen-micro
RUN go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /opt/protos