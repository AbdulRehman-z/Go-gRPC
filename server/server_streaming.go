package main

import (
	pb "grpc/proto"
	"log"
)

func (s *welcome) ServerStream(request *pb.CitiesArray , stream pb.WelcomeService_ServerStreamServer) error {
	log.Println("Received this on server: ",request.Cities)

	for _ ,name := range request.Cities {
		response := &pb.WelcomeResponse{Message: "That's a" + name}

        if err := stream.Send(response); err != nil {
			log.Println("Error: ",err)
		}
	}
   return nil
}