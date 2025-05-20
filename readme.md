# E-Commerce Platform with gRPC

## Overview

This project implements an e-commerce platform using microservices architecture with gRPC for inter-service communication. The platform includes the following services:

1. **API Gateway**: Exposes RESTful endpoints and forwards requests to backend services via gRPC.
2. **Inventory Service**: Manages products and categories.
3. **Order Service**: Handles orders and payments.
4. **User Service**: Manages user registration, authentication, and profiles.

## Setup Instructions

### Prerequisites

- Go 1.20+
- Protocol Buffers Compiler (`protoc`)
- MongoDB or any other database of your choice

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd e-commerce-platform
   ```

2. Generate gRPC stubs:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/*.proto
   ```

3. Start each service:
   - **Inventory Service**:
     ```bash
     go run inventory-service/cmd/main.go
     ```
   - **User Service**:
     ```bash
     go run user-service/cmd/main.go
     ```
   - **Order Service**:
     ```bash
     go run order-service/cmd/main.go
     ```

4. Start the API Gateway:
   ```bash
   go run api-gateway/cmd/main.go
   ```

5. Test the endpoints using a REST client like Postman.

## Endpoints

- **POST /products**: Create a new product.
- **POST /users/register**: Register a new user.
- **POST /orders**: Place a new order.

## Notes

- Ensure all services are running before testing.
- Update configuration files with your database credentials.