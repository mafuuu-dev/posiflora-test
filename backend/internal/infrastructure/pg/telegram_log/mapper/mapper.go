package pg_tgl_mapper

import (
	"backend/core/pkg/pgscan"
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/domain/telegram_log/entity"
)

func ToEntity(row pgscan.Scannable) (tgl_entity.TelegramLog, error) {
	var log tgl_entity.TelegramLog
	err := row.Scan(
		&log.ID,
		&log.ShopID,
		&log.OrderID,
		&log.Message,
		&log.Status,
		&log.Error,
		&log.SentAt,
	)

	return log, err
}

func ToStatsDto(row pgscan.Scannable) (tgl_dto.StatsDTO, error) {
	var stats tgl_dto.StatsDTO
	err := row.Scan(
		&stats.LastSentAt,
		&stats.SentCount,
		&stats.FailedCount,
	)

	return stats, err
}
