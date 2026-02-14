package pg_tgi

import (
	"backend/core/pkg/repository"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/infrastructure/pg/telegram_integration/query"
)

type TelegramIntegrationRepository struct {
	*repository.Repository
}

func NewRepository(factory *repository.Factory) *TelegramIntegrationRepository {
	return &TelegramIntegrationRepository{
		Repository: factory.Instance(),
	}
}

func (r *TelegramIntegrationRepository) GetIntegration(shopID int64) (*tgi_entity.TelegramIntegration, error) {
	return pg_tgi_query.NewGetIntegration(r.Query()).
		Set(shopID).
		Execute()
}

func (r *TelegramIntegrationRepository) UpsertIntegration(
	shopID int64,
	botToken string,
	chatID string,
	isEnabled bool,
) (*tgi_entity.TelegramIntegration, error) {
	return pg_tgi_query.NewUpsertIntegration(r.Query()).
		Set(shopID, botToken, chatID, isEnabled).
		Execute()
}
