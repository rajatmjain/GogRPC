package main

import (
	pb "GogRPC/proto"
	"log"
	"time"
)

func (s *helloServer) callSayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer)(error){
	log.Printf("got request with names: %v",req.Names)

	for _, name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello"+name,
		}
		if err := stream.SendMsg(res); err!=nil{
			return err
		}
		time.Sleep(3*time.Second)
		
	}
	return nil
}