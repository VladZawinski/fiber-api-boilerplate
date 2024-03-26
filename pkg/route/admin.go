package route

import (
	"futbook/pkg/middleware"
	"futbook/platform/logger"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private route.
func AdminRoutes(a *fiber.App) {
	// Admin route group
	adminRoute := a.Group("/api/v1/admin", middleware.JWTProtected(), middleware.IsAdmin)
	// User
	adminRoute.Post("/", func(c *fiber.Ctx) error {
		logger.GetLogger().Debug("Hello")
		c.SendStatus(200)
		return nil
	})
	// adminRoute.Get("/", controller.GetUsers)
	// adminRoute.Get("/:id", controller.GetUser)
	// adminRoute.Put("/:id", controller.UpdateUser)
	// adminRoute.Delete("/:id", controller.DeleteUser)

	// // Book
	// route := a.Group("/api/v1/books", middleware.JWTProtected())
	// route.Post("/", controller.CreateBook)
	// route.Put("/:id", controller.UpdateBook)
	// route.Delete("/:id", controller.DeleteBook)

}
