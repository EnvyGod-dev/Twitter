package routes

import (
	"twitter/handlers"
	"twitter/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(hd *handlers.Handlers) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.CustomOverErrMsg,
		BodyLimit:    100 * 1024 * 1024,
	})

	app.Use(cors.New(
		cors.Config{
			AllowMethods: "GET,POST,PATCH,DELETE,PUT",
			AllowOrigins: "*",
		},
	))

	TwitterRoutes(app, hd)
	return app
}
