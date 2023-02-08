package main

import (
	"context"
	"fmt"
	pb "grpc/proto"
	"time"
)

func CallWelcome(client pb.WelcomeServiceClient)   {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
    defer cancel()

	response, err := client.Welcome(ctx,&pb.WelcomeRequest{Name: "Kamal"})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Println(response)
}