package routes

import (
	"backend/core/pkg/scope"
	"backend/internal/api/routes/maintenance"
	"backend/internal/api/routes/shops"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, scope *scope.Scope) {
	api := app.Group("/api")

	maintenance.RegisterRoutes(api, scope)
	shops.RegisterRoutes(api, scope)
}
