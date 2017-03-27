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

func (s *taskServer) Run(ctx context.Context, req *proto.RunRequest) (*proto.RunResponse, error) {
	log.Printf("Received Task.Run with method: %s", req.Method)

	rsp := new(proto.RunResponse)
	switch req.Method {
	case "GET":
		rsp.StatusCode = proto.RunResponse_OK
	case "POST":
		rsp.StatusCode = proto.RunResponse_CREATED
	default:
		rsp.StatusCode = proto.RunResponse_FAILED
	}

	return rsp, nil
}
