package entity

import "errors"

type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Stock       int     `bson:"stock"`
	Category    string  `bson:"category"`
}

type ProductFilter struct {
	Name     string
	Category string
	MinPrice float64
	MaxPrice float64
	Page     int
	Limit    int
}

var (
	ErrProductNotFound = errors.New("product not found")
)
