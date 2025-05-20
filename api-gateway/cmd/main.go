package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/internal/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	pbinv "api-gateway/proto/inventory"
	pborder "api-gateway/proto/order"
	pbuser "api-gateway/proto/user"
)

func main() {
	cfg := config.NewConfig()

	// Connect to gRPC services
	inventoryConn, err := grpc.Dial(cfg.InventoryServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to inventory service: %v", err)
	}
	defer inventoryConn.Close()

	orderConn, err := grpc.Dial(cfg.OrderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to order service: %v", err)
	}
	defer orderConn.Close()

	userConn, err := grpc.Dial(cfg.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	// Create gRPC clients
	inventoryClient := pbinv.NewInventoryServiceClient(inventoryConn)
	orderClient := pborder.NewOrderServiceClient(orderConn)
	userClient := pbuser.NewUserServiceClient(userConn)

	// Setup Gin
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware())
	// router.Use(middleware.AuthMiddleware()) // Uncomment if you want auth

	h := handler.NewGatewayHandler(inventoryClient, orderClient, userClient)

	// Product routes
	router.POST("/products", h.CreateProduct)
	router.GET("/products/:id", h.GetProduct)
	router.GET("/products", h.ListProducts)
	router.PUT("/products/:id", h.UpdateProduct)
	// Order routes
	router.POST("/orders", h.CreateOrder)
	router.GET("/orders/:id", h.GetOrder)
	router.GET("/orders", h.ListOrders) // <-- Add this line
	// User routes
	router.POST("/users/register", h.RegisterUser)
	router.POST("/users/login", h.AuthenticateUser)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/debug/routes", func(c *gin.Context) {
		c.JSON(200, gin.H{"routes": router.Routes()})
	})

	// Graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("API Gateway running on port %s", cfg.HTTPPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down API Gateway...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
