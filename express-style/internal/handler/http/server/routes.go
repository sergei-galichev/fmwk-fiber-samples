package server

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *server) SetupRoutes() {
	//api := s.app.Group("/api/v1", logger.New(), s.AuthRequire())
	api := s.app.Group("/api/v1/", logger.New())
	//api.Static("/", "./web/static")

	// some test routes
	//api.Use(
	//	func(ctx *fiber.Ctx) error {
	//		return ctx.Bind(
	//			fiber.Map{
	//				"Title": "My custom title",
	//			},
	//		)
	//	},
	//)
	//api.Get(
	//	"/bind", func(ctx *fiber.Ctx) error {
	//		return ctx.Render("base.tmpl", fiber.Map{})
	//	},
	//)
	api.Get("/settings", s.GetSettings)
	api.Get("/params/*/*/*", s.PrintAllParams)
	api.Get("/info", s.Info)
	//

	api.Get("/", s.GetAllProducts)
	api.Get("/:id", s.GetSingleProduct)
	api.Post("/", s.CreateProduct)
	api.Delete("/:id", s.DeleteProduct)
}
