package main

import (
	"log"
	"net"

	proto "github.com/LOG-ED/go-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterTaskServer(s, newTaskServer())

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(l)
}
