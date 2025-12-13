package main

import (
	"context"

	pb "github.com/Rushi2398/learn-gRPC/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
