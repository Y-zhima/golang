FROM golang:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git wget protobuf

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
RUN unzip protoc-3.6.1-linux-x86_64.zip -d /usr/local/ -x bin/protoc
RUN rm protoc-3.6.1-linux-x86_64.zip

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
RUN go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
RUN go get -u github.com/golang/mock/gomock && go install github.com/golang/mock/mockgen

RUN git clone https://github.com/grpc-ecosystem/grpc-gateway.git /go/src/github.com/grpc-ecosystem/grpc-gateway
RUN git clone https://github.com/mwitkow/go-proto-validators.git /go/src/github.com/mwitkow/go-proto-validators

WORKDIR /opt/protos