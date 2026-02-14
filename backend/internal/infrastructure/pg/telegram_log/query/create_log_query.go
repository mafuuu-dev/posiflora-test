package pg_tgl_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_log/entity"
	"backend/internal/domain/telegram_log/enum"
	"backend/internal/infrastructure/pg/telegram_log/mapper"
)

type CreateLog struct {
	*query.Query
	shopID  int64
	orderID int64
	message string
	status  tgl_enum.DBType
	error   string
}

func NewCreateLog(factory *query.Factory) *CreateLog {
	return &CreateLog{
		Query: factory.Instance(),
	}
}

func (q *CreateLog) Set(shopID int64, orderID int64, message string, status tgl_enum.DBType, error string) *CreateLog {
	q.shopID = shopID
	q.orderID = orderID
	q.message = message
	q.status = status
	q.error = error
	return q
}

func (q *CreateLog) Execute() (*tgl_entity.TelegramLog, error) {
	row := q.QueryRow(q, q.shopID, q.orderID, q.message, string(q.status), q.error)
	return pgscan.ScanOne[tgl_entity.TelegramLog](row, pg_tgl_mapper.ToEntity)
}

func (q *CreateLog) Sql() string {
	return `
		INSERT INTO telegram_send_log (shop_id, order_id, message, status, error)
		SELECT $1::BIGINT, $2::BIGINT, $3::TEXT, $4::TELEGRAM_SEND_STATUS, $5::TEXT
		RETURNING id, shop_id, order_id, message, status, error, sent_at
	`
}
