package router

import (
	"github.com/gofiber/fiber/v2"
	"alignspace/config"
)

func SetupRoutes(app *fiber.App) {
	routes := config.GetRoutes()

	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}
}
