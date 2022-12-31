package main

import (
	pb "GogRPC/proto"
	"context"
)

func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "Hello gRPC",
	},nil
}
