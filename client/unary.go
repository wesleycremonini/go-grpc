package main

import (
	"context"
	"log"
	"time"

	"wesleycremonini/grpc/proto"
)

func callSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
