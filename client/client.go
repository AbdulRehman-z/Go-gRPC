package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const (
	port = ":8000"
)


func main() {

	connection,err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkFatalError(err)

	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(connection)
	

}

func checkFatalError(err error){
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}