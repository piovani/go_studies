package cmd

import (
	"context"
	"log"
	"net"

	"github.com/piovani/go_studies/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{Message: "Hello, " + in.GetName()}, nil
}

var Api = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteAPI()
	},
}

func ExecuteAPI() {
	println("Running")

	listerner, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})
	if err = s.Serve(listerner); err != nil {
		log.Fatal(err)
	}
}
