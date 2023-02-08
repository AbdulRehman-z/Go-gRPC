package main

import (
	"context"
	pb "grpc/proto"
)

func(s *welcome) Welcome(ctx context.Context, request *pb.WelcomeRequest) (*pb.WelcomeResponse, error)  {
	return  &pb.WelcomeResponse{Message: "Welcome " + request.Name}, nil
}