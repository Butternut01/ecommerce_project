package entity

import "errors"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (os OrderStatus) IsValid() bool {
	switch os {
	case OrderStatusPending, OrderStatusCompleted, OrderStatusCancelled:
		return true
	default:
		return false
	}
}

type OrderItem struct {
	ProductID string  `bson:"product_id"`
	Quantity  int     `bson:"quantity"`
	Price     float64 `bson:"price"`
}

type Order struct {
	ID        string      `bson:"_id,omitempty"`
	UserID    string      `bson:"user_id"`
	Items     []OrderItem `bson:"items"`
	Total     float64     `bson:"total"`
	Status    OrderStatus `bson:"status"`
	CreatedAt int64       `bson:"created_at"`
	UpdatedAt int64       `bson:"updated_at"`
}

type OrderFilter struct {
    UserID      string
    Status      OrderStatus
    Page        int
    Limit       int
    CreatedAfter int64  // Add this new field
}

var (
	ErrOrderNotFound = errors.New("order not found")
)