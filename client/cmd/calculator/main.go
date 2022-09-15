package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/tagpro/grpc-example/pkg/tagpro/services/calculator"
)

// connects with the Calculator gRPC server
func main() {
	// Create a connection to the server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	add(ctx, conn)
}

func add(ctx context.Context, conn *grpc.ClientConn) {
	// Create a client
	client := v1.NewCalculatorClient(conn)

	// Create a context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Call the Add method
	req := &v1.AddRequest{A: 1, B: 2}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalf("error while calling Add RPC: %v", err)
	}
	log.Printf("Response from Add: %v", res.Result)
}
