package postgres

import (
	"database/sql"
	"express-style/pkg/database"
	//_ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
)

var _ database.Storage = (*postgresStorage)(nil)

type postgresStorage struct {
	session db.Session
}

func NewStorage() (*postgresStorage, error) {
	dsn := os.Getenv("LOC_DSN")
	connDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, errors.New("postgres: database connection error")
	}

	session, err := postgresql.New(connDB)
	if err != nil {
		return nil, errors.New("postgres: database connection error")
	}

	err = session.Ping()
	if err != nil {
		return nil, errors.New("postgres: error ping database")
	}

	return &postgresStorage{
		session: session,
	}, nil
}

func (s *postgresStorage) Session() db.Session {
	return s.session
}

func (s *postgresStorage) Close() error {
	err := s.session.Close()
	if err != nil {
		return errors.New("")
	}

	return nil
}

func (s *postgresStorage) CreateProductTable() error {
	_, err := s.session.SQL().Query(
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			amount integer,
			product_name text,
			description text,
			category text,
			customer_id integer
		);
	`,
	)
	if err != nil {
		return errors.New("postgres: error create product table")
	}

	_, err = s.session.SQL().Query(
		`CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			first_name text,
			last_name text,
			age integer
		);
	`,
	)

	return nil
}
