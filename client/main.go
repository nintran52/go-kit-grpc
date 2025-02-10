package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nintran52/go-kit-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send request to the server
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go-Kit"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.GetMessage())
}
