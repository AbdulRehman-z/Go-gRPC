package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
)


const (
	port = ":8000"
)


func main() {

	connection,err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkFatalError(err)

	//defer func(connection *grpc.ClientConn) {
	//	err := connection.Close()
	//	if err != nil {
	//		fmt.Printf("Error: %s", err)
	//	}
	//}(connection)



defer func(connection *grpc.ClientConn) {
	err := connection.Close()
	if err != nil {
		fmt.Printf("Errormlmwdklaskl: %s", err)
	}
}(connection)

   client := pb.NewWelcomeServiceClient(connection)
   namesOfCities := &pb.CitiesArray{
	   Cities: []string{"Karachi","lahore"},
   }

   // unary Api
   // CallWelcome(client)

   // serverSideStreaming Api
   //serverSideStreaming(client, namesOfCities)

   // clientSideStreaming
   //clientSideStreaming(client,namesOfCities)

   // bidirectionalStreaming
   bidirectionalClientStreaming(client,namesOfCities)
}

func checkFatalError(err error){
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}