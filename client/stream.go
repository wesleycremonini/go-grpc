package main

import (
	"context"
	"io"
	"log"
	"time"

	"wesleycremonini/grpc/proto"
)

func callSayHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}

		log.Printf("Bidirectional Streaming finished")
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
}