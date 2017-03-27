package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/LOG-ED/go-grpc/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	taskEndpoint = flag.String("task_endpoint", "srv:9090", "endpoint of Task service")
)

// Run starts an HTTP server
func Run(address string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterTaskHandlerFromEndpoint(ctx, mux, *taskEndpoint, dialOpts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(address, mux)
}

func main() {
	flag.Parse()

	if err := Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
