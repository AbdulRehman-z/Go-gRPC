package main

import (
	pb "grpc/proto"
	"io"
	"log"
)

func  (s *welcome) BidirectionalStream(stream pb.WelcomeService_BidirectionalStreamServer) error  {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.WelcomeResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}