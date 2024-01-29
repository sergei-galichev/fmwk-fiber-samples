package product

import (
	"express-style/internal/repository/product/dao"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
	"github.com/upper/db/v4"
)

func (r *repository) InsertDummyData() error {
	for i := 0; i < 10_000; {
		prod := gofakeit.Product()

		_, err := r.storage.Session().SQL().InsertInto("products").Values(
			&dao.ProductDAO{
				Amount:      int64(gofakeit.IntRange(5, 70)),
				Name:        prod.Name,
				Description: prod.Description,
				Category:    prod.Categories[0],
			},
		).Exec()
		if err != nil {
			log.Error(err)
		}

		i++
	}

	return nil
}

func (r *repository) CreateProduct(model *dao.ProductDAO) error {
	_, err := r.storage.Session().SQL().InsertInto("products").Values(model).Exec()
	if err != nil {
		return errors.New("repo: error creating product")
	}
	return nil
}

func (r *repository) GetAllProducts() ([]*dao.ProductDAO, error) {
	result := r.storage.Session().Collection("products").Find()
	defer r.closeResultSet(result)

	var prods []*dao.ProductDAO
	err := result.All(&prods)
	if err != nil {
		return nil, errors.New("repo: error get all products")
	}
	return prods, nil
}

func (r *repository) GetSingleProduct(id int64) (*dao.ProductDAO, error) {
	var d dao.ProductDAO
	err := r.storage.Session().Get(&d, db.Cond{"id": id})
	if err != nil {
		return nil, errors.New("repo: error get single product")
	}

	return &d, nil
}

func (r *repository) UpdateProduct(model *dao.ProductDAO) error {
	return nil
}

func (r *repository) DeleteProduct(id int64) error {
	return nil
}

func (r *repository) closeResultSet(result db.Result) {
	err := result.Close()
	if err != nil {
		log.Error("repo: error close result set")
	}
}
