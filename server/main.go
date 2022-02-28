package main

import (
	"context"
	"log"
	"net"

	pb "github.com/axyan/grpc-tutorial/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTokenServer
}

func (s *server) Generate(ctx context.Context, id *pb.TokenRequest) (*pb.TokenString, error) {
	log.Printf("Received: %v", id.String())
	return &pb.TokenString{Token: "token here"}, nil
}

func (s *server) Validate(ctx context.Context, token *pb.TokenString) (*pb.ParsedToken, error) {
	log.Printf("Received: %v", token.String())
	return &pb.ParsedToken{Valid: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTokenServer(s, &server{})
	log.Printf("lisening on: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
