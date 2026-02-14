package pg_tgi_mapper

import (
	"backend/core/pkg/pgscan"
	"backend/internal/domain/telegram_integration/entity"
)

func ToEntity(row pgscan.Scannable) (tgi_entity.TelegramIntegration, error) {
	var integration tgi_entity.TelegramIntegration
	err := row.Scan(
		&integration.ID,
		&integration.ShopID,
		&integration.BotToken,
		&integration.ChatID,
		&integration.IsEnabled,
		&integration.CreatedAt,
		&integration.UpdatedAt,
	)

	return integration, err
}
