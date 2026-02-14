package shops

import (
	"backend/core/pkg/scope"
	"backend/internal/api/routes/shops/orders"
	"backend/internal/api/routes/shops/telegram"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, scope *scope.Scope) {
	group := router.Group("/shops/:shop_id<\\d+>")

	orders.RegisterRoutes(group, scope)
	telegram.RegisterRoutes(group, scope)
}
