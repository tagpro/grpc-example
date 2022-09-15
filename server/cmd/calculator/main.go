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

const port = "50051"

func main() {
	// Boilerplate to start service
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

	server := grpc.NewServer()
	reflection.Register(server)

	// Connect service code to listener port
	// Register API v1
	calculatorService := calculator.NewCalculatorService()
	v1.RegisterCalculatorServer(server, calculatorService)

	log.Printf("listening on port %s", exposedPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed at: %v", err)
	}
}
