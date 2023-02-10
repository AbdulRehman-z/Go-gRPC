package main

import (
	pb "grpc/proto"
	"io"
	"log"
)

func  (s *welcome) BidirectionalStreamServer(stream pb.WelcomeService_BidirectionalStreamServer) error  {
	for {
		request,err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream ended")
		}
		log.Println("Stream received from client: ", request.Name)

		response := &pb.WelcomeResponse{Message: "Capital" + request.Name}

		if err := stream.Send(response); err != nil {
			panic(err)
		}
		log.Println("Stream sent to client: ", request.Name)
	}
}