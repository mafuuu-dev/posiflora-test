package pg_tgl

import (
	"backend/core/pkg/repository"
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/domain/telegram_log/entity"
	"backend/internal/domain/telegram_log/enum"
	"backend/internal/infrastructure/pg/telegram_log/query"
)

type TelegramLogRepository struct {
	*repository.Repository
}

func NewRepository(factory *repository.Factory) *TelegramLogRepository {
	return &TelegramLogRepository{
		Repository: factory.Instance(),
	}
}

func (r *TelegramLogRepository) GetStats(shopID int64) (*tgl_dto.StatsDTO, error) {
	return pg_tgl_query.NewGetStats(r.Query()).
		Set(shopID).
		Execute()
}

func (r *TelegramLogRepository) CreateLog(
	shopID int64,
	orderID int64,
	message string,
	status tgl_enum.DBType,
	error string,
) (*tgl_entity.TelegramLog, error) {
	return pg_tgl_query.NewCreateLog(r.Query()).
		Set(shopID, orderID, message, status, error).
		Execute()
}
