package product

import (
	"express-style/internal/domain/product"
	"express-style/internal/service/product/converter"
)

func (s *service) InsertDummyData() error {
	return s.productRepository.InsertDummyData()
}

func (s *service) CreateProduct(model *product.Product) error {
	return s.productRepository.CreateProduct(converter.DomainToDAO(model))
}

func (s *service) GetAllProducts() ([]*product.Product, error) {
	return converter.DAOsToDomains(s.productRepository.GetAllProducts())
}

func (s *service) GetProducts(pageSize, pageNum uint) ([]*product.Product, uint64, uint, error) {
	d, count, pages, err := s.productRepository.GetProducts(pageSize, pageNum)
	p, _ := converter.DAOsToDomains(d, nil)
	return p, count, pages, err
}

func (s *service) GetSingleProduct(id int64) (*product.Product, error) {
	d, err := s.productRepository.GetSingleProduct(id)
	return converter.DAOToDomain(d), err
}

func (s *service) UpdateProduct(model *product.Product) error {
	return nil
}

func (s *service) DeleteProduct(id int64) error {
	return nil
}
