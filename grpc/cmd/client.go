package cmd

import (
	"context"
	"fmt"

	"github.com/piovani/go_studies/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Client = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteClient()
	},
}

func ExecuteClient() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: "Marcos",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
