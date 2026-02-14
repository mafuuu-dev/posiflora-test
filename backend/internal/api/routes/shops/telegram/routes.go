package telegram

import (
	"backend/core/pkg/scope"
	"backend/internal/api/handlers/shops/telegram"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, scope *scope.Scope) {
	group := router.Group("/telegram")

	group.Get("/status", handler_telegram.Status(scope)...)
	group.Post("/connect", handler_telegram.Connect(scope)...)
}
