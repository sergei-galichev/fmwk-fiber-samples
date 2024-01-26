package dao

import (
	"github.com/upper/db/v4"
)

var (
	_ = db.Record(&ProductDAO{})
)

type ProductDAO struct {
	ID          int64  `db:"id,omitempty"`
	Amount      int64  `db:"amount"`
	Name        string `db:"product_name"`
	Description string `db:"description"`
	Category    string `db:"category"`
}

func (*ProductDAO) Store(session db.Session) db.Store {
	return session.Collection("products")
}
