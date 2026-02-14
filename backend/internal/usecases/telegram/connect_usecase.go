package app_telegram

import (
	"backend/core/pkg/scope"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/domain/telegram_integration/repository"
	"backend/internal/infrastructure/pg/telegram_integration"
)

type repositoriesConnectIntegration struct {
	TelegramIntegration tgi_repository.TelegramIntegrationRepository
}

type ConnectIntegration struct {
	scope        *scope.Scope
	repositories repositoriesConnectIntegration
}

func NewConnectIntegration(scope *scope.Scope) *ConnectIntegration {
	return &ConnectIntegration{
		scope: scope,
		repositories: repositoriesConnectIntegration{
			TelegramIntegration: pg_tgi.NewRepository(scope.Support.Factory.Repository),
		},
	}
}

func (u *ConnectIntegration) Change(
	shopID int64,
	botToken string,
	chatID string,
	isEnabled bool,
) (*tgi_entity.TelegramIntegration, error) {
	return u.repositories.TelegramIntegration.UpsertIntegration(shopID, botToken, chatID, isEnabled)
}
