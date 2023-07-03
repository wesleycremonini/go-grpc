package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"wesleycremonini/grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	names := &proto.NamesList{
		Names: []string{"Wesley", "Macaco"},
	}

	// callSayHello(client)
	callSayHelloBidirectionalStream(client, names)
}
