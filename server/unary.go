package main

import (
	"context"

	"wesleycremonini/grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
