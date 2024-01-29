package server

import (
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *server) SetupRoutes() {
	main := s.app.Group("/")
	main.Get("/logout", s.Logout)
	main.Get(
		"/metrics", monitor.New(
			monitor.Config{
				Title: "My service Metrics Page",
			},
		),
	)

	api := s.app.Group("/api/v1/")

	api.Use(
		recoverFiber.New(),
		s.AuthRequire(),
		pprof.New(
			pprof.Config{
				Prefix: "/profiler",
			},
		),
	)

	api.Get("/settings", s.GetSettings)
	api.Get("/params/*/*/*", s.PrintAllParams)
	api.Get("/info", s.Info)

	api.Post("/dummy", s.InsertDummyData)
	api.Get("/", s.GetAllProducts)
	api.Get("/:id", s.GetSingleProduct)
	api.Post("/", s.CreateProduct)
	api.Delete("/:id", s.DeleteProduct)
}
