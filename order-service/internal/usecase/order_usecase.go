package usecase

import (
	"order-service/internal/entity"
	"order-service/internal/repository"
	"time"
)

type OrderUseCase interface {
	CreateOrder(order *entity.Order) error
	GetOrder(id string) (*entity.Order, error)
	UpdateOrderStatus(id string, status entity.OrderStatus) error
	ListOrders(filter entity.OrderFilter) ([]entity.Order, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}

func (uc *orderUseCase) CreateOrder(order *entity.Order) error {
	// Calculate total if not set
	if order.Total == 0 {
		for _, item := range order.Items {
			order.Total += item.Price * float64(item.Quantity)
		}
	}
	
	// Set timestamps
	now := time.Now().Unix()
	order.CreatedAt = now
	order.UpdatedAt = now
	
	return uc.orderRepo.Create(order)
}


func (uc *orderUseCase) GetOrder(id string) (*entity.Order, error) {
    return uc.orderRepo.FindByID(id)
}

func (uc *orderUseCase) UpdateOrderStatus(id string, status entity.OrderStatus) error {
    return uc.orderRepo.UpdateStatus(id, status)
}

func (uc *orderUseCase) ListOrders(filter entity.OrderFilter) ([]entity.Order, error) {
    return uc.orderRepo.FindAll(filter)
}
