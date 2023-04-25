package main

import (
	"context"
	"log"
	"net"

	"github.com/piovani/go_studies/grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	println("Running")

	listerner, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})
	if err = s.Serve(listerner); err != nil {
		log.Fatal(err)
	}

}
