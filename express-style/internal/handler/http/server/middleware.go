package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"
)

func (s *server) AuthRequire() func(*fiber.Ctx) error {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	cfg := basicauth.Config{
		Users: map[string]string{
			user: pass,
		},
	}
	return basicauth.New(cfg)
}
