package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
)

type StorageDB struct {
	DB *sql.DB
}

func NewStorage() *StorageDB {
	return &StorageDB{}
}

func (s *StorageDB) Connect() error {
	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return errors.New("error parsing string to uint")
	}
	s.DB, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			port,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return errors.New("failed to connect to database")
	}

	if err = s.DB.Ping(); err != nil {
		return errors.New("failed to ping database")
	}

	s.CreateProductTable()
	log.Info("Database connection is opened!")
	return nil
}
