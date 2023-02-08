package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
	"net"
)

const (
	port = ":8000"
)

type welcome struct {
	pb.WelcomeServiceServer
}

func main() {
	listen,err := net.Listen("tcp", port)
	checkFatalError(err)

	grpcServer := grpc.NewServer()
	// Register the service
	pb.RegisterWelcomeServiceServer(grpcServer, &welcome{})
	err = grpcServer.Serve(listen)
	checkFatalError(err)
}

func checkFatalError(err error)  {
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}





