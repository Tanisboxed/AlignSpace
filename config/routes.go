package config

import (
	"github.com/gofiber/fiber/v2"
	"alignspace/internal/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

func GetRoutes() []Route {
	return []Route{
		/********** Misc Routes **********/
		{
			Method:  "GET",
			Path:    "/health",
			Handler: handlers.HealthCheck,
		},
	}
}
