package orders

import (
	"backend/core/pkg/scope"
	"backend/internal/api/handlers/shops/orders"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, scope *scope.Scope) {
	group := router.Group("/orders")

	group.Post("/", handler_orders.Create(scope)...)
}
