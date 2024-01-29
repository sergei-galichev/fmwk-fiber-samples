package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"
)

func (s *server) AuthRequire() fiber.Handler {

	u := os.Getenv("U_NAME")
	p := os.Getenv("U_PASS")

	cfg := basicauth.Config{
		Users: map[string]string{
			u: p,
		},
		Authorizer: func(user string, pass string) bool {
			if user == "" || pass == "" {
				return false
			} else if user == u && pass == p {
				return true
			} else {
				return false
			}
		},
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				&fiber.Map{
					"success": false,
					"message": "Unauthorized",
				},
			)
		},
	}

	return basicauth.New(cfg)
}
