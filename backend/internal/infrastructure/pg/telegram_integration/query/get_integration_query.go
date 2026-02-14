package pg_tgi_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/infrastructure/pg/telegram_integration/mapper"
)

type GetIntegration struct {
	*query.Query
	shopID int64
}

func NewGetIntegration(factory *query.Factory) *GetIntegration {
	return &GetIntegration{
		Query: factory.Instance(),
	}
}

func (q *GetIntegration) Set(shopID int64) *GetIntegration {
	q.shopID = shopID
	return q
}

func (q *GetIntegration) Execute() (*tgi_entity.TelegramIntegration, error) {
	row := q.QueryRow(q, q.shopID)
	return pgscan.ScanOne[tgi_entity.TelegramIntegration](row, pg_tgi_mapper.ToEntity)
}

func (q *GetIntegration) Sql() string {
	return `
		SELECT id, shop_id, bot_token, chat_id, is_enabled, created_at, updated_at FROM telegram_integrations 
		WHERE shop_id = $1
		LIMIT 1
	`
}
