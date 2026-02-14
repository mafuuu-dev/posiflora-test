package handler_telegram

import (
	"backend/core/pkg/errorsx"
	"backend/core/pkg/request"
	"backend/core/pkg/response"
	"backend/core/pkg/scope"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/services/request"
	"backend/internal/usecases/telegram"

	"github.com/gofiber/fiber/v2"
)

type StatusHandler struct {
	*request.Handler
}

func Status(scope *scope.Scope) []fiber.Handler {
	handler := &StatusHandler{
		Handler: request.NewHandler(scope),
	}

	return handler.Instance(handler)
}

func (h *StatusHandler) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID, err := service_request.GetShopIDFromParams(c)
		if err != nil {
			return response.NotValidRequest(c, err)
		}

		statusIntegration := app_telegram.NewStatusIntegration(h.SC())
		integration, stats, err := statusIntegration.Get(*shopID)
		if err != nil {
			return response.NotValidRequest(c, errorsx.Wrapf(err, "Error getting telegram integration status"))
		}

		return response.Success(c, h.response(*integration, *stats))
	}
}

func (h *StatusHandler) response(integration tgi_entity.TelegramIntegration, stats tgl_dto.StatsDTO) any {
	var lastSentAt string
	if stats.LastSentAt != nil {
		lastSentAt = stats.LastSentAt.Format("2006-01-02 15:04:05")
	}

	return fiber.Map{
		"integration": fiber.Map{
			"is_enabled": integration.IsEnabled,
			"chat_id":    integration.ChatID,
		},
		"stats": fiber.Map{
			"last_sent_at": lastSentAt,
			"sent_count":   stats.SentCount,
			"failed_count": stats.FailedCount,
		},
	}
}
