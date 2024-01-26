package app

import (
	repositories "express-style/internal/repository"
	productRepo "express-style/internal/repository/product"
	services "express-style/internal/service"
	productService "express-style/internal/service/product"
	"express-style/pkg/database"
	"express-style/pkg/database/postgres"
	"github.com/gofiber/fiber/v2/log"
)

type ServiceProvider struct {
	productService services.ProductService
	productRepo    repositories.ProductRepository
	storage        database.Storage
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) Storage() database.Storage {
	if sp.storage == nil {
		var err error
		sp.storage, err = postgres.NewStorage()
		if err != nil {
			log.Fatal(err)
		}
		err = sp.storage.CreateProductTable()
		if err != nil {
			log.Fatal(err)
		}
	}
	return sp.storage
}

func (sp *ServiceProvider) ProductRepository() repositories.ProductRepository {
	if sp.productRepo == nil {
		sp.productRepo = productRepo.NewRepository(sp.Storage())
	}
	return sp.productRepo
}

func (sp *ServiceProvider) ProductService() services.ProductService {
	if sp.productService == nil {
		sp.productService = productService.NewService(sp.ProductRepository())
	}
	return sp.productService
}
