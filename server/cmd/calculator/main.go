package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/tagpro/grpc-example/pkg/tagpro/services/calculator"
	"github.com/tagpro/grpc-example/server/pkg/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Boilerplate to start service
const port = "50051"

func main() {

	// Create a listener on TCP port
	exposedPort := os.Getenv("PORT")
	if exposedPort == "" {
		exposedPort = port
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", exposedPort))
	if err != nil {
		log.Fatal("unable to listen on port ", exposedPort)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Println(err)
		}
	}(lis)

	// Create a gRPC server object
	server := grpc.NewServer()

	// Attach server reflection. This helps clients like Postman or bloomRPC to identify the methods in service
	reflection.Register(server)

	// Register API with gRPC server
	calculatorService := calculator.NewCalculatorService()
	v1.RegisterCalculatorServer(server, calculatorService)

	log.Printf("listening on port %s", exposedPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed at: %v", err)
	}
}
