package main

import (
	pb "GogRPC/proto"
	"context"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient,names *pb.NameList){
	var message *pb.HelloResponse
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil{
		log.Fatalf("Greeting failure: %v",err)
	}
	
	for{
		err := stream.RecvMsg(message)
		if err == io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Streaming failure: %v",err)
		}
		log.Printf(message.Message)
	}
	log.Printf("Streaming finished.")
}