package tgl_repository

import (
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/domain/telegram_log/entity"
	"backend/internal/domain/telegram_log/enum"
)

type TelegramLogRepository interface {
	GetStats(shopID int64) (*tgl_dto.StatsDTO, error)
	GetLog(shopID int64, orderID int64) (*tgl_entity.TelegramLog, error)

	CreateLog(
		shopID int64,
		orderID int64,
		message string,
		status tgl_enum.DBType,
		error string,
	) (*tgl_entity.TelegramLog, error)
}
