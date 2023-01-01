package main

import (
	pb "GogRPC/proto"
	"context"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err!=nil{
		log.Fatalf("Couldn't send names: %v",err)
	}

	waitc := make(chan struct{})

	go func() {
		for{
			message, err:= stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Streaming error: %v",err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req);err!=nil{
			log.Fatalf("Sending request error: %v",err)
		}
		time.Sleep(2*time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}