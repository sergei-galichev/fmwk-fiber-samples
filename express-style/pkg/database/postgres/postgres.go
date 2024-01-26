package postgres

import (
	"express-style/pkg/database"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
	"strings"
)

var _ database.Storage = (*postgresStorage)(nil)

type postgresStorage struct {
	session db.Session
}

func NewStorage() (*postgresStorage, error) {
	var host strings.Builder
	//host.WriteString(os.Getenv("PG_HOST"))
	host.WriteString("localhost")
	host.WriteString(":")
	//host.WriteString(os.Getenv("PG_PORT"))
	host.WriteString(os.Getenv("PG_EXT_PORT"))
	log.Info(host.String())

	url := postgresql.ConnectionURL{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		Host:     host.String(),
		Database: os.Getenv("PG_DB_NAME"),
		Options: map[string]string{
			"sslmode": os.Getenv("PG_SSL"),
		},
	}

	//db.LC().SetLevel(db.LogLevelDebug)

	session, err := postgresql.Open(url)
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
			product_name text UNIQUE,
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
