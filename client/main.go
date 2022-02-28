package main

import (
	"context"
	"log"
	"time"

	pb "github.com/axyan/grpc-tutorial/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTokenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Generate(ctx, &pb.TokenRequest{Id: "random id here"})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	log.Printf("received message: %v", r.GetToken())

	x, err := c.Validate(ctx, &pb.TokenString{Token: "random token string"})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	log.Printf("received message: %v", x.GetValid())
}
