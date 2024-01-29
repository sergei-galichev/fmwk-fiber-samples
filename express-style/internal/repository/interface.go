package repository

import (
	"express-style/internal/repository/product/dao"
)

type ProductRepository interface {
	InsertDummyData() error
	CreateProduct(model *dao.ProductDAO) error
	GetAllProducts() ([]*dao.ProductDAO, error)
	GetSingleProduct(id int64) (*dao.ProductDAO, error)
	UpdateProduct(model *dao.ProductDAO) error
	DeleteProduct(id int64) error
}
