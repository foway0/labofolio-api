package main

import (
	"context"
	"github.com/foway0/labofolio-web-bff/grpc_specs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	grpc_spec.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *grpc_spec.HelloRequest) (*grpc_spec.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &grpc_spec.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_spec.RegisterGreeterServer(s, &server{})
	// local only???
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
