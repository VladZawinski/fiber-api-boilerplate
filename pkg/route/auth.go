package route

import (
	"futbook/app/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	// Create route group.
	route := a.Group("/api/v1/auth/")
	route.Post("login", controller.Login)
	route.Post("register", controller.Register)
	// route.Post("/token/new", controller.GetNewAccessToken)
	// route.Get("/books", controller.GetBooks)
	// route.Get("/books/:id", controller.GetBook)

}
