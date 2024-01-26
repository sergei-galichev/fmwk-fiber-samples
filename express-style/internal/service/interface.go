package service

import (
	"express-style/internal/domain/product"
)

type ProductService interface {
	CreateProduct(model *product.Product) error
	GetAllProducts() ([]*product.Product, error)
	GetSingleProduct(id int64) (*product.Product, error)
	UpdateProduct(model *product.Product) error
	DeleteProduct(id int64) error
}
