package converter

import (
	"express-style/internal/domain/product"
	"express-style/internal/repository/product/dao"
)

func DomainToDAO(model *product.Product) *dao.ProductDAO {
	return &dao.ProductDAO{
		ID:          model.ID,
		Amount:      model.Amount,
		Name:        model.Name,
		Description: model.Description,
		Category:    model.Category,
	}
}

func DAOToDomain(model *dao.ProductDAO) *product.Product {
	return &product.Product{
		ID:          model.ID,
		Amount:      model.Amount,
		Name:        model.Name,
		Description: model.Description,
		Category:    model.Category,
	}
}

func DAOsToDomains(dao []*dao.ProductDAO, err error) ([]*product.Product, error) {
	var models []*product.Product
	for _, d := range dao {
		models = append(models, DAOToDomain(d))
	}
	return models, err
}
