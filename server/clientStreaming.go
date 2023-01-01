package main

import (
	pb "GogRPC/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer)error{
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err!=nil{
			log.Fatalf("\nClient Streaming failure: %v\n",err)
		}
		log.Printf("Got request with name: %v",req.Name)
		messages = append(messages, "Hello",req.Name)
	}
}