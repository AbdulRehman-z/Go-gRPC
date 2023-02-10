package main

import (
	pb "grpc/proto"
	"io"
	"log"
)

func (s *welcome) ClientStream(stream pb.WelcomeService_ClientStreamServer) error {
	var cities []string
	for {
		req,err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CitiesArray{Cities: cities})
		}
		if err != nil {
			panic(err)
		}
		log.Println("Request received from client: ",req.Name)
		cities = append(cities, req.Name)
	}
}
