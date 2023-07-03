package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"wesleycremonini/grpc/proto"
)

type helloServer struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
