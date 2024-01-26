package product

import (
	"express-style/internal/repository"
	services "express-style/internal/service"
)

var _ services.ProductService = (*service)(nil)

type service struct {
	productRepository repository.ProductRepository
}

func NewService(productRepository repository.ProductRepository) *service {
	return &service{
		productRepository: productRepository,
	}
}
