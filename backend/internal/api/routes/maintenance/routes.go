package maintenance

import (
	"backend/core/pkg/scope"
	"backend/internal/api/handlers/maintenance"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, scope *scope.Scope) {
	group := router.Group("/maintenance")

	group.Get("/health", handler_maintenance.Health(scope)...)
}
