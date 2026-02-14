package app_telegram

import (
	"backend/core/pkg/scope"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/domain/telegram_integration/repository"
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/domain/telegram_log/repository"
	"backend/internal/infrastructure/pg/telegram_integration"
	"backend/internal/infrastructure/pg/telegram_log"
)

type repositoriesStatusIntegration struct {
	TelegramIntegration tgi_repository.TelegramIntegrationRepository
	TelegramLog         tgl_repository.TelegramLogRepository
}

type StatusIntegration struct {
	scope        *scope.Scope
	repositories repositoriesStatusIntegration
}

func NewStatusIntegration(scope *scope.Scope) *StatusIntegration {
	return &StatusIntegration{
		scope: scope,
		repositories: repositoriesStatusIntegration{
			TelegramIntegration: pg_tgi.NewRepository(scope.Support.Factory.Repository),
			TelegramLog:         pg_tgl.NewRepository(scope.Support.Factory.Repository),
		},
	}
}

func (u *StatusIntegration) Get(shopID int64) (*tgi_entity.TelegramIntegration, *tgl_dto.StatsDTO, error) {
	integration, err := u.repositories.TelegramIntegration.GetIntegration(shopID)
	if err != nil {
		return nil, nil, err
	}
	if integration == nil {
		integration = &tgi_entity.TelegramIntegration{IsEnabled: false}
	}

	stats, err := u.repositories.TelegramLog.GetStats(shopID)
	if err != nil {
		return nil, nil, err
	}

	return integration, stats, nil
}
