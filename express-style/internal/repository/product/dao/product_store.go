package dao

import (
	"github.com/pkg/errors"
	"github.com/upper/db/v4"
)

var (
	_ = db.Store(&ProductStore{})
)

type ProductStore struct {
	db.Collection
}

func Products(session db.Session) *ProductStore {
	return &ProductStore{
		session.Collection("products"),
	}
}

func (p *ProductStore) FindByName(name string) (*ProductDAO, error) {
	var d ProductDAO

	if err := p.Find(db.Cond{"product_name": name}).One(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (p *ProductStore) GetProducts(pageSize, pageNum uint) (d []*ProductDAO, count uint64, pages uint, err error) {
	res := p.Find().Paginate(pageSize)
	if err = res.Page(pageNum).All(&d); err != nil {
		return nil, 0, 0, errors.New("Products: failed to get products")
	}

	count, err = res.TotalEntries()
	if err != nil {
		return nil, 0, 0, errors.New("Products: failed to total entries")
	}
	pages, err = res.TotalPages()
	if err != nil {
		return nil, 0, 0, errors.New("Products: failed to total pages")
	}

	return
}
