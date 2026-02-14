package tgi_repository

import "backend/internal/domain/telegram_integration/entity"

type TelegramIntegrationRepository interface {
	GetIntegration(shopID int64) (*tgi_entity.TelegramIntegration, error)

	UpsertIntegration(
		shopID int64,
		botToken string,
		chatID string,
		isEnabled bool,
	) (*tgi_entity.TelegramIntegration, error)
}
