//package main
//
//import (
//	"io"
//	"log"
//	"time"
//)

package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
)

func bidirectionalClientStreaming(client pb.WelcomeServiceClient, names *pb.CitiesArray) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Cities{
		req := &pb.WelcomeRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
	}

	err = stream.CloseSend()
	if err != nil {
		return
	}
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}





//func  bidirectionalClientStreaming(client pb.WelcomeServiceClient, cities *proto.CitiesArray)  {
//	stream,err := client.BidirectionalStream(context.Background())
//	if err != nil {
//		log.Println("Error: ",err)
//	}
//
//	for _,city := range cities.Cities {
//		req := &pb.WelcomeRequest{Name:city}
//		if err := stream.Send(req); err != nil {
//			log.Println("Error: ",err)
//		}
//		log.Println("Request send with city name: ",city)
//	}
//
//	res ,err := stream.Recv()
//	log.Println("Response from server: ",res.Message)
//    if err == io.EOF {
//		log.Println("Stream ended")
//	}
//
//}