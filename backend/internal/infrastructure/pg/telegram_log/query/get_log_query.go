package pg_tgl_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_log/entity"
	"backend/internal/infrastructure/pg/telegram_log/mapper"
)

type GetLog struct {
	*query.Query
	shopID  int64
	orderID int64
}

func NewGetLog(factory *query.Factory) *GetLog {
	return &GetLog{
		Query: factory.Instance(),
	}
}

func (q *GetLog) Set(shopID int64, orderID int64) *GetLog {
	q.shopID = shopID
	q.orderID = orderID
	return q
}

func (q *GetLog) Execute() (*tgl_entity.TelegramLog, error) {
	row := q.QueryRow(q, q.shopID, q.orderID)
	return pgscan.ScanOne[tgl_entity.TelegramLog](row, pg_tgl_mapper.ToEntity)
}

func (q *GetLog) Sql() string {
	return `
		SELECT id, shop_id, order_id, message, status, error, sent_at
		FROM telegram_send_log
		WHERE shop_id = $1 AND order_id = $2
		LIMIT 1
	`
}
