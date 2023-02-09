package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
)

func serverSideStreaming(client pb.WelcomeServiceClient, names *pb.CitiesArray)  {
	log.Println("streaming data to the server started")

   stream,err  := client.ServerStream(context.Background(),names)

	if err!=nil {
		log.Println("Error: ",err)
	}

   for {
	 res,err := stream.Recv()

	   if err == io.EOF {
		   break
	   }
	   if err != nil {
		   log.Println("Error: ",err)
	   }

	   log.Println("Names of cities: ", res)
   }


}
