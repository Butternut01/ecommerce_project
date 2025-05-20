package usecase

import (
	"context"
	"log"
	"time"

	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
)

type ProductUseCase struct {
	productRepo repository.ProductRepository
	cacheRepo   repository.ProductCacheRepository
}

func NewProductUseCase(
	productRepo repository.ProductRepository,
	cacheRepo repository.ProductCacheRepository,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
		cacheRepo:   cacheRepo,
	}
}

func (uc *ProductUseCase) CreateProduct(product *entity.Product) error {
	return uc.productRepo.Create(product)
}

func (uc *ProductUseCase) GetProduct(id string) (*entity.Product, error) {
	ctx := context.Background()

	// Try to get product from cache
	cachedProduct, err := uc.cacheRepo.GetProduct(ctx, id)
	if err != nil {
		log.Printf("Cache get error for product %s: %v", id, err)
	}
	if cachedProduct != nil {
		return cachedProduct, nil
	}

	// Cache miss - get from DB
	product, err := uc.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Cache the product for 5 minutes
	if err := uc.cacheRepo.SetProduct(ctx, product, 5*time.Minute); err != nil {
		log.Printf("Failed to cache product %s: %v", id, err)
	}

	return product, nil
}

func (uc *ProductUseCase) UpdateProduct(product *entity.Product) error {
	if err := uc.productRepo.Update(product); err != nil {
		return err
	}

	// Invalidate cache
	ctx := context.Background()
	if err := uc.cacheRepo.DeleteProduct(ctx, product.ID); err != nil {
		log.Printf("Failed to invalidate cache for product %s: %v", product.ID, err)
	}

	return nil
}

func (uc *ProductUseCase) DeleteProduct(id string) error {
	if err := uc.productRepo.Delete(id); err != nil {
		return err
	}

	// Invalidate cache
	ctx := context.Background()
	if err := uc.cacheRepo.DeleteProduct(ctx, id); err != nil {
		log.Printf("Failed to invalidate cache for product %s: %v", id, err)
	}

	return nil
}

func (uc *ProductUseCase) ListProducts(filter entity.ProductFilter) ([]entity.Product, error) {
	return uc.productRepo.FindAll(filter)
}
