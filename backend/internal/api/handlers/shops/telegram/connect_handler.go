package handler_telegram

import (
	"backend/core/pkg/errorsx"
	"backend/core/pkg/request"
	"backend/core/pkg/response"
	"backend/core/pkg/scope"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/services/request"
	"backend/internal/usecases/telegram"

	"github.com/gofiber/fiber/v2"
)

type connectHandler struct {
	BotToken  string `json:"bot_token" validate:"required_if=IsEnabled true"`
	ChatID    string `json:"chat_id" validate:"required_if=IsEnabled true"`
	IsEnabled *bool  `json:"is_enabled" validate:"required"`
}

type ConnectHandler struct {
	*request.Handler
}

func Connect(scope *scope.Scope) []fiber.Handler {
	handler := &ConnectHandler{
		Handler: request.NewHandler(scope),
	}

	return handler.Instance(handler)
}

func (h *ConnectHandler) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID, err := service_request.GetShopIDFromParams(c)
		if err != nil {
			return response.NotValidRequest(c, err)
		}

		model := &connectHandler{}
		if err := h.Validator().Validate(c, model); err != nil {
			return response.NotValidRequest(c, err)
		}

		connectIntegration := app_telegram.NewConnectIntegration(h.SC())
		integration, err := connectIntegration.Change(*shopID, model.BotToken, model.ChatID, *model.IsEnabled)
		if err != nil {
			return response.NotValidRequest(c, errorsx.Wrapf(err, "Error connecting telegram integration"))
		}

		return response.Success(c, h.response(*integration))
	}
}

func (h *ConnectHandler) response(integration tgi_entity.TelegramIntegration) any {
	return fiber.Map{
		"integration": fiber.Map{
			"bot_token":  integration.BotToken,
			"chat_id":    integration.ChatID,
			"is_enabled": integration.IsEnabled,
		},
	}
}
