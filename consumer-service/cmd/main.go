package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "consumer-service/proto"
	pborder "consumer-service/proto"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Connect to Inventory Service
	inventoryConn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}
	defer inventoryConn.Close()
	inventoryClient := pb.NewInventoryServiceClient(inventoryConn)

	// Connect to Order Service
	orderConn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderConn.Close()
	orderClient := pborder.NewOrderServiceClient(orderConn)

	// Subscribe to order.created events
	_, err = nc.Subscribe("order.created", func(msg *nats.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var order pborder.OrderResponse
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Failed to unmarshal order: %v", err)
			return
		}

		log.Printf("Processing order %s with %d items", order.Id, len(order.Items))

		// 1. Validate inventory
		var insufficientItems []string
		for _, item := range order.Items {
			product, err := inventoryClient.GetProduct(ctx, &pb.GetProductRequest{Id: item.ProductId})
			if err != nil {
				log.Printf("Product %s not found: %v", item.ProductId, err)
				insufficientItems = append(insufficientItems, item.ProductId)
				continue
			}
			if product.Stock < item.Quantity {
				log.Printf("Insufficient stock for %s (has %d, needs %d)",
					item.ProductId, product.Stock, item.Quantity)
				insufficientItems = append(insufficientItems, item.ProductId)
			}
		}

		if len(insufficientItems) > 0 {
			updateOrderStatus(orderClient, ctx, order.Id, "failed", "insufficient stock for: "+strings.Join(insufficientItems, ","))
			return
		}

		// 2. Update inventory atomically
		for _, item := range order.Items {
			product, err := inventoryClient.GetProduct(ctx, &pb.GetProductRequest{Id: item.ProductId})
			if err != nil || product.Stock < item.Quantity {
				log.Printf("Race condition or missing product: %s", item.ProductId)
				updateOrderStatus(orderClient, ctx, order.Id, "failed", "inventory conflict for "+item.ProductId)
				return
			}

			_, err = inventoryClient.UpdateProduct(ctx, &pb.ProductRequest{
				Id:    item.ProductId,
				Stock: product.Stock - item.Quantity,
			})
			if err != nil {
				log.Printf("Failed to update product %s: %v", item.ProductId, err)
				updateOrderStatus(orderClient, ctx, order.Id, "failed", "update failed for "+item.ProductId)
				return
			}
			log.Printf("Updated stock for %s to %d", item.ProductId, product.Stock-item.Quantity)
		}

		// 3. Finalize order
		updateOrderStatus(orderClient, ctx, order.Id, "completed", "")
		log.Printf("Order %s completed successfully", order.Id)
	})

	if err != nil {
		log.Fatalf("Failed to subscribe to NATS: %v", err)
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Consumer service running. Waiting for events...")
	<-quit
	log.Println("Shutting down consumer service...")
}

func updateOrderStatus(client pborder.OrderServiceClient, ctx context.Context, orderID, status, reason string) {
	_, err := client.UpdateOrderStatus(ctx, &pborder.UpdateOrderStatusRequest{
		Id:     orderID,
		Status: status,
	})
	if err != nil {
		log.Printf("Failed to update order %s status to %s: %v", orderID, status, err)
	}
}
