package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"order-service/internal/config"
	"order-service/internal/controller"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	pb "order-service/proto"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Connect to MongoDB client
	client, err := config.ConnectMongoDBClient(cfg.MongoDBURI) // Updated to get mongo.Client
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Select the database
	db := client.Database(cfg.MongoDBName)

	// Initialize repository with both db and client for transactions
	orderRepo := repository.NewOrderRepository(db, client)

	// Initialize use case
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	orderController := controller.NewOrderController(orderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderController)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":"+cfg.ServerPort)
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
	log.Printf("Order Service is running on port %s", cfg.ServerPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error serving gRPC: %v", err)
	}
}
