package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	user "go_grpc/proto/gen"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	usr := user.NewUserServiceClient(conn)

	fmt.Print("Enter the UserName: ")

	var inputUser string
	fmt.Scanln(&inputUser)

	response, err := usr.GetUser(context.Background(), &user.UserRequest{Username: inputUser})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	log.Printf("Reponse from service: %v", response)
}
