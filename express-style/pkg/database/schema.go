package database

import (
	"github.com/gofiber/fiber/v2/log"
)

func (s *StorageDB) CreateProductTable() {
	_, err := s.DB.Query(
		`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			amount integer,
			name text UNIQUE,
			description text,
			category text NOT NULL
		);
	`,
	)
	if err != nil {
		log.Fatal(err)
	}
}
