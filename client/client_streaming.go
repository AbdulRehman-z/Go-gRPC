package main

import (
	"context"
	pb "grpc/proto"
	"log"
)

func clientSideStreaming(client pb.WelcomeServiceClient,names *pb.CitiesArray)  {
	log.Println("Client Side Streaming Started")
	stream,err := client.ClientStream(context.Background())

	if err != nil {
		log.Fatalf("Error: %s",err)
	}

	for _ ,name := range names.Cities {
		req := &pb.WelcomeRequest{Name:name}
		if err := stream.Send(req); err != nil {
          log.Fatalf("Error: %s",err)
		}
		log.Println("Request send with city name: ",name)
	}

	res ,err := stream.CloseAndRecv()
	log.Println("Response from server: ",res.Cities)
	if err != nil {
		log.Fatalf("Error: %s",err)
	}
}