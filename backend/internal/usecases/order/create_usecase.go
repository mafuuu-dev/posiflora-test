package app_order

import (
	"backend/core/pkg/scope"
	"backend/internal/domain/order/entity"
	"backend/internal/domain/order/repository"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/domain/telegram_integration/repository"
	"backend/internal/domain/telegram_log/entity"
	"backend/internal/domain/telegram_log/enum"
	"backend/internal/domain/telegram_log/repository"
	"backend/internal/domain/telegram_log/service"
	"backend/internal/infrastructure/pg/order"
	"backend/internal/infrastructure/pg/telegram_integration"
	"backend/internal/infrastructure/pg/telegram_log"
	"backend/internal/services/telegram"
)

type repositoriesCreateOrder struct {
	Order               order_repository.OrderRepository
	TelegramIntegration tgi_repository.TelegramIntegrationRepository
	TelegramLog         tgl_repository.TelegramLogRepository
}

type servicesCreateOrder struct {
	TelegramSender *service_telegram.TelegramSender
}

type CreateOrder struct {
	scope        *scope.Scope
	repositories repositoriesCreateOrder
	services     servicesCreateOrder
}

func NewCreateOrder(scope *scope.Scope) *CreateOrder {
	return &CreateOrder{
		scope: scope,
		repositories: repositoriesCreateOrder{
			Order:               pg_order.NewRepository(scope.Support.Factory.Repository),
			TelegramIntegration: pg_tgi.NewRepository(scope.Support.Factory.Repository),
			TelegramLog:         pg_tgl.NewRepository(scope.Support.Factory.Repository),
		},
		services: servicesCreateOrder{
			TelegramSender: service_telegram.NewTelegramSender(scope),
		},
	}
}

func (u *CreateOrder) Create(
	shopID int64,
	number string,
	total int64,
	customerName string,
) (*order_entity.Order, string, error) {
	order, err := u.repositories.Order.CreateOrder(shopID, number, total, customerName)
	if err != nil {
		return nil, string(tgl_enum.Error), err
	}

	integration, err := u.repositories.TelegramIntegration.GetIntegration(shopID)
	if err != nil {
		return order, string(tgl_enum.Error), err
	}

	if integration == nil || !integration.IsEnabled {
		return order, string(tgl_enum.Skipped), nil
	}

	log, err := u.sendMessage(*order, *integration)
	if err != nil {
		return order, string(tgl_enum.Error), err
	}

	return order, log.Status, nil
}

func (u *CreateOrder) sendMessage(
	o order_entity.Order,
	i tgi_entity.TelegramIntegration,
) (*tgl_entity.TelegramLog, error) {
	if log, err := u.repositories.TelegramLog.GetLog(o.ShopID, o.ID); err != nil || log != nil {
		return log, err
	}

	message := tgl_service.MakeLogMessage(o.Number, o.Total, o.CustomerName)
	if err := u.services.TelegramSender.SendMessage(i.BotToken, i.ChatID, message); err != nil {
		return u.repositories.TelegramLog.CreateLog(o.ShopID, o.ID, message, tgl_enum.Failed, err.Error())
	}

	return u.repositories.TelegramLog.CreateLog(o.ShopID, o.ID, message, tgl_enum.Sent, "")
}
