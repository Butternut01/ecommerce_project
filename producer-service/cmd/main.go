package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "producer-service/proto"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Connect to Order Service
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer conn.Close()

	orderClient := pb.NewOrderServiceClient(conn)

	// Track published orders
	var publishedOrders sync.Map
	lastChecked := time.Now().Add(-1 * time.Minute) // Start by checking recent orders

	// Create a context for the stream
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start order processing loop
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Get orders created since last check
				req := &pb.ListOrdersRequest{
					Status: "pending",
					Limit:  100,
					// Add timestamp filter to only get new orders
					CreatedAfter: lastChecked.Unix(),
				}

				// Update last checked time
				lastChecked = time.Now()

				stream, err := orderClient.ListOrders(ctx, req)
				if err != nil {
					log.Printf("Failed to get orders: %v", err)
					continue
				}

				for {
					res, err := stream.Recv()
					if err != nil {
						// EOF is normal when stream ends
						break
					}

					for _, order := range res.Orders {
						if _, loaded := publishedOrders.LoadOrStore(order.Id, true); !loaded {
							orderData, err := json.Marshal(order)
							if err != nil {
								log.Printf("Failed to marshal order: %v", err)
								continue
							}

							if err := nc.Publish("order.created", orderData); err != nil {
								log.Printf("Failed to publish event: %v", err)
								publishedOrders.Delete(order.Id) // Allow retry
							} else {
								log.Printf("Published order %s", order.Id)
							}
						}
					}
				}
			}
		}
	}()

	log.Println("Producer service running...")
	<-quit
	log.Println("Shutting down producer-service...")
}