package server

import (
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *server) SetupRoutes() {
	main := s.app.Group("/")
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

	api.Post("/product/dummy", s.InsertDummyData)
	api.Get("/product/all", s.GetAllProducts)
	api.Get("/product/:page_size/:page_num", s.GetProducts)
	api.Get("/product/:id", s.GetSingleProduct)
	api.Post("/product/", s.CreateProduct)
	api.Delete("/product/:id", s.DeleteProduct)
}
