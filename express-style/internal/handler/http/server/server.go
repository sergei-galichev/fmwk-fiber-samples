package server

import (
	"express-style/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pkg/errors"
	"net"
	"os"
)

type server struct {
	app            *fiber.App
	productService service.ProductService
}

func NewServer(productService service.ProductService) *server {
	app := fiber.New()

	log.SetLevel(log.LevelInfo)
	app.Use(
		logger.New(
			logger.Config{
				Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
			},
		),
		cors.New(
			cors.Config{
				AllowHeaders: fiber.HeaderAuthorization,
			},
		),
	)

	return &server{
		app:            app,
		productService: productService,
	}
}

func (s *server) Run() error {
	addr := net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT"))
	err := s.app.Listen(addr)
	if err != nil {
		return errors.New("Server: error starting")
	}

	return nil
}

func (s *server) Shutdown() error {
	return s.app.Shutdown()
}
