FROM golang:1.7

RUN apt-get update && apt-get install -y unzip vim \
    && rm -rf /var/lib/apt/lists/*

ENV PROTOBUF_VERSION 3.2.0

RUN curl -fsSL https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-linux-x86_64.zip -o protobuf.zip

RUN unzip protobuf.zip

RUN go get -v github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis | true

RUN go get -v github.com/golang/protobuf/protoc-gen-go

RUN mkdir -p /go/src/github.com/LOG-ED/go-grpc

COPY . /go/src/github.com/LOG-ED/go-grpc

WORKDIR /go/src/github.com/LOG-ED/go-grpc