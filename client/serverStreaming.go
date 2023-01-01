package main

import (
	pb "GogRPC/proto"
	"context"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil{
		log.Fatalf("Greeting failure: %v",err)
	}
	for{
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("\nServer Streaming failure: %v\n",err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished.")
}