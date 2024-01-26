package database

import (
	"github.com/upper/db/v4"
)

type Storage interface {
	Session() db.Session
	Close() error
	CreateProductTable() error
}
