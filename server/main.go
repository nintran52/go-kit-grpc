package main

import (
	"log"
	"net"

	"github.com/nintran52/go-kit-grpc/endpoint"
	pb "github.com/nintran52/go-kit-grpc/proto"
	"github.com/nintran52/go-kit-grpc/service"
	"github.com/nintran52/go-kit-grpc/transport"
	"google.golang.org/grpc"
)

func main() {
	// Create service instance
	svc := service.NewGreeterService()
	sayHelloEndpoint := endpoint.MakeSayHelloEndpoint(svc)

	// Initialize gRPC server
	grpcServer := transport.NewGRPCServer(sayHelloEndpoint)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, grpcServer)

	log.Println("gRPC server is listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
