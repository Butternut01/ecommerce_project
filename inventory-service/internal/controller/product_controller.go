package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto"
)

type ProductController struct {
	pb.UnimplementedInventoryServiceServer
	productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
	}
}

func (c *ProductController) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product := &entity.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Stock:       int(req.GetStock()),
		Category:    req.GetCategory(),
	}

	if err := c.productUseCase.CreateProduct(product); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		Category:    product.Category,
	}, nil
}

func (c *ProductController) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	product, err := c.productUseCase.GetProduct(req.GetId())
	if err != nil {
		if errors.Is(err, entity.ErrProductNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		Category:    product.Category,
	}, nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product := &entity.Product{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Stock:       int(req.GetStock()),
		Category:    req.GetCategory(),
	}

	if err := c.productUseCase.UpdateProduct(product); err != nil {
		if errors.Is(err, entity.ErrProductNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		Category:    product.Category,
	}, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if err := c.productUseCase.DeleteProduct(req.GetId()); err != nil {
		if errors.Is(err, entity.ErrProductNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}

	return &pb.DeleteProductResponse{Success: true}, nil
}

func (c *ProductController) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	filter := entity.ProductFilter{
		Name:     req.GetName(),
		Category: req.GetCategory(),
		MinPrice: req.GetMinPrice(),
		MaxPrice: req.GetMaxPrice(),
		Page:     int(req.GetPage()),
		Limit:    int(req.GetLimit()),
	}

	products, err := c.productUseCase.ListProducts(filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	var productResponses []*pb.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &pb.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       int32(product.Stock),
			Category:    product.Category,
		})
	}

	return &pb.ListProductsResponse{Products: productResponses}, nil
}