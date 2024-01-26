package product

import (
	"express-style/internal/domain/product"
	"express-style/internal/service/product/converter"
)

func (r *service) CreateProduct(model *product.Product) error {
	return r.productRepository.CreateProduct(converter.DomainToDAO(model))
}

func (r *service) GetAllProducts() ([]*product.Product, error) {
	return converter.DAOsToDomains(r.productRepository.GetAllProducts())
}

func (r *service) GetSingleProduct(id int64) (*product.Product, error) {
	return nil, nil
}

func (r *service) UpdateProduct(model *product.Product) error {
	return nil
}

func (r *service) DeleteProduct(id int64) error {
	return nil
}
