package main

import (
	"flag"
	"log"

	proto "github.com/LOG-ED/go-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address = flag.String("task_endpoint", "srv:9090", "endpoint of Task service")
)

func main() {
	flag.Parse()
	conn, _ := grpc.Dial(*address, grpc.WithInsecure())
	defer conn.Close()
	c := proto.NewTaskClient(conn)

	r, err := c.Run(context.Background(), &proto.RunRequest{Method: "GET"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("output: %v", r)
}
