package server

import (
	"fmt"
	"futbook/pkg/config"
	"futbook/pkg/middleware"
	"futbook/pkg/route"
	"futbook/platform/database"
	"futbook/platform/logger"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	appCfg := config.AppCfg()
	// connect to DB
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("failed database setup. error: %v", err)
	}
	logger.SetUpLogger()
	logr := logger.GetLogger()
	// Define Fiber config & app.
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)
	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	route.GeneralRoute(app)
	route.PublicRoutes(app)
	route.AuthRoutes(app)
	route.UserRoute(app)
	route.AdminRoutes(app)
	route.NotFoundRoute(app)

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Error("Oops... server is not running! error: %v", err)
	}
}
