package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"order-service/internal/entity"
	"order-service/internal/usecase"
	pb "order-service/proto"
)

type OrderController struct {
	pb.UnimplementedOrderServiceServer
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
	}
}

func (c *OrderController) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order := &entity.Order{
		UserID: req.GetUserId(),
		Total:  req.GetTotal(),
		Status: entity.OrderStatusPending,
	}

	for _, item := range req.GetItems() {
		order.Items = append(order.Items, entity.OrderItem{
			ProductID: item.GetProductId(),
			Quantity:  int(item.GetQuantity()),
			Price:     item.GetPrice(),
		})
	}

	if err := c.orderUseCase.CreateOrder(order); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	return convertOrderToResponse(order), nil
}

func (c *OrderController) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := c.orderUseCase.GetOrder(req.GetId())
	if err != nil {
		if errors.Is(err, entity.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, "order not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get order: %v", err)
	}

	return convertOrderToResponse(order), nil
}

func (c *OrderController) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.OrderResponse, error) {
	orderStatus := entity.OrderStatus(req.GetStatus())
	if !orderStatus.IsValid() {
		return nil, status.Errorf(codes.InvalidArgument, "invalid order status")
	}

	if err := c.orderUseCase.UpdateOrderStatus(req.GetId(), orderStatus); err != nil {
		if errors.Is(err, entity.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, "order not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update order status: %v", err)
	}

	// Fetch updated order to return
	order, err := c.orderUseCase.GetOrder(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch updated order: %v", err)
	}

	return convertOrderToResponse(order), nil
}

func (c *OrderController) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	filter := entity.OrderFilter{
		UserID: req.GetUserId(),
		Status: entity.OrderStatus(req.GetStatus()),
		Page:   int(req.GetPage()),
		Limit:  int(req.GetLimit()),
	}

	orders, err := c.orderUseCase.ListOrders(filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders: %v", err)
	}

	var responses []*pb.OrderResponse
	for _, order := range orders {
		responses = append(responses, convertOrderToResponse(&order))
	}

	return &pb.ListOrdersResponse{Orders: responses}, nil
}

func convertOrderToResponse(order *entity.Order) *pb.OrderResponse {
	var items []*pb.OrderItem
	for _, item := range order.Items {
		items = append(items, &pb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		})
	}

	return &pb.OrderResponse{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     items,
		Total:     order.Total,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}