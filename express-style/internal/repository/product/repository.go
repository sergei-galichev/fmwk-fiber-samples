package product

import (
	repositories "express-style/internal/repository"
	"express-style/pkg/database"
)

var _ repositories.ProductRepository = (*repository)(nil)

type repository struct {
	storage database.Storage
}

func NewRepository(storage database.Storage) *repository {
	return &repository{
		storage: storage,
	}
}
