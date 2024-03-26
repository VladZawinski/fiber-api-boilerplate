package route

import (
	"futbook/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App) {
	// Create route group.
	route := a.Group("/api/v1/", middleware.JWTProtected())
	route.Post("play/", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return nil
	})
	// route.Post("/token/new", controller.GetNewAccessToken)
	// route.Get("/books", controller.GetBooks)
	// route.Get("/books/:id", controller.GetBook)

}
