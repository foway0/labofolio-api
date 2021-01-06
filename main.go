package main

import (
	"context"
	"github.com/foway0/labofolio-web-bff/hoge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	hoge.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hoge.HelloRequest) (*hoge.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hoge.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hoge.RegisterGreeterServer(s, &server{})
	// local only???
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
