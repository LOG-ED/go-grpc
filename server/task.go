package main

import (
	"log"

	proto "github.com/LOG-ED/go-grpc/proto"
	"golang.org/x/net/context"
)

type taskServer struct{}

func newTaskServer() proto.TaskServer {
	return new(taskServer)
}

func (s *taskServer) Run(ctx context.Context, msg *proto.RunRequest) (*proto.RunResponse, error) {
	log.Println(msg)
	resp := proto.RunResponse{StatusCode: proto.RunResponse_OK}
	return &resp, nil
}
