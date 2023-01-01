package main

import (
	"log"

	pb "GogRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main(){
	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Client did not connect: %v",err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Bruce Wayne", "Clark Kent","Diana Prince","Barry Allen"},
	}
	//callSayHello(client)
	//callSayHelloServerStreaming(client,names)
	callSayHelloClientStreaming(client,names)
	//callSayHelloBidirectionalStreaming(client,names)

}