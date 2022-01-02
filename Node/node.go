package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Hexfall/Increment/increment"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 5080, "port")

func main() {
	flag.Parse()

	// gRPC set-up.
	log.Printf("Attempting to listen on port: %d\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d. Error: %v", *port, err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	node := increment.NewNode()
	increment.RegisterCommunicationServer(grpcServer, &node)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC client on port %d. Error: %v", *port, err)
	}
}
