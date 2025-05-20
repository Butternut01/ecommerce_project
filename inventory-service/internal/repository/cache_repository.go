package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"inventory-service/internal/entity"
)

type ProductCacheRepository interface {
	GetProduct(ctx context.Context, id string) (*entity.Product, error)
	SetProduct(ctx context.Context, product *entity.Product, expiration time.Duration) error
	DeleteProduct(ctx context.Context, id string) error
}

type productCacheRepository struct {
	client *redis.Client
}

func NewProductCacheRepository(client *redis.Client) ProductCacheRepository {
	return &productCacheRepository{client: client}
}

func (r *productCacheRepository) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	val, err := r.client.Get(ctx, "product:"+id).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var product entity.Product
	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productCacheRepository) SetProduct(ctx context.Context, product *entity.Product, expiration time.Duration) error {
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, "product:"+product.ID, data, expiration).Err()
}

func (r *productCacheRepository) DeleteProduct(ctx context.Context, id string) error {
	return r.client.Del(ctx, "product:"+id).Err()
}