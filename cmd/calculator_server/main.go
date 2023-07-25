package main

import (
	"log"
	"net"

	pb "github.com/demoS11/calc-service/pkg/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// main is the entry point of the calculator server application.
func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server instance
	baseServer := grpc.NewServer()
	pb.RegisterCalculatorServer(baseServer, NewServer())

	reflection.Register(baseServer)

	log.Printf("Starting server on %s", lis.Addr().String())
	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
