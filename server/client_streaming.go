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

//
//goroutine 34 [running]:
//main.(*welcome).ClientStream(0xc000055b10?, {0xb64700, 0xc000210190})
//C:/Users/Abdul Rehman/Desktop/grpc/server/client_streaming.go:20 +0x102
//grpc/proto._WelcomeService_ClientStream_Handler({0xa515e0?, 0xc00004fc20}, {0xb63128?, 0xc000240000})
//C:/Users/Abdul Rehman/Desktop/grpc/proto/proto_grpc.pb.go:204 +0x9f
//google.golang.org/grpc.(*Server).processStreamingRPC(0xc0001021e0, {0xb64a38, 0xc0000841a0}, 0xc00023a000, 0xc000106a50, 0xe1e9c0, 0x0)
//C:/Users/Abdul Rehman/go/pkg/mod/google.golang.org/grpc@v1.52.3/server.go:1620 +0x11e7
//google.golang.org/grpc.(*Server).handleStream(0xc0001021e0, {0xb64a38, 0xc0000841a0}, 0xc00023a000, 0x0)
//C:/Users/Abdul Rehman/go/pkg/mod/google.golang.org/grpc@v1.52.3/server.go:1708 +0x9ea
//google.golang.org/grpc.(*Server).serveStreams.func1.2()
//C:/Users/Abdul Rehman/go/pkg/mod/google.golang.org/grpc@v1.52.3/server.go:965 +0x98
//created by google.golang.org/grpc.(*Server).serveStreams.func1
//C:/Users/Abdul Rehman/go/pkg/mod/google.golang.org/grpc@v1.52.3/server.go:963 +0x28a
//exit status 2