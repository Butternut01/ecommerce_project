package repository

import (
	"context"
	"time"

	"order-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository interface {
	Create(order *entity.Order) error
	FindByID(id string) (*entity.Order, error)
	UpdateStatus(id string, status entity.OrderStatus) error
	FindAll(filter entity.OrderFilter) ([]entity.Order, error)
}

type orderRepository struct {
	collection *mongo.Collection
	client     *mongo.Client // For transaction support
}

func NewOrderRepository(db *mongo.Database, client *mongo.Client) OrderRepository {
	return &orderRepository{
		collection: db.Collection("orders"),
		client:     client,
	}
}

func (r *orderRepository) Create(order *entity.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	session, err := r.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		order.CreatedAt = time.Now().Unix()
		order.UpdatedAt = order.CreatedAt

		if order.Total == 0 {
			for _, item := range order.Items {
				order.Total += item.Price * float64(item.Quantity)
			}
		}

		_, err := r.collection.InsertOne(sessCtx, order)
		if err != nil {
			return nil, err
		}

		// Add other transactional logic here if needed

		return nil, nil
	})

	return err
}

func (r *orderRepository) FindByID(id string) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order entity.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) UpdateStatus(id string, status entity.OrderStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now().Unix(),
		},
	}

	_, err = r.collection.UpdateByID(ctx, objectID, update)
	return err
}

func (r *orderRepository) FindAll(filter entity.OrderFilter) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := bson.M{}
	if filter.UserID != "" {
		query["user_id"] = filter.UserID
	}
	if filter.Status != "" {
		query["status"] = filter.Status
	}
	if filter.CreatedAfter > 0 {
		query["created_at"] = bson.M{"$gte": filter.CreatedAfter}
	}

	opts := options.Find()
	if filter.Limit > 0 {
		opts.SetLimit(int64(filter.Limit))
		if filter.Page > 0 {
			opts.SetSkip(int64((filter.Page - 1) * filter.Limit))
		}
	}

	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []entity.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
