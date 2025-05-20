package main

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"

	"inventory-service/internal/config"
	"inventory-service/internal/controller"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	defer redisClient.Close()

	// Test Redis connection
	ctx := context.Background()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Connect to MongoDB
	db, err := config.ConnectMongoDB(cfg.MongoDBURI)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	cacheRepo := repository.NewProductCacheRepository(redisClient)

	// Initialize use case with Redis caching
	productUseCase := usecase.NewProductUseCase(productRepo, cacheRepo)

	// Initialize gRPC server and controller
	grpcServer := grpc.NewServer()
	productController := controller.NewProductController(*productUseCase)
	pb.RegisterInventoryServiceServer(grpcServer, productController)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":"+cfg.ServerPort)
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
	log.Printf("Inventory Service is running on port %s", cfg.ServerPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error serving gRPC: %v", err)
	}
}
